package enums

type MessageState uint8

const (
	MessageStateInvalid MessageState = iota
	MessageReceived
	MessageProcessing
	MessageProcessed
)

func (s MessageState) String() string {
	switch s {
	case MessageReceived:
		return "MessageReceived"
	case MessageProcessing:
		return "MessageProcessing"
	case MessageProcessed:
		return "MessageProcessed"
	}
	return "MessageStateInvalid"
}
