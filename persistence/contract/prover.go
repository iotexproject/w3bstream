package contract

import (
	"bytes"
	"container/list"
	"context"
	"log/slog"
	"math"
	"math/big"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/smartcontracts/go/blocknumber"
	"github.com/machinefi/sprout/smartcontracts/go/multicall"
	"github.com/machinefi/sprout/smartcontracts/go/prover"
)

var (
	operatorSetTopicHash     = crypto.Keccak256Hash([]byte("OperatorSet(uint256,address)"))
	nodeTypeUpdatedTopicHash = crypto.Keccak256Hash([]byte("NodeTypeUpdated(uint256,uint256)"))
	proverPausedTopicHash    = crypto.Keccak256Hash([]byte("ProverPaused(uint256)"))
	proverResumedTopicHash   = crypto.Keccak256Hash([]byte("ProverResumed(uint256)"))

	emptyAddress = common.Address{}
)

type BlockProver struct {
	BlockNumber uint64
	Provers     map[uint64]*Prover
}

type Prover struct {
	ID              uint64
	OperatorAddress common.Address
	BlockNumber     uint64
	Paused          *bool
	NodeTypes       uint64
}

func (ps *BlockProver) Merge(diff *BlockProver) {
	ps.BlockNumber = diff.BlockNumber
	for id, p := range ps.Provers {
		diffP, ok := diff.Provers[id]
		if ok {
			p.Merge(diffP)
		}
	}
	for id, p := range diff.Provers {
		if _, ok := ps.Provers[id]; !ok {
			np := &Prover{}
			np.Merge(p)
			ps.Provers[id] = np
		}
	}
}

func (p *Prover) Merge(diff *Prover) {
	if diff.ID != 0 {
		p.ID = diff.ID
	}
	if !bytes.Equal(diff.OperatorAddress[:], emptyAddress[:]) {
		p.OperatorAddress = diff.OperatorAddress
	}
	if diff.BlockNumber != 0 {
		p.BlockNumber = diff.BlockNumber
	}
	if diff.Paused != nil {
		p.Paused = diff.Paused
	}
	if diff.NodeTypes != 0 {
		p.NodeTypes = diff.NodeTypes
	}
}

type blockProvers struct {
	mu       sync.Mutex
	capacity uint64
	blocks   *list.List
}

func (c *blockProvers) provers(blockNumber uint64) *BlockProver {
	c.mu.Lock()
	defer c.mu.Unlock()

	np := &BlockProver{Provers: map[uint64]*Prover{}}

	for e := c.blocks.Front(); e != nil; e = e.Next() {
		ep := e.Value.(*BlockProver)
		if blockNumber > ep.BlockNumber {
			break
		}
		np.Merge(ep)
	}
	return np
}

func (c *blockProvers) add(diff *BlockProver) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.blocks.PushBack(diff)

	if uint64(c.blocks.Len()) > c.capacity {
		h := c.blocks.Front()
		np := &BlockProver{Provers: map[uint64]*Prover{}}
		np.Merge(h.Value.(*BlockProver))
		np.Merge(h.Next().Value.(*BlockProver))
		c.blocks.Remove(h.Next())
		c.blocks.Remove(h)
		c.blocks.PushFront(np)
	}
}

func ListAndWatchProver(chainEndpoint, contractAddress string, tracebackLength uint64) (<-chan *BlockProver, error) {
	ch := make(chan *BlockProver, 10)
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
	targetBlockNumber := latestBlockNumber - tracebackLength
	// TODO delete
	// if err := listProver(ch, instance, targetBlockNumber); err != nil {
	// 	return nil, err
	// }

	topics := []common.Hash{operatorSetTopicHash, nodeTypeUpdatedTopicHash, proverResumedTopicHash}
	watchProver(ch, client, instance, 3*time.Second, contractAddress, topics, 1000, targetBlockNumber)

	return ch, nil
}

