package dispatcher

import (
	"context"
	"fmt"
	"log/slog"
	"sync/atomic"
	"time"

	"github.com/iotexproject/w3bstream/metrics"
	"github.com/iotexproject/w3bstream/p2p"
	"github.com/iotexproject/w3bstream/task"
)

type dispatchedTask struct {
	dispatchedTime time.Time
	finished       atomic.Bool
	timeOut        func(s *task.StateLog)
	cancel         context.CancelFunc
	waitTime       time.Duration
	task           *task.Task
	pubSubs        *p2p.PubSubs
	handler        *taskStateHandler
}

func (t *dispatchedTask) handleState(s *task.StateLog) {
	if t.handler.handle(t.dispatchedTime, s, t.task) {
		t.cancel()
		t.finished.Store(true)
	}
}

func (t *dispatchedTask) runWatchdog(ctx context.Context) {
	retryChan := time.After(t.waitTime)
	timeoutChan := time.After(2 * t.waitTime)
	for {
		select {
		case <-ctx.Done():
			slog.Info("task finished", "project_id", t.task.ProjectID, "task_id", t.task.ID)
			return
		case <-retryChan:
			slog.Info("retry task", "project_id", t.task.ProjectID, "task_id", t.task.ID, "wait_time", t.waitTime)
			metrics.RetryTaskNumMtc(t.task.ProjectID, t.task.ProjectVersion)

			if err := t.pubSubs.Publish(t.task.ProjectID, &p2p.Data{Task: t.task}); err != nil {
				slog.Error("failed to publish p2p data", "project_id", t.task.ProjectID, "task_id", t.task.ID)
			}
		case <-timeoutChan:
			slog.Info("task timeout", "project_id", t.task.ProjectID, "task_id", t.task.ID, "wait_time", 2*t.waitTime)
			metrics.TimeoutTaskNumMtc(t.task.ProjectID, t.task.ProjectVersion)

			t.timeOut(&task.StateLog{
				TaskID:    t.task.ID,
				State:     task.StateFailed,
				Comment:   fmt.Sprintf("task timeout, number of retries %v, total waiting time %v", 1, 2*t.waitTime),
				CreatedAt: time.Now(),
			})
		}
	}
}

func newDispatchedTask(task *task.Task, timeOut func(s *task.StateLog), pubSubs *p2p.PubSubs, handler *taskStateHandler) *dispatchedTask {
	ctx, cancel := context.WithCancel(context.Background())
	t := &dispatchedTask{
		dispatchedTime: time.Now(),
		finished:       atomic.Bool{},
		timeOut:        timeOut,
		cancel:         cancel,
		waitTime:       5 * time.Minute, // TODO define wait time config
		task:           task,
		pubSubs:        pubSubs,
		handler:        handler,
	}
	go t.runWatchdog(ctx)
	return t
}
