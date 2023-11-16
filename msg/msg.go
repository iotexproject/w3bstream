package msg

type Msg struct {
	// ID message unique id for tracing message status
	ID             string `json:"id"`
	ProjectID      uint64 `json:"projectID"`
	ProjectVersion string `json:"projectVersion"`
	Data           string `json:"data"`
}

type FetchStrategy string

const (
	FIFO FetchStrategy = "fifo"
)
