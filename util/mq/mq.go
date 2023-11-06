package mq

import "github.com/machinefi/w3bstream-mainnet/msg"

type MQ interface {
	Enqueue(*msg.Msg) error
	Dequeue() (*msg.Msg, error)
	// will block caller
	Watch(func(*msg.Msg))
}
