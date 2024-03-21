package task

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/output"
	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/project"
)

type Datasource interface {
	Retrieve(nextTaskID uint64) (*Task, error)
}

type Persistence interface {
	Create(tl *TaskStateLog) error
}

type ProjectManager interface {
	Get(projectID uint64, version string) (*project.Project, error)
	GetAllProjectID() []uint64
	GetNotify() <-chan uint64
}

type Dispatcher struct {
	datasource                Datasource
	persistence               Persistence
	projectManager            ProjectManager
	pubSubs                   *p2p.PubSubs
	operatorPrivateKeyECDSA   string
	operatorPrivateKeyED25519 string
}

// will block caller
func (d *Dispatcher) Dispatch(nextTaskID uint64) {
	ticker := time.NewTicker(3 * time.Second)

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
	t, err := d.datasource.Retrieve(nextTaskID)
	if err != nil {
		return 0, errors.Wrap(err, "failed to retrieve task from data source")
	}
	if t == nil {
		return nextTaskID, nil
	}
	if err := d.pubSubs.Add(t.ProjectID); err != nil {
		return 0, errors.Wrapf(err, "failed to add project pubsub, project_id %v", t.ProjectID)
	}
	data, err := json.Marshal(&p2pData{
		Task: t,
	})
	if err != nil {
		return 0, errors.Wrapf(err, "failed to marshal p2p data, project_id %v", t.ProjectID)
	}
	if err := d.pubSubs.Publish(t.ProjectID, data); err != nil {
		return 0, errors.Wrapf(err, "failed to publish data, project_id %v", t.ProjectID)
	}
	slog.Debug("dispatched a task", "project_id", t.ProjectID, "task_id", t.ID)
	return t.ID + 1, nil
}

func (d *Dispatcher) handleP2PData(rawdata []byte, topic *pubsub.Topic) {
	data := p2pData{}
	if err := json.Unmarshal(rawdata, &data); err != nil {
		slog.Error("failed to unmarshal p2p data", "error", err)
		return
	}
	if data.TaskStateLog == nil {
		return
	}

	l := data.TaskStateLog
	if err := d.persistence.Create(l); err != nil {
		slog.Error("failed to create task state log", "error", err, "taskID", l.Task.ID)
		return
	}
	if l.State != TaskStateProved {
		return
	}

	p, err := d.projectManager.Get(l.Task.ProjectID, l.Task.ProjectVersion)
	if err != nil {
		slog.Error("failed to get project", "error", err, "project_id", l.Task.ProjectID, "project_version", l.Task.ProjectVersion)
		return
	}

	output, err := output.New(&p.Config.Output, d.operatorPrivateKeyECDSA, d.operatorPrivateKeyED25519)
	if err != nil {
		slog.Error("failed to init output", "error", err)
		if err := d.persistence.Create(&TaskStateLog{
			Task:      l.Task,
			State:     TaskStateFailed,
			Comment:   err.Error(),
			CreatedAt: time.Now(),
		}); err != nil {
			slog.Error("failed to create failed task state", "error", err, "taskID", l.Task.ID)
		}
		return
	}

	slog.Debug("output proof", "outputter", fmt.Sprintf("%T", output))

	outRes, err := output.Output(l.Task.ProjectID, l.Task.Data, l.Result)
	if err != nil {
		slog.Error("failed to output", "error", err, "taskID", l.Task.ID)
		if err := d.persistence.Create(&TaskStateLog{
			Task:      l.Task,
			State:     TaskStateFailed,
			Comment:   err.Error(),
			CreatedAt: time.Now(),
		}); err != nil {
			slog.Error("failed to create failed task state", "error", err, "taskID", l.Task.ID)
		}
		return
	}

	if err := d.persistence.Create(&TaskStateLog{
		Task:      l.Task,
		State:     TaskStateOutputted,
		Result:    []byte(outRes),
		CreatedAt: time.Now(),
	}); err != nil {
		slog.Error("failed to create outputted task state", "error", err, "taskID", l.Task.ID)
	}
}

func NewDispatcher(persistence Persistence, projectManager ProjectManager, datasource Datasource, bootNodeMultiaddr, operatorPrivateKey, operatorPrivateKeyED25519 string, iotexChainID int) (*Dispatcher, error) {
	d := &Dispatcher{
		datasource:                datasource,
		persistence:               persistence,
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
