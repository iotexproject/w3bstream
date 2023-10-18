package mq

import "github.com/machinefi/w3bstream-mainnet/pkg/msg"

type MQ interface {
	Enqueue(*msg.Msg) error
	Dequeue() (*msg.Msg, error)
}
