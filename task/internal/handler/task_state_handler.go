package handler

import (
	"log/slog"
	"time"

	"github.com/machinefi/sprout/output"
	"github.com/machinefi/sprout/persistence/contract"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/types"
)

type SaveTaskStateLog func(s *types.TaskStateLog, t *types.Task) error

type Project func(projectID uint64) (*project.Project, error)

type LatestProvers func() []*contract.Prover

type TaskStateHandler struct {
	latestProvers             LatestProvers // optional, will be nil in local model
	saveTaskStateLog          SaveTaskStateLog
	project                   Project
	operatorPrivateKeyECDSA   string
	operatorPrivateKeyED25519 string
}

func (h *TaskStateHandler) Handle(s *types.TaskStateLog, t *types.Task) (finished bool) {
	// TODO dispatcher will send a failed TaskStateLog when timeout, without signature. maybe dispatcher need a sig also
	// if h.latestProvers != nil && s.Signature != "" {
	// 	ps := h.latestProvers()
	// 	signerAddress, err := s.SignerAddress(t)
	// 	if err != nil {
	// 		slog.Error("failed to get task state log signer address", "error", err, "task_id", s.TaskID)
	// 		return
	// 	}
	// 	legal := false
	// 	for _, p := range ps {
	// 		if p.OperatorAddress == signerAddress {
	// 			legal = true
	// 			break
	// 		}
	// 	}
	// 	if !legal {
	// 		slog.Error("failed to verify task state log signature", "task_id", s.TaskID, "signer_address", signerAddress.String())
	// 		return
	// 	}
	// }
	if err := h.saveTaskStateLog(s, t); err != nil {
		slog.Error("failed to create task state log", "error", err, "task_id", s.TaskID)
		return
	}
	if s.State == types.TaskStateFailed {
		return true
	}

	if s.State != types.TaskStateProved {
		return
	}
	p, err := h.project(t.ProjectID)
	if err != nil {
		slog.Error("failed to get project", "error", err, "project_id", t.ProjectID)
		return
	}
	c, err := p.DefaultConfig()
	if err != nil {
		slog.Error("failed to get project config", "error", err, "project_id", t.ProjectID, "project_version", p.DefaultVersion)
		return
	}

	output, err := output.New(&c.Output, h.operatorPrivateKeyECDSA, h.operatorPrivateKeyED25519)
	if err != nil {
		slog.Error("failed to init output", "error", err, "project_id", t.ProjectID)
		if err := h.saveTaskStateLog(&types.TaskStateLog{
			TaskID:    s.TaskID,
			State:     types.TaskStateFailed,
			Comment:   err.Error(),
			CreatedAt: time.Now(),
		}, t); err != nil {
			slog.Error("failed to create failed task state", "error", err, "task_id", s.TaskID)
			return
		}
		return true
	}

	outRes, err := output.Output(t, s.Result)
	if err != nil {
		slog.Error("failed to output", "error", err, "task_id", s.TaskID)
		if err := h.saveTaskStateLog(&types.TaskStateLog{
			TaskID:    s.TaskID,
			State:     types.TaskStateFailed,
			Comment:   err.Error(),
			CreatedAt: time.Now(),
		}, t); err != nil {
			slog.Error("failed to create failed task state", "error", err, "task_id", s.TaskID)
			return
		}
		return true
	}

	if err := h.saveTaskStateLog(&types.TaskStateLog{
		TaskID:    s.TaskID,
		State:     types.TaskStateOutputted,
		Comment:   "output type: " + string(c.Output.Type),
		Result:    []byte(outRes),
		CreatedAt: time.Now(),
	}, t); err != nil {
		slog.Error("failed to create outputted task state", "error", err, "task_id", s.TaskID)
		return
	}
	return true
}

func NewTaskStateHandler(saveTaskStateLog SaveTaskStateLog, latestProvers LatestProvers, project Project, operatorPrivateKeyECDSA, operatorPrivateKeyED25519 string) *TaskStateHandler {
	return &TaskStateHandler{
		latestProvers:             latestProvers,
		saveTaskStateLog:          saveTaskStateLog,
		project:                   project,
		operatorPrivateKeyECDSA:   operatorPrivateKeyECDSA,
		operatorPrivateKeyED25519: operatorPrivateKeyED25519,
	}
}
