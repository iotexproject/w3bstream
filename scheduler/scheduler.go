package scheduler

import (
	"container/list"
	"context"
	"log/slog"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/utils/contract"
	"github.com/machinefi/sprout/utils/distance"
	"github.com/machinefi/sprout/utils/hash"
)

type HandleProjectProvers func(projectID uint64, proverIDs []uint64)

type GetCachedProjectIDs func() []uint64

type projectOffset struct {
	scheduledBlockNumber atomic.Uint64
	projectIDs           sync.Map // projectID(uint64)->true
}

type scheduler struct {
	contractProvers      *contractProvers
	contractProjects     *contractProjects
	projectOffsets       *sync.Map // project offset in epoch offset(uint64) -> *projectOffset
	epoch                uint64
	pubSubs              *p2p.PubSubs // TODO define interface
	chainHead            chan uint64
	proverID             uint64
	handleProjectProvers HandleProjectProvers
}

func (s *scheduler) schedule() {
	for head := range s.chainHead {
		for blockNumber := head - s.epoch + 1; blockNumber <= head; blockNumber++ {
			offset := blockNumber % s.epoch
			projects, ok := s.projectOffsets.Load(offset)
			if !ok {
				continue
			}
			if projects.(*projectOffset).scheduledBlockNumber.Load() == blockNumber {
				continue
			}

			proverIDs := []uint64{}
			for id := range s.contractProvers.get(blockNumber).Provers {
				proverIDs = append(proverIDs, id)
			}
			scheduled := true

			projects.(*projectOffset).projectIDs.Range(func(key, value any) bool {
				projectID := key.(uint64)
				slog.Info("a new epoch has arrived", "block_number", blockNumber, "project_id", projectID)

				amount := uint64(1) // TODO fetch amount from project attr
				if amount > uint64(len(proverIDs)) {
					slog.Error("no enough resource for the project", "required_prover_amount", amount, "current_prover_amount", len(proverIDs), "project_id", projectID)
					scheduled = false
					return false
				}

				projectProverIDs := distance.GetMinNLocation(proverIDs, projectID, amount)

				isMy := false
				for _, p := range projectProverIDs {
					if p == s.proverID {
						isMy = true
					}
				}
				if !isMy {
					slog.Info("the project not scheduled to this prover", "project_id", projectID)
					s.pubSubs.Delete(projectID)
					return true
				}
				s.handleProjectProvers(projectID, projectProverIDs)
				s.pubSubs.Add(projectID)
				slog.Info("the project scheduled to this prover", "project_id", projectID)
				return true
			})
			if scheduled {
				projects.(*projectOffset).scheduledBlockNumber.Store(blockNumber)
			}
		}
	}
}

func watchChainHead(head chan<- uint64, chainEndpoint string) error {
	client, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		return errors.Wrapf(err, "failed to dial chain endpoint %s", chainEndpoint)
	}
	currentHead := uint64(0)

	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for range ticker.C {
			latestBlockNumber, err := client.BlockNumber(context.Background())
			if err != nil {
				slog.Error("failed to query the latest block number", "error", err)
				continue
			}
			if latestBlockNumber > currentHead {
				head <- latestBlockNumber
				currentHead = latestBlockNumber
			}
		}
	}()
	return nil
}

func Run(epoch uint64, chainEndpoint, proverContractAddress, projectContractAddress, projectFileDirectory string, proverID uint64,
	pubSubs *p2p.PubSubs, handleProjectProvers HandleProjectProvers, getProjectIDs GetCachedProjectIDs) error {

	if projectFileDirectory != "" {
		dummySchedule(pubSubs, handleProjectProvers, getProjectIDs)
		return nil
	}

	contractProvers := &contractProvers{
		epoch: epoch,
		datas: list.New(),
	}

	proverCh, err := contract.ListAndWatchProver(chainEndpoint, proverContractAddress)
	if err != nil {
		return err
	}
	go func() {
		for p := range proverCh {
			slog.Info("get a new provers", "block_number", p.BlockNumber)
			contractProvers.set(p)
		}
	}()

	projectOffsets := &sync.Map{}
	contractProjects := &contractProjects{
		epoch: epoch,
		datas: list.New(),
	}
	projectCh, err := contract.ListAndWatchProject(chainEndpoint, projectContractAddress)
	if err != nil {
		return err
	}
	go func() {
		for p := range projectCh {
			slog.Info("get a new projects", "block_number", p.BlockNumber)
			contractProjects.set(p)

			for projectID := range p.Projects {
				projectIDHash := hash.Sum256Uint64(projectID)
				offset := new(big.Int).SetBytes(projectIDHash[:]).Uint64() % epoch

				projects, _ := projectOffsets.LoadOrStore(offset, &projectOffset{})
				projects.(*projectOffset).projectIDs.Store(projectID, true)
			}
		}
	}()

	chainHead := make(chan uint64)
	if err := watchChainHead(chainHead, chainEndpoint); err != nil {
		return err
	}

	s := &scheduler{
		contractProvers:      contractProvers,
		contractProjects:     contractProjects,
		projectOffsets:       projectOffsets,
		epoch:                epoch,
		pubSubs:              pubSubs,
		chainHead:            chainHead,
		proverID:             proverID,
		handleProjectProvers: handleProjectProvers,
	}
	go s.schedule()
	return nil
}

func dummySchedule(pubSubs *p2p.PubSubs, handleProjectProvers HandleProjectProvers, getProjectIDs GetCachedProjectIDs) {
	s := &scheduler{
		pubSubs:              pubSubs,
		handleProjectProvers: handleProjectProvers,
	}

	projectIDs := getProjectIDs()
	for _, id := range projectIDs {
		s.handleProjectProvers(id, []uint64{})
		s.pubSubs.Add(id)
		slog.Info("the project scheduled to this prover", "project_id", id)
	}
}
