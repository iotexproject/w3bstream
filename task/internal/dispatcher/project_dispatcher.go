package dispatcher

import (
	"log/slog"
	"time"

	"github.com/machinefi/sprout/datasource"
	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/task/internal/handler"
	"github.com/machinefi/sprout/types"
	"github.com/pkg/errors"
)

type NewDatasource func(datasourceURI string) (datasource.Datasource, error)

type Publish func(projectID uint64, data *p2p.Data) error

type FetchProcessedTaskID func(projectID uint64) (uint64, error)

type UpsertProcessedTask func(projectID, taskID uint64) error

type ProjectDispatcher struct {
	window       *window
	waitInterval time.Duration
	startTaskID  uint64
	projectID    uint64
	datasource   datasource.Datasource
	publish      Publish
}

func (d *ProjectDispatcher) Handle(s *types.TaskStateLog) {
	d.window.consume(s)
}

func (d *ProjectDispatcher) run() {
	nextTaskID := d.startTaskID
	for {
		next, err := d.dispatch(nextTaskID)
		if err != nil {
			slog.Error("failed to dispatch task", "error", err, "project_id", d.projectID)
			continue
		}
		if nextTaskID == next {
			time.Sleep(d.waitInterval)
		}
		nextTaskID = next
	}
}

func (d *ProjectDispatcher) dispatch(nextTaskID uint64) (uint64, error) {
	t, err := d.datasource.Retrieve(d.projectID, nextTaskID)
	if err != nil {
		return 0, errors.Wrap(err, "failed to retrieve task from data source")
	}
	if t == nil {
		return nextTaskID, nil
	}
	d.window.produce(t)

	if err := d.publish(t.ProjectID, &p2p.Data{Task: t}); err != nil {
		return 0, errors.Wrapf(err, "failed to publish data, project_id %v, task_id %v", t.ProjectID, t.ID)
	}
	slog.Debug("dispatched a task", "project_id", t.ProjectID, "task_id", t.ID)
	return t.ID + 1, nil
}

func NewProjectDispatcher(fetch FetchProcessedTaskID, upsert UpsertProcessedTask, datasourceURI string, newDatasource NewDatasource, projectMeta *project.ProjectMeta, publish Publish, handler *handler.TaskStateHandler) (*ProjectDispatcher, error) {
	nextTaskID, err := fetch(projectMeta.ProjectID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to fetch next task_id, project_id %v", projectMeta.ProjectID)
	}
	datasource, err := newDatasource(datasourceURI)
	if err != nil {
		return nil, errors.Wrap(err, "failed to new task retriever")
	}
	windowSize := projectMeta.ProverAmount
	if windowSize == 0 {
		windowSize = 1
	}
	window := newWindow(windowSize, publish, handler, upsert)
	d := &ProjectDispatcher{
		window:       window,
		waitInterval: 3 * time.Second,
		startTaskID:  nextTaskID,
		datasource:   datasource,
		projectID:    projectMeta.ProjectID,
		publish:      publish,
	}
	go d.run()
	return d, nil
}
