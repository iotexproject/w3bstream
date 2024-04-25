package task

import (
	"log/slog"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/persistence/contract"
	internaldispatcher "github.com/machinefi/sprout/task/internal/dispatcher"
	"github.com/machinefi/sprout/task/internal/handler"
	"github.com/machinefi/sprout/types"
)

type LatestProjects func() []*contract.Project

type Persistence interface {
	Create(tl *types.TaskStateLog, t *types.Task) error
	ProcessedTaskID(projectID uint64) (uint64, error)
	UpsertProcessedTask(projectID, taskID uint64) error
}

type Datasource interface {
	Retrieve(projectID, nextTaskID uint64) (*types.Task, error)
}

type ProjectIDs func() []uint64

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
	pd.(*internaldispatcher.ProjectDispatcher).Handle(s)
}

func RunDispatcher(persistence Persistence, newDatasource internaldispatcher.NewDatasource,
<<<<<<< HEAD
	getProject handler.Project, bootNodeMultiaddr, operatorPrivateKey, operatorPrivateKeyED25519 string, sequencerPubKey []byte,
	iotexChainID int, projectNotification <-chan *contract.Project, latestProjects LatestProjects, latestProvers handler.LatestProvers) error {
=======
	getProject handler.Project, bootNodeMultiaddr, operatorPrivateKey, operatorPrivateKeyED25519 string,
	iotexChainID int, projectNotification <-chan *contract.Project, latestProjects LatestProjects) error {
>>>>>>> origin/develop
	projectDispatchers := &sync.Map{}
	d := &dispatcher{projectDispatchers: projectDispatchers}

	ps, err := p2p.NewPubSubs(d.handleP2PData, bootNodeMultiaddr, iotexChainID)
	if err != nil {
		return err
	}

<<<<<<< HEAD
	taskStateHandler := handler.NewTaskStateHandler(persistence.Create, latestProvers, getProject, operatorPrivateKey, operatorPrivateKeyED25519)

	return dispatch(persistence, newDatasource, getProject, projectDispatchers, ps, taskStateHandler,
		projectNotification, latestProjects, sequencerPubKey)
}

func RunLocalDispatcher(persistence Persistence, newDatasource internaldispatcher.NewDatasource,
	getProjectIDs ProjectIDs, getProject handler.Project, operatorPrivateKey, operatorPrivateKeyED25519, bootNodeMultiaddr string, sequencerPubKey []byte, iotexChainID int) error {
=======
	taskStateHandler := handler.NewTaskStateHandler(persistence.Create, getProject, operatorPrivateKey, operatorPrivateKeyED25519)

	return dispatch(persistence, newDatasource, getProject, projectDispatchers, ps, taskStateHandler,
		projectNotification, latestProjects)
}

func RunLocalDispatcher(persistence Persistence, newDatasource internaldispatcher.NewDatasource,
	getProjectIDs ProjectIDs, getProject handler.Project, operatorPrivateKey, operatorPrivateKeyED25519, bootNodeMultiaddr string, iotexChainID int) error {
>>>>>>> origin/develop
	projectDispatchers := &sync.Map{}
	d := &dispatcher{projectDispatchers: projectDispatchers}

	ps, err := p2p.NewPubSubs(d.handleP2PData, bootNodeMultiaddr, iotexChainID)
	if err != nil {
		return err
	}

<<<<<<< HEAD
	taskStateHandler := handler.NewTaskStateHandler(persistence.Create, nil, getProject, operatorPrivateKey, operatorPrivateKeyED25519)
=======
	taskStateHandler := handler.NewTaskStateHandler(persistence.Create, getProject, operatorPrivateKey, operatorPrivateKeyED25519)
>>>>>>> origin/develop

	projectIDs := getProjectIDs()
	for _, id := range projectIDs {
		_, ok := projectDispatchers.Load(id)
		if ok {
			continue
		}
		p, err := getProject(id)
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
		pd, err := internaldispatcher.NewProjectDispatcher(persistence.ProcessedTaskID,
<<<<<<< HEAD
			persistence.UpsertProcessedTask, p.DatasourceURI, newDatasource, cp, ps.Publish, taskStateHandler, sequencerPubKey)
=======
			persistence.UpsertProcessedTask, p.DatasourceURI, newDatasource, cp, ps.Publish, taskStateHandler)
>>>>>>> origin/develop
		if err != nil {
			return errors.Wrapf(err, "failed to new project dispatcher, project_id %v", id)
		}
		projectDispatchers.Store(id, pd)
	}
	return nil
}

<<<<<<< HEAD
func setProjectDispatcher(persistence Persistence, newDatasource internaldispatcher.NewDatasource, projectDispatchers *sync.Map, p *contract.Project, getProject handler.Project, ps *p2p.PubSubs, handler *handler.TaskStateHandler, sequencerPubKey []byte) {
=======
func setProjectDispatcher(persistence Persistence, newDatasource internaldispatcher.NewDatasource, projectDispatchers *sync.Map, p *contract.Project, getProject handler.Project, ps *p2p.PubSubs, handler *handler.TaskStateHandler) {
>>>>>>> origin/develop
	if p.Uri != "" {
		_, ok := projectDispatchers.Load(p.ID)
		if ok {
			return
		}
		pf, err := getProject(p.ID)
		if err != nil {
			slog.Error("failed to get project", "project_id", p.ID, "error", err)
			return
		}
		if err := ps.Add(p.ID); err != nil {
			slog.Error("failed to add pubsubs", "project_id", p.ID, "error", err)
			return
		}
		pd, err := internaldispatcher.NewProjectDispatcher(persistence.ProcessedTaskID,
<<<<<<< HEAD
			persistence.UpsertProcessedTask, pf.DatasourceURI, newDatasource, p, ps.Publish, handler, sequencerPubKey)
=======
			persistence.UpsertProcessedTask, pf.DatasourceURI, newDatasource, p, ps.Publish, handler)
>>>>>>> origin/develop
		if err != nil {
			slog.Error("failed to new project dispatcher", "project_id", p.ID, "error", err)
			return
		}
		projectDispatchers.Store(p.ID, pd)
		slog.Info("a new project dispatcher started", "project_id", p.ID)
	}
}

func dispatch(persistence Persistence, newDatasource internaldispatcher.NewDatasource, getProject handler.Project,
	projectDispatchers *sync.Map, ps *p2p.PubSubs, handler *handler.TaskStateHandler,
<<<<<<< HEAD
	projectNotification <-chan *contract.Project, latestProjects LatestProjects, sequencerPubKey []byte) error {

	projects := latestProjects()
	for _, p := range projects {
		setProjectDispatcher(persistence, newDatasource, projectDispatchers, p, getProject, ps, handler, sequencerPubKey)
=======
	projectNotification <-chan *contract.Project, latestProjects LatestProjects) error {

	projects := latestProjects()
	for _, p := range projects {
		setProjectDispatcher(persistence, newDatasource, projectDispatchers, p, getProject, ps, handler)
>>>>>>> origin/develop
	}

	go func() {
		for p := range projectNotification {
			slog.Info("get new project contract event", "project_id", p.ID, "block_number", p.BlockNumber)
<<<<<<< HEAD
			setProjectDispatcher(persistence, newDatasource, projectDispatchers, p, getProject, ps, handler, sequencerPubKey)
=======
			setProjectDispatcher(persistence, newDatasource, projectDispatchers, p, getProject, ps, handler)
>>>>>>> origin/develop
		}
	}()
	return nil
}
