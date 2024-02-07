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
	_
	TaskStateProved
	_
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
	case TaskStateProved:
		return "proved"
	case TaskStateOutputted:
		return "outputted"
	case TaskStateFailed:
		return "failed"
	default:
		return "invalid"
	}
}
