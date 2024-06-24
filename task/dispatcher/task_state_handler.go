package dispatcher

import (
	"log/slog"
	"time"

	"github.com/machinefi/sprout/metrics"
	"github.com/machinefi/sprout/output"
	"github.com/machinefi/sprout/task"
)

type taskStateHandler struct {
	contract                  Contract // optional, will be nil in local model
	persistence               Persistence
	projectManager            ProjectManager
	operatorPrivateKeyECDSA   string
	operatorPrivateKeyED25519 string
	contractWhitelist         string
}

func (h *taskStateHandler) handle(dispatchedTime time.Time, s *task.StateLog, t *task.Task) (finished bool) {
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
	if err := h.persistence.Create(s, t); err != nil {
		slog.Error("failed to create task state log", "error", err, "task_id", s.TaskID)
		return
	}
	if s.State == task.StateFailed {
		metrics.FailedTaskNumMtc(t.ProjectID, t.ProjectVersion)
		metrics.TaskFinalStateNumMtc(t.ProjectID, t.ProjectVersion, task.StateFailed.String())
		return true
	}

	if s.State != task.StateProved {
		return
	}
	p, err := h.projectManager.Project(t.ProjectID)
	if err != nil {
		slog.Error("failed to get project", "error", err, "project_id", t.ProjectID)
		return
	}
	c, err := p.DefaultConfig()
	if err != nil {
		slog.Error("failed to get project config", "error", err, "project_id", t.ProjectID, "project_version", p.DefaultVersion)
		return
	}

	output, err := output.New(&c.Output, h.operatorPrivateKeyECDSA, h.operatorPrivateKeyED25519, h.contractWhitelist)
	if err != nil {
		slog.Error("failed to init output", "error", err, "project_id", t.ProjectID)
		metrics.FailedTaskNumMtc(t.ProjectID, t.ProjectVersion)
		metrics.TaskFinalStateNumMtc(t.ProjectID, t.ProjectVersion, task.StateFailed.String())

		if err := h.persistence.Create(&task.StateLog{
			TaskID:    s.TaskID,
			State:     task.StateFailed,
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
		metrics.FailedTaskNumMtc(t.ProjectID, t.ProjectVersion)
		metrics.TaskFinalStateNumMtc(t.ProjectID, t.ProjectVersion, task.StateFailed.String())

		if err := h.persistence.Create(&task.StateLog{
			TaskID:    s.TaskID,
			State:     task.StateFailed,
			Comment:   err.Error(),
			CreatedAt: time.Now(),
		}, t); err != nil {
			slog.Error("failed to create failed task state", "error", err, "task_id", s.TaskID)
			return
		}
		return true
	}

	metrics.TaskDurationMtc(t.ProjectID, t.ProjectVersion, float64(time.Now().UnixNano())/1e9-float64(dispatchedTime.UnixNano())/1e9)
	metrics.SucceedTaskNumMtc(t.ProjectID, t.ProjectVersion)
	metrics.TaskFinalStateNumMtc(t.ProjectID, t.ProjectVersion, task.StateOutputted.String())

	if err := h.persistence.Create(&task.StateLog{
		TaskID:    s.TaskID,
		State:     task.StateOutputted,
		Comment:   "output type: " + string(c.Output.Type),
		Result:    []byte(outRes),
		CreatedAt: time.Now(),
	}, t); err != nil {
		slog.Error("failed to create outputted task state", "error", err, "task_id", s.TaskID)
		return
	}
	return true
}

func newTaskStateHandler(persistence Persistence, contract Contract, projectManager ProjectManager,
	operatorPrivateKeyECDSA, operatorPrivateKeyED25519, contractWhitelist string) *taskStateHandler {
	return &taskStateHandler{
		contract:                  contract,
		persistence:               persistence,
		projectManager:            projectManager,
		operatorPrivateKeyECDSA:   operatorPrivateKeyECDSA,
		operatorPrivateKeyED25519: operatorPrivateKeyED25519,
		contractWhitelist:         contractWhitelist,
	}
}
