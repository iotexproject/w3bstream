package dispatcher

import (
	"context"
	"fmt"
	"log/slog"
	"sync/atomic"
	"time"

	"github.com/machinefi/sprout/output"
	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/types"
)

type dispatcherTask struct {
	finished                  atomic.Bool
	cancel                    context.CancelFunc
	waitTime                  time.Duration
	task                      *types.Task
	persistence               Persistence
	projectConfigManager      ProjectConfigManager
	publisher                 Publisher
	operatorPrivateKeyECDSA   string
	operatorPrivateKeyED25519 string
}

func (t *dispatcherTask) handleState(s *types.TaskStateLog) {
	if err := t.persistence.Create(s); err != nil {
		slog.Error("failed to create task state log", "error", err, "task_id", t.task.ID)
		return
	}
	if s.State == types.TaskStateFailed {
		t.finish()
		return
	}

	if s.State != types.TaskStateProved {
		return
	}
	p, err := t.projectConfigManager.Get(t.task.ProjectID, t.task.ProjectVersion)
	if err != nil {
		slog.Error("failed to get project config", "error", err, "project_id", t.task.ProjectID, "project_version", t.task.ProjectVersion)
		return
	}

	output, err := output.New(&p.Output, t.operatorPrivateKeyECDSA, t.operatorPrivateKeyED25519)
	if err != nil {
		slog.Error("failed to init output", "error", err, "project_id", t.task.ProjectID)
		if err := t.persistence.Create(&types.TaskStateLog{
			TaskID:    t.task.ID,
			State:     types.TaskStateFailed,
			Comment:   err.Error(),
			CreatedAt: time.Now(),
		}); err != nil {
			slog.Error("failed to create failed task state", "error", err, "task_id", t.task.ID)
			return
		}
		t.finish()
		return
	}

	outRes, err := output.Output(t.task.ProjectID, t.task.Data, s.Result)
	if err != nil {
		slog.Error("failed to output", "error", err, "task_id", t.task.ID)
		if err := t.persistence.Create(&types.TaskStateLog{
			TaskID:    t.task.ID,
			State:     types.TaskStateFailed,
			Comment:   err.Error(),
			CreatedAt: time.Now(),
		}); err != nil {
			slog.Error("failed to create failed task state", "error", err, "task_id", t.task.ID)
			return
		}
		t.finish()
		return
	}

	if err := t.persistence.Create(&types.TaskStateLog{
		TaskID:    t.task.ID,
		State:     types.TaskStateOutputted,
		Comment:   "output type: " + string(p.Output.Type),
		Result:    []byte(outRes),
		CreatedAt: time.Now(),
	}); err != nil {
		slog.Error("failed to create outputted task state", "error", err, "task_id", t.task.ID)
		return
	}
	t.finish()
}

func (t *dispatcherTask) finish() {
	t.cancel()
	t.finished.Store(true)
}

func (t *dispatcherTask) runWatchdog(ctx context.Context) {
	retryChan := time.After(t.waitTime)
	timeoutChan := time.After(2 * t.waitTime)
	for {
		select {
		case <-ctx.Done():
			slog.Info("task finished", "task_id", t.task.ID, "project_id", t.task.ProjectID)
			return
		case <-retryChan:
			if err := t.publisher.Publish(t.task.ProjectID, &p2p.Data{Task: t.task}); err != nil {
				slog.Error("failed to publish p2p data", "project_id", t.task.ProjectID, "task_id", t.task.ID)
			}
		case <-timeoutChan:
			t.handleState(&types.TaskStateLog{
				TaskID:    t.task.ID,
				State:     types.TaskStateFailed,
				Comment:   fmt.Sprintf("task timeout, number of retries %v, total waiting time %v", 1, 2*t.waitTime),
				CreatedAt: time.Now(),
			})
		}
	}
}

func newDispatcherTask(task *types.Task, persistence Persistence, projectConfigManager ProjectConfigManager, publisher Publisher, operatorPrivateKeyECDSA, operatorPrivateKeyED25519 string) *dispatcherTask {
	ctx, cancel := context.WithCancel(context.Background())
	t := &dispatcherTask{
		finished:                  atomic.Bool{},
		cancel:                    cancel,
		waitTime:                  5 * time.Minute, // TODO define wait time config
		task:                      task,
		persistence:               persistence,
		projectConfigManager:      projectConfigManager,
		publisher:                 publisher,
		operatorPrivateKeyECDSA:   operatorPrivateKeyECDSA,
		operatorPrivateKeyED25519: operatorPrivateKeyED25519,
	}
	go t.runWatchdog(ctx)
	return t
}
