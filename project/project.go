package project

import (
	"github.com/machinefi/w3bstream-mainnet/msg"
	"github.com/machinefi/w3bstream-mainnet/output"
	"github.com/machinefi/w3bstream-mainnet/vm"
)

type Project struct {
	ID     uint64 `json:"id"`
	Config Config `json:"config"`
}

type Config struct {
	Code             []byte            `json:"code"`
	CodeExpParam     string            `json:"codeExpParam,omitempty"`
	MsgFetchStrategy msg.FetchStrategy `json:"messageFetchStrategy"`
	VMType           vm.Type           `json:"vmType"`
	OutputType       output.Type       `json:"outputType"`
	OutputAddress    string            `json:"outputAddress,omitempty"`
}
