package types

import "time"

type Message struct {
	ID             string `json:"id"`
	ProjectID      uint64 `json:"projectID"`
	ProjectVersion string `json:"projectVersion"`
	Data           string `json:"data"`
}

type MessageWithTime struct {
	Message
	CreatedAt time.Time
}
