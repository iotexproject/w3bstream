package project

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
	"github.com/ethereum/go-ethereum/event"
	"github.com/pkg/errors"
)

type monitor struct {
	client    *ethclient.Client
	latest    int64
	step      int64
	interval  time.Duration
	addresses []common.Address
	topics    [][]common.Hash
	events    map[common.Hash]chan *types.Log
	errs      chan error
	stop      chan struct{}
}

var _ event.Subscription = (*monitor)(nil)

func (m *monitor) Err() <-chan error {
	return m.errs
}

func (m *monitor) Unsubscribe() {
	m.stop <- struct{}{}
}

func (m *monitor) Events(topic string) (<-chan *types.Log, bool) {
	ch, ok := m.events[crypto.Keccak256Hash([]byte(topic))]
	return ch, ok
}

func (m *monitor) mustEvents(topic string) <-chan *types.Log {
	ch, ok := m.Events(topic)
	if !ok {
		panic("event not subscribed " + topic)
	}
	return ch
}

func (m *monitor) doRun() bool {
	query := ethereum.FilterQuery{
		Addresses: m.addresses,
		Topics:    m.topics,
	}
	select {
	case <-m.stop:
		slog.Info("monitor stopped")
		return true
	default:
		latestBlk, err := m.client.BlockNumber(context.Background())
		if err != nil {
			slog.Error("query latest block number", "msg", err)
			time.Sleep(m.interval)
			return false
		}
		if uint64(m.latest) > latestBlk {
			time.Sleep(m.interval)
			return false
		}
		query.FromBlock = big.NewInt(m.latest)
		query.ToBlock = big.NewInt(min(m.latest+m.step, int64(latestBlk)))

		logs, err := m.client.FilterLogs(context.Background(), query)
		if err != nil {
			slog.Error("failed to filter logs", "msg", err)
			time.Sleep(m.interval)
			return false
		}
		m.latest = query.ToBlock.Int64()
		if len(logs) == 0 {
			goto TryLater
		}
		slog.Info("filter logs", "count", len(logs))
		for _, l := range logs {
			topic := l.Topics[0]
			if _, ok := m.events[topic]; !ok {
				continue
			}
			m.events[topic] <- &l
		}
	TryLater:
		if query.ToBlock.Int64()-query.FromBlock.Int64() < m.step {
			time.Sleep(m.interval)
		}
	}
	return false
}

func (m *monitor) run() {
	for {
		if finished := m.doRun(); finished {
			return
		}
	}
}

func newMonitor(chainEndpoint string, addresses []string, topics []string, from, step int64, interval time.Duration) (*monitor, error) {
	client, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to dial chain endpoint %s", chainEndpoint)
	}

	m := &monitor{
		client:   client,
		latest:   from,
		step:     step,
		interval: interval,
		events:   nil,
		errs:     make(chan error),
		stop:     make(chan struct{}),
	}
	m.addresses = make([]common.Address, 0, len(addresses))
	for _, addr := range addresses {
		m.addresses = append(m.addresses, common.HexToAddress(addr))
	}

	m.topics = make([][]common.Hash, 1)
	m.events = make(map[common.Hash]chan *types.Log)
	for i := range topics {
		topic := crypto.Keccak256Hash([]byte(topics[i]))
		m.topics[0] = append(m.topics[0], topic)
		m.events[topic] = make(chan *types.Log, 100)
	}

	return m, nil
}

func newDefaultMonitor(chainEndpoint string, addresses []string, topics []string) (*monitor, error) {
	client, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to dial chain endpoint %s", chainEndpoint)
	}
	defer client.Close()

	latestBlk, err := client.BlockNumber(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get current block number")
	}

	return newMonitor(chainEndpoint, addresses, topics, int64(latestBlk), 100000, time.Second*10)
}
