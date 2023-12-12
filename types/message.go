package types

type Message struct {
	ID        string `json:"id"`
	ProjectID uint64 `json:"projectID"`
	Data      string `json:"data"`
}

type MessageState uint8

const (
	MessageStateInvalid MessageState = iota
	MessageStateReceived
	MessageStateFetched
	MessageStateProving
	MessageStateProved
	MessageStateOutputting
	MessageStateOutputted
	MessageStateFailed
)

func (s MessageState) String() string {
	switch s {
	case MessageStateReceived:
		return "received"
	case MessageStateFetched:
		return "fetched"
	case MessageStateProving:
		return "proving"
	case MessageStateProved:
		return "proved"
	case MessageStateOutputting:
		return "outputting"
	case MessageStateOutputted:
		return "outputted"
	case MessageStateFailed:
		return "failed"
	default:
		return "invalid"
	}
}
