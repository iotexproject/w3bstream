package msg

type Msg struct {
	ProjectID      uint64 `json:"projectID"`
	ProjectVersion string `json:"projectVersion"`
	Data           string `json:"data"`
}

type FetchStrategy string

const (
	FIFO FetchStrategy = "fifo"
)
