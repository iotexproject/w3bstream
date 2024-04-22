package scheduler

import (
	"container/list"
	"context"
	"log/slog"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/persistence/contract"
	"github.com/machinefi/sprout/util/distance"
	"github.com/machinefi/sprout/util/hash"
)

type HandleProjectProvers func(projectID uint64, proverIDs []uint64)

type ProjectIDs func() []uint64

type projectOffset struct {
	scheduledBlockNumber atomic.Uint64
	projectIDs           sync.Map // projectID(uint64)->true
}

type scheduler struct {
	contractProver       *contractProver
	contractProject      *contractProject
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
			for id := range s.contractProver.blockProver(blockNumber).Provers {
				proverIDs = append(proverIDs, id)
			}
			scheduled := true

			projects.(*projectOffset).projectIDs.Range(func(key, value any) bool {
				projectID := key.(uint64)
				slog.Info("a new epoch has arrived", "project_id", projectID, "block_number", blockNumber)

				amount := uint64(1) // TODO fetch amount from project attr
				if amount > uint64(len(proverIDs)) {
					slog.Error("no enough resource for the project", "project_id", projectID, "required_prover_amount", amount, "current_prover_amount", len(proverIDs))
					scheduled = false
					return false
				}

				projectProverIDs := distance.Sort(proverIDs, projectID)
				projectProverIDs = projectProverIDs[:amount]

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
				if err := s.pubSubs.Add(projectID); err != nil {
					slog.Error("failed to add pubsubs", "project_id", projectID, "error", err)
					scheduled = false
					return false
				}
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
	pubSubs *p2p.PubSubs, handleProjectProvers HandleProjectProvers, getProjectIDs ProjectIDs) error {

	if projectFileDirectory != "" {
		dummySchedule(pubSubs, handleProjectProvers, getProjectIDs)
		return nil
	}

	contractProver := &contractProver{
		epoch:  epoch,
		blocks: list.New(),
	}

	proverCh, err := contract.ListAndWatchProver(chainEndpoint, proverContractAddress, epoch)
	if err != nil {
		return err
	}
	go func() {
		for p := range proverCh {
			slog.Info("get new prover contract events", "block_number", p.BlockNumber)
			contractProver.add(p)
		}
	}()

	projectOffsets := &sync.Map{}
	contractProject := &contractProject{
		epoch:  epoch,
		blocks: list.New(),
	}
	projectCh, err := contract.ListAndWatchProject(chainEndpoint, projectContractAddress, epoch)
	if err != nil {
		return err
	}
	go func() {
		for p := range projectCh {
			slog.Info("get new project contract events", "block_number", p.BlockNumber)
			contractProject.add(p)

			for projectID := range p.Projects {
				offset := hash.Keccak256Uint64(projectID).Big().Uint64() % epoch

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
		contractProver:       contractProver,
		contractProject:      contractProject,
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

func dummySchedule(pubSubs *p2p.PubSubs, handleProjectProvers HandleProjectProvers, projectIDs ProjectIDs) {
	s := &scheduler{
		pubSubs:              pubSubs,
		handleProjectProvers: handleProjectProvers,
	}

	ids := projectIDs()
	for _, id := range ids {
		s.handleProjectProvers(id, []uint64{})
		if err := s.pubSubs.Add(id); err != nil {
			slog.Error("failed to add pubsubs", "project_id", id, "error", err)
			continue
		}
		slog.Info("the project scheduled to this prover", "project_id", id)
	}
}
