package mq

import "github.com/machinefi/sprout/msg"

type MQ interface {
	Enqueue(*msg.Msg) error
	Dequeue() (*msg.Msg, error)
	// will block caller
	Watch(func(*msg.Msg))
}
