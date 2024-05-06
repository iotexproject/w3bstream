package dispatcher

import (
	"log/slog"
	"sync"

	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/task"
)

type window struct {
	cond        *sync.Cond
	front       int
	rear        int
	tasks       []*dispatchedTask
	pubSubs     *p2p.PubSubs
	handler     *taskStateHandler
	persistence Persistence
}

func (w *window) consume(s *task.StateLog) {
	w.cond.L.Lock()
	defer w.cond.Broadcast()
	defer w.cond.L.Unlock()

	t := w.getTask(s.TaskID)
	if t == nil {
		slog.Error("failed to get task in processing window", "task_id", s.TaskID)
		return
	}
	t.handleState(s)
	w.deQueue()
}

func (w *window) produce(t *task.Task) {
	w.cond.L.Lock()
	for w.isFull() {
		w.cond.Wait()
	}

	dt := newDispatchedTask(t, w.consume, w.pubSubs, w.handler)
	w.enQueue(dt)

	w.cond.L.Unlock()
}

func (w *window) getTask(taskID uint64) *dispatchedTask {
	for _, t := range w.tasks {
		if t != nil && t.task.ID == taskID {
			return t
		}
	}
	return nil
}

func (w *window) enQueue(value *dispatchedTask) {
	w.tasks[w.rear] = value
	w.rear = (w.rear + 1) % len(w.tasks)
}

func (w *window) deQueue() {
	for !w.isEmpty() {
		if t := w.tasks[w.front]; t.finished.Load() {
			w.front = (w.front + 1) % len(w.tasks)
			if err := w.persistence.UpsertProcessedTask(t.task.ProjectID, t.task.ID); err != nil {
				slog.Error("failed to upsert processed task", "project_id", t.task.ProjectID, "task_id", t.task.ID)
			}
		} else {
			return
		}
	}
}

func (w *window) isEmpty() bool {
	return w.rear == w.front
}

func (w *window) isFull() bool {
	return (w.rear+1)%len(w.tasks) == w.front
}

func newWindow(size uint64, pubSubs *p2p.PubSubs, handler *taskStateHandler, persistence Persistence) *window {
	return &window{
		cond:        sync.NewCond(&sync.Mutex{}),
		tasks:       make([]*dispatchedTask, size+1),
		pubSubs:     pubSubs,
		handler:     handler,
		persistence: persistence,
	}
}
