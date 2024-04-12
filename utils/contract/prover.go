package contract

import (
	"context"
	"log/slog"
	"math/big"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/smartcontracts/go/prover"
)

var (
	operatorSetTopicHash     = crypto.Keccak256Hash([]byte("OperatorSet(uint256,address)"))
	nodeTypeUpdatedTopicHash = crypto.Keccak256Hash([]byte("NodeTypeUpdated(uint256,uint256)"))
	proverPausedTopicHash    = crypto.Keccak256Hash([]byte("ProverPaused(uint256)"))
	proverResumedTopicHash   = crypto.Keccak256Hash([]byte("ProverResumed(uint256)"))
)

type Prover struct {
	ID              uint64
	OperatorAddress string
	BlockNumber     uint64
	Paused          *bool
	NodeTypes       uint64
}

func (p *Prover) Merge(diff *Prover) {
	if diff.ID != 0 {
		p.ID = diff.ID
	}
	if diff.OperatorAddress != "" {
		p.OperatorAddress = diff.OperatorAddress
	}
	if diff.BlockNumber != 0 {
		p.BlockNumber = diff.BlockNumber
	}
	if diff.Paused != nil {
		paused := *diff.Paused
		p.Paused = &paused
	}
	if diff.NodeTypes != 0 {
		p.NodeTypes = diff.NodeTypes
	}
}

func ListAndWatchProver(chainEndpoint, contractAddress string) (<-chan *Prover, error) {
	ch := make(chan *Prover, 10)
	client, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial chain endpoint")
	}

	instance, err := prover.NewProver(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to new prover contract instance")
	}

	latestBlockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to query the latest block number")
	}
	watchProver(ch, client, instance, 3*time.Second, contractAddress,
		[]common.Hash{operatorSetTopicHash, nodeTypeUpdatedTopicHash, proverResumedTopicHash}, 1000, latestBlockNumber)
	if err := listProver(ch, instance, latestBlockNumber); err != nil {
		return nil, err
	}
	return ch, nil
}

func listProver(ch chan<- *Prover, instance *prover.Prover, targetBlockNumber uint64) error {
	for id := uint64(1); ; id++ {
		mp, err := instance.Operator(&bind.CallOpts{}, new(big.Int).SetUint64(id))
		if err != nil {
			return errors.Wrapf(err, "failed to get operator from chain, prover_id %v", id)
		}
		if mp.String() == "" {
			return nil
		}

		isPaused, err := instance.IsPaused(&bind.CallOpts{}, new(big.Int).SetUint64(id))
		if err != nil {
			return errors.Wrapf(err, "failed to get prover pause status from chain, prover_id %v", id)
		}
		nodeTypes, err := instance.NodeType(&bind.CallOpts{}, new(big.Int).SetUint64(id))
		if err != nil {
			return errors.Wrapf(err, "failed to get prover nodeTypes from chain, prover_id %v", id)
		}

		ch <- &Prover{
			ID:              id,
			OperatorAddress: mp.String(),
			BlockNumber:     targetBlockNumber,
			Paused:          &isPaused,
			NodeTypes:       nodeTypes.Uint64(),
		}
	}
}

