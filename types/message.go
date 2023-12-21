package types

type Message struct {
	ID        string `json:"id"`
	ProjectID uint64 `json:"projectID"`
	Data      string `json:"data"`
}
