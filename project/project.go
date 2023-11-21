package project

import (
	"github.com/machinefi/sprout/message"
	"github.com/machinefi/sprout/output"
	"github.com/machinefi/sprout/vm"
)

type Project struct {
	ID     uint64 `json:"id"`
	Config Config `json:"config"`
}

type Config struct {
	Code             string                `json:"code"`
	CodeExpParam     string                `json:"codeExpParam,omitempty"`
	MsgFetchStrategy message.FetchStrategy `json:"messageFetchStrategy"`
	VMType           vm.Type               `json:"vmType"`
	OutputType       output.Type           `json:"outputType"`
	OutputAddress    string                `json:"outputAddress,omitempty"`
}
