package server_test

import (
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/vm/server"
)

func TestMgr(t *testing.T) {
	r := require.New(t)
	m := server.NewMgr()
	p := gomonkey.NewPatches()
	defer p.Reset()

	p = p.ApplyFuncReturn(server.NewInstance, &server.Instance{}, nil)

	i, err := m.Acquire(1, "any", "any", "any")
	r.NoError(err)

	m.Release(1, i)

	i2, err := m.Acquire(1, "any", "any", "any")
	r.NoError(err)
	r.Equal(i, i2)
}
