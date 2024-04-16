package project

import (
	"crypto/sha256"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/require"
)

func Test_newCache(t *testing.T) {
	r := require.New(t)

	t.Run("FailedToCreateDir", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(os.MkdirAll, errors.New(t.Name()))
		_, err := newCache("")
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(os.MkdirAll, nil)
		_, err := newCache("")
		r.NoError(err)
	})
}

func TestCache_get(t *testing.T) {
	r := require.New(t)

	c := &cache{}
	t.Run("FailedToRead", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(os.ReadFile, nil, errors.New(t.Name()))
		d := c.get(uint64(0), nil)
		r.Empty(d)
	})

	t.Run("FailedToValidate", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(os.ReadFile, []byte("data"), nil)

		d := c.get(uint64(0), []byte("data"))
		r.Empty(d)
	})

	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(os.ReadFile, []byte("data"), nil)

		data := []byte("data")
		h := sha256.New()
		h.Write(data)
		hash := h.Sum(nil)
		d := c.get(uint64(0), hash)
		r.Equal(data, d)
	})
}

func TestCache_getPath(t *testing.T) {
	r := require.New(t)

	c := &cache{
		dir: "test",
	}
	path := c.getPath(uint64(0))
	r.Equal(fmt.Sprintf("%s/%d", c.dir, 0), path)
}
