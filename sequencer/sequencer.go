package sequencer

import "github.com/machinefi/sprout/message"

type Sequencer interface {
	Save(msg *message.Message) (msgID uint64, err error)
	Fetch(projectID, afterMsgID uint64, strategy message.FetchStrategy) ([]*message.Message, error)
}
