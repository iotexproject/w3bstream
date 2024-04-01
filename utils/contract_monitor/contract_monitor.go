package contract_monitor

import (
	"context"
	"log/slog"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

func NewContractMonitor(endpoint, address, topic string, from, step uint64, interval time.Duration) (*ContractMonitor, error) {
	if from == 0 {
		from = 1
	}
	if step == 0 {
		step = 100000
	}

	var (
		err     error
		monitor = &ContractMonitor{
			err:      make(chan error),
			stop:     make(chan struct{}),
			latest:   from,
			step:     step,
			events:   make(chan *types.Log, 32),
			interval: interval,
		}
	)

	monitor.client, err = ethclient.Dial(endpoint)
	if err != nil {
		slog.Error("failed to dail chain endpoint", "endpoint", endpoint, "error", err)
		return nil, errors.Wrap(err, "failed to dail chain endpoint")
	}

	monitor.address = []common.Address{common.HexToAddress(address)}
	monitor.topics = [][]common.Hash{{crypto.Keccak256Hash([]byte(topic))}}

	return monitor, nil
}

func NewDefaultContractMonitor(endpoint, address, topic string) (*ContractMonitor, error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dail chain endpoint")
	}
	defer client.Close()

	latest, err := client.BlockNumber(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to query the latest block number")
	}

	return NewContractMonitor(endpoint, address, topic, latest, 100000, time.Second*10)
}

type ContractMonitor struct {
	client   *ethclient.Client
	address  []common.Address
	events   chan *types.Log
	topics   [][]common.Hash
	latest   uint64
	step     uint64
	interval time.Duration
	err      chan error
	stop     chan struct{}
}

func (m *ContractMonitor) Err() <-chan error {
	return m.err
}

func (m *ContractMonitor) Unsubscribe() {
	m.stop <- struct{}{}
}

func (m *ContractMonitor) Start() {
	slog.Info("start monitoring contract events", "topic", m.topics[0][0].String(), "address", m.address)

	ctx, cancel := context.WithCancel(context.Background())

	for {
		select {
		case <-m.stop:
			cancel()
			m.client.Close()
			close(m.events)
			m.err <- errors.New("monitoring stopped")
			return
		default:
			queried, err := m.query(ctx)
			if err != nil {
				slog.Warn("failed to query", "error", err)
			}
			if queried == -1 {
				time.Sleep(m.interval)
				continue
			}
			slog.Info("event queried", "latest_block", m.latest, "queried_events", queried)
		}
	}
}

func (m *ContractMonitor) Events() <-chan *types.Log {
	return m.events
}

func (m *ContractMonitor) query(ctx context.Context) (queried int, err error) {
	var latest uint64

	latest, err = m.client.BlockNumber(ctx)
	if err != nil {
		err = errors.Wrap(err, "failed to query latest block number")
		return
	}

	if m.latest >= latest {
		return -1, nil
	}

	var (
		query = ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(m.latest)),
			ToBlock:   big.NewInt(int64(min(m.latest+m.step, latest))),
			Addresses: m.address,
			Topics:    m.topics,
		}
		logs []types.Log
	)

	logs, err = m.client.FilterLogs(ctx, query)
	if err != nil {
		err = errors.Wrap(err, "failed to query logs")
		return
	}

	queried = len(logs)
	m.latest = query.ToBlock.Uint64() + 1

	for i := range logs {
		m.events <- &logs[i]
	}
	return
}
