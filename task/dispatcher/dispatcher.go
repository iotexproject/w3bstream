package dispatcher

import (
	"log/slog"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/datasource"
	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/persistence/contract"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/task"
)

type NewDatasource func(datasourceURI string) (datasource.Datasource, error)

type Contract interface {
	LatestProjects() []*contract.Project
	LatestProvers() []*contract.Prover
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

func RunDispatcher(persistence Persistence, newDatasource NewDatasource,
	projectManager ProjectManager, bootNodeMultiaddr, operatorPrivateKey, operatorPrivateKeyED25519 string, sequencerPubKey []byte,
	iotexChainID int, projectNotification <-chan *contract.Project, contract Contract) error {
	projectDispatchers := &sync.Map{}
	d := &dispatcher{projectDispatchers: projectDispatchers}

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
	if p.Uri != "" {
		_, ok := projectDispatchers.Load(p.ID)
		if ok {
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
