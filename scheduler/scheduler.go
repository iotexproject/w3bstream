package scheduler

import (
	"log/slog"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/persistence/contract"
	"github.com/machinefi/sprout/util/distance"
	"github.com/machinefi/sprout/util/hash"
)

type HandleProjectProvers func(projectID uint64, proverIDs []uint64)

type ProjectIDs func() []uint64

type ContractProject func(projectID, blockNumber uint64) *contract.Project
type LatestProjects func() []*contract.Project
type ContractProvers func(blockNumber uint64) []*contract.Prover

type projectOffset struct {
	scheduledBlockNumber atomic.Uint64
	projectIDs           sync.Map // projectID(uint64)->true
}

type scheduler struct {
	chainHead            <-chan uint64
	contractProject      ContractProject
	contractProvers      ContractProvers
	projectOffsets       *sync.Map // project offset in epoch offset(uint64) -> *projectOffset
	epoch                uint64
	pubSubs              *p2p.PubSubs // TODO define interface
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
			for _, p := range s.contractProvers(blockNumber) {
				if !*p.Paused {
					proverIDs = append(proverIDs, p.ID)
				}
			}
			scheduled := true

			projects.(*projectOffset).projectIDs.Range(func(key, value any) bool {
				projectID := key.(uint64)
				slog.Info("a new epoch has arrived", "project_id", projectID, "block_number", blockNumber)

				cp := s.contractProject(projectID, blockNumber)
				if cp == nil {
					slog.Error("failed to find project from contract", "project_id", projectID)
					scheduled = false
					return false
				}

				amount := uint64(1)
				if v, ok := cp.Attributes[contract.RequiredProverAmountHash]; ok {
					n, err := strconv.ParseUint(string(v), 10, 64)
					if err != nil {
						slog.Error("failed to parse project required prover amount", "project_id", projectID)
						scheduled = false
						return false
					}
					amount = n
				}

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

func storeProject(epoch uint64, projectOffsets *sync.Map, p *contract.Project) {
	offset := hash.Keccak256Uint64(p.ID).Big().Uint64() % epoch

	projects, _ := projectOffsets.LoadOrStore(offset, &projectOffset{})
	projects.(*projectOffset).projectIDs.Store(p.ID, true)
}

func Run(epoch uint64, proverID uint64, pubSubs *p2p.PubSubs, handleProjectProvers HandleProjectProvers, chainHead <-chan uint64, projectNotification <-chan *contract.Project, contractProject ContractProject, latestProjects LatestProjects, contractProvers ContractProvers) error {
	projectOffsets := &sync.Map{}

	ps := latestProjects()
	for _, p := range ps {
		storeProject(epoch, projectOffsets, p)
	}
	go func() {
		for p := range projectNotification {
			slog.Info("get new project contract events", "block_number", p.BlockNumber)
			storeProject(epoch, projectOffsets, p)
		}
	}()

	s := &scheduler{
		contractProvers:      contractProvers,
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

func RunLocal(pubSubs *p2p.PubSubs, handleProjectProvers HandleProjectProvers, projectIDs ProjectIDs) {
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
