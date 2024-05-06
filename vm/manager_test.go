package vm

import (
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/require"
)

func TestMgr(t *testing.T) {
	r := require.New(t)
	m := newManager()
	p := gomonkey.NewPatches()
	defer p.Reset()

	p.ApplyFuncReturn(newInstance, &instance{}, nil)

	i, err := m.acquire(1, "any", "any", "any")
	r.NoError(err)

	m.release(1, i)

	i2, err := m.acquire(1, "any", "any", "any")
	r.NoError(err)
	r.Equal(i, i2)
}
