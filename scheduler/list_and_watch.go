package scheduler

import (
	"bytes"
	"context"
	"log/slog"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/persistence/prover"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/project/contracts"
	"github.com/machinefi/sprout/utils/hash"
)

func newProverContractInstance(chainEndpoint, proverContractAddress string) (*prover.Prover, error) {
	client, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to dial chain endpoint %s", chainEndpoint)
	}
	instance, err := prover.NewProver(common.HexToAddress(proverContractAddress), client)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to new prover contract instance")
	}
	return instance, nil
}

func listAndFillProver(provers *sync.Map, chainEndpoint, proverContractAddress string) error {
	instance, err := newProverContractInstance(chainEndpoint, proverContractAddress)
	if err != nil {
		return err
	}

	for i := uint64(1); ; i++ {
		prover, err := instance.Provers(nil, i)
		if err != nil {
			return errors.Wrapf(err, "failed to get prover from chain, prover_id %v", i)
		}
		if prover.Id == "" {
			break
		}
		provers.Store(prover.Id, true)
	}
	return nil
}

func watchProver(provers *sync.Map, chainEndpoint, proverContractAddress string) error {
	instance, err := newProverContractInstance(chainEndpoint, proverContractAddress)
	if err != nil {
		return err
	}

	events := make(chan *prover.ProverProverUpserted, 10)
	sub, err := instance.WatchProverUpserted(nil, events, nil)
	if err != nil {
		return errors.Wrapf(err, "failed to watch prover upserted event")
	}
	go func() {
		for {
			select {
			case err := <-sub.Err():
				slog.Error("got an error when watching prover upserted event", "error", err)
			case e := <-events:
				provers.Store(string(e.ProverID[:]), true)
			}
		}
	}()
	return nil
}

func newProjectContractInstance(chainEndpoint, projectContractAddress string) (*contracts.Contracts, error) {
	client, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to dial chain endpoint %s", chainEndpoint)
	}
	instance, err := contracts.NewContracts(common.HexToAddress(projectContractAddress), client)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to new project contract instance")
	}
	return instance, nil
}

func listAndFillProject(projectOffsets *sync.Map, epoch uint64, chainEndpoint, projectContractAddress string) error {
	instance, err := newProjectContractInstance(chainEndpoint, projectContractAddress)
	if err != nil {
		return err
	}

	emptyHash := [32]byte{}
	for projectID := uint64(1); ; projectID++ {
		mp, err := instance.Projects(nil, projectID)
		if err != nil {
			return errors.Wrapf(err, "failed to get project meta from chain, project_id %v", projectID)
		}
		if mp.Uri == "" || bytes.Equal(mp.Hash[:], emptyHash[:]) {
			return nil
		}
		setProjectOffsets(projectOffsets, epoch, &project.ProjectMeta{
			ProjectID: projectID,
			Uri:       mp.Uri,
			Hash:      mp.Hash,
		})
	}
}

func watchProject(projectOffsets *sync.Map, epoch uint64, chainEndpoint, projectContractAddress string) error {
	instance, err := newProjectContractInstance(chainEndpoint, projectContractAddress)
	if err != nil {
		return err
	}

	events := make(chan *contracts.ContractsProjectUpserted, 10)
	sub, err := instance.WatchProjectUpserted(nil, events, nil)
	if err != nil {
		return errors.Wrapf(err, "failed to watch project upserted event")
	}
	go func() {
		for {
			select {
			case err := <-sub.Err():
				slog.Error("got an error when watching project upserted event", "error", err)
			case e := <-events:
				setProjectOffsets(projectOffsets, epoch, &project.ProjectMeta{
					ProjectID: e.ProjectId,
					Uri:       e.Uri,
					Hash:      e.Hash,
				})
			}
		}
	}()
	return nil
}

func setProjectOffsets(projectOffsets *sync.Map, epoch uint64, meta *project.ProjectMeta) {
	projectIDHash := hash.Sum256Uint64(meta.ProjectID)
	offset := new(big.Int).SetBytes(projectIDHash[:]).Uint64() % epoch
	projectOffsets.Store(offset, meta)
}

func watchChainHead(head chan<- *types.Header, chainEndpoint string) error {
	client, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		return errors.Wrapf(err, "failed to dial chain endpoint %s", chainEndpoint)
	}
	sub, err := client.SubscribeNewHead(context.Background(), head)
	if err != nil {
		return errors.Wrap(err, "failed to watch chain head")
	}
	go func() {
		for err := range sub.Err() {
			slog.Error("got an error when watching chain head", "error", err)
		}
	}()
	return nil
}