func listProver(client *ethclient.Client, proverContractAddress, blockNumberContractAddress, multiCallContractAddress common.Address) ([]*Prover, uint64, uint64, error) {
	multiCallInstance, err := multicall.NewMulticall(multiCallContractAddress, client)
	if err != nil {
		return nil, 0, 0, errors.Wrap(err, "failed to new multi call contract instance")
	}
	blockNumberABI, err := abi.JSON(strings.NewReader(blocknumber.BlocknumberMetaData.ABI))
	if err != nil {
		return nil, 0, 0, errors.Wrap(err, "failed to decode block number contract abi")
	}
	proverABI, err := abi.JSON(strings.NewReader(prover.ProverMetaData.ABI))
	if err != nil {
		return nil, 0, 0, errors.Wrap(err, "failed to decode prover contract abi")
	}
	blockNumberCallData, err := blockNumberABI.Pack("blockNumber")
	if err != nil {
		return nil, 0, 0, errors.Wrap(err, "failed to pack block number call data")
	}
	ps := []*Prover{}
	minBlockNumber := uint64(math.MaxUint64)
	maxBlockNumber := uint64(0)
	for proverID := uint64(1); ; proverID++ {
		operatorCallData, err := proverABI.Pack("operator", new(big.Int).SetUint64(proverID))
		if err != nil {
			return nil, 0, 0, errors.Wrap(err, "failed to pack prover operator call data")
		}
		isPausedCallData, err := proverABI.Pack("isPaused", new(big.Int).SetUint64(proverID))
		if err != nil {
			return nil, 0, 0, errors.Wrap(err, "failed to pack prover is paused call data")
		}
		nodeTypeCallData, err := proverABI.Pack("nodeType", new(big.Int).SetUint64(proverID))
		if err != nil {
			return nil, 0, 0, errors.Wrap(err, "failed to pack prover node type call data")
		}

		result, err := multiCallInstance.MultiCall(
			nil,
			[]common.Address{
				blockNumberContractAddress,
				proverContractAddress,
				proverContractAddress,
				proverContractAddress,
			},
			[][]byte{
				blockNumberCallData,
				operatorCallData,
				isPausedCallData,
				nodeTypeCallData,
			},
		)
		if err != nil {
			return nil, 0, 0, errors.Wrapf(err, "failed to multi call, prover_id %v", proverID)
		}

		out, err := blockNumberABI.Unpack("blockNumber", result[0])
		if err != nil {
			return nil, 0, 0, errors.Wrapf(err, "failed to unpack block number result, prover_id %v", proverID)
		}
		preBlockNumber := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
		blockNumber := preBlockNumber.Uint64() - 1

		minBlockNumber = min(minBlockNumber, blockNumber)
		maxBlockNumber = max(maxBlockNumber, blockNumber)

		if len(result[1]) == 0 || len(result[2]) == 0 || len(result[3]) == 0 {
			break
		}

		out, err = proverABI.Unpack("operator", result[1])
		if err != nil {
			return nil, 0, 0, errors.Wrapf(err, "failed to unpack prover operator result, prover_id %v", proverID)
		}
		operator := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

		out, err = proverABI.Unpack("isPaused", result[2])
		if err != nil {
			return nil, 0, 0, errors.Wrapf(err, "failed to unpack prover is paused result, prover_id %v", proverID)
		}
		isPaused := *abi.ConvertType(out[0], new(bool)).(*bool)

		out, err = proverABI.Unpack("nodeType", result[3])
		if err != nil {
			return nil, 0, 0, errors.Wrapf(err, "failed to unpack prover node type result, prover_id %v", proverID)
		}
		nodeType := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		ps = append(ps, &Prover{
			ID:              proverID,
			BlockNumber:     blockNumber,
			Paused:          &isPaused,
			OperatorAddress: operator,
			NodeTypes:       nodeType.Uint64(),
		})
	}
	return ps, minBlockNumber, maxBlockNumber, nil
}