func watchProver(ch chan<- *Prover, client *ethclient.Client, instance *prover.Prover, interval time.Duration,
	contractAddress string, topics []common.Hash, step, startBlockNumber uint64) {
	queriedBlockNumber := startBlockNumber
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(contractAddress)},
		Topics: [][]common.Hash{{
			(topics[0]),
			(topics[1]),
			(topics[2]),
			(topics[3]),
		}},
	}
	ticker := time.NewTicker(interval)

	operatorSetMap := make(map[uint64]map[uint64]*prover.ProverOperatorSet) // proverID -> (blockID -> *prover.ProverOperatorSet)
	operatorSetIndex := uint(0)
	nodeTypeUpdatedMap := make(map[uint64]map[uint64]*prover.ProverNodeTypeUpdated) // proverID -> (blockID -> *prover.ProverOperatorSet)
	nodeTypeUpdatedIndex := uint(0)
	proverPausedMap := make(map[uint64]map[uint64]*prover.ProverProverPaused) // proverID -> (blockID -> *prover.ProverOperatorSet)
	proverPausedIndex := uint(0)
	proverResumedMap := make(map[uint64]map[uint64]*prover.ProverProverResumed) // proverID -> (blockID -> *prover.ProverProverResumed)
	proverResumedIndex := uint(0)

	go func() {
		for range ticker.C {
			from := queriedBlockNumber + 1
			to := from + step

			latestBlockNumber, err := client.BlockNumber(context.Background())
			if err != nil {
				slog.Error("failed to query the latest block number", "error", err)
				continue
			}
			if to > latestBlockNumber {
				to = latestBlockNumber
			}
			if from > to {
				continue
			}
			query.FromBlock = new(big.Int).SetUint64(from)
			query.ToBlock = new(big.Int).SetUint64(to)
			logs, err := client.FilterLogs(context.Background(), query)
			if err != nil {
				slog.Error("failed to filter contract logs", "error", err)
				continue
			}
			for i, evLog := range logs {
				eventSignature := evLog.Topics[0].Hex()
				switch eventSignature {
				case operatorSetTopicHash.Hex():
					ev, err := instance.ParseOperatorSet(logs[i])
					if err != nil {
						slog.Error("failed to parse prover operator set event", "error", err)
						continue
					}
					if evLog.TxIndex >= operatorSetIndex {
						operatorSetIndex = evLog.TxIndex
						operatorSetMap[ev.Id.Uint64()] = map[uint64]*prover.ProverOperatorSet{evLog.BlockNumber: ev}
					}
				case nodeTypeUpdatedTopicHash.Hex():
					ev, err := instance.ParseNodeTypeUpdated(logs[i])
					if err != nil {
						slog.Error("failed to parse prover nodeType updated event", "error", err)
						continue
					}
					if evLog.TxIndex >= nodeTypeUpdatedIndex {
						nodeTypeUpdatedIndex = evLog.TxIndex
						nodeTypeUpdatedMap[ev.Id.Uint64()] = map[uint64]*prover.ProverNodeTypeUpdated{evLog.BlockNumber: ev}
					}
				case proverPausedTopicHash.Hex():
					ev, err := instance.ParseProverPaused(logs[i])
					if err != nil {
						slog.Error("failed to parse prover paused event", "error", err)
						continue
					}
					if evLog.TxIndex >= proverPausedIndex {
						proverPausedIndex = evLog.TxIndex
						proverPausedMap[ev.Id.Uint64()] = map[uint64]*prover.ProverProverPaused{evLog.BlockNumber: ev}
					}
				case proverResumedTopicHash.Hex():
					ev, err := instance.ParseProverResumed(logs[i])
					if err != nil {
						slog.Error("failed to parse prover resumed event", "error", err)
						continue
					}
					if evLog.TxIndex >= proverResumedIndex {
						proverResumedIndex = evLog.TxIndex
						proverResumedMap[ev.Id.Uint64()] = map[uint64]*prover.ProverProverResumed{evLog.BlockNumber: ev}
					}
				default:
					slog.Error("not support parse event", "event", eventSignature)
				}
			}
			for _, blockMap := range operatorSetMap {
				blockIds := make([]uint64, 0, len(blockMap))
				for k := range blockMap {
					blockIds = append(blockIds, k)
				}
				sort.Slice(blockIds, func(i, j int) bool { return blockIds[i] < blockIds[j] })
				for _, blockId := range blockIds {
					flag := true
					ch <- &Prover{
						ID:              blockMap[blockId].Id.Uint64(),
						OperatorAddress: blockMap[blockId].Operator.String(),
						BlockNumber:     blockId,
						Paused:          &flag,
						NodeTypes:       0,
					}
				}
			}
			for _, blockMap := range nodeTypeUpdatedMap {
				blockIds := make([]uint64, 0, len(blockMap))
				for k := range blockMap {
					blockIds = append(blockIds, k)
				}
				sort.Slice(blockIds, func(i, j int) bool { return blockIds[i] < blockIds[j] })
				for _, blockId := range blockIds {
					ch <- &Prover{
						ID:          blockMap[blockId].Id.Uint64(),
						BlockNumber: blockId,
						NodeTypes:   blockMap[blockId].Typ.Uint64(),
					}
				}
			}
			for _, blockMap := range proverPausedMap {
				blockIds := make([]uint64, 0, len(blockMap))
				for k := range blockMap {
					blockIds = append(blockIds, k)
				}
				sort.Slice(blockIds, func(i, j int) bool { return blockIds[i] < blockIds[j] })
				for _, blockId := range blockIds {
					flag := true
					ch <- &Prover{
						ID:          blockMap[blockId].Id.Uint64(),
						BlockNumber: blockId,
						Paused:      &flag,
					}
				}
			}
			for _, blockMap := range proverResumedMap {
				blockIds := make([]uint64, 0, len(blockMap))
				for k := range blockMap {
					blockIds = append(blockIds, k)
				}
				sort.Slice(blockIds, func(i, j int) bool { return blockIds[i] < blockIds[j] })
				for _, blockId := range blockIds {
					flag := false
					ch <- &Prover{
						ID:          blockMap[blockId].Id.Uint64(),
						BlockNumber: blockId,
						Paused:      &flag,
					}
				}
			}
			queriedBlockNumber = to
		}
	}()
}
