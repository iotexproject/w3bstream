package dispatcher

import (
	"context"
	"fmt"
	"log/slog"
	"sync/atomic"
	"time"

	"github.com/machinefi/sprout/metrics"
	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/task/internal/handler"
	"github.com/machinefi/sprout/types"
)

type dispatcherTask struct {
	dispatchedTime time.Time
	finished       atomic.Bool
	timeOut        func(s *types.TaskStateLog)
	cancel         context.CancelFunc
	waitTime       time.Duration
	task           *types.Task
	publish        Publish
	handler        *handler.TaskStateHandler
}

func (t *dispatcherTask) handleState(s *types.TaskStateLog) {
	if t.handler.Handle(t.dispatchedTime, s, t.task) {
		t.cancel()
		t.finished.Store(true)
	}
}

func (t *dispatcherTask) runWatchdog(ctx context.Context) {
	retryChan := time.After(t.waitTime)
	timeoutChan := time.After(2 * t.waitTime)
	for {
		select {
		case <-ctx.Done():
			slog.Info("task finished", "project_id", t.task.ProjectID, "task_id", t.task.ID)
			return
		case <-retryChan:
			slog.Info("retry task", "project_id", t.task.ProjectID, "task_id", t.task.ID, "wait_time", t.waitTime)
			metrics.RetryTaskNumMtc(t.task.ProjectID, t.task.ID, t.task.ProjectVersion)

			if err := t.publish(t.task.ProjectID, &p2p.Data{Task: t.task}); err != nil {
				slog.Error("failed to publish p2p data", "project_id", t.task.ProjectID, "task_id", t.task.ID)
			}
		case <-timeoutChan:
			slog.Info("task timeout", "project_id", t.task.ProjectID, "task_id", t.task.ID, "wait_time", 2*t.waitTime)
			metrics.TimeoutTaskNumMtc(t.task.ProjectID, t.task.ID, t.task.ProjectVersion)

			t.timeOut(&types.TaskStateLog{
				TaskID:    t.task.ID,
				State:     types.TaskStateFailed,
				Comment:   fmt.Sprintf("task timeout, number of retries %v, total waiting time %v", 1, 2*t.waitTime),
				CreatedAt: time.Now(),
			})
		}
	}
}

func newDispatcherTask(task *types.Task, timeOut func(s *types.TaskStateLog), publish Publish, handler *handler.TaskStateHandler) *dispatcherTask {
	ctx, cancel := context.WithCancel(context.Background())
	t := &dispatcherTask{
		dispatchedTime: time.Now(),
		finished:       atomic.Bool{},
		timeOut:        timeOut,
		cancel:         cancel,
		waitTime:       5 * time.Minute, // TODO define wait time config
		task:           task,
		publish:        publish,
		handler:        handler,
	}
	go t.runWatchdog(ctx)
	return t
}
