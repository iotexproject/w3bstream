package sequencer

import "github.com/machinefi/sprout/msg"

type Sequencer interface {
	Save(msg *msg.Msg) (msgID uint64, err error)
	Fetch(projectID, afterMsgID uint64, strategy msg.FetchStrategy) ([]*msg.Msg, error)
}
