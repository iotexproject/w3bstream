package project

import (
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	testeth "github.com/machinefi/sprout/testutil/eth"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestManager(t *testing.T) {
	require := require.New(t)
	p := gomonkey.NewPatches()

	t.Run("NewManagerDialChainFailed", func(t *testing.T) {
		testeth.EthclientDial(p, nil, errors.New(t.Name()))
		defer p.Reset()

		_, err := NewMonitor("", []string{}, []string{}, 1, 100, 3*time.Second)
		require.ErrorContains(err, t.Name())
	})
	t.Run("GetNotExist", func(t *testing.T) {
		m := &Manager{}
		_, err := m.Get(1, "0.1")
		require.ErrorContains(err, "project config not exist")
	})
	t.Run("GetSuccess", func(t *testing.T) {
		m := &Manager{
			pool: map[key]*Config{getKey(1, "0.1"): {}},
		}
		_, err := m.Get(1, "0.1")
		require.NoError(err)
	})
	t.Run("SetSuccess", func(t *testing.T) {
		m := &Manager{
			pool:       map[key]*Config{},
			projectIDs: map[uint64]bool{},
		}
		m.Set(1, "0.1", &Config{})
	})
	t.Run("GetAllSuccess", func(t *testing.T) {
		m := &Manager{
			projectIDs: map[uint64]bool{1: true},
		}
		ids := m.GetAllProjectID()
		require.Equal(len(ids), 1)
		require.Equal(ids[0], uint64(1))
	})
}
