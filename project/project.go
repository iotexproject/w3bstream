package project

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"

	"github.com/machinefi/sprout/output"
	"github.com/machinefi/sprout/utils/ipfs"
	"github.com/machinefi/sprout/vm"
)

var (
	errEmptyConfigData    = errors.New("empty config data")
	errInvalidProjectCode = errors.New("invalid project code")
	errUnsupportedVMType  = errors.New("unsupported vm type")
)

type Config struct {
	Datasource     string        `json:"datasource"`
	DefaultVersion string        `json:"defaultVersion"`
	Versions       []*ConfigData `json:"versions"`
	defaultData    *ConfigData
	versions       map[string]*ConfigData
}

func (c *Config) Validate() error {
	if len(c.Versions) == 0 {
		return errEmptyConfigData
	}

	// remove invalid version data and set default
	var firstVersionData *ConfigData
	for i, v := range c.Versions {
		if err := v.Validate(); err != nil {
			c.Versions = append(c.Versions[:i], c.Versions[i+1:]...)
			continue
		}
		if c.versions == nil {
			c.versions = make(map[string]*ConfigData)
		}
		if _, ok := c.versions[v.Version]; ok {
			continue // to avoid overwrite exist version data
		}
		c.versions[v.Version] = v
		if firstVersionData == nil {
			firstVersionData = v
		}
	}

	// ensure contains at least one valid version data
	if firstVersionData == nil {
		return errEmptyConfigData
	}

	c.defaultData = c.versions[c.DefaultVersion]
	if c.defaultData == nil {
		c.defaultData = firstVersionData
	}

	// inherit default values
	if c.defaultData.DatasourceURI == "" {
		c.defaultData.DatasourceURI = c.Datasource
	}

	return nil
}

func (c *Config) DefaultConfigData() *ConfigData {
	return c.defaultData
}

func (c *Config) GetConfigDataByVersion(version string) *ConfigData {
	return c.versions[version]
}

type ConfigData struct {
	DatasourceURI   string            `json:"datasourceURI"`
	VMType          vm.Type           `json:"vmType"`
	Output          output.Config     `json:"output"`
	Aggregation     AggregationConfig `json:"aggregation"`
	ResourceRequest ResourceRequest   `json:"resourceRequest"`
	Version         string            `json:"version"`
	CodeExpParam    string            `json:"codeExpParam,omitempty"`
	Code            string            `json:"code"`
}

func (d *ConfigData) Validate() error {
	if len(d.Code) == 0 {
		return errInvalidProjectCode
	}
	switch d.VMType {
	default:
		return errUnsupportedVMType
	case vm.Halo2, vm.Wasm, vm.Risc0, vm.ZKwasm:
		return nil
	}
}

type AggregationConfig struct {
	Amount uint `json:"amount,omitempty"`
}

type ResourceRequest struct {
	ProverAmount uint `json:"proverAmount,omitempty"`
}

type Project struct {
	Config  *Config
	Provers []string
}

type ProjectMeta struct {
	ProjectID    uint64
	Uri          string
	Hash         [32]byte
	Paused       bool
	ProverAmount uint // TODO change this after contract code updated
}

func (m *ProjectMeta) GetConfigData(ipfsEndpoint string) ([]byte, error) {
	u, err := url.Parse(m.Uri)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse project uri %s", m.Uri)
	}

	var data []byte
	switch u.Scheme {
	case "http", "https":
		resp, _err := http.Get(m.Uri)
		if _err != nil {
			return nil, errors.Wrapf(_err, "failed to fetch project config, project_id %d, uri %s", m.ProjectID, m.Uri)
		}
		defer resp.Body.Close()
		// TODO network error should try again
		data, err = io.ReadAll(resp.Body)

	case "ipfs":
		// ipfs url: ipfs://${endpoint}/${cid}
		sh := ipfs.NewIPFS(u.Host)
		cid := strings.Split(strings.Trim(u.Path, "/"), "/")
		data, err = sh.Cat(cid[0])

	default:
		// fetch content by ipfs cid with default endpoint
		sh := ipfs.NewIPFS(ipfsEndpoint)
		cid := strings.Split(strings.Trim(u.Path, "/"), "/")
		data, err = sh.Cat(cid[0])
	}

	if err != nil {
		return nil, errors.Wrapf(err, "failed to read project config, project_id %d, uri %s", m.ProjectID, m.Uri)
	}

	h := sha256.New()
	if _, err := h.Write(data); err != nil {
		return nil, errors.Wrap(err, "failed to generate project config hash")
	}
	if !bytes.Equal(h.Sum(nil), m.Hash[:]) {
		return nil, errors.New("failed to validate project config hash")
	}

	return data, nil
}

func convertConfigs(data []byte) (*Config, error) {
	c := &Config{}
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal project config")
	}
	if err := c.Validate(); err != nil {
		return nil, err
	}
	return c, nil
}
