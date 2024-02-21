package project

import (
	"testing"

	"github.com/machinefi/sprout/types"
	"github.com/stretchr/testify/assert"
)

func TestGetOutputDefault(t *testing.T) {
	c := &Config{}
	_, err := c.GetOutput("", "")
	assert.NoError(t, err)
}

func TestGetOutputStdout(t *testing.T) {
	c := &Config{
		Output: OutputConfig{
			Type: types.OutputStdout,
		},
	}
	_, err := c.GetOutput("", "")
	assert.NoError(t, err)
}

func TestGetOutputEthereum(t *testing.T) {
	c := &Config{
		Output: OutputConfig{
			Type: types.OutputEthereumContract,
			Ethereum: EthereumConfig{
				ContractAbiJSON: `[{"constant":true,"inputs":[],"name":"getProof","outputs":[{"internalType":"bytes","name":"","type":"bytes"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"bytes","name":"_proof","type":"bytes"}],"name":"setProof","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`,
			},
		},
	}
	_, err := c.GetOutput("c47bbade736b0f82788aa6eaa06140cdf41a544707edef944299642e0d708cab", "")
	assert.NoError(t, err)
}

func TestGetOutputSolana(t *testing.T) {
	c := &Config{
		Output: OutputConfig{
			Type:   types.OutputSolanaProgram,
			Solana: SolanaConfig{},
		},
	}
	_, err := c.GetOutput("", "308edd7fca562182adbffaa59264a138d9e04f9f3adbda2c80ef1ca71b7dcfa4")
	assert.NoError(t, err)
}
