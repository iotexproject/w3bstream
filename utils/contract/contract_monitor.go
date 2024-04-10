package contract

import (
	"bytes"
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

	"github.com/machinefi/sprout/smartcontracts/go/project"
	"github.com/machinefi/sprout/smartcontracts/go/prover"
)

const (
	RequiredProverAmount = "RequiredProverAmount"
	VmType               = "VmType"

	attributeSetTopic         = "AttributeSet(uint256,bytes32,bytes)"
	projectPausedTopic        = "ProjectPaused(uint256)"
	projectResumedTopic       = "ProjectResumed(uint256)"
	projectConfigUpdatedTopic = "ProjectConfigUpdated(uint256,string,bytes32)"

	operatorSetTopic     = "OperatorSet(uint256,address)"
	nodeTypeUpdatedTopic = "NodeTypeUpdated(uint256,uint256)"
	proverPausedTopic    = "ProverPaused(uint256)"
	proverResumedTopic   = "ProverResumed(uint256)"
)

type Project struct {
	ID          uint64
	BlockNumber uint64
	Paused      *bool
	Uri         string
	Hash        [32]byte
	Attributes  map[[32]byte][]byte
}

type Prover struct {
	ID              uint64
	OperatorAddress string
	BlockNumber     uint64
	Paused          *bool
	NodeTypes       uint64
}

func ListAndWatchProject(chainEndpoint, contractAddress string) (<-chan *Project, error) {
	ch := make(chan *Project, 10)
	client, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial chain endpoint")
	}

	instance, err := project.NewProject(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to new project contract instance")
	}

	latestBlockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to query the latest block number")
	}
	watchProject(ch, client, instance, 3*time.Second, contractAddress,
		[]string{attributeSetTopic, projectPausedTopic, projectResumedTopic, projectConfigUpdatedTopic}, 1000, latestBlockNumber)
	if err := listProject(ch, instance, latestBlockNumber); err != nil {
		return nil, err
	}
	return ch, nil
}

func listProject(ch chan<- *Project, instance *project.Project, targetBlockNumber uint64) error {
	emptyHash := [32]byte{}
	for projectID := uint64(1); ; projectID++ {
		mp, err := instance.Config(&bind.CallOpts{}, new(big.Int).SetUint64(projectID))
		if err != nil {
			return errors.Wrapf(err, "failed to get project meta from chain, project_id %v", projectID)
		}
		if mp.Uri == "" || bytes.Equal(mp.Hash[:], emptyHash[:]) {
			return nil
		}

		isPaused, err := instance.IsPaused(&bind.CallOpts{}, new(big.Int).SetUint64(projectID))
		if err != nil {
			return errors.Wrapf(err, "failed to get project pause status from chain, project_id %v", projectID)
		}

		proverAmt, err := instance.Attributes(&bind.CallOpts{}, new(big.Int).SetUint64(projectID),
			crypto.Keccak256Hash([]byte(RequiredProverAmount)))
		if err != nil {
			return errors.Wrapf(err, "failed to get project attributes from chain, project_id %v, key %s", projectID, RequiredProverAmount)
		}
		vmType, err := instance.Attributes(&bind.CallOpts{}, new(big.Int).SetUint64(projectID),
			crypto.Keccak256Hash([]byte(VmType)))
		if err != nil {
			return errors.Wrapf(err, "failed to get project attributes from chain, project_id %v, key %s", projectID, vmType)
		}
		attributes := make(map[[32]byte][]byte)
		attributes[crypto.Keccak256Hash([]byte(RequiredProverAmount))] = proverAmt
		attributes[crypto.Keccak256Hash([]byte(VmType))] = vmType

		ch <- &Project{
			ID:          projectID,
			BlockNumber: targetBlockNumber,
			Paused:      &isPaused,
			Uri:         mp.Uri,
			Hash:        mp.Hash,
			Attributes:  attributes,
		}
	}
}

