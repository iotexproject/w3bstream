package project

import (
	"github.com/machinefi/sprout/output"
	"github.com/machinefi/sprout/types"
)

type Config struct {
	Code         string            `json:"code"`
	CodeExpParam string            `json:"codeExpParam,omitempty"`
	VMType       types.VM          `json:"vmType"`
	Output       output.Config     `json:"output"`
	Aggregation  AggregationConfig `json:"aggregation"`
	Version      string            `json:"version"`
}

type AggregationConfig struct {
	Amount uint `json:"amount,omitempty"`
}

func (c *Config) GetOutput(privateKeyECDSA, privateKeyED25519 string) (output.Output, error) {
	return c.Output.SetPrivateKey(privateKeyECDSA, privateKeyED25519).Output()
}
