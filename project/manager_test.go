package project

import (
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/project/contracts"
	"github.com/machinefi/sprout/testutil"
)

func TestNewManager(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	t.Run("DialChainFailed", func(t *testing.T) {
		p = testutil.EthClientDial(p, nil, errors.New(t.Name()))

		_, err := NewManager("", "", "", "")
		r.ErrorContains(err, t.Name())
	})
	t.Run("NewContractsFailed", func(t *testing.T) {
		p = testutil.EthClientDial(p, ethclient.NewClient(&rpc.Client{}), nil)
		p = p.ApplyFuncReturn(contracts.NewContracts, nil, errors.New(t.Name()))

		_, err := NewManager("", "", "", "")
		r.ErrorContains(err, t.Name())
	})
	t.Run("NewDefaultMonitorFailed", func(t *testing.T) {
		p = p.ApplyFuncReturn(contracts.NewContracts, nil, nil)
		p = p.ApplyPrivateMethod(&Manager{}, "fillProjectPool", func() {})
		p = p.ApplyFuncReturn(NewDefaultMonitor, nil, errors.New(t.Name()))

		_, err := NewManager("", "", "", "")
		r.ErrorContains(err, t.Name())
	})
	t.Run("Success", func(t *testing.T) {
		p = p.ApplyFuncReturn(NewDefaultMonitor, &Monitor{}, nil)
		p = p.ApplyPrivateMethod(&Monitor{}, "run", func() {})
		p = p.ApplyMethodReturn(&Monitor{}, "MustEvents", make(chan *types.Log))
		p = p.ApplyPrivateMethod(&Manager{}, "watchProjectRegistrar", func(<-chan *types.Log, event.Subscription) {})

		_, err := NewManager("", "", "", "")
		r.NoError(err)
	})
}

type testSubscription struct {
	errChain chan error
}

func (s testSubscription) Err() <-chan error {
	return s.errChain
}

func (s testSubscription) Unsubscribe() {}

func TestManager_Get(t *testing.T) {
	r := require.New(t)

	t.Run("NotExist", func(t *testing.T) {
		m := &Manager{}
		_, err := m.Get(1, "0.1")
		r.ErrorContains(err, "project config not exist")
	})
	t.Run("Success", func(t *testing.T) {
		m := &Manager{
			pool: map[key]*Config{getKey(1, "0.1"): {}},
		}
		_, err := m.Get(1, "0.1")
		r.NoError(err)
	})
}

func TestManager_Set(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		m := &Manager{
			pool:       map[key]*Config{},
			projectIDs: map[uint64]bool{},
		}
		m.Set(1, "0.1", &Config{})
	})
}

func TestManager_GetAllProjectID(t *testing.T) {
	r := require.New(t)

	t.Run("Success", func(t *testing.T) {
		m := &Manager{
			projectIDs: map[uint64]bool{1: true},
		}
		ids := m.GetAllProjectID()
		r.Equal(len(ids), 1)
		r.Equal(ids[0], uint64(1))
	})
}

func TestManager_GetNotify(t *testing.T) {
	r := require.New(t)

	t.Run("Success", func(t *testing.T) {
		m := &Manager{
			projectIDs: map[uint64]bool{1: true},
			notify:     make(chan uint64, 1),
		}
		notify := m.GetNotify()
		m.notify <- uint64(1)
		d := <-notify
		r.Equal(d, uint64(1))
	})
}

func TestManager_doProjectRegistrarWatch(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	t.Run("Success", func(t *testing.T) {
		p = p.ApplyMethodReturn(&contracts.ContractsFilterer{}, "ParseProjectUpserted", &contracts.ContractsProjectUpserted{ProjectId: 1}, nil)
		p = p.ApplyMethodReturn(&ProjectMeta{}, "GetConfigs", []*Config{{}}, nil)

		m := &Manager{
			projectIDs: map[uint64]bool{},
			pool:       map[key]*Config{},
			notify:     make(chan uint64, 1),
			instance:   &contracts.Contracts{},
		}

		errChain := make(chan error)
		logChain := make(chan *types.Log, 1)
		logChain <- &types.Log{}

		m.doProjectRegistrarWatch(logChain, testSubscription{errChain})
		notify := m.GetNotify()
		d := <-notify
		r.Equal(d, uint64(1))
	})
}

func TestManager_doProjectPoolFill(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	t.Run("ReadChainFailed", func(t *testing.T) {
		p = p.ApplyMethodReturn(&contracts.ContractsCaller{}, "Projects", nil, errors.New(t.Name()))

		m := &Manager{
			projectIDs: map[uint64]bool{},
			pool:       map[key]*Config{},
			instance:   &contracts.Contracts{},
		}
		finished := m.doProjectPoolFill(1)
		r.False(finished)
		r.Equal(len(m.GetAllProjectID()), 0)
	})
	t.Run("Finished", func(t *testing.T) {
		p = p.ApplyMethodReturn(&contracts.ContractsCaller{}, "Projects", struct {
			Uri    string
			Hash   [32]byte
			Paused bool
		}{}, nil)

		m := &Manager{
			projectIDs: map[uint64]bool{},
			pool:       map[key]*Config{},
			instance:   &contracts.Contracts{},
		}
		finished := m.doProjectPoolFill(1)
		r.True(finished)
		r.Equal(len(m.GetAllProjectID()), 0)
	})
	t.Run("GetConfigFailed", func(t *testing.T) {
		p = p.ApplyMethodReturn(&contracts.ContractsCaller{}, "Projects", struct {
			Uri    string
			Hash   [32]byte
			Paused bool
		}{
			Uri:  "test",
			Hash: [32]byte{1},
		}, nil)

		m := &Manager{
			projectIDs: map[uint64]bool{},
			pool:       map[key]*Config{},
			instance:   &contracts.Contracts{},
		}
		finished := m.doProjectPoolFill(1)
		r.False(finished)
		r.Equal(len(m.GetAllProjectID()), 0)
	})
	t.Run("Success", func(t *testing.T) {
		p = p.ApplyMethodReturn(&contracts.ContractsCaller{}, "Projects", struct {
			Uri    string
			Hash   [32]byte
			Paused bool
		}{
			Uri:  "test",
			Hash: [32]byte{1},
		}, nil)
		p = p.ApplyMethodReturn(&ProjectMeta{}, "GetConfigs", []*Config{{}}, nil)

		m := &Manager{
			projectIDs: map[uint64]bool{},
			pool:       map[key]*Config{},
			instance:   &contracts.Contracts{},
		}
		finished := m.doProjectPoolFill(1)
		r.False(finished)
		r.Equal(len(m.GetAllProjectID()), 1)
	})
}

func TestManager_fillProjectPool(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	t.Run("Success", func(t *testing.T) {
		p = p.ApplyPrivateMethod(&Manager{}, "doProjectPoolFill", func(uint64) bool { return true })

		m := &Manager{
			projectIDs: map[uint64]bool{},
			pool:       map[key]*Config{},
			instance:   &contracts.Contracts{},
		}
		m.fillProjectPool()
		r.Equal(len(m.GetAllProjectID()), 0)
	})
}
