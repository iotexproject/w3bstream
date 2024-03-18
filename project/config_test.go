package project

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/project_config.json
var content []byte

func TestConfig_GetOutput(t *testing.T) {
	r := require.New(t)

	c := &Config{}
	r.NoError(json.Unmarshal(content, c))

	o, err := c.GetOutput("any", "any")
	r.NoError(err)
	r.Equal(o.Type(), c.Output.Type)
}
