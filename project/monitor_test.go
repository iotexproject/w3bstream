package project

import (
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestNewMonitor(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	t.Run("FailedToDialChain", func(t *testing.T) {
		p = p.ApplyFuncReturn(ethclient.Dial, nil, errors.New(t.Name()))

		_, err := NewMonitor("", []string{}, []string{}, 1, 100, 3*time.Second)
		r.ErrorContains(err, t.Name())
	})
	t.Run("Success", func(t *testing.T) {
		p = p.ApplyFuncReturn(ethclient.Dial, nil, nil)

		_, err := NewMonitor("", []string{"0x02feBE78F3A740b3e9a1CaFAA1b23a2ac0793D26"}, []string{"ProjectUpserted(uint64,string,bytes32)"}, 1, 100, 3*time.Second)
		r.NoError(err)
	})
}

func TestNewDefaultMonitor(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	t.Run("FailedToDialChain", func(t *testing.T) {
		p = p.ApplyFuncReturn(ethclient.Dial, nil, errors.New(t.Name()))

		_, err := NewDefaultMonitor("", []string{}, []string{})
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToGetBlockNumber", func(t *testing.T) {
		p = p.ApplyFuncReturn(ethclient.Dial, ethclient.NewClient(nil), nil)
		p = p.ApplyMethodReturn(&ethclient.Client{}, "Close")
		p = p.ApplyMethodReturn(&ethclient.Client{}, "BlockNumber", uint64(0), errors.New(t.Name()))

		_, err := NewDefaultMonitor("", []string{}, []string{})
		r.ErrorContains(err, t.Name())
	})
	t.Run("Success", func(t *testing.T) {
		p = p.ApplyMethodReturn(&ethclient.Client{}, "BlockNumber", uint64(100), nil)
		p = p.ApplyFuncReturn(NewMonitor, &Monitor{}, nil)

		_, err := NewDefaultMonitor("", []string{}, []string{})
		r.NoError(err)
	})
}

func TestMonitor_Err(t *testing.T) {
	r := require.New(t)

	t.Run("Success", func(t *testing.T) {
		m := &Monitor{
			err: make(chan error, 1),
		}
		res := m.Err()
		err := errors.New(t.Name())
		m.err <- err
		r.Equal(<-res, err)
	})
}

func TestMonitor_Unsubscribe(t *testing.T) {
	r := require.New(t)

	t.Run("Success", func(t *testing.T) {
		m := &Monitor{
			stop: make(chan struct{}, 1),
		}
		m.Unsubscribe()
		r.NotNil(<-m.stop)
	})
}

func TestMonitor_Events(t *testing.T) {
	r := require.New(t)

	t.Run("Success", func(t *testing.T) {
		topic := "test"
		m := &Monitor{
			events: map[common.Hash]chan *types.Log{crypto.Keccak256Hash([]byte(topic)): make(chan *types.Log)},
		}
		_, ok := m.Events(topic)
		r.True(ok)
	})
}

func TestMonitor_MustEvents(t *testing.T) {
	r := require.New(t)

	t.Run("Success", func(t *testing.T) {
		topic := "test"
		m := &Monitor{
			events: map[common.Hash]chan *types.Log{crypto.Keccak256Hash([]byte(topic)): make(chan *types.Log)},
		}
		ch := m.MustEvents(topic)
		r.NotNil(ch)
	})
}

func TestMonitor_doRun(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	t.Run("Stopped", func(t *testing.T) {
		m := &Monitor{
			stop: make(chan struct{}, 1),
		}
		m.Unsubscribe()
		finished := m.doRun()
		r.True(finished)
	})
	t.Run("FailedToGetBlockNumber", func(t *testing.T) {
		p = p.ApplyMethodReturn(&ethclient.Client{}, "BlockNumber", uint64(100), errors.New(t.Name()))
		p = p.ApplyFuncReturn(time.Sleep)

		m := &Monitor{
			stop: make(chan struct{}, 1),
		}
		finished := m.doRun()
		r.False(finished)
	})
	t.Run("BlockNumberBehind", func(t *testing.T) {
		p = p.ApplyMethodReturn(&ethclient.Client{}, "BlockNumber", uint64(100), nil)

		m := &Monitor{
			latest: 1000,
			stop:   make(chan struct{}, 1),
		}
		finished := m.doRun()
		r.False(finished)
	})
	t.Run("FailedToFilterLogs", func(t *testing.T) {
		p = p.ApplyMethodReturn(&ethclient.Client{}, "FilterLogs", nil, errors.New(t.Name()))

		m := &Monitor{
			latest: 1,
			stop:   make(chan struct{}, 1),
		}
		finished := m.doRun()
		r.False(finished)
	})
	t.Run("FilterLogsEmpty", func(t *testing.T) {
		p = p.ApplyMethodReturn(&ethclient.Client{}, "FilterLogs", nil, nil)

		m := &Monitor{
			latest: 1,
			stop:   make(chan struct{}, 1),
		}
		finished := m.doRun()
		r.False(finished)
	})
	t.Run("FilterLogsEmptyResult", func(t *testing.T) {
		p = p.ApplyMethodReturn(&ethclient.Client{}, "FilterLogs", []types.Log{{Topics: []common.Hash{crypto.Keccak256Hash([]byte("0"))}}}, nil)

		m := &Monitor{
			latest: 1,
			stop:   make(chan struct{}, 1),
		}
		finished := m.doRun()
		r.False(finished)
	})
}

func TestMonitor_run(t *testing.T) {
	p := gomonkey.NewPatches()
	defer p.Reset()

	t.Run("Success", func(t *testing.T) {
		p = p.ApplyPrivateMethod(&Monitor{}, "doRun", func() bool { return true })

		m := &Monitor{}
		m.run()
	})
}
