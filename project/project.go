package project

import (
	"github.com/holiman/uint256"
	msgtype "github.com/machinefi/w3bstream-mainnet/msg/types"
	outtype "github.com/machinefi/w3bstream-mainnet/output/types"
	vmtype "github.com/machinefi/w3bstream-mainnet/vm/types"
)

type Project struct {
	ID     uint256.Int `json:"id"`
	Config Config      `json:"config"`
}

type Config struct {
	Code             []byte                `json:"code"`
	CodeExpParam     string                `json:"codeExpParam,omitempty"`
	MsgFetchStrategy msgtype.FetchStrategy `json:"messageFetchStrategy"`
	VMType           vmtype.Type           `json:"vmType"`
	OutputType       outtype.Type          `json:"outputType"`
	OutputAddress    string                `json:"outputAddress,omitempty"`
}
