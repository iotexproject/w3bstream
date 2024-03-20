package task

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/vm"
)

type VMHandler interface {
	Handle(projectID uint64, vmtype vm.Type, code string, expParam string, data [][]byte) ([]byte, error)
}

type Processor struct {
	vmHandler      VMHandler
	projectManager ProjectManager
	ps             *p2p.PubSubs
}

func (r *Processor) handleP2PData(data []byte, topic *pubsub.Topic) {
	d := p2pData{}
	if err := json.Unmarshal(data, &d); err != nil {
		slog.Error("failed to unmarshal p2p data", "error", err)
		return
	}
	if d.Task == nil {
		return
	}

	t := d.Task
	slog.Debug("get a new task", "task_id", t.ID)
	r.reportSuccess(t, TaskStateDispatched, nil, topic)

	config, err := r.projectManager.Get(t.ProjectID, t.ProjectVersion)
	if err != nil {
		slog.Error("failed to get project", "error", err, "project_id", t.ProjectID, "project_version", t.ProjectVersion)
		r.reportFail(t, err, topic)
		return
	}

	res, err := r.vmHandler.Handle(t.ProjectID, config.VMType, config.Code, config.CodeExpParam, t.Data)
	if err != nil {
		slog.Error("failed to generate proof", "error", err)
		r.reportFail(t, err, topic)
		return
	}
	r.reportSuccess(t, TaskStateProved, res, topic)
}

func (r *Processor) reportFail(t *Task, err error, topic *pubsub.Topic) {
	d, err := json.Marshal(&p2pData{
		TaskStateLog: &TaskStateLog{
			Task:      *t,
			State:     TaskStateFailed,
			Comment:   err.Error(),
			CreatedAt: time.Now(),
		},
	})
	if err != nil {
		slog.Error("failed to marshal p2p task state log data to json", "error", err, "task_id", t.ID)
		return
	}
	if err := topic.Publish(context.Background(), d); err != nil {
		slog.Error("failed to publish task state log data to p2p network", "error", err, "task_id", t.ID)
	}
}

func (r *Processor) reportSuccess(t *Task, state TaskState, result []byte, topic *pubsub.Topic) {
	d, err := json.Marshal(&p2pData{
		TaskStateLog: &TaskStateLog{
			Task:      *t,
			State:     state,
			Result:    result,
			CreatedAt: time.Now(),
		},
	})
	if err != nil {
		slog.Error("failed to marshal p2p task state log data to json", "error", err, "task_id", t.ID)
		return
	}
	if err := topic.Publish(context.Background(), d); err != nil {
		slog.Error("failed to publish task state log data to p2p network", "error", err, "task_id", t.ID)
	}
}

func (r *Processor) Run() {
	// TODO project load & delete
}

func NewProcessor(vmHandler VMHandler, projectManager ProjectManager, bootNodeMultiaddr string, iotexChainID int) (*Processor, error) {
	p := &Processor{
		vmHandler:      vmHandler,
		projectManager: projectManager,
	}

	ps, err := p2p.NewPubSubs(p.handleP2PData, bootNodeMultiaddr, iotexChainID)
	if err != nil {
		return nil, err
	}
	p.ps = ps

	for _, id := range projectManager.GetAllProjectID() {
		if err := ps.Add(id); err != nil {
			return nil, errors.Wrapf(err, "add project %d pubsub failed", id)
		}
		slog.Debug("processor project added", "projectID", id)
	}

	notify := projectManager.GetNotify()
	go func() {
		for id := range notify {
			if err := ps.Add(id); err != nil {
				slog.Error("add project pubsub failed", "projectID", id, "error", err)
			}
			slog.Debug("processor project added", "projectID", id)
		}
	}()

	return p, nil
}