func listProver1(ch chan<- *BlockProver, instance *prover.Prover, targetBlockNumber uint64) error {
	provers := map[uint64]*Prover{}
	for id := uint64(1); ; id++ {
		mp, err := instance.Operator(nil, new(big.Int).SetUint64(id))
		if err != nil {
			if strings.Contains(err.Error(), "execution reverted: ERC721: invalid token ID") {
				break
			}
			return errors.Wrapf(err, "failed to get operator from chain, prover_id %v", id)
		}

		isPaused, err := instance.IsPaused(nil, new(big.Int).SetUint64(id))
		if err != nil {
			return errors.Wrapf(err, "failed to get prover pause status from chain, prover_id %v", id)
		}
		nodeTypes, err := instance.NodeType(nil, new(big.Int).SetUint64(id))
		if err != nil {
			return errors.Wrapf(err, "failed to get prover nodeTypes from chain, prover_id %v", id)
		}

		provers[id] = &Prover{
			ID:              id,
			OperatorAddress: mp,
			BlockNumber:     targetBlockNumber,
			Paused:          &isPaused,
			NodeTypes:       nodeTypes.Uint64(),
		}
	}
	ch <- &BlockProver{
		BlockNumber: targetBlockNumber,
		Provers:     provers,
	}
	return nil
}

func watchProver(ch chan<- *BlockProver, client *ethclient.Client, instance *prover.Prover, interval time.Duration,
	contractAddress string, topics []common.Hash, step, startBlockNumber uint64) {
	queriedBlockNumber := startBlockNumber
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(contractAddress)},
		Topics: [][]common.Hash{{
			(topics[0]),
			(topics[1]),
			(topics[2]),
		}},
	}
	ticker := time.NewTicker(interval)

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
			// TODO delete
			processProverLogs(nil, logs, instance)
		}
	}()
}

func processProverLogs(add func(*BlockProver), logs []types.Log, instance *prover.Prover) error {
	sort.Slice(logs, func(i, j int) bool {
		if logs[i].BlockNumber != logs[j].BlockNumber {
			return logs[i].BlockNumber < logs[j].BlockNumber
		}
		return logs[i].TxIndex < logs[j].TxIndex
	})
	psMap := map[uint64]*BlockProver{}

	for _, l := range logs {
		ps, ok := psMap[l.BlockNumber]
		if !ok {
			ps = &BlockProver{
				BlockNumber: l.BlockNumber,
				Provers:     map[uint64]*Prover{},
			}
		}
		switch l.Topics[0] {
		case operatorSetTopicHash:
			e, err := instance.ParseOperatorSet(l)
			if err != nil {
				return errors.Wrap(err, "failed to parse prover operator set event")
			}

			p, ok := ps.Provers[e.Id.Uint64()]
			if !ok {
				p = &Prover{ID: e.Id.Uint64()}
			}
			p.OperatorAddress = e.Operator
			ps.Provers[e.Id.Uint64()] = p

		case nodeTypeUpdatedTopicHash:
			e, err := instance.ParseNodeTypeUpdated(l)
			if err != nil {
				return errors.Wrap(err, "failed to parse prover node type updated event")
			}

			p, ok := ps.Provers[e.Id.Uint64()]
			if !ok {
				p = &Prover{ID: e.Id.Uint64()}
			}
			p.NodeTypes = e.Typ.Uint64()
			ps.Provers[e.Id.Uint64()] = p

		case proverPausedTopicHash:
			e, err := instance.ParseProverPaused(l)
			if err != nil {
				return errors.Wrap(err, "failed to parse prover paused event")
			}

			p, ok := ps.Provers[e.Id.Uint64()]
			if !ok {
				p = &Prover{ID: e.Id.Uint64()}
			}
			paused := true
			p.Paused = &paused
			ps.Provers[e.Id.Uint64()] = p

		case proverResumedTopicHash:
			e, err := instance.ParseProverResumed(l)
			if err != nil {
				return errors.Wrap(err, "failed to parse prover resumed event")
			}

			p, ok := ps.Provers[e.Id.Uint64()]
			if !ok {
				p = &Prover{ID: e.Id.Uint64()}
			}
			paused := false
			p.Paused = &paused
			ps.Provers[e.Id.Uint64()] = p
		}
		psMap[l.BlockNumber] = ps
	}

	psSlice := []*BlockProver{}
	for _, p := range psMap {
		psSlice = append(psSlice, p)
	}
	sort.Slice(psSlice, func(i, j int) bool {
		return psSlice[i].BlockNumber < psSlice[j].BlockNumber
	})

	for _, p := range psSlice {
		add(p)
	}
	return nil
}
