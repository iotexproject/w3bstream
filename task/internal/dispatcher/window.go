package dispatcher

import (
	"log/slog"
	"sync"

	"github.com/machinefi/sprout/task/internal/handler"
	"github.com/machinefi/sprout/types"
)

type window struct {
	cond    *sync.Cond
	front   int
	rear    int
	tasks   []*dispatcherTask
	publish Publish
	handler *handler.TaskStateHandler
}

func (w *window) consume(s *types.TaskStateLog) {
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

func (w *window) produce(t *types.Task) {
	w.cond.L.Lock()
	for w.isFull() {
		w.cond.Wait()
	}

	dt := newDispatcherTask(t, w.publish, w.handler)
	w.enQueue(dt)

	w.cond.L.Unlock()
}

func (w *window) getTask(taskID uint64) *dispatcherTask {
	for _, t := range w.tasks {
		if t != nil && t.task.ID == taskID {
			return t
		}
	}
	return nil
}

func (w *window) enQueue(value *dispatcherTask) {
	w.tasks[w.rear] = value
	w.rear = (w.rear + 1) % len(w.tasks)
}

func (w *window) deQueue() {
	for !w.isEmpty() {
		if w.tasks[w.front].finished.Load() {
			w.front = (w.front + 1) % len(w.tasks)
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

func newWindow(size uint, publish Publish, handler *handler.TaskStateHandler) *window {
	return &window{
		cond:    sync.NewCond(&sync.Mutex{}),
		tasks:   make([]*dispatcherTask, size+1),
		publish: publish,
		handler: handler,
	}
}
