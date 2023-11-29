package project

import (
	"github.com/machinefi/sprout/types"
)

type Project struct {
	ID     uint64 `json:"id"`
	Config Config `json:"config"`
}

type Config struct {
	Code         string   `json:"code"`
	CodeExpParam string   `json:"codeExpParam,omitempty"`
	VMType       types.VM `json:"vmType"`
	Output       struct {
		Type     types.Output `json:"type"`
		Ethereum struct {
			ChainName       string `json:"chainName"`
			ContractAddress string `json:"contractAddress"`
		} `json:"ethereum,omitempty"`
		Solana struct {
			ChainName      string `json:"chainName"`
			ProgramID      string `json:"programID"`
			StateAccountPK string `json:"stateAccountPK"`
		} `json:"solana,omitempty"`
	} `json:"output,omitempty"`
}
