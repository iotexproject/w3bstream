package dispatcher

import (
	"log/slog"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/datasource"
	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/persistence/contract"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/scheduler"
	"github.com/machinefi/sprout/task"
)

type NewDatasource func(datasourceURI string) (datasource.Datasource, error)

type Contract interface {
	LatestProjects() []*contract.Project
	LatestProvers() []*contract.Prover
	Project(projectID, blockNumber uint64) *contract.Project
}

type ProjectManager interface {
	ProjectIDs() []uint64
	Project(projectID uint64) (*project.Project, error)
}

type Persistence interface {
	Create(tl *task.StateLog, t *task.Task) error
	ProcessedTaskID(projectID uint64) (uint64, error)
	UpsertProcessedTask(projectID, taskID uint64) error
}

type dispatcher struct {
	projectDispatchers *sync.Map // projectID(uint64) -> *ProjectDispatcher
	projectOffsets     *scheduler.ProjectEpochOffsets
}

func (d *dispatcher) handleP2PData(data *p2p.Data, topic *pubsub.Topic) {
	if data.TaskStateLog == nil {
		return
	}
	s := data.TaskStateLog

	pd, ok := d.projectDispatchers.Load(s.ProjectID)
	if !ok {
		slog.Error("the project dispatcher not exist", "project_id", s.ProjectID)
		return
	}
	pd.(*projectDispatcher).handle(s)
}

func (d *dispatcher) setWindowSize(chainHead <-chan uint64, c Contract) {
	for head := range chainHead {
		ps := d.projectOffsets.Projects(head)
		for _, p := range ps {
			cp := c.Project(p.ID, head)
			if cp == nil {
				slog.Error("the contract project not exist when set window size", "project_id", p.ID, "block_number", head)
				continue
			}
			pd, ok := d.projectDispatchers.Load(p.ID)
			if !ok {
				slog.Error("the project dispatcher not exist when set window size", "project_id", p.ID)
				continue
			}

			proverAmount := uint64(1)
			if v, ok := cp.Attributes[contract.RequiredProverAmountHash]; ok {
				n, err := strconv.ParseUint(string(v), 10, 64)
				if err != nil {
					slog.Error("failed to parse project required prover amount when set window size", "project_id", p.ID)
					continue
				}
				proverAmount = n
			}
			pd.(*projectDispatcher).window.size.Store(proverAmount)
		}
	}
}

func RunDispatcher(persistence Persistence, newDatasource NewDatasource,
	projectManager ProjectManager, bootNodeMultiaddr, operatorPrivateKey, operatorPrivateKeyED25519 string, sequencerPubKey []byte,
	iotexChainID int, projectNotification <-chan *contract.Project, chainHead <-chan uint64, contract Contract, projectOffsets *scheduler.ProjectEpochOffsets) error {
	projectDispatchers := &sync.Map{}
	d := &dispatcher{
		projectDispatchers: projectDispatchers,
		projectOffsets:     projectOffsets,
	}

	ps, err := p2p.NewPubSubs(d.handleP2PData, bootNodeMultiaddr, iotexChainID)
	if err != nil {
		return err
	}

	taskStateHandler := newTaskStateHandler(persistence, contract, projectManager, operatorPrivateKey, operatorPrivateKeyED25519)

	return dispatch(persistence, newDatasource, projectManager, projectDispatchers, ps, taskStateHandler,
		projectNotification, contract, sequencerPubKey)
}

func RunLocalDispatcher(persistence Persistence, newDatasource NewDatasource,
	projectManager ProjectManager, operatorPrivateKey, operatorPrivateKeyED25519, bootNodeMultiaddr string, sequencerPubKey []byte, iotexChainID int) error {
	projectDispatchers := &sync.Map{}
	d := &dispatcher{projectDispatchers: projectDispatchers}

	ps, err := p2p.NewPubSubs(d.handleP2PData, bootNodeMultiaddr, iotexChainID)
	if err != nil {
		return err
	}

	taskStateHandler := newTaskStateHandler(persistence, nil, projectManager, operatorPrivateKey, operatorPrivateKeyED25519)

	projectIDs := projectManager.ProjectIDs()
	for _, id := range projectIDs {
		_, ok := projectDispatchers.Load(id)
		if ok {
			continue
		}
		p, err := projectManager.Project(id)
		if err != nil {
			return errors.Wrapf(err, "failed to get project, project_id %v", id)
		}
		if err := ps.Add(id); err != nil {
			return errors.Wrapf(err, "failed to add pubsubs, project_id %v", id)
		}
		cp := &contract.Project{
			ID:         id,
			Attributes: map[common.Hash][]byte{},
		}
		pd, err := newProjectDispatcher(persistence, p.DatasourceURI, newDatasource, cp, ps, taskStateHandler, sequencerPubKey)
		if err != nil {
			return errors.Wrapf(err, "failed to new project dispatcher, project_id %v", id)
		}
		projectDispatchers.Store(id, pd)
	}
	return nil
}

func setProjectDispatcher(persistence Persistence, newDatasource NewDatasource, projectDispatchers *sync.Map, p *contract.Project, projectManager ProjectManager, ps *p2p.PubSubs, handler *taskStateHandler, sequencerPubKey []byte) {
	ep, ok := projectDispatchers.Load(p.ID)
	if ok {
		if p.Paused != nil {
			ep.(*projectDispatcher).paused.Store(*p.Paused)
		}
		return
	}
	pf, err := projectManager.Project(p.ID)
	if err != nil {
		slog.Error("failed to get project", "project_id", p.ID, "error", err)
		return
	}
	if err := ps.Add(p.ID); err != nil {
		slog.Error("failed to add pubsubs", "project_id", p.ID, "error", err)
		return
	}
	pd, err := newProjectDispatcher(persistence, pf.DatasourceURI, newDatasource, p, ps, handler, sequencerPubKey)
	if err != nil {
		slog.Error("failed to new project dispatcher", "project_id", p.ID, "error", err)
		return
	}
	projectDispatchers.Store(p.ID, pd)
	slog.Info("a new project dispatcher started", "project_id", p.ID)
}

func dispatch(persistence Persistence, newDatasource NewDatasource, projectManager ProjectManager,
	projectDispatchers *sync.Map, ps *p2p.PubSubs, handler *taskStateHandler,
	projectNotification <-chan *contract.Project, contract Contract, sequencerPubKey []byte) error {

	projects := contract.LatestProjects()
	for _, p := range projects {
		setProjectDispatcher(persistence, newDatasource, projectDispatchers, p, projectManager, ps, handler, sequencerPubKey)
	}

	go func() {
		for p := range projectNotification {
			slog.Info("get new project contract event", "project_id", p.ID, "block_number", p.BlockNumber)
			setProjectDispatcher(persistence, newDatasource, projectDispatchers, p, projectManager, ps, handler, sequencerPubKey)
		}
	}()
	return nil
}
