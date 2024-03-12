package task

import (
	"fmt"
	"log/slog"
	"time"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/datasource"
	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/persistence"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/types"
)

type Dispatcher struct {
	source                    datasource.Datasource
	pubSubs                   *p2p.PubSubs
	pg                        *persistence.Postgres
	projectManager            *project.Manager
	operatorPrivateKeyECDSA   string
	operatorPrivateKeyED25519 string
}

// will block caller
func (d *Dispatcher) Dispatch() {
	ticker := time.NewTicker(3 * time.Second)
	nextTaskID := uint64(0)

	for range ticker.C {
		next, err := d.dispatchTask(nextTaskID)
		if err != nil {
			slog.Error("failed to dispatch task", "error", err)
			continue
		}
		nextTaskID = next
	}
}

func (d *Dispatcher) dispatchTask(nextTaskID uint64) (uint64, error) {
	t, err := d.source.Retrieve(nextTaskID)
	if err != nil {
		return 0, errors.Wrap(err, "failed to retrieve task from data source")
	}
	if t == nil {
		return nextTaskID, nil
	}
	if err := d.pubSubs.Add(t.ProjectID); err != nil {
		return 0, errors.Wrapf(err, "failed to add project pubsub, project_id %v", t.ProjectID)
	}
	if err := d.pubSubs.Publish(t.ProjectID, &p2p.Data{
		Task: t,
	}); err != nil {
		return 0, errors.Wrapf(err, "failed to publish data, project_id %v", t.ProjectID)
	}
	slog.Debug("dispatched a task", "project_id", t.ProjectID, "task_id", t.ID)
	return t.ID + 1, nil
}

func (d *Dispatcher) handleP2PData(data *p2p.Data, topic *pubsub.Topic) {
	if data.TaskStateLog == nil {
		return
	}
	l := data.TaskStateLog
	if err := d.pg.Create(l); err != nil {
		slog.Error("failed to create task state log", "error", err, "taskID", l.Task.ID)
		return
	}
	if l.State != types.TaskStateProved {
		return
	}

	config, err := d.projectManager.Get(l.Task.ProjectID, l.Task.ProjectVersion)
	if err != nil {
		slog.Error("failed to get project", "error", err, "project_id", l.Task.ProjectID, "project_version", l.Task.ProjectVersion)
		return
	}

	output, err := config.GetOutput(d.operatorPrivateKeyECDSA, d.operatorPrivateKeyED25519)
	if err != nil {
		slog.Error("failed to init output", "error", err)
		if err := d.pg.Create(&types.TaskStateLog{
			Task:      l.Task,
			State:     types.TaskStateFailed,
			Comment:   err.Error(),
			CreatedAt: time.Now(),
		}); err != nil {
			slog.Error("failed to create failed task state", "error", err, "taskID", l.Task.ID)
		}
		return
	}

	slog.Debug("output proof", "outputter", fmt.Sprintf("%T", output))

	outRes, err := output.Output(&l.Task, l.Result)
	if err != nil {
		slog.Error("failed to output", "error", err, "taskID", l.Task.ID)
		if err := d.pg.Create(&types.TaskStateLog{
			Task:      l.Task,
			State:     types.TaskStateFailed,
			Comment:   err.Error(),
			CreatedAt: time.Now(),
		}); err != nil {
			slog.Error("failed to create failed task state", "error", err, "taskID", l.Task.ID)
		}
		return
	}

	if err := d.pg.Create(&types.TaskStateLog{
		Task:      l.Task,
		State:     types.TaskStateOutputted,
		Result:    []byte(outRes),
		CreatedAt: time.Now(),
	}); err != nil {
		slog.Error("failed to create outputted task state", "error", err, "taskID", l.Task.ID)
	}
}

func NewDispatcher(pg *persistence.Postgres, projectManager *project.Manager, bootNodeMultiaddr, operatorPrivateKey, operatorPrivateKeyED25519 string, iotexChainID int, source datasource.Datasource) (*Dispatcher, error) {
	d := &Dispatcher{
		source:                    source,
		pg:                        pg,
		projectManager:            projectManager,
		operatorPrivateKeyECDSA:   operatorPrivateKey,
		operatorPrivateKeyED25519: operatorPrivateKeyED25519,
	}
	ps, err := p2p.NewPubSubs(d.handleP2PData, bootNodeMultiaddr, iotexChainID)
	if err != nil {
		return nil, err
	}
	d.pubSubs = ps

	return d, nil
}
