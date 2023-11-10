package project

import (
	"github.com/holiman/uint256"
	"github.com/machinefi/w3bstream-mainnet/msg"
	"github.com/machinefi/w3bstream-mainnet/output"
	"github.com/machinefi/w3bstream-mainnet/vm"
)

type Project struct {
	ID     uint256.Int `json:"id"`
	Config Config      `json:"config"`
}

type Config struct {
	Code             []byte            `json:"code"`
	MsgFetchStrategy msg.FetchStrategy `json:"messageFetchStrategy"`
	VMType           vm.Type           `json:"vmType"`
	OutputType       output.Type       `json:"outputType"`
	OutputAddress    string            `json:"outputAddress,omitempty"`
}
