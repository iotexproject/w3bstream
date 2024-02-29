package project

import (
	"runtime"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/machinefi/sprout/project/contracts"
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
		p.ApplyFuncReturn(ethclient.Dial, nil, errors.New(t.Name()))
		defer p.Reset()

		_, err := NewManager("", "", "")
		require.ErrorContains(err, t.Name())
	})
	t.Run("NewManagerNewContractsFailed", func(t *testing.T) {
		p.ApplyFuncReturn(ethclient.Dial, nil, nil)
		p.ApplyFuncReturn(contracts.NewContracts, nil, errors.New(t.Name()))
		defer p.Reset()

		_, err := NewManager("", "", "")
		require.ErrorContains(err, t.Name())
	})
	// t.Run("NewManagerSuccess", func(t *testing.T) {
	// 	p.ApplyFuncReturn(ethclient.Dial, ethclient.NewClient(&rpc.Client{}), nil)
	// 	p.ApplyFuncReturn(contracts.NewContracts, nil, nil)
	// 	p.ApplyPrivateMethod(&Manager{}, "fillProjectPool", func() {})
	// 	p.ApplyMethodReturn(&ethclient.Client{}, "BlockNumber", uint64(0), nil)
	// 	//p.ApplyPrivateMethod(&Monitor{}, "run", func() {})
	// 	//p.ApplyPrivateMethod(&Manager{}, "watchProjectRegistrar", func(<-chan *types.Log, event.Subscription) {})
	// 	defer p.Reset()

	// 	_, err := NewManager("", "", "")
	// 	require.NoError(err)
	// })
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
