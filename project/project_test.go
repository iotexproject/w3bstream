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

	"github.com/machinefi/sprout/utils/ipfs"
)

func TestProjectMeta_GetConfigs_init(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	t.Run("InvalidUri", func(t *testing.T) {
		p = p.ApplyFuncReturn(url.Parse, nil, errors.New(t.Name()))

		_, err := (&Meta{}).GetProjectRawData("")
		r.ErrorContains(err, t.Name())
	})
}

func TestProjectMeta_GetConfigs_http(t *testing.T) {
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

		_, err := pm.GetProjectRawData("")
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToIOReadAll", func(t *testing.T) {
		p = p.ApplyFuncReturn(http.Get, &http.Response{
			Body: io.NopCloser(bytes.NewReader(jc)),
		}, nil)
		p = p.ApplyFuncReturn(io.ReadAll, nil, errors.New(t.Name()))

		_, err := pm.GetProjectRawData("")
		r.ErrorContains(err, t.Name())
	})
	t.Run("HashMismatch", func(t *testing.T) {
		p = p.ApplyFuncReturn(io.ReadAll, jc, nil)

		npm := *pm
		npm.Hash = [32]byte{}
		_, err := npm.GetProjectRawData("")
		r.ErrorContains(err, "failed to validate project hash")
	})
	t.Run("Success", func(t *testing.T) {
		_, err := pm.GetProjectRawData("")
		r.NoError(err)
	})
}

func TestProjectMeta_GetConfigs_ipfs(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	pm := &Meta{
		Uri: "ipfs://test.com/123",
	}
	t.Run("FailedToGetIPFS", func(t *testing.T) {
		p = p.ApplyMethodReturn(&ipfs.IPFS{}, "Cat", nil, errors.New(t.Name()))

		_, err := pm.GetProjectRawData("")
		r.ErrorContains(err, t.Name())
	})
}

func TestProjectMeta_GetConfigs_default(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	pm := &Meta{
		Uri: "test.com/123",
	}

	t.Run("FailedToGetIPFS", func(t *testing.T) {
		p = p.ApplyMethodReturn(&ipfs.IPFS{}, "Cat", nil, errors.New(t.Name()))

		_, err := pm.GetProjectRawData("")
		r.ErrorContains(err, t.Name())
	})
}
