package types

type Message struct {
	ID        string `json:"id"`
	ProjectID uint64 `json:"projectID"`
	Data      string `json:"data"`
}

type Task struct {
	ID       string     `json:"id"`
	Messages []*Message `json:"messages"`
}

type TaskState uint8

const (
	TaskStateInvalid TaskState = iota
	TaskStateReceived
	TaskStateFetched
	TaskStateProving
	TaskStateProved
	TaskStateOutputting
	TaskStateOutputted
	TaskStateFailed
)

func (s TaskState) String() string {
	switch s {
	case TaskStateReceived:
		return "received"
	case TaskStateFetched:
		return "fetched"
	case TaskStateProving:
		return "proving"
	case TaskStateProved:
		return "proved"
	case TaskStateOutputting:
		return "outputting"
	case TaskStateOutputted:
		return "outputted"
	case TaskStateFailed:
		return "failed"
	default:
		return "invalid"
	}
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
