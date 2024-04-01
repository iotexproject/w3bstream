package types

import "time"

type Task struct {
	ID             uint64   `json:"id"`
	ProjectID      uint64   `json:"projectID"`
	ProjectVersion string   `json:"projectVersion"`
	Data           [][]byte `json:"data"`
}

type TaskState uint8

const (
	TaskStateInvalid TaskState = iota
	TaskStatePacked
	TaskStateDispatched
	_
	TaskStateProved
	_
	TaskStateOutputted
	TaskStateFailed
)

type TaskStateLog struct {
	TaskID    uint64
	State     TaskState
	Comment   string
	Result    []byte
	CreatedAt time.Time
}

func (s TaskState) String() string {
	switch s {
	case TaskStatePacked:
		return "packed"
	case TaskStateDispatched:
		return "dispatched"
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

type p2pData struct {
	Task         *Task         `json:"task,omitempty"`
	TaskStateLog *TaskStateLog `json:"taskStateLog,omitempty"`
}
