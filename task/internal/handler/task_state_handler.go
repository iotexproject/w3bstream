package handler

import (
	"log/slog"
	"strconv"
	"time"

	"github.com/machinefi/sprout/metrics"
	"github.com/machinefi/sprout/output"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/types"
)

type SaveTaskStateLog func(s *types.TaskStateLog, t *types.Task) error

type Project func(projectID uint64) (*project.Project, error)

type TaskStateHandler struct {
	saveTaskStateLog          SaveTaskStateLog
	project                   Project
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
		metrics.FailedTaskNumMtc(strconv.FormatUint(t.ProjectID, 10), t.ProjectVersion)
		metrics.TaskFinalStateMtc(strconv.FormatUint(t.ProjectID, 10), t.ProjectVersion,
			strconv.FormatUint(t.ID, 10), types.TaskStateFailed.String())

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
		metrics.FailedTaskNumMtc(strconv.FormatUint(t.ProjectID, 10), t.ProjectVersion)
		metrics.TaskFinalStateMtc(strconv.FormatUint(t.ProjectID, 10), t.ProjectVersion,
			strconv.FormatUint(t.ID, 10), types.TaskStateFailed.String())

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

	metrics.TaskEndTimeMtc(strconv.FormatUint(t.ProjectID, 10), t.ProjectVersion, strconv.FormatUint(t.ID, 10))
	metrics.SucceedTaskNumMtc(strconv.FormatUint(t.ProjectID, 10), t.ProjectVersion)
	metrics.TaskFinalStateMtc(strconv.FormatUint(t.ProjectID, 10), t.ProjectVersion,
		strconv.FormatUint(t.ID, 10), types.TaskStateOutputted.String())

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

func NewTaskStateHandler(saveTaskStateLog SaveTaskStateLog, project Project, operatorPrivateKeyECDSA, operatorPrivateKeyED25519 string) *TaskStateHandler {
	return &TaskStateHandler{
		saveTaskStateLog:          saveTaskStateLog,
		project:                   project,
		operatorPrivateKeyECDSA:   operatorPrivateKeyECDSA,
		operatorPrivateKeyED25519: operatorPrivateKeyED25519,
	}
}
