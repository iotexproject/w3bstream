package project

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"

	"github.com/machinefi/sprout/types"
)

type ProjectMeta struct {
	ProjectID uint64
	Uri       string
	Hash      [32]byte
	Paused    bool
}

func (m *ProjectMeta) ProjectConfig() (*Config, error) {
	rsp, err := http.Get(m.Uri)
	if err != nil {
		return nil, errors.Wrapf(err, "fetch project config: [id:%d] [uri:%s]", m.ProjectID, m.Uri)
	}
	defer rsp.Body.Close()
	content, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "read project config: [id:%d] [uri:%s]", m.ProjectID, m.Uri)
	}
	c := &Config{}
	if err = json.Unmarshal(content, c); err != nil {
		return nil, errors.Wrapf(err, "parse project config: [id:%d] [uri:%s]", m.ProjectID, m.Uri)
	}
	// simple validation
	if len(c.Code) == 0 || c.VMType == "" {
		return nil, errors.Errorf("invalid project config: [id:%d] [uri:%s]", m.ProjectID, m.Uri)
	}
	// TODO validate hash
	if c.Output.Type == "" {
		c.Output.Type = types.OutputStdout
	}
	return c, nil
}
