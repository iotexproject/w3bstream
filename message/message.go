package message

type Message struct {
	ID             string `json:"id"`
	ProjectID      uint64 `json:"projectID"`
	ProjectVersion string `json:"projectVersion"`
	Data           string `json:"data"`
}

type FetchStrategy string

const (
	FIFO FetchStrategy = "fifo"
)
