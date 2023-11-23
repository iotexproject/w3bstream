package mq

import (
	"github.com/machinefi/sprout/proto"
)

type MQ interface {
	Enqueue(*proto.Message) error
	Dequeue() (*proto.Message, error)
	// will block caller
	Watch(func(*proto.Message))
}
