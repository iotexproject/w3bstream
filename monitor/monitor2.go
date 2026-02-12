package monitor

import (
	"context"
	"log/slog"
	"math/big"
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/iotexproject/w3bstream/monitor/event"
)

type (
	Monitor struct {
		eventMap  map[eventID]event.EventInterface
		addresses []common.Address
		topics    []common.Hash
		client    *ethclient.Client

		startHeight int64
		cancel      context.CancelFunc
		wg          sync.WaitGroup
	}

	eventID [52]byte // address + topic
)

const (
	//TODO
	_scanInterval = 500
)

func NewMonitor(events []event.EventInterface, client *ethclient.Client) *Monitor {
	m := &Monitor{
		client:      client,
		startHeight: -1,
	}

	m.eventMap = make(map[eventID]event.EventInterface)
	addressMap := make(map[common.Address]struct{})
	topicMap := make(map[common.Hash]struct{})
	for _, e := range events {
		addressMap[e.Contract()] = struct{}{}
		topicMap[e.Topic()] = struct{}{}
		m.eventMap[m.eventID(e.Contract(), e.Topic())] = e
	}

	m.addresses = make([]common.Address, 0, len(addressMap))
	for a := range addressMap {
		m.addresses = append(m.addresses, a)
	}
	m.topics = make([]common.Hash, 0, len(topicMap))
	for t := range topicMap {
		m.topics = append(m.topics, t)
	}

	return m
}

func (m *Monitor) SetStartHeight(height int64) {
	m.startHeight = height
}

func (m *Monitor) Start() error {
	header, err := m.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return err
	}
	if m.startHeight < 0 {
		m.startHeight = header.Number.Int64()
	}

	// Sync events from startHeight to the latest block
	if m.startHeight < header.Number.Int64() {
		slog.Info("scanning events", "from", m.startHeight, "to", header.Number.Int64())

		var (
			minHeight = uint64(m.startHeight)
			maxHeight = header.Number.Uint64()
			length    = uint64(_scanInterval)
		)
		for from := minHeight; from <= maxHeight; from += length {
			to := from + length - 1
			if to > maxHeight {
				to = maxHeight
			}

			m.scanEvents(from, to)

			if to == maxHeight {
				break
			}
		}
		slog.Info("contract data synchronization completed", "current_height", maxHeight)
	}

	ctx, cancel := context.WithCancel(context.Background())
	m.cancel = cancel

	go m.monitor(ctx)

	return nil
}

func (m *Monitor) Stop() error {
	m.cancel()
	m.wg.Wait()
	return nil
}

func (m *Monitor) eventID(contract common.Address, topic common.Hash) eventID {
	b := make([]byte, 0, len(contract.Bytes())+len(topic.Bytes()))
	b = append(b, contract.Bytes()...)
	b = append(b, topic.Bytes()...)
	var id eventID
	copy(id[:], b)
	return id
}

func (m *Monitor) monitor(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second) // TODO:
	m.wg.Add(1)
	defer m.wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			header, err := m.client.HeaderByNumber(context.Background(), nil)
			if err != nil {
				slog.Error("failed to retrieve latest block header", "err", err)
				continue
			}
			targetHeight := min(header.Number.Uint64(), uint64(m.startHeight)+_scanInterval)

			if err := m.scanEvents(uint64(m.startHeight), targetHeight); err != nil {
				slog.Error("failed to scan events", "err", err)
				continue
			}

			m.startHeight = int64(targetHeight)
		}
	}

}

func (m *Monitor) scanEvents(from uint64, to uint64) error {
	query := ethereum.FilterQuery{
		Addresses: m.addresses,
		Topics:    [][]common.Hash{m.topics},
		FromBlock: new(big.Int).SetUint64(from),
		ToBlock:   new(big.Int).SetUint64(to),
	}
	logs, err := m.client.FilterLogs(context.Background(), query)
	if err != nil {
		return err
	}
	sort.Slice(logs, func(i, j int) bool {
		if logs[i].BlockNumber != logs[j].BlockNumber {
			return logs[i].BlockNumber < logs[j].BlockNumber
		}
		return logs[i].TxIndex < logs[j].TxIndex
	})

	for _, log := range logs {
		e := m.eventMap[m.eventID(log.Address, log.Topics[0])]
		if e == nil {
			continue
		}
		if err := e.HandleEvent(log); err != nil {
			return err
		}
	}

	return nil
}
