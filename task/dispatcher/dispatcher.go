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

type Dispatcher struct {
	local                     bool
	projectDispatchers        *sync.Map // projectID(uint64) -> *ProjectDispatcher
	projectOffsets            *scheduler.ProjectEpochOffsets
	pubSubs                   *p2p.PubSubs
	persistence               Persistence
	newDatasource             NewDatasource
	projectManager            ProjectManager
	bootNodeMultiaddr         string
	operatorPrivateKey        string
	operatorPrivateKeyED25519 string
	sequencerPubKey           []byte
	iotexChainID              int
	projectNotification       <-chan *contract.Project
	chainHeadNotification     <-chan uint64
	contract                  Contract
	taskStateHandler          *taskStateHandler
}

func (d *Dispatcher) handleP2PData(data *p2p.Data, topic *pubsub.Topic) {
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

func (d *Dispatcher) setWindowSize(head uint64) {
	ps := d.projectOffsets.Projects(head)
	for _, p := range ps {
		cp := d.contract.Project(p.ID, head)
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

func New(persistence Persistence, newDatasource NewDatasource,
	projectManager ProjectManager, bootNodeMultiaddr, operatorPrivateKey, operatorPrivateKeyED25519 string,
	sequencerPubKey []byte, iotexChainID int, projectNotification <-chan *contract.Project, chainHeadNotification <-chan uint64,
	contract Contract, projectOffsets *scheduler.ProjectEpochOffsets) (*Dispatcher, error) {

	projectDispatchers := &sync.Map{}
	taskStateHandler := newTaskStateHandler(persistence, contract, projectManager, operatorPrivateKey, operatorPrivateKeyED25519)
	d := &Dispatcher{
		local:                     false,
		persistence:               persistence,
		newDatasource:             newDatasource,
		projectManager:            projectManager,
		bootNodeMultiaddr:         bootNodeMultiaddr,
		operatorPrivateKey:        operatorPrivateKey,
		operatorPrivateKeyED25519: operatorPrivateKeyED25519,
		sequencerPubKey:           sequencerPubKey,
		iotexChainID:              iotexChainID,
		projectNotification:       projectNotification,
		chainHeadNotification:     chainHeadNotification,
		contract:                  contract,
		projectOffsets:            projectOffsets,
		projectDispatchers:        projectDispatchers,
		taskStateHandler:          taskStateHandler,
	}
	ps, err := p2p.NewPubSubs(d.handleP2PData, bootNodeMultiaddr, iotexChainID)
	if err != nil {
		return nil, err
	}
	d.pubSubs = ps
	return d, nil
}

func NewLocal(persistence Persistence, newDatasource NewDatasource,
	projectManager ProjectManager, operatorPrivateKey, operatorPrivateKeyED25519, bootNodeMultiaddr string,
	sequencerPubKey []byte, iotexChainID int) (*Dispatcher, error) {

	projectDispatchers := &sync.Map{}
	taskStateHandler := newTaskStateHandler(persistence, nil, projectManager, operatorPrivateKey, operatorPrivateKeyED25519)
	d := &Dispatcher{
		local:              true,
		projectDispatchers: projectDispatchers,
	}
	ps, err := p2p.NewPubSubs(d.handleP2PData, bootNodeMultiaddr, iotexChainID)
	if err != nil {
		return nil, err
	}

	projectIDs := projectManager.ProjectIDs()
	for _, id := range projectIDs {
		_, ok := projectDispatchers.Load(id)
		if ok {
			continue
		}
		p, err := projectManager.Project(id)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get project, project_id %v", id)
		}
		if err := ps.Add(id); err != nil {
			return nil, errors.Wrapf(err, "failed to add pubsubs, project_id %v", id)
		}
		cp := &contract.Project{
			ID:         id,
			Attributes: map[common.Hash][]byte{},
		}
		pd, err := newProjectDispatcher(persistence, p.DatasourceURI, newDatasource, cp, ps, taskStateHandler, sequencerPubKey)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to new project dispatcher, project_id %v", id)
		}
		projectDispatchers.Store(id, pd)
	}
	return d, nil
}

func (d *Dispatcher) setProjectDispatcher(p *contract.Project) {
	ep, ok := d.projectDispatchers.Load(p.ID)
	if ok {
		if p.Paused != nil {
			ep.(*projectDispatcher).paused.Store(*p.Paused)
		}
		return
	}
	if p.Uri == "" {
		return
	}
	pf, err := d.projectManager.Project(p.ID)
	if err != nil {
		slog.Error("failed to get project", "project_id", p.ID, "error", err)
		return
	}
	if err := d.pubSubs.Add(p.ID); err != nil {
		slog.Error("failed to add pubsubs", "project_id", p.ID, "error", err)
		return
	}
	pd, err := newProjectDispatcher(d.persistence, pf.DatasourceURI, d.newDatasource, p, d.pubSubs, d.taskStateHandler, d.sequencerPubKey)
	if err != nil {
		slog.Error("failed to new project dispatcher", "project_id", p.ID, "error", err)
		return
	}
	d.projectDispatchers.Store(p.ID, pd)
	slog.Info("a new project dispatcher started", "project_id", p.ID)
}

func (d *Dispatcher) Run() {
	if d.local {
		return
	}
	projects := d.contract.LatestProjects()
	for _, p := range projects {
		d.setProjectDispatcher(p)
	}

	go func() {
		for p := range d.projectNotification {
			slog.Info("get new project contract event", "project_id", p.ID, "block_number", p.BlockNumber)
			d.setProjectDispatcher(p)
		}
	}()
	go func() {
		for head := range d.chainHeadNotification {
			d.setWindowSize(head)
		}
	}()
}
