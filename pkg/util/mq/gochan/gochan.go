package gochan

import (
	"github.com/machinefi/w3bstream-mainnet/pkg/msg"
	"github.com/machinefi/w3bstream-mainnet/pkg/util/mq"
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

func New() mq.MQ {
	return &queue{
		q: make(chan *msg.Msg, defaultQueueSize),
	}
}
