package project

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/util/ipfs"
)

var (
	errEmptyConfig = errors.New("config is empty")
	errEmptyCode   = errors.New("code is empty")
)

// TODO: prefer protobuf for serialization and deserialization
type Project struct {
	DefaultVersion string    `json:"defaultVersion,omitempty"`
	Configs        []*Config `json:"config"`
}

type Meta struct {
	ProjectID *big.Int
	Uri       string
	Hash      [32]byte
}

type Attribute struct {
	Paused                bool
	RequestedProverAmount uint64
}

type Config struct {
	Version  string `json:"version"`
	VMTypeID uint64 `json:"vmTypeID"`
	Code     string `json:"code"`
	Metadata string `json:"metadata,omitempty"`
}

func (p *Project) Config(version string) (*Config, error) {
	if len(version) == 0 {
		return p.DefaultConfig()
	}
	for _, c := range p.Configs {
		if c.Version == version {
			return c, nil
		}
	}
	return nil, errors.New("project config not exist")
}

func (p *Project) DefaultConfig() (*Config, error) {
	if len(p.DefaultVersion) > 0 {
		return p.Config(p.DefaultVersion)
	}
	return p.Configs[0], nil
}

func (p *Project) Marshal() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Project) Unmarshal(data []byte) error {
	if err := json.Unmarshal(data, p); err != nil {
		return errors.Wrap(err, "failed to unmarshal project")
	}

	if len(p.Configs) == 0 {
		return errEmptyConfig
	}
	for _, c := range p.Configs {
		if err := c.validate(); err != nil {
			return err
		}
	}
	return nil
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
