package p2p

import (
	"github.com/machinefi/sprout/types"
)

type Data struct {
	Request  *RequestData  `json:"request,omitempty"`
	Message  *MessageData  `json:"message,omitempty"`
	Response *ResponseData `json:"response,omitempty"`
}

type RequestData struct {
	ProjectID uint64 `json:"projectID,omitempty"`
}

type MessageData struct {
	Messages []*types.Message `json:"messages,omitempty"`
}

type ResponseData struct {
	MessageIDs []string           `json:"messageIDs,omitempty"`
	State      types.MessageState `json:"state,omitempty"`
	Comment    string             `json:"comment,omitempty"`
}
