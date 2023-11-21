package mq

import "github.com/machinefi/sprout/message"

type MQ interface {
	Enqueue(*message.Message) error
	Dequeue() (*message.Message, error)
	// will block caller
	Watch(func(*message.Message))
}
