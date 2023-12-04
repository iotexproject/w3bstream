package project

import (
	"github.com/machinefi/sprout/types"
)

type Project struct {
	ID     uint64 `json:"id"`
	Config Config `json:"config"`
}

type Config struct {
	Code          string       `json:"code"`
	CodeExpParam  string       `json:"codeExpParam,omitempty"`
	VMType        types.VM     `json:"vmType"`
	OutputType    types.Output `json:"outputType"`
	OutputAddress string       `json:"outputAddress,omitempty"`
}
