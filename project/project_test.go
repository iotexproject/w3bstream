package project

import (
	"crypto/sha256"
	"encoding/json"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/machinefi/sprout/types"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestGetOutput(t *testing.T) {
	require := require.New(t)

	t.Run("default", func(t *testing.T) {
		c := &Config{}
		_, err := c.GetOutput("", "")
		require.NoError(err)
	})
	t.Run("stdout", func(t *testing.T) {
		c := &Config{
			Output: OutputConfig{
				Type: types.OutputStdout,
			},
		}
		_, err := c.GetOutput("", "")
		require.NoError(err)
	})
	t.Run("ethereum", func(t *testing.T) {
		c := &Config{
			Output: OutputConfig{
				Type: types.OutputEthereumContract,
				Ethereum: EthereumConfig{
					ContractAbiJSON: `[{"constant":true,"inputs":[],"name":"getProof","outputs":[{"internalType":"bytes","name":"","type":"bytes"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"bytes","name":"_proof","type":"bytes"}],"name":"setProof","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`,
				},
			},
		}
		_, err := c.GetOutput("c47bbade736b0f82788aa6eaa06140cdf41a544707edef944299642e0d708cab", "")
		require.NoError(err)
	})
	t.Run("solana", func(t *testing.T) {
		c := &Config{
			Output: OutputConfig{
				Type:   types.OutputSolanaProgram,
				Solana: SolanaConfig{},
			},
		}
		_, err := c.GetOutput("", "308edd7fca562182adbffaa59264a138d9e04f9f3adbda2c80ef1ca71b7dcfa4")
		require.NoError(err)
	})
}

func TestGetConfigsHttp(t *testing.T) {
	require := require.New(t)

	cs := []*Config{
		{
			Code:    "i am code",
			VMType:  types.VMHalo2,
			Version: "0.1",
		},
	}
	jc, err := json.Marshal(cs)
	require.NoError(err)

	h := sha256.New()
	_, err = h.Write(jc)
	require.NoError(err)
	hash := h.Sum(nil)

	pm := &ProjectMeta{
		ProjectID: 1,
		Uri:       "https://localhost/project_config",
		Hash:      [32]byte(hash),
	}

	t.Run("success", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder("GET", "https://localhost/project_config", httpmock.NewBytesResponder(200, jc))

		resultConfigs, err := pm.GetConfigs("")
		require.NoError(err)
		require.Equal(len(resultConfigs), len(cs))
		require.Equal(resultConfigs[0].Code, "i am code")
	})
	t.Run("http error", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder("GET", "https://localhost/project_config", httpmock.NewErrorResponder(errors.New("http error")))

		_, err := pm.GetConfigs("")
		require.ErrorContains(err, "http error")
	})
	t.Run("hash mismatch", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder("GET", "https://localhost/project_config", httpmock.NewBytesResponder(200, jc))

		npm := *pm
		npm.Hash = [32]byte{}
		_, err := npm.GetConfigs("")
		require.ErrorContains(err, "validate project config hash failed")
	})
}
