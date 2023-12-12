package types

import "time"

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

type TaskStateLog struct {
	TaskID    string
	State     TaskState
	Comment   string
	CreatedAt time.Time
}

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
