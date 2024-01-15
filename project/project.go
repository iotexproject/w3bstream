package project

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"

	"github.com/machinefi/sprout/output"
	"github.com/machinefi/sprout/types"
	"github.com/machinefi/sprout/utils/ipfs"
)

type Config struct {
	Code         string       `json:"code"`
	CodeExpParam string       `json:"codeExpParam,omitempty"`
	VMType       types.VM     `json:"vmType"`
	Output       OutputConfig `json:"output,omitempty"`
	Version      string       `json:"version"`
}

type OutputConfig struct {
	Type types.Output `json:"type"`

	Ethereum struct {
		ChainEndpoint   string `json:"chainEndpoint"`
		ContractAddress string `json:"contractAddress"`
		ContractMethod  string `json:"contractMethod"`
		ContractAbiJSON string `json:"contractAbiJSON"`
	} `json:"ethereum,omitempty"`

	Solana struct {
		ChainEndpoint  string `json:"chainEndpoint"`
		ProgramID      string `json:"programID"`
		StateAccountPK string `json:"stateAccountPK"`
	} `json:"solana,omitempty"`
}

func (c *Config) GetOutput(privateKeyECDSA, privateKeyED25519 string) (output.Output, error) {
	outConf := c.Output

	switch outConf.Type {
	case types.OutputEthereumContract:
		ethConf := outConf.Ethereum
		return output.NewEthereum(ethConf.ChainEndpoint, privateKeyECDSA, ethConf.ContractAddress, ethConf.ContractAbiJSON, ethConf.ContractMethod)
	case types.OutputSolanaProgram:
		solConf := outConf.Solana
		return output.NewSolanaProgram(solConf.ChainEndpoint, solConf.ProgramID, privateKeyED25519, solConf.StateAccountPK)
	default:
		return output.NewStdout(), nil
	}
}

type ProjectMeta struct {
	ProjectID uint64
	Uri       string
	Hash      [32]byte
	Paused    bool
}

func (m *ProjectMeta) GetConfigs(ipfsEndpoint string) ([]*Config, error) {
	var (
		content []byte
		err     error
	)
	if u, _err := url.Parse(m.Uri); _err == nil && (u.Scheme == "http" || u.Scheme == "https") {
		// fetch content from ipfs gateway
		resp, _err := http.Get(m.Uri)
		if _err != nil {
			return nil, errors.Wrapf(err, "fetch project config failed, projectID %d, uri %s", m.ProjectID, m.Uri)
		}
		defer resp.Body.Close()

		// TODO network error should try again
		content, err = io.ReadAll(resp.Body)
	} else {
		// fetch content by ipfs cid
		sh := ipfs.NewIPFS(ipfsEndpoint)
		content, err = sh.Cat(m.Uri)
	}

	if err != nil {
		return nil, errors.Wrapf(err, "read project config failed, projectID %d, uri %s", m.ProjectID, m.Uri)
	}

	h := sha256.New()
	_, err = h.Write(content)
	if err != nil {
		return nil, errors.Wrap(err, "failed to validate hash")
	}

	sha256sum := h.Sum(nil)
	if !bytes.Equal(sha256sum, m.Hash[:]) {
		return nil, errors.Wrap(err, "failed to validate hash, not equal expect")
	}

	cs := []*Config{}
	// TODO parsing error should skip
	if err = json.Unmarshal(content, &cs); err != nil {
		return nil, errors.Wrapf(err, "parse project config failed, projectID %d, uri %s", m.ProjectID, m.Uri)
	}

	if len(cs) == 0 {
		return nil, errors.Errorf("empty project config, projectID %d, uri %s", m.ProjectID, m.Uri)
	}
	for _, c := range cs {
		if c.Code == "" || c.VMType == "" || c.Version == "" {
			return nil, errors.Errorf("invalid project config, projectID %d, uri %s", m.ProjectID, m.Uri)
		}
	}

	return cs, nil
}
