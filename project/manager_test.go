package project

import (
	"runtime"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	testeth "github.com/machinefi/sprout/testutil/eth"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestManager(t *testing.T) {
	if runtime.GOOS == `darwin` {
		return
	}
	require := require.New(t)
	p := gomonkey.NewPatches()

	t.Run("NewManagerDialChainFailed", func(t *testing.T) {
		testeth.EthclientDial(p, nil, errors.New(t.Name()))
		defer p.Reset()

		_, err := NewManager("", "", "")
		require.ErrorContains(err, t.Name())
	})
	t.Run("NewManagerNewContractsFailed", func(t *testing.T) {
		testeth.EthclientDial(p, nil, nil)
		testeth.ProjectRegistrarContract(p, nil, errors.New(t.Name()))
		defer p.Reset()

		_, err := NewManager("", "", "")
		require.ErrorContains(err, t.Name())
	})
	t.Run("NewManagerNewDefaultMonitorFailed", func(t *testing.T) {
		testeth.EthclientDial(p, nil, nil)
		testeth.ProjectRegistrarContract(p, nil, nil)
		p.ApplyPrivateMethod(&Manager{}, "fillProjectPool", func() {})
		p.ApplyFuncReturn(NewDefaultMonitor, nil, errors.New(t.Name()))
		defer p.Reset()

		_, err := NewManager("", "", "")
		require.ErrorContains(err, t.Name())
	})
	t.Run("NewManagerSuccess", func(t *testing.T) {
		testeth.EthclientDial(p, nil, nil)
		testeth.ProjectRegistrarContract(p, nil, nil)
		p.ApplyPrivateMethod(&Manager{}, "fillProjectPool", func() {})
		p.ApplyFuncReturn(NewDefaultMonitor, nil, &Monitor{})
		p.ApplyPrivateMethod(&Monitor{}, "run", func() {})
		p.ApplyPrivateMethod(&Manager{}, "watchProjectRegistrar", func(<-chan *types.Log, event.Subscription) {})
		defer p.Reset()

		_, err := NewManager("", "", "")
		require.NoError(err)
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
