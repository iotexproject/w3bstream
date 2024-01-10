package task

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/types"
	"github.com/machinefi/sprout/vm"
)

type Processor struct {
	vmHandler                 *vm.Handler
	projectManager            *project.Manager
	operatorPrivateKeyECDSA   string
	operatorPrivateKeyED25519 string
	ps                        *p2p.PubSubs
}

func (r *Processor) handleP2PData(d *p2p.Data, topic *pubsub.Topic) {
	if d.Task == nil {
		return
	}
	tid := d.Task.ID
	ms := d.Task.Messages
	slog.Debug("get new task", "task_id", tid)
	r.reportSuccess(tid, types.TaskStateFetched, "", topic)

	config, err := r.projectManager.Get(ms[0].ProjectID, ms[0].ProjectVersion)
	if err != nil {
		slog.Error("get project failed", "error", err)
		r.reportFail(tid, err, topic)
		return
	}

	r.reportSuccess(tid, types.TaskStateProving, "", topic)
	res, err := r.vmHandler.Handle(ms, config.VMType, config.Code, config.CodeExpParam)
	if err != nil {
		slog.Error("proof failed", "error", err)
		r.reportFail(tid, err, topic)
		return
	}
	slog.Debug("proof result", "proof_result", string(res))
	r.reportSuccess(tid, types.TaskStateProved, string(res), topic)

	output, err := config.GetOutput(r.operatorPrivateKeyECDSA, r.operatorPrivateKeyED25519)
	if err != nil {
		err = errors.Wrap(err, "fail to init output")
		slog.Error(err.Error())
		r.reportFail(tid, err, topic)
		return
	}

	slog.Debug("output proof", "outputter", fmt.Sprintf("%T", output))

	r.reportSuccess(tid, types.TaskStateOutputting, "output proof", topic)
	outRes, err := output.Output(d.Task, res)
	if err != nil {
		slog.Error(err.Error())
		r.reportFail(tid, err, topic)
		return
	}

	r.reportSuccess(tid, types.TaskStateOutputted, fmt.Sprintf("output result: %s", outRes), topic)
	slog.Debug("output success", "result", outRes)
}

func (r *Processor) reportFail(taskID string, err error, topic *pubsub.Topic) {
	j, err := json.Marshal(&p2p.Data{
		TaskStateLog: &types.TaskStateLog{
			TaskID:    taskID,
			State:     types.TaskStateFailed,
			Comment:   err.Error(),
			CreatedAt: time.Now(),
		},
	})
	if err != nil {
		slog.Error("json marshal p2p task state log data failed", "error", err, "taskID", taskID)
		return
	}
	if err := topic.Publish(context.Background(), j); err != nil {
		slog.Error("publish task state log data to p2p network failed", "error", err, "taskID", taskID)
	}
}

func (r *Processor) reportSuccess(taskID string, state types.TaskState, comment string, topic *pubsub.Topic) {
	j, err := json.Marshal(&p2p.Data{
		TaskStateLog: &types.TaskStateLog{
			TaskID:    taskID,
			State:     state,
			Comment:   comment,
			CreatedAt: time.Now(),
		},
	})
	if err != nil {
		slog.Error("json marshal p2p task state log data failed", "error", err, "taskID", taskID)
		return
	}
	if err := topic.Publish(context.Background(), j); err != nil {
		slog.Error("publish task state log data to p2p network failed", "error", err, "taskID", taskID)
	}
}

func (r *Processor) Run() {
	// TODO project load & delete
}

func NewProcessor(vmHandler *vm.Handler, projectManager *project.Manager, operatorPrivateKey, operatorPrivateKeyED25519, bootNodeMultiaddr string, iotexChainID int) (*Processor, error) {
	p := &Processor{
		vmHandler:                 vmHandler,
		operatorPrivateKeyECDSA:   operatorPrivateKey,
		operatorPrivateKeyED25519: operatorPrivateKeyED25519,
		projectManager:            projectManager,
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
		slog.Debug("project added", "project_id", id)
	}
	return p, nil
}
