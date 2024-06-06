package dispatcher

import (
	"container/list"
	"log/slog"
	"sync"
	"sync/atomic"

	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/task"
)

type window struct {
	cond        *sync.Cond
	size        *atomic.Uint64
	tasks       *list.List
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

func (w *window) setSize(size uint64) {
	w.size.Store(size)
	w.cond.Broadcast()
}

func (w *window) getTask(taskID uint64) *dispatchedTask {
	for e := w.tasks.Front(); e != nil; e = e.Next() {
		t := e.Value.(*dispatchedTask)
		if t.task.ID == taskID {
			return t
		}
	}
	return nil
}

func (w *window) enQueue(value *dispatchedTask) {
	w.tasks.PushBack(value)
}

func (w *window) deQueue() {
	for !w.isEmpty() {
		if t := w.tasks.Front().Value.(*dispatchedTask); t.finished.Load() {
			w.tasks.Remove(w.tasks.Front())
			if err := w.persistence.UpsertProcessedTask(t.task.ProjectID, t.task.ID); err != nil {
				slog.Error("failed to upsert processed task", "project_id", t.task.ProjectID, "task_id", t.task.ID)
			}
		} else {
			return
		}
	}
}

func (w *window) isEmpty() bool {
	return w.tasks.Len() == 0
}

func (w *window) isFull() bool {
	return uint64(w.tasks.Len()) >= w.size.Load()
}

func newWindow(size uint64, pubSubs *p2p.PubSubs, handler *taskStateHandler, persistence Persistence) *window {
	s := &atomic.Uint64{}
	s.Store(size)
	return &window{
		cond:        sync.NewCond(&sync.Mutex{}),
		size:        s,
		tasks:       list.New(),
		pubSubs:     pubSubs,
		handler:     handler,
		persistence: persistence,
	}
}
