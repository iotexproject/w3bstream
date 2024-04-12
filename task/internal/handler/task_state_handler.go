package handler

import (
	"log/slog"
	"time"

	"github.com/machinefi/sprout/output"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/types"
)

type SaveTaskStateLog func(s *types.TaskStateLog, t *types.Task) error

type GetProject func(projectID uint64) (*project.Project, error)

type TaskStateHandler struct {
	saveTaskStateLog          SaveTaskStateLog
	getProject                GetProject
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
	p, err := h.getProject(t.ProjectID)
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

func NewTaskStateHandler(saveTaskStateLog SaveTaskStateLog, getProject GetProject, operatorPrivateKeyECDSA, operatorPrivateKeyED25519 string) *TaskStateHandler {
	return &TaskStateHandler{
		saveTaskStateLog:          saveTaskStateLog,
		getProject:                getProject,
		operatorPrivateKeyECDSA:   operatorPrivateKeyECDSA,
		operatorPrivateKeyED25519: operatorPrivateKeyED25519,
	}
}
