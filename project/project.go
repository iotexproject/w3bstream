package project

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/machinefi/sprout/types"
	"github.com/pkg/errors"
)

type (
	Project struct {
		ID     uint64 `json:"id"`
		Config Config `json:"config"`
	}

	Config struct {
		Code         string       `json:"code"`
		CodeExpParam string       `json:"codeExpParam,omitempty"`
		VMType       types.VM     `json:"vmType"`
		Output       OutputConfig `json:"output,omitempty"`
	}

	// OutputConfig is the config for output
	OutputConfig struct {
		Type types.Output `json:"type"`

		Ethereum struct {
			ChainName       string `json:"chainName"`
			ContractAddress string `json:"contractAddress"`
		} `json:"ethereum,omitempty"`

		Solana struct {
			ChainName      string `json:"chainName"`
			ProgramID      string `json:"programID"`
			StateAccountPK string `json:"stateAccountPK"`
		} `json:"solana,omitempty"`
	}
)

type ProjectMeta struct {
	ProjectID uint64
	Uri       string
	Hash      [32]byte
	Paused    bool
}

func (m *ProjectMeta) GetProject() (*Project, error) {
	resp, err := http.Get(m.Uri)
	if err != nil {
		return nil, errors.Wrapf(err, "fetch project config failed, projectID %d, uri %s", m.ProjectID, m.Uri)
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "read project config failed, projectID %d, uri %s", m.ProjectID, m.Uri)
	}
	c := &Config{}
	if err = json.Unmarshal(content, c); err != nil {
		return nil, errors.Wrapf(err, "parse project config failed, projectID %d, uri %s", m.ProjectID, m.Uri)
	}
	// simple validation
	if len(c.Code) == 0 || c.VMType == "" {
		return nil, errors.Errorf("invalid project config, projectID %d, uri %s", m.ProjectID, m.Uri)
	}
	// TODO validate hash
	return &Project{
		ID:     m.ProjectID,
		Config: *c,
	}, nil
}
