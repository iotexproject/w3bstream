package task

import (
	"fmt"
	"log/slog"
	"time"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/persistence"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/types"
)

type Dispatcher struct {
	pubSubs                   *p2p.PubSubs
	pg                        *persistence.Postgres
	projectManager            *project.Manager
	operatorPrivateKeyECDSA   string
	operatorPrivateKeyED25519 string
}

// will block caller
func (d *Dispatcher) Dispatch() {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if err := d.pubTask(); err != nil {
			slog.Error("failed to dispatch task", "error", err)
		}
	}
}

func (d *Dispatcher) pubTask() error {
	t, err := d.pg.Fetch()
	if err != nil {
		return errors.Wrapf(err, "failed to get task")
	}
	if t == nil {
		return errors.New("get task nil")
	}

	if err := d.pubSubs.Add(t.ProjectID); err != nil {
		return errors.Wrapf(err, "failed to add project pubsub, projectID %d", t.ProjectID)
	}

	slog.Debug("dispatch project task", "projectID", t.ProjectID, "taskID", t.ID)
	if err := d.pubSubs.Publish(t.ProjectID, &p2p.Data{
		Task: t,
	}); err != nil {
		return errors.Wrapf(err, "failed to publish data, projectID %d", t.ProjectID)
	}
	return nil
}

func (d *Dispatcher) handleP2PData(data *p2p.Data, topic *pubsub.Topic) {
	if data.TaskStateLog == nil {
		return
	}
	l := data.TaskStateLog
	if err := d.pg.UpdateState(l.TaskID, l.State, l.Comment, l.CreatedAt); err != nil {
		slog.Error("update task state failed", "error", err, "taskID", l.TaskID)
		return
	}
	if l.State != types.TaskStateProved {
		return
	}

	task, err := d.pg.FetchByID(l.TaskID)
	if err != nil {
		slog.Error("fetch task failed", "error", err, "taskID", l.TaskID)
		return
	}
	config, err := d.projectManager.Get(task.ProjectID, task.ProjectVersion)
	if err != nil {
		slog.Error("get project failed", "error", err, "projectID", task.ProjectID, "projectVersion", task.ProjectVersion)
		return
	}

	output, err := config.GetOutput(d.operatorPrivateKeyECDSA, d.operatorPrivateKeyED25519)
	if err != nil {
		slog.Error("init output failed", "error", err)
		if err := d.pg.UpdateState(l.TaskID, types.TaskStateFailed, err.Error(), time.Now()); err != nil {
			slog.Error("update task state to statefailed failed", "error", err, "taskID", l.TaskID)
		}
		return
	}

	slog.Debug("output proof", "outputter", fmt.Sprintf("%T", output))

	outRes, err := output.Output(task, []byte(l.Comment))
	if err != nil {
		slog.Error("output failed", "error", err)
		if err := d.pg.UpdateState(l.TaskID, types.TaskStateFailed, err.Error(), time.Now()); err != nil {
			slog.Error("update task state to statefailed failed", "error", err, "taskID", l.TaskID)
		}
		return
	}

	if err := d.pg.UpdateState(l.TaskID, types.TaskStateOutputted, fmt.Sprintf("output result: %s", outRes), time.Now()); err != nil {
		slog.Error("update task state to outputted failed", "error", err, "taskID", l.TaskID)
	}
}

func NewDispatcher(pg *persistence.Postgres, projectManager *project.Manager, bootNodeMultiaddr, operatorPrivateKey, operatorPrivateKeyED25519 string, iotexChainID int) (*Dispatcher, error) {
	d := &Dispatcher{
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
