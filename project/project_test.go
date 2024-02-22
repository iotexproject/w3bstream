package project

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"io"
	"net/http"
	"runtime"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/machinefi/sprout/testutil"
	"github.com/machinefi/sprout/types"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestGetOutput(t *testing.T) {
	require := require.New(t)

	t.Run("Default", func(t *testing.T) {
		c := &Config{}
		_, err := c.GetOutput("", "")
		require.NoError(err)
	})
	t.Run("Stdout", func(t *testing.T) {
		c := &Config{
			Output: OutputConfig{
				Type: types.OutputStdout,
			},
		}
		_, err := c.GetOutput("", "")
		require.NoError(err)
	})
	t.Run("Ethereum", func(t *testing.T) {
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
	t.Run("Solana", func(t *testing.T) {
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
	if runtime.GOOS == `darwin` {
		return
	}
	require := require.New(t)
	p := gomonkey.NewPatches()

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

	t.Run("InvalidUri", func(t *testing.T) {
		testutil.URLParse(p, nil, errors.New(t.Name()))
		_, err := pm.GetConfigs("")
		require.ErrorContains(err, t.Name())
	})
	t.Run("GetHTTPFailed", func(t *testing.T) {
		testutil.HttpGet(p, nil, errors.New(t.Name()))
		_, err := pm.GetConfigs("")
		require.ErrorContains(err, t.Name())
	})
	t.Run("IOReadAllFailed", func(t *testing.T) {
		testutil.HttpGet(p, &http.Response{
			Body: io.NopCloser(bytes.NewReader(jc)),
		}, nil)
		testutil.IoReadAll(p, nil, errors.New(t.Name()))
		_, err := pm.GetConfigs("")
		require.ErrorContains(err, t.Name())
	})
	t.Run("HashMismatch", func(t *testing.T) {
		testutil.HttpGet(p, &http.Response{
			Body: io.NopCloser(bytes.NewReader(jc)),
		}, nil)

		npm := *pm
		npm.Hash = [32]byte{}
		_, err := npm.GetConfigs("")
		require.ErrorContains(err, "validate project config hash failed")
	})
	t.Run("Success", func(t *testing.T) {
		testutil.HttpGet(p, &http.Response{
			Body: io.NopCloser(bytes.NewReader(jc)),
		}, nil)

		resultConfigs, err := pm.GetConfigs("")
		require.NoError(err)
		require.Equal(len(resultConfigs), len(cs))
		require.Equal(resultConfigs[0].Code, "i am code")
	})
}
