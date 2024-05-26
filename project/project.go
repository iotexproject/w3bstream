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
	"github.com/machinefi/sprout/util/ipfs"
	"github.com/machinefi/sprout/vm"
)

var (
	errEmptyConfig       = errors.New("config is empty")
	errEmptyCode         = errors.New("code is empty")
	errUnsupportedVMType = errors.New("unsupported vm type")
)

type Project struct {
	DatasourceURI  string    `json:"datasourceURI,omitempty"`
	DefaultVersion string    `json:"defaultVersion"`
	Versions       []*Config `json:"versions"`
}

type Meta struct {
	ProjectID uint64
	Uri       string
	Hash      [32]byte
}

type Attribute struct {
	Paused                bool
	RequestedProverAmount uint64
}

type Config struct {
	Version      string        `json:"version"`
	VMType       vm.Type       `json:"vmType"`
	Output       output.Config `json:"output"`
	CodeExpParam string        `json:"codeExpParam,omitempty"`
	Code         string        `json:"code"`
}

func (p *Project) Config(version string) (*Config, error) {
	for _, c := range p.Versions {
		if c.Version == version {
			return c, nil
		}
	}
	return nil, errors.New("project config not exist")
}

func (p *Project) DefaultConfig() (*Config, error) {
	return p.Config(p.DefaultVersion)
}

func (c *Config) validate() error {
	if len(c.Code) == 0 {
		return errEmptyCode
	}
	switch c.VMType {
	default:
		return errUnsupportedVMType
	case vm.Halo2, vm.Wasm, vm.Risc0, vm.ZKwasm:
		return nil
	}
}

func (m *Meta) FetchProjectRawData(ipfsEndpoint string) ([]byte, error) {
	u, err := url.Parse(m.Uri)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse project uri %s", m.Uri)
	}

	var data []byte
	switch u.Scheme {
	case "http", "https":
		resp, _err := http.Get(m.Uri)
		if _err != nil {
			return nil, errors.Wrapf(_err, "failed to fetch project, uri %s", m.Uri)
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
		return nil, errors.Wrapf(err, "failed to read project, uri %s", m.Uri)
	}

	h := sha256.New()
	if _, err := h.Write(data); err != nil {
		return nil, errors.Wrap(err, "failed to generate project hash")
	}
	if !bytes.Equal(h.Sum(nil), m.Hash[:]) {
		return nil, errors.New("failed to validate project hash")
	}

	return data, nil
}

func convertProject(projectRawData []byte) (*Project, error) {
	p := &Project{}
	if err := json.Unmarshal(projectRawData, &p); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal project")
	}

	if len(p.Versions) == 0 {
		return nil, errEmptyConfig
	}
	for _, c := range p.Versions {
		if err := c.validate(); err != nil {
			return nil, err
		}
	}
	return p, nil
}
