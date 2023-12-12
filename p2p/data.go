package p2p

import (
	"github.com/machinefi/sprout/types"
)

type Data struct {
	Task         *types.Task         `json:"task,omitempty"`
	TaskStateLog *types.TaskStateLog `json:"taskStateLog,omitempty"`
}
