package p2p

import "github.com/machinefi/sprout/proto"

type DataType string

const (
	Request  DataType = "request"
	Response DataType = "response"
	Message  DataType = "message"
)

type Data struct {
	Type      DataType             `json:"type"`
	ProjectID uint64               `json:"projectID,omitempty"`
	Messages  []*proto.Message     `json:"messages,omitempty"`
	Report    *proto.ReportRequest `json:"response,omitempty"`
}
