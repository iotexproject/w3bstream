package gochan

import (
	"github.com/machinefi/sprout/msg"
	"github.com/machinefi/sprout/util/mq"
)

const defaultQueueSize = 4096

type queue struct {
	q chan *msg.Msg
}

func (q *queue) Enqueue(msg *msg.Msg) error {
	select {
	case q.q <- msg:
		return nil
	default:
		return mq.ErrMQFull
	}
}

func (q *queue) Dequeue() (*msg.Msg, error) {
	select {
	case m := <-q.q:
		return m, nil
	default:
		return nil, mq.ErrMQEmpty
	}
}

func (q *queue) Watch(h func(*msg.Msg)) {
	for m := range q.q {
		h(m)
	}
}

func New() mq.MQ {
	return &queue{
		q: make(chan *msg.Msg, defaultQueueSize),
	}
}
