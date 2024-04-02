package handler

import (
	"log/slog"
	"time"

	"github.com/machinefi/sprout/output"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/types"
)

type SaveTaskStateLog func(s *types.TaskStateLog, t *types.Task) error

type GetProjectConfig func(projectID uint64, version string) (*project.Config, error)

type TaskStateHandler struct {
	saveTaskStateLog          SaveTaskStateLog
	getProjectConfig          GetProjectConfig
	operatorPrivateKeyECDSA   string
	operatorPrivateKeyED25519 string
}

func (h *TaskStateHandler) Handle(s *types.TaskStateLog, t *types.Task) (finished bool) {
	// TODO Verify TaskStateLog Signature
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
	p, err := h.getProjectConfig(t.ProjectID, t.ProjectVersion)
	if err != nil {
		slog.Error("failed to get project config", "error", err, "project_id", t.ProjectID, "project_version", t.ProjectVersion)
		return
	}

	output, err := output.New(&p.Output, h.operatorPrivateKeyECDSA, h.operatorPrivateKeyED25519)
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
		Comment:   "output type: " + string(p.Output.Type),
		Result:    []byte(outRes),
		CreatedAt: time.Now(),
	}, t); err != nil {
		slog.Error("failed to create outputted task state", "error", err, "task_id", s.TaskID)
		return
	}
	return true
}

func NewTaskStateHandler(saveTaskStateLog SaveTaskStateLog, getProjectConfig GetProjectConfig, operatorPrivateKeyECDSA, operatorPrivateKeyED25519 string) *TaskStateHandler {
	return &TaskStateHandler{
		saveTaskStateLog:          saveTaskStateLog,
		getProjectConfig:          getProjectConfig,
		operatorPrivateKeyECDSA:   operatorPrivateKeyECDSA,
		operatorPrivateKeyED25519: operatorPrivateKeyED25519,
	}
}
