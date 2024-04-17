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

	"github.com/machinefi/sprout/util/ipfs"
	"github.com/machinefi/sprout/vm"
)

func TestProjectMeta_GetProjectRawData_init(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	t.Run("InvalidUri", func(t *testing.T) {
		p = p.ApplyFuncReturn(url.Parse, nil, errors.New(t.Name()))

		_, err := (&Meta{}).FetchProjectRawData("")
		r.ErrorContains(err, t.Name())
	})
}

func TestProjectMeta_GetProjectRawData_http(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	c := Project{}
	jc, err := json.Marshal(c)
	r.NoError(err)

	h := sha256.New()
	_, err = h.Write(jc)
	r.NoError(err)
	hash := h.Sum(nil)

	pm := &Meta{
		Uri:  "https://test.com/project_config",
		Hash: [32]byte(hash),
	}

	t.Run("FailedToGetHTTP", func(t *testing.T) {
		p = p.ApplyFuncReturn(http.Get, nil, errors.New(t.Name()))

		_, err := pm.FetchProjectRawData("")
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToIOReadAll", func(t *testing.T) {
		p = p.ApplyFuncReturn(http.Get, &http.Response{
			Body: io.NopCloser(bytes.NewReader(jc)),
		}, nil)
		p = p.ApplyFuncReturn(io.ReadAll, nil, errors.New(t.Name()))

		_, err := pm.FetchProjectRawData("")
		r.ErrorContains(err, t.Name())
	})
	t.Run("HashMismatch", func(t *testing.T) {
		p = p.ApplyFuncReturn(io.ReadAll, jc, nil)

		npm := *pm
		npm.Hash = [32]byte{}
		_, err := npm.FetchProjectRawData("")
		r.ErrorContains(err, "failed to validate project hash")
	})
	t.Run("Success", func(t *testing.T) {
		_, err := pm.FetchProjectRawData("")
		r.NoError(err)
	})
}

func TestProjectMeta_GetProjectRawData_ipfs(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	pm := &Meta{
		Uri: "ipfs://test.com/123",
	}
	t.Run("FailedToGetIPFS", func(t *testing.T) {
		p = p.ApplyMethodReturn(&ipfs.IPFS{}, "Cat", nil, errors.New(t.Name()))

		_, err := pm.FetchProjectRawData("")
		r.ErrorContains(err, t.Name())
	})
}

func TestProjectMeta_GetProjectRawData_default(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	pm := &Meta{
		Uri: "test.com/123",
	}

	t.Run("FailedToGetIPFS", func(t *testing.T) {
		p = p.ApplyMethodReturn(&ipfs.IPFS{}, "Cat", nil, errors.New(t.Name()))

		_, err := pm.FetchProjectRawData("")
		r.ErrorContains(err, t.Name())
	})
}

func TestProject_GetConfig(t *testing.T) {
	r := require.New(t)

	conf := &Config{
		Version: "0.1",
	}
	project := &Project{
		Versions: []*Config{conf},
	}

	t.Run("Success", func(t *testing.T) {
		c, err := project.GetConfig("0.1")
		r.NoError(err)
		r.Equal(conf, c)
	})

	t.Run("NotExist", func(t *testing.T) {
		_, err := project.GetConfig("0.3")
		r.ErrorContains(err, "project config not exist")
	})
}

func TestProject_Validate(t *testing.T) {
	r := require.New(t)

	project := &Project{}

	t.Run("EmptyConfig", func(t *testing.T) {
		err := project.Validate()
		r.EqualError(err, errEmptyConfig.Error())
	})

	t.Run("FailedToValidate", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		project.Versions = []*Config{&Config{}}
		p = p.ApplyMethodReturn(&Config{}, "Validate", errors.New(t.Name()))

		err := project.Validate()
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		project.Versions = []*Config{&Config{}}
		p = p.ApplyMethodReturn(&Config{}, "Validate", nil)

		err := project.Validate()
		r.NoError(err)
	})
}

func TestConfig_Validate(t *testing.T) {
	r := require.New(t)

	config := &Config{
		VMType: vm.Halo2,
		Code:   "testCode",
	}

	t.Run("EmptyCode", func(t *testing.T) {
		c := *config
		c.Code = ""
		err := c.Validate()
		r.EqualError(err, errEmptyCode.Error())
	})

	t.Run("UnsupportedVMType", func(t *testing.T) {
		c := *config
		c.VMType = "test"
		err := c.Validate()
		r.EqualError(err, errUnsupportedVMType.Error())
	})

	t.Run("Success", func(t *testing.T) {
		err := config.Validate()
		r.NoError(err)
	})
}

func TestConvertProject(t *testing.T) {
	r := require.New(t)

	t.Run("FailedToUnmarshal", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(json.Unmarshal, errors.New(t.Name()))
		_, err := convertProject(nil)
		r.ErrorContains(err, t.Name())
	})

	t.Run("FailedToValidate", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(json.Unmarshal, nil)
		p = p.ApplyMethodReturn(&Project{}, "Validate", errors.New(t.Name()))
		_, err := convertProject(nil)
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(json.Unmarshal, nil)
		p = p.ApplyMethodReturn(&Project{}, "Validate", nil)
		project, err := convertProject(nil)
		r.NoError(err)
		r.Empty(project)
	})
}
