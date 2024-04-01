package task

import (
	"bytes"
	"log/slog"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/project/contracts"
	"github.com/machinefi/sprout/utils/hash"
)

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