func watchProject(ch chan<- *Project, client *ethclient.Client, instance *project.Project, interval time.Duration,
	contractAddress string, topics []string, step, startBlockNumber uint64) {
	queriedBlockNumber := startBlockNumber
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(contractAddress)},
		Topics: [][]common.Hash{{
			crypto.Keccak256Hash([]byte(topics[0])),
			crypto.Keccak256Hash([]byte(topics[1])),
			crypto.Keccak256Hash([]byte(topics[2])),
			crypto.Keccak256Hash([]byte(topics[3])),
		}},
	}
	ticker := time.NewTicker(interval)

	attributeSetMap := make(map[uint64]map[uint64]map[uint]*project.ProjectAttributeSet) // projectID -> (blockID -> (txIndex -> *project.ProjectAttributeSet))

	projectPausedMap := make(map[uint64]map[uint64]*project.ProjectProjectPaused) // projectID -> (blockID -> *project.ProjectProjectPaused)
	projectPausedIndex := uint(0)
	projectResumedMap := make(map[uint64]map[uint64]*project.ProjectProjectResumed) // projectID -> (blockID -> *project.ProjectProjectResumed)
	projectResumedIndex := uint(0)
	projectConfigUpdatedMap := make(map[uint64]map[uint64]*project.ProjectProjectConfigUpdated) // projectID -> (blockID -> *project.ProjectProjectConfigUpdated)
	projectConfigUpdatedIndex := uint(0)

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
			for _, evLog := range logs {
				eventSignature := evLog.Topics[0].Hex()
				switch eventSignature {
				case crypto.Keccak256Hash([]byte(attributeSetTopic)).Hex():
					ev, err := instance.ParseAttributeSet(evLog)
					if err != nil {
						slog.Error("failed to parse project attribute set event", "error", err)
						continue
					}
					attributeSetMap[ev.ProjectId.Uint64()] = map[uint64]map[uint]*project.ProjectAttributeSet{
						evLog.BlockNumber: {evLog.TxIndex: ev},
					}
				case crypto.Keccak256Hash([]byte(projectPausedTopic)).Hex():
					ev, err := instance.ParseProjectPaused(evLog)
					if err != nil {
						slog.Error("failed to parse project paused event", "error", err)
						continue
					}
					if evLog.TxIndex >= projectPausedIndex {
						projectPausedIndex = evLog.TxIndex
						projectPausedMap[ev.ProjectId.Uint64()] = map[uint64]*project.ProjectProjectPaused{evLog.BlockNumber: ev}
					}
				case crypto.Keccak256Hash([]byte(projectResumedTopic)).Hex():
					ev, err := instance.ParseProjectResumed(evLog)
					if err != nil {
						slog.Error("failed to parse project resumed event", "error", err)
						continue
					}
					if evLog.TxIndex >= projectResumedIndex {
						projectResumedIndex = evLog.TxIndex
						projectResumedMap[ev.ProjectId.Uint64()] = map[uint64]*project.ProjectProjectResumed{evLog.BlockNumber: ev}
					}
				case crypto.Keccak256Hash([]byte(projectConfigUpdatedTopic)).Hex():
					ev, err := instance.ParseProjectConfigUpdated(evLog)
					if err != nil {
						slog.Error("failed to parse project config updated event", "error", err)
						continue
					}
					if evLog.TxIndex >= projectConfigUpdatedIndex {
						projectConfigUpdatedIndex = evLog.TxIndex
						projectConfigUpdatedMap[ev.ProjectId.Uint64()] = map[uint64]*project.ProjectProjectConfigUpdated{evLog.BlockNumber: ev}
					}
				default:
					slog.Error("not support parse event", "event", eventSignature)
				}
			}
			for projectId, blockMap := range attributeSetMap {
				blockIds := make([]uint64, 0, len(blockMap))
				for k := range blockMap {
					blockIds = append(blockIds, k)
				}
				sort.Slice(blockIds, func(i, j int) bool { return blockIds[i] < blockIds[j] })
				for _, blockId := range blockIds {
					txIndexMap := blockMap[blockId]
					txIndexs := make([]uint, 0, len(txIndexMap))
					for k := range txIndexMap {
						txIndexs = append(txIndexs, k)
					}
					sort.Slice(txIndexs, func(i, j int) bool { return txIndexs[i] < txIndexs[j] })
					attributes := make(map[[32]byte][]byte)
					for _, txIndex := range txIndexs {
						attributes[txIndexMap[txIndex].Key] = txIndexMap[txIndex].Value
					}
					ch <- &Project{
						ID:          projectId,
						BlockNumber: blockId,
						Attributes:  attributes,
					}
				}
			}
			for _, blockMap := range projectPausedMap {
				blockIds := make([]uint64, 0, len(blockMap))
				for k := range blockMap {
					blockIds = append(blockIds, k)
				}
				sort.Slice(blockIds, func(i, j int) bool { return blockIds[i] < blockIds[j] })
				for _, blockId := range blockIds {
					flag := true
					ch <- &Project{
						ID:          blockMap[blockId].ProjectId.Uint64(),
						BlockNumber: blockId,
						Paused:      &flag,
					}
				}
			}
			for _, blockMap := range projectResumedMap {
				blockIds := make([]uint64, 0, len(blockMap))
				for k := range blockMap {
					blockIds = append(blockIds, k)
				}
				sort.Slice(blockIds, func(i, j int) bool { return blockIds[i] < blockIds[j] })
				for _, blockId := range blockIds {
					flag := false
					ch <- &Project{
						ID:          blockMap[blockId].ProjectId.Uint64(),
						BlockNumber: blockId,
						Paused:      &flag,
					}
				}
			}
			for _, blockMap := range projectConfigUpdatedMap {
				blockIds := make([]uint64, 0, len(blockMap))
				for k := range blockMap {
					blockIds = append(blockIds, k)
				}
				sort.Slice(blockIds, func(i, j int) bool { return blockIds[i] < blockIds[j] })
				for _, blockId := range blockIds {
					flag := true
					ch <- &Project{
						ID:          blockMap[blockId].ProjectId.Uint64(),
						BlockNumber: blockId,
						Paused:      &flag,
						Uri:         blockMap[blockId].Uri,
						Hash:        blockMap[blockId].Hash,
					}
				}
			}
			queriedBlockNumber = to
		}
	}()
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
		[]string{operatorSetTopic, nodeTypeUpdatedTopic, projectPausedTopic, proverResumedTopic}, 1000, latestBlockNumber)
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
	contractAddress string, topics []string, step, startBlockNumber uint64) {
	queriedBlockNumber := startBlockNumber
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(contractAddress)},
		Topics: [][]common.Hash{{
			crypto.Keccak256Hash([]byte(topics[0])),
			crypto.Keccak256Hash([]byte(topics[1])),
			crypto.Keccak256Hash([]byte(topics[2])),
			crypto.Keccak256Hash([]byte(topics[3])),
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
				case crypto.Keccak256Hash([]byte(operatorSetTopic)).Hex():
					ev, err := instance.ParseOperatorSet(logs[i])
					if err != nil {
						slog.Error("failed to parse prover operator set event", "error", err)
						continue
					}
					if evLog.TxIndex >= operatorSetIndex {
						operatorSetIndex = evLog.TxIndex
						operatorSetMap[ev.Id.Uint64()] = map[uint64]*prover.ProverOperatorSet{evLog.BlockNumber: ev}
					}
				case crypto.Keccak256Hash([]byte(nodeTypeUpdatedTopic)).Hex():
					ev, err := instance.ParseNodeTypeUpdated(logs[i])
					if err != nil {
						slog.Error("failed to parse prover nodeType updated event", "error", err)
						continue
					}
					if evLog.TxIndex >= nodeTypeUpdatedIndex {
						nodeTypeUpdatedIndex = evLog.TxIndex
						nodeTypeUpdatedMap[ev.Id.Uint64()] = map[uint64]*prover.ProverNodeTypeUpdated{evLog.BlockNumber: ev}
					}
				case crypto.Keccak256Hash([]byte(proverPausedTopic)).Hex():
					ev, err := instance.ParseProverPaused(logs[i])
					if err != nil {
						slog.Error("failed to parse prover paused event", "error", err)
						continue
					}
					if evLog.TxIndex >= proverPausedIndex {
						proverPausedIndex = evLog.TxIndex
						proverPausedMap[ev.Id.Uint64()] = map[uint64]*prover.ProverProverPaused{evLog.BlockNumber: ev}
					}
				case crypto.Keccak256Hash([]byte(proverResumedTopic)).Hex():
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
