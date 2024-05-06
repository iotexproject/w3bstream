package dispatcher

import (
	"log/slog"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/machinefi/sprout/datasource"
	"github.com/machinefi/sprout/metrics"
	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/persistence/contract"
	"github.com/machinefi/sprout/task"
)

type projectDispatcher struct {
	window          *window
	waitInterval    time.Duration
	startTaskID     uint64
	projectID       uint64
	datasource      datasource.Datasource
	pubSubs         *p2p.PubSubs
	sequencerPubKey []byte
}

func (d *projectDispatcher) handle(s *task.StateLog) {
	d.window.consume(s)
}

func (d *projectDispatcher) run() {
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

func (d *projectDispatcher) dispatch(nextTaskID uint64) (uint64, error) {
	t, err := d.datasource.Retrieve(d.projectID, nextTaskID)
	if err != nil {
		return 0, errors.Wrap(err, "failed to retrieve task from data source")
	}
	if t == nil {
		return nextTaskID, nil
	}
	if err := t.VerifySignature(d.sequencerPubKey); err != nil {
		return 0, errors.Wrap(err, "failed to verify task signature")
	}

	d.window.produce(t)

	metrics.DispatchedTaskNumMtc(d.projectID, t.ProjectVersion)

	if err := d.pubSubs.Publish(t.ProjectID, &p2p.Data{Task: t}); err != nil {
		return 0, errors.Wrapf(err, "failed to publish data, project_id %v, task_id %v", t.ProjectID, t.ID)
	}
	slog.Debug("dispatched a task", "project_id", t.ProjectID, "task_id", t.ID)
	return t.ID + 1, nil
}

func newProjectDispatcher(persistence Persistence, datasourceURI string, newDatasource NewDatasource, p *contract.Project, pubSubs *p2p.PubSubs, handler *taskStateHandler, sequencerPubKey []byte) (*projectDispatcher, error) {
	processedTaskID, err := persistence.ProcessedTaskID(p.ID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to fetch next task_id, project_id %v", p.ID)
	}
	datasource, err := newDatasource(datasourceURI)
	if err != nil {
		return nil, errors.Wrap(err, "failed to new task retriever")
	}

	proverAmount := uint64(1)
	if v, ok := p.Attributes[contract.RequiredProverAmountHash]; ok {
		n, err := strconv.ParseUint(string(v), 10, 64)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to parse project required prover amount, project_id %v", p.ID)
		}
		proverAmount = n
	}

	window := newWindow(proverAmount, pubSubs, handler, persistence)
	d := &projectDispatcher{
		window:          window,
		waitInterval:    3 * time.Second,
		startTaskID:     processedTaskID + 1,
		datasource:      datasource,
		projectID:       p.ID,
		pubSubs:         pubSubs,
		sequencerPubKey: sequencerPubKey,
	}
	go d.run()
	return d, nil
}
