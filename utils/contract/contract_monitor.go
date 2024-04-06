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

	"github.com/machinefi/sprout/persistence/prover"
	"github.com/machinefi/sprout/project/contracts"
)

const (
	ProjectUpsertedTopic = "ProjectUpserted(uint64,string,bytes32)"
	ProverUpsertedTopic  = "ProverUpserted(string)"
)

type Project struct {
	Uri         string
	Hash        [32]byte
	ID          uint64
	BlockNumber uint64
}

type Prover struct {
	ID          string
	BlockNumber uint64
}

func ListAndWatchProject(chainEndpoint, contractAddress string) (<-chan *Project, error) {
	ch := make(chan *Project, 10)
	client, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial chain endpoint")
	}

	instance, err := contracts.NewContracts(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to new project contract instance")
	}

	latestBlockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to query the latest block number")
	}
	watchProject(ch, client, instance, 3*time.Second, contractAddress, ProjectUpsertedTopic, 1000, latestBlockNumber)
	if err := listProject(ch, instance, latestBlockNumber); err != nil {
		return nil, err
	}
	return ch, nil
}

func listProject(ch chan<- *Project, instance *contracts.Contracts, targetBlockNumber uint64) error {
	emptyHash := [32]byte{}
	for projectID := uint64(1); ; projectID++ {
		mp, err := instance.Projects(&bind.CallOpts{
			BlockNumber: new(big.Int).SetUint64(targetBlockNumber),
		}, projectID)
		if err != nil {
			return errors.Wrapf(err, "failed to get project meta from chain, project_id %v", projectID)
		}
		if mp.Uri == "" || bytes.Equal(mp.Hash[:], emptyHash[:]) {
			return nil
		}
		ch <- &Project{
			Uri:         mp.Uri,
			Hash:        mp.Hash,
			ID:          projectID,
			BlockNumber: targetBlockNumber,
		}
	}
}

func watchProject(ch chan<- *Project, client *ethclient.Client, instance *contracts.Contracts, interval time.Duration, contractAddress, topic string, step, startBlockNumber uint64) {
	queriedBlockNumber := startBlockNumber
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(contractAddress)},
		Topics:    [][]common.Hash{{crypto.Keccak256Hash([]byte(topic))}},
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
			for i := range logs {
				ev, err := instance.ParseProjectUpserted(logs[i])
				if err != nil {
					slog.Error("failed to parse project upserted event", "error", err)
					continue
				}
				ch <- &Project{
					Uri:         ev.Uri,
					Hash:        ev.Hash,
					ID:          ev.ProjectId,
					BlockNumber: logs[i].BlockNumber,
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
	watchProver(ch, client, instance, 3*time.Second, contractAddress, ProverUpsertedTopic, 1000, latestBlockNumber)
	if err := listProver(ch, instance, latestBlockNumber); err != nil {
		return nil, err
	}
	return ch, nil
}

func listProver(ch chan<- *Prover, instance *prover.Prover, targetBlockNumber uint64) error {
	for id := uint64(1); ; id++ {
		mp, err := instance.Provers(&bind.CallOpts{
			BlockNumber: new(big.Int).SetUint64(targetBlockNumber),
		}, id)
		if err != nil {
			return errors.Wrapf(err, "failed to get project meta from chain, project_id %v", id)
		}
		if mp.Id == "" {
			return nil
		}
		ch <- &Prover{
			ID:          mp.Id,
			BlockNumber: targetBlockNumber,
		}
	}
}

func watchProver(ch chan<- *Prover, client *ethclient.Client, instance *prover.Prover, interval time.Duration, contractAddress, topic string, step, startBlockNumber uint64) {
	queriedBlockNumber := startBlockNumber
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(contractAddress)},
		Topics:    [][]common.Hash{{crypto.Keccak256Hash([]byte(topic))}},
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
			for i := range logs {
				ev, err := instance.ParseProverUpserted(logs[i])
				if err != nil {
					slog.Error("failed to parse prover upserted event", "error", err)
					continue
				}
				ch <- &Prover{
					ID:          ev.ProverID.String(),
					BlockNumber: logs[i].BlockNumber,
				}
			}
			queriedBlockNumber = to
		}
	}()
}
