package task

import (
	"log/slog"
	"sync"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/persistence/contract"
	"github.com/machinefi/sprout/project"
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
	getProject handler.Project, bootNodeMultiaddr, operatorPrivateKey, operatorPrivateKeyED25519 string,
	iotexChainID int, projectNotification <-chan *contract.Project, latestProjects LatestProjects) error {
	projectDispatchers := &sync.Map{}
	d := &dispatcher{projectDispatchers: projectDispatchers}

	ps, err := p2p.NewPubSubs(d.handleP2PData, bootNodeMultiaddr, iotexChainID)
	if err != nil {
		return err
	}

	taskStateHandler := handler.NewTaskStateHandler(persistence.Create, getProject, operatorPrivateKey, operatorPrivateKeyED25519)

	return dispatch(persistence, newDatasource, getProject, projectDispatchers, ps, taskStateHandler,
		projectNotification, latestProjects)
}

func RunLocalDispatcher(persistence Persistence, newDatasource internaldispatcher.NewDatasource,
	getProjectIDs ProjectIDs, getProject handler.Project, operatorPrivateKey, operatorPrivateKeyED25519, bootNodeMultiaddr string, iotexChainID int) error {
	projectDispatchers := &sync.Map{}
	d := &dispatcher{projectDispatchers: projectDispatchers}

	ps, err := p2p.NewPubSubs(d.handleP2PData, bootNodeMultiaddr, iotexChainID)
	if err != nil {
		return err
	}

	taskStateHandler := handler.NewTaskStateHandler(persistence.Create, getProject, operatorPrivateKey, operatorPrivateKeyED25519)

	projectIDs := getProjectIDs()
	for _, id := range projectIDs {
		pm := &project.Meta{
			ProjectID: id,
		}
		_, ok := projectDispatchers.Load(id)
		if ok {
			continue
		}
		p, err := getProject(id)
		if err != nil {
			return errors.Wrapf(err, "failed to get project, project_id %v", id)
		}
		ps.Add(id)
		pd, err := internaldispatcher.NewProjectDispatcher(persistence.ProcessedTaskID,
			persistence.UpsertProcessedTask, p.DatasourceURI, newDatasource, pm, ps.Publish, taskStateHandler)
		if err != nil {
			return errors.Wrapf(err, "failed to new project dispatcher, project_id %v", id)
		}
		projectDispatchers.Store(id, pd)
	}
	return nil
}

func setProjectDispatcher(persistence Persistence, newDatasource internaldispatcher.NewDatasource, projectDispatchers *sync.Map, p *contract.Project, getProject handler.Project, ps *p2p.PubSubs, handler *handler.TaskStateHandler) {
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
		pm := &project.Meta{
			ProjectID: p.ID,
			Uri:       p.Uri,
			Hash:      p.Hash,
		}
		pd, err := internaldispatcher.NewProjectDispatcher(persistence.ProcessedTaskID,
			persistence.UpsertProcessedTask, pf.DatasourceURI, newDatasource, pm, ps.Publish, handler)
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
	projectNotification <-chan *contract.Project, latestProjects LatestProjects) error {

	projects := latestProjects()
	for _, p := range projects {
		setProjectDispatcher(persistence, newDatasource, projectDispatchers, p, getProject, ps, handler)
	}

	go func() {
		for p := range projectNotification {
			slog.Info("get new project contract event", "project_id", p.ID, "block_number", p.BlockNumber)
			setProjectDispatcher(persistence, newDatasource, projectDispatchers, p, getProject, ps, handler)
		}
	}()
	return nil
}
