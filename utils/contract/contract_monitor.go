package contract

import (
	"bytes"
	"context"
	"log/slog"
	"math/big"
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
					ch <- &Project{
						ID:          ev.ProjectId.Uint64(),
						BlockNumber: evLog.BlockNumber,
						Attributes:  map[[32]byte][]byte{ev.Key: ev.Value},
					}
				case crypto.Keccak256Hash([]byte(projectPausedTopic)).Hex():
					ev, err := instance.ParseProjectPaused(evLog)
					if err != nil {
						slog.Error("failed to parse project paused event", "error", err)
						continue
					}
					flag := true
					ch <- &Project{
						ID:          ev.ProjectId.Uint64(),
						BlockNumber: evLog.BlockNumber,
						Paused:      &flag,
					}
				case crypto.Keccak256Hash([]byte(projectResumedTopic)).Hex():
					ev, err := instance.ParseProjectResumed(evLog)
					if err != nil {
						slog.Error("failed to parse project resumed event", "error", err)
						continue
					}
					flag := false
					ch <- &Project{
						ID:          ev.ProjectId.Uint64(),
						BlockNumber: evLog.BlockNumber,
						Paused:      &flag,
					}
				case crypto.Keccak256Hash([]byte(projectConfigUpdatedTopic)).Hex():
					ev, err := instance.ParseProjectConfigUpdated(evLog)
					if err != nil {
						slog.Error("failed to parse project config updated event", "error", err)
						continue
					}
					flag := true
					ch <- &Project{
						ID:          ev.ProjectId.Uint64(),
						BlockNumber: evLog.BlockNumber,
						Paused:      &flag,
						Uri:         ev.Uri,
						Hash:        ev.Hash,
					}
				default:
					slog.Error("not support parse event", "event", eventSignature)
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
					flag := true
					ch <- &Prover{
						OperatorAddress: ev.Operator.String(),
						BlockNumber:     logs[i].BlockNumber,
						Paused:          &flag,
						NodeTypes:       0,
					}
				case crypto.Keccak256Hash([]byte(nodeTypeUpdatedTopic)).Hex():
					ev, err := instance.ParseNodeTypeUpdated(logs[i])
					if err != nil {
						slog.Error("failed to parse prover nodeType updated event", "error", err)
						continue
					}
					ch <- &Prover{
						OperatorAddress: ev.Id.String(),
						BlockNumber:     logs[i].BlockNumber,
						NodeTypes:       ev.Typ.Uint64(),
					}
				case crypto.Keccak256Hash([]byte(proverPausedTopic)).Hex():
					ev, err := instance.ParseProverPaused(logs[i])
					if err != nil {
						slog.Error("failed to parse prover paused event", "error", err)
						continue
					}
					flag := true
					ch <- &Prover{
						OperatorAddress: ev.Id.String(),
						BlockNumber:     logs[i].BlockNumber,
						Paused:          &flag,
					}
				case crypto.Keccak256Hash([]byte(proverResumedTopic)).Hex():
					ev, err := instance.ParseProverResumed(logs[i])
					if err != nil {
						slog.Error("failed to parse prover resumed event", "error", err)
						continue
					}
					flag := false
					ch <- &Prover{
						OperatorAddress: ev.Id.String(),
						BlockNumber:     logs[i].BlockNumber,
						Paused:          &flag,
					}

				default:
					slog.Error("not support parse event", "event", eventSignature)
				}
			}
			queriedBlockNumber = to
		}
	}()
}
