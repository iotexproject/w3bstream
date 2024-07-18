package p2p

import (
	pubsub "github.com/libp2p/go-libp2p-pubsub"

	"github.com/iotexproject/w3bstream/task"
)

type Data struct {
	Task         *task.Task     `json:"task,omitempty"`
	TaskStateLog *task.StateLog `json:"taskStateLog,omitempty"`
}

type HandleSubscriptionMessage func(*Data, *pubsub.Topic)
