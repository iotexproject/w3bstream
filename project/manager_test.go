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
	t.Run("GetNotifySuccess", func(t *testing.T) {
		m := &Manager{
			projectIDs: map[uint64]bool{1: true},
			notify:     make(chan uint64, 1),
		}
		notify := m.GetNotify()
		m.notify <- uint64(1)
		d := <-notify
		require.Equal(d, uint64(1))
	})
	t.Run("WatchProjectRegistrarSuccess", func(t *testing.T) {
		m := &Manager{
			projectIDs: map[uint64]bool{1: true},
			notify:     make(chan uint64, 1),
			instance:   &contracts.Contracts{},
		}
		p.ApplyMethodReturn(&contracts.Contracts{}, "ParseProjectUpserted", &contracts.ContractsProjectUpserted{ProjectId: 1}, nil)
		p.ApplyMethodReturn(&ProjectMeta{}, "GetConfigs", []*Config{{}}, nil)
		defer p.Reset()

		notify := m.GetNotify()
		m.notify <- uint64(1)
		d := <-notify
		require.Equal(d, uint64(1))
	})
	t.Run("FillProjectPoolEmpty", func(t *testing.T) {
		m := &Manager{
			projectIDs: map[uint64]bool{},
			pool:       map[key]*Config{},
			instance:   &contracts.Contracts{},
		}
		p.ApplyMethodReturn(&contracts.Contracts{}, "Projects", &struct {
			Uri    string
			Hash   [32]byte
			Paused bool
		}{}, nil)
		defer p.Reset()

		require.Equal(len(m.GetAllProjectID()), 0)
	})
	t.Run("FillProjectPoolSuccess", func(t *testing.T) {
		instance := &contracts.Contracts{}
		m := &Manager{
			projectIDs: map[uint64]bool{},
			pool:       map[key]*Config{},
			instance:   instance,
		}
		data := []struct {
			Uri    string
			Hash   [32]byte
			Paused bool
		}{
			{
				Uri:  "test",
				Hash: [32]byte{1},
			},
			{},
		}
		p.ApplyMethodSeq(instance, "Projects", []gomonkey.OutputCell{
			{Values: gomonkey.Params{data[0], nil}},
			{Values: gomonkey.Params{data[1], nil}},
		})
		p.ApplyMethodReturn(&ProjectMeta{}, "GetConfigs", []*Config{{}}, nil)
		defer p.Reset()

		require.Equal(len(m.GetAllProjectID()), 1)
	})
}
