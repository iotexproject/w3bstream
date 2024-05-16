package scheduler

import (
	"log/slog"
	"strconv"

	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/persistence/contract"
	"github.com/machinefi/sprout/util/distance"
)

type HandleProjectProvers func(projectID uint64, proverIDs []uint64)

type ProjectIDs func() []uint64
type ContractProject func(projectID, blockNumber uint64) *contract.Project
type ContractProvers func(blockNumber uint64) []*contract.Prover

type scheduler struct {
	chainHead            <-chan uint64
	contractProject      ContractProject
	contractProvers      ContractProvers
	projectOffsets       *ProjectEpochOffsets
	epoch                uint64
	pubSubs              *p2p.PubSubs
	proverID             uint64
	handleProjectProvers HandleProjectProvers
}

func (s *scheduler) schedule() {
	for head := range s.chainHead {
		for blockNumber := head - s.epoch + 1; blockNumber <= head; blockNumber++ {
			projects := s.projectOffsets.Projects(blockNumber)
			if len(projects) == 0 {
				continue
			}
			scheduled := true
			for _, p := range projects {
				if p.ScheduledBlockNumber != blockNumber {
					scheduled = false
					break
				}
			}
			if scheduled {
				continue
			}

			proverIDs := []uint64{}
			for _, p := range s.contractProvers(blockNumber) {
				if !*p.Paused {
					proverIDs = append(proverIDs, p.ID)
				}
			}

			for _, p := range projects {
				if p.ScheduledBlockNumber == blockNumber {
					continue
				}
				projectID := p.ID
				slog.Info("a new epoch has arrived", "project_id", projectID, "block_number", blockNumber)

				cp := s.contractProject(projectID, blockNumber)
				if cp == nil {
					slog.Error("failed to find project from contract", "project_id", projectID)
					continue
				}

				amount := uint64(1)
				if v, ok := cp.Attributes[contract.RequiredProverAmountHash]; ok {
					n, err := strconv.ParseUint(string(v), 10, 64)
					if err != nil {
						slog.Error("failed to parse project required prover amount", "project_id", projectID)
						continue
					}
					amount = n
				}

				if amount > uint64(len(proverIDs)) {
					slog.Error("no enough resource for the project", "project_id", projectID, "required_prover_amount", amount, "current_prover_amount", len(proverIDs))
					continue
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
					continue
				}
				s.handleProjectProvers(projectID, projectProverIDs)
				if err := s.pubSubs.Add(projectID); err != nil {
					slog.Error("failed to add pubsubs", "project_id", projectID, "error", err)
					continue
				}
				s.projectOffsets.setBlockNumber(projectID, blockNumber)
				slog.Info("the project scheduled to this prover", "project_id", projectID)
			}
		}
	}
}

func Run(epoch uint64, proverID uint64, pubSubs *p2p.PubSubs, handleProjectProvers HandleProjectProvers, chainHead <-chan uint64, contractProject ContractProject, contractProvers ContractProvers, projectOffsets *ProjectEpochOffsets) error {
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
