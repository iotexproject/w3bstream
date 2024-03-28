package contract_monitor

import (
	"context"
	"errors"
	"testing"
	"time"

	. "github.com/agiledragon/gomonkey/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultContractMonitor(t *testing.T) {
	r := require.New(t)

	t.Run("DailChainEndpoint", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(ethclient.Dial, nil, errors.New(t.Name()))
		monitor, err := NewDefaultContractMonitor("any", "any", "any")

		r.Nil(monitor)
		r.ErrorContains(err, t.Name())
	})

	t.Run("QueryLatestBlockNumber", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(ethclient.Dial, &ethclient.Client{}, nil)
		p = p.ApplyMethodReturn(&ethclient.Client{}, "Close")
		p = p.ApplyMethodReturn(&ethclient.Client{}, "BlockNumber", uint64(0), errors.New(t.Name()))

		monitor, err := NewDefaultContractMonitor("any", "any", "any")

		r.Nil(monitor)
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(ethclient.Dial, &ethclient.Client{}, nil)
		p = p.ApplyMethodReturn(&ethclient.Client{}, "Close")
		p = p.ApplyMethodReturn(&ethclient.Client{}, "BlockNumber", uint64(1000), nil)
		p = p.ApplyFuncReturn(NewContractMonitor, &ContractMonitor{}, nil)

		monitor, err := NewDefaultContractMonitor("any", "any", "any")

		r.NotNil(monitor)
		r.NoError(err)
	})
}

func TestNewContractMonitor(t *testing.T) {
	r := require.New(t)

	t.Run("DailChainEndpoint", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(ethclient.Dial, nil, errors.New(t.Name()))
		monitor, err := NewContractMonitor("any", "any", "any", 0, 10, time.Second)

		r.Nil(monitor)
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(ethclient.Dial, &ethclient.Client{}, nil)
		monitor, err := NewContractMonitor("any", "any", "any", 0, 10, time.Second)

		r.NotNil(monitor)
		r.NoError(err)
	})
}

func TestContractMonitor_query(t *testing.T) {
	r := require.New(t)

	p := NewPatches()
	defer p.Reset()

	p = p.ApplyFuncReturn(ethclient.Dial, &ethclient.Client{}, nil)

	monitor, err := NewContractMonitor("any", "any", "any", 100, 10, time.Second)
	r.NoError(err)

	t.Run("QueryLatestBlockNumber", func(t *testing.T) {
		p = p.ApplyMethodReturn(&ethclient.Client{}, "BlockNumber", uint64(0), errors.New(t.Name()))

		queried, err := monitor.query(context.Background())
		r.Equal(queried, 0)
		r.ErrorContains(err, t.Name())
	})

	t.Run("CheckIfNeedFilterLogs", func(t *testing.T) {
		p = p.ApplyMethodReturn(&ethclient.Client{}, "BlockNumber", uint64(100), nil)

		queried, err := monitor.query(context.Background())
		r.Equal(queried, 0)
		r.NoError(err)
	})

	t.Run("FilterLogs", func(t *testing.T) {
		p = p.ApplyMethodReturn(&ethclient.Client{}, "BlockNumber", uint64(200), nil)
		p = p.ApplyMethodReturn(&ethclient.Client{}, "FilterLogs", nil, errors.New(t.Name()))

		queried, err := monitor.query(context.Background())
		r.Equal(queried, 0)
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		logs := []types.Log{{}, {}}

		t.Run("InStepRange", func(t *testing.T) {
			latest := monitor.latest + monitor.step - 1

			p = p.ApplyMethodReturn(&ethclient.Client{}, "BlockNumber", latest, nil)
			p = p.ApplyMethodReturn(&ethclient.Client{}, "FilterLogs", logs, nil)

			queried, err := monitor.query(context.Background())
			r.Equal(queried, len(logs))
			r.NoError(err)
			r.Equal(monitor.latest, latest+1)
		})

		t.Run("OutStepRange", func(t *testing.T) {
			current := monitor.latest
			latest := monitor.latest + monitor.step + 1

			p = p.ApplyMethodReturn(&ethclient.Client{}, "BlockNumber", latest, nil)
			p = p.ApplyMethodReturn(&ethclient.Client{}, "FilterLogs", logs, nil)

			queried, err := monitor.query(context.Background())
			r.Equal(queried, len(logs))
			r.NoError(err)
			r.Equal(monitor.latest, current+monitor.step+1)
		})
	})
}

func TestContractMonitor_Start(t *testing.T) {
	r := require.New(t)

	t.Run("LoopQueryLogs", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		times := 0

		p = p.ApplyPrivateMethod(&ContractMonitor{}, "query",
			func(monitor *ContractMonitor, ctx context.Context) (int, error) {
				time.Sleep(200 * time.Millisecond)
				times++
				if times%3 == 1 {
					return 0, errors.New("any")
				} else if times%3 == 2 {
					return 0, nil
				} else {
					return 10, nil
				}
			},
		)
		p = p.ApplyMethodReturn(&ethclient.Client{}, "Close")

		m := &ContractMonitor{
			client: &ethclient.Client{},
			events: make(chan *types.Log),
			topics: [][]common.Hash{{crypto.Keccak256Hash([]byte("ProjectUpserted(uint64,string,bytes32)"))}},
			stop:   make(chan struct{}),
			err:    make(chan error),
		}

		go m.Start()

		time.Sleep(2 * time.Second)
		m.Unsubscribe()
		time.Sleep(time.Second)
		r.ErrorContains(<-m.Err(), "monitoring stopped")
	})
}
