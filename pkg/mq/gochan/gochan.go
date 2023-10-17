package gochan

import "github.com/machinefi/w3bstream-mainnet/pkg/mq"

type queue struct {
	q chan *mq.Msg
}

func (q *queue) Enqueue(msg *mq.Msg) error {
	select {
	case q.q <- msg:
		return nil
	default:
		return mq.ErrMQFull
	}
}

func (q *queue) Dequeue() (*mq.Msg, error) {
	select {
	case m := <-q.q:
		return m, nil
	default:
		return nil, mq.ErrMQEmpty
	}
}

func New(queueSize int) mq.MQ {
	return &queue{
		q: make(chan *mq.Msg, queueSize),
	}
}
