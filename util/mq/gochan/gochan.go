package gochan

import (
	"github.com/machinefi/sprout/proto"
	"github.com/machinefi/sprout/util/mq"
)

const defaultQueueSize = 4096

type queue struct {
	q chan *proto.Message
}

func (q *queue) Enqueue(msg *proto.Message) error {
	select {
	case q.q <- msg:
		return nil
	default:
		return mq.ErrMQFull
	}
}

func (q *queue) Dequeue() (*proto.Message, error) {
	select {
	case m := <-q.q:
		return m, nil
	default:
		return nil, mq.ErrMQEmpty
	}
}

func (q *queue) Watch(h func(*proto.Message)) {
	for m := range q.q {
		h(m)
	}
}

func New() mq.MQ {
	return &queue{
		q: make(chan *proto.Message, defaultQueueSize),
	}
}
