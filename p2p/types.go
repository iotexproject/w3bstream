package p2p

import (
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/machinefi/sprout/types"
)

type Data struct {
	Task         *types.Task         `json:"task,omitempty"`
	TaskStateLog *types.TaskStateLog `json:"taskStateLog,omitempty"`
}

type HandleSubscriptionMessage func(*Data, *pubsub.Topic)
