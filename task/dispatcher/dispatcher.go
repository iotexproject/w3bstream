package dispatcher

import (
	"log/slog"
	"strconv"
	"sync"
	"time"

	pubsub "github.com/libp2p/go-libp2p-pubsub"

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
	LatestProject(projectID uint64) *contract.Project
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
	defaultDatasourceURI      string
	bootNodeMultiaddr         string
	operatorPrivateKey        string
	operatorPrivateKeyED25519 string
	sequencerPubKey           []byte
	iotexChainID              int
	projectNotification       <-chan uint64
	chainHeadNotification     <-chan uint64
	contract                  Contract
	taskStateHandler          *taskStateHandler
	windowSizeSetInterval     time.Duration
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

func (d *Dispatcher) setRequiredProverAmount(head uint64) {
	ps := d.projectOffsets.Projects(head)
	for _, p := range ps {
		cp := d.contract.Project(p.ID, head)
		if cp == nil {
			slog.Error("the contract project not exist when set window size", "project_id", p.ID, "block_number", head)
			continue
		}
		pd, ok := d.projectDispatchers.Load(p.ID)
		if !ok {
			continue
		}

		proverAmount := uint64(1)
		if v, ok := cp.Attributes[contract.RequiredProverAmount]; ok {
			n, err := strconv.ParseUint(string(v), 10, 64)
			if err != nil {
				slog.Error("failed to parse project required prover amount when set window size", "project_id", p.ID)
				continue
			}
			proverAmount = n
		}
		pd.(*projectDispatcher).requiredProverAmount.Store(proverAmount)
	}
}

func (d *Dispatcher) setProjectDispatcher(pid uint64) {
	cp := d.contract.LatestProject(pid)
	if cp == nil {
		slog.Error("the contract project not exist", "project_id", pid)
		return
	}
	ep, ok := d.projectDispatchers.Load(pid)
	if ok {
		ep.(*projectDispatcher).paused.Store(cp.Paused)
		return
	}
	if cp.Uri == "" {
		return
	}

	pf, err := d.projectManager.Project(cp.ID)
	if err != nil {
		slog.Error("failed to get project", "project_id", cp.ID, "error", err)
		return
	}
	if err := d.pubSubs.Add(cp.ID); err != nil {
		slog.Error("failed to add pubsubs", "project_id", cp.ID, "error", err)
		return
	}
	uri := pf.DatasourceURI
	if uri == "" {
		uri = d.defaultDatasourceURI
	}
	pd, err := newProjectDispatcher(d.persistence, uri, d.newDatasource, cp, d.pubSubs, d.taskStateHandler, d.sequencerPubKey)
	if err != nil {
		slog.Error("failed to new project dispatcher", "project_id", cp.ID, "error", err)
		return
	}
	d.projectDispatchers.Store(cp.ID, pd)
	slog.Info("a new project dispatcher started", "project_id", cp.ID)
}

func (d *Dispatcher) setWindowSize() {
	ticker := time.NewTicker(d.windowSizeSetInterval)
	for range ticker.C {
		provers := d.contract.LatestProvers()
		proverAmount := uint64(0)
		for _, p := range provers {
			if p.Paused {
				continue
			}
			proverAmount++
		}
		d.projectDispatchers.Range(func(k, v interface{}) bool {
			pd := v.(*projectDispatcher)
			if pd.idle.Load() || proverAmount == 0 {
				pd.window.setSize(0)
				return true
			}
			size := proverAmount
			if size > pd.requiredProverAmount.Load() {
				size = pd.requiredProverAmount.Load()
			}
			proverAmount -= size
			pd.window.setSize(size)
			return true
		})
	}
}

func (d *Dispatcher) Run() {
	if d.local {
		return
	}
	projects := d.contract.LatestProjects()
	for _, p := range projects {
		d.setProjectDispatcher(p.ID)
	}

	go func() {
		for pid := range d.projectNotification {
			slog.Info("get new project contract event", "project_id", pid)
			d.setProjectDispatcher(pid)
		}
	}()
	go func() {
		for head := range d.chainHeadNotification {
			d.setRequiredProverAmount(head)
		}
	}()
	go d.setWindowSize()
}

func New(persistence Persistence, newDatasource NewDatasource,
	projectManager ProjectManager, defaultDatasourceURI, bootNodeMultiaddr, operatorPrivateKey, operatorPrivateKeyED25519 string,
	sequencerPubKey []byte, iotexChainID int, projectNotification <-chan uint64, chainHeadNotification <-chan uint64,
	contract Contract, projectOffsets *scheduler.ProjectEpochOffsets) (*Dispatcher, error) {

	projectDispatchers := &sync.Map{}
	taskStateHandler := newTaskStateHandler(persistence, contract, projectManager, operatorPrivateKey, operatorPrivateKeyED25519)
	d := &Dispatcher{
		local:                     false,
		persistence:               persistence,
		newDatasource:             newDatasource,
		projectManager:            projectManager,
		defaultDatasourceURI:      defaultDatasourceURI,
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
		windowSizeSetInterval:     5 * time.Second,
	}
	ps, err := p2p.NewPubSubs(d.handleP2PData, bootNodeMultiaddr, iotexChainID)
	if err != nil {
		return nil, err
	}
	d.pubSubs = ps
	return d, nil
}
