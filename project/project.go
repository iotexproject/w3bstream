package project

import (
	"bytes"
	"crypto/sha256"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/mailru/easyjson"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/output"
	"github.com/iotexproject/w3bstream/util/ipfs"
)

var (
	errEmptyConfig = errors.New("config is empty")
	errEmptyCode   = errors.New("code is empty")
)

type Project struct {
	DatasourceURI    string    `json:"datasourceURI,omitempty"`
	DatasourcePubKey string    `json:"datasourcePublicKey,omitempty"`
	DefaultVersion   string    `json:"defaultVersion"`
	Versions         []*Config `json:"versions"`
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
	VMTypeID     uint64        `json:"vmTypeID"`
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
	return nil
}

func (m *Meta) FetchProjectFile() ([]byte, error) {
	u, err := url.Parse(m.Uri)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse project file uri %s", m.Uri)
	}

	var data []byte
	switch u.Scheme {
	case "http", "https":
		resp, _err := http.Get(m.Uri)
		if _err != nil {
			return nil, errors.Wrapf(_err, "failed to fetch project file, uri %s", m.Uri)
		}
		defer resp.Body.Close()
		data, err = io.ReadAll(resp.Body)

	case "ipfs":
		// ipfs url: ipfs://${endpoint}/${cid}
		sh := ipfs.NewIPFS(u.Host)
		cid := strings.Split(strings.Trim(u.Path, "/"), "/")
		data, err = sh.Cat(cid[0])

	default:
		return nil, errors.Errorf("invalid project file uri %s", m.Uri)
	}

	if err != nil {
		return nil, errors.Wrapf(err, "failed to read project file, uri %s", m.Uri)
	}

	h := sha256.New()
	if _, err := h.Write(data); err != nil {
		return nil, errors.Wrap(err, "failed to generate project file hash")
	}
	if !bytes.Equal(h.Sum(nil), m.Hash[:]) {
		return nil, errors.New("failed to validate project file hash")
	}

	return data, nil
}

func convertProject(projectFile []byte) (*Project, error) {
	p := &Project{}
	if err := easyjson.Unmarshal(projectFile, p); err != nil {
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
