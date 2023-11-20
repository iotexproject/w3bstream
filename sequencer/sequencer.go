package sequencer

import "github.com/machinefi/sprout/msg"

type Sequencer interface {
	Save(msg *msg.Msg) error
	Fetch(projectID uint64, strategy msg.FetchStrategy) ([]*msg.Msg, error)
}
