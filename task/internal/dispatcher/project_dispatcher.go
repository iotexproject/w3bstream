package dispatcher

import (
	"log/slog"
	"time"

	"github.com/pkg/errors"

	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/types"
)

type ProjectDispatcher struct {
	window        *window
	waitInterval  time.Duration
	startTaskID   uint64
	datasourceURI string
	datasource    Datasource
	persistence   Persistence
	projectID     uint64
	publisher     Publisher
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

	if err := d.publisher.Publish(t.ProjectID, &p2p.Data{Task: t}); err != nil {
		return 0, errors.Wrapf(err, "failed to publish data, project_id %v, task_id %v", t.ProjectID, t.ID)
	}
	slog.Debug("dispatched a task", "project_id", t.ProjectID, "task_id", t.ID)
	return t.ID + 1, nil
}

func NewProjectDispatcher(datasourceURI string, newDatasource NewDatasource, projectMeta *project.ProjectMeta, publisher Publisher, persistence Persistence, projectConfigManager ProjectConfigManager, operatorPrivateKeyECDSA, operatorPrivateKeyED25519 string) (*ProjectDispatcher, error) {
	nextTaskID, err := persistence.FetchNextTaskID(projectMeta.ProjectID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to fetch next task_id, project_id %v", projectMeta.ProjectID)
	}
	datasource, err := newDatasource(datasourceURI)
	if err != nil {
		return nil, errors.Wrap(err, "failed to new data source")
	}
	windowSize := projectMeta.ProverAmount
	if windowSize == 0 {
		windowSize = 1
	}
	window := newWindow(windowSize, persistence, projectConfigManager, publisher, operatorPrivateKeyECDSA, operatorPrivateKeyED25519)
	d := &ProjectDispatcher{
		window:        window,
		waitInterval:  3 * time.Second,
		startTaskID:   nextTaskID,
		datasourceURI: datasourceURI,
		datasource:    datasource,
		persistence:   persistence,
		projectID:     projectMeta.ProjectID,
		publisher:     publisher,
	}
	go d.run()
	return d, nil
}
