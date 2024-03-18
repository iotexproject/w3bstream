package project

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/types"
	"github.com/machinefi/sprout/utils/ipfs"
)

func TestConfig_GetOutput(t *testing.T) {
	r := require.New(t)

	t.Run("Default", func(t *testing.T) {
		c := &Config{}
		_, err := c.GetOutput("", "")
		r.NoError(err)
	})
	t.Run("Stdout", func(t *testing.T) {
		c := &Config{
			Output: OutputConfig{
				Type: types.OutputStdout,
			},
		}
		_, err := c.GetOutput("", "")
		r.NoError(err)
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
		r.NoError(err)
	})
	t.Run("Solana", func(t *testing.T) {
		c := &Config{
			Output: OutputConfig{
				Type:   types.OutputSolanaProgram,
				Solana: SolanaConfig{},
			},
		}
		_, err := c.GetOutput("", "308edd7fca562182adbffaa59264a138d9e04f9f3adbda2c80ef1ca71b7dcfa4")
		r.NoError(err)
	})
}

func TestProjectMeta_GetConfigs_init(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	t.Run("InvalidUri", func(t *testing.T) {
		p = p.ApplyFuncReturn(url.Parse, nil, errors.New(t.Name()))

		_, err := (&ProjectMeta{}).GetConfigData("")
		r.ErrorContains(err, t.Name())
	})
}

func TestProjectMeta_GetConfigs_http(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	cs := []*Config{
		{
			Code:    "i am code",
			VMType:  types.VMHalo2,
			Version: "0.1",
		},
	}
	jc, err := json.Marshal(cs)
	r.NoError(err)

	h := sha256.New()
	_, err = h.Write(jc)
	r.NoError(err)
	hash := h.Sum(nil)

	pm := &ProjectMeta{
		ProjectID: 1,
		Uri:       "https://test.com/project_config",
		Hash:      [32]byte(hash),
	}

	t.Run("FailedToGetHTTP", func(t *testing.T) {
		p = p.ApplyFuncReturn(http.Get, nil, errors.New(t.Name()))

		_, err := pm.GetConfigData("")
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToIOReadAll", func(t *testing.T) {
		p = p.ApplyFuncReturn(http.Get, &http.Response{
			Body: io.NopCloser(bytes.NewReader(jc)),
		}, nil)
		p = p.ApplyFuncReturn(io.ReadAll, nil, errors.New(t.Name()))

		_, err := pm.GetConfigData("")
		r.ErrorContains(err, t.Name())
	})
	t.Run("HashMismatch", func(t *testing.T) {
		p = p.ApplyFuncReturn(io.ReadAll, jc, nil)

		npm := *pm
		npm.Hash = [32]byte{}
		_, err := npm.GetConfigData("")
		r.ErrorContains(err, "failed to validate project config hash")
	})
	t.Run("Success", func(t *testing.T) {
		_, err := pm.GetConfigData("")
		r.NoError(err)
	})
}

func TestProjectMeta_GetConfigs_ipfs(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	pm := &ProjectMeta{
		Uri: "ipfs://test.com/123",
	}
	t.Run("FailedToGetIPFS", func(t *testing.T) {
		p = p.ApplyMethodReturn(&ipfs.IPFS{}, "Cat", nil, errors.New(t.Name()))

		_, err := pm.GetConfigData("")
		r.ErrorContains(err, t.Name())
	})
}

func TestProjectMeta_GetConfigs_default(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	pm := &ProjectMeta{
		Uri: "test.com/123",
	}

	t.Run("FailedToGetIPFS", func(t *testing.T) {
		p = p.ApplyMethodReturn(&ipfs.IPFS{}, "Cat", nil, errors.New(t.Name()))

		_, err := pm.GetConfigData("")
		r.ErrorContains(err, t.Name())
	})
}
