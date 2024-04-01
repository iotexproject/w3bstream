package scheduler

import (
	"log/slog"
	"sync"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/utils/distance"
)

type HandleProjectProvers func(projectID uint64, provers []string)

type scheduler struct {
	provers              *sync.Map // proverPrivateKey(string) -> true(bool)
	projectOffsets       *sync.Map // project offset in interval(uint64) -> projectMeta(*project.ProjectMeta)
	epoch                uint64
	pubSubs              *p2p.PubSubs // TODO define interface
	chainHead            chan *types.Header
	proverPrivateKey     string
	handleProjectProvers HandleProjectProvers
}

func (s *scheduler) schedule() {
	for head := range s.chainHead {
		offset := s.epoch - (head.Number.Uint64() % s.epoch)
		metaValue, ok := s.projectOffsets.Load(offset)
		if !ok {
			continue
		}
		meta := metaValue.(*project.ProjectMeta)
		provers := s.getAllProver()

		amount := meta.ProverAmount
		if amount == 0 {
			amount = 1
		}
		if amount > uint(len(provers)) {
			slog.Error("no enough resource for the project", "require prover amount", amount, "current prover", len(provers), "project_id", meta.ProjectID)
			continue
		}

		projectProvers := distance.GetMinNLocation(provers, meta.ProjectID, uint64(amount))

		isMy := false
		for _, p := range projectProvers {
			if p == s.proverPrivateKey {
				isMy = true
			}
		}
		if !isMy {
			slog.Info("the project not scheduld to this prover", "project_id", meta.ProjectID)
			s.pubSubs.Delete(meta.ProjectID)
			continue
		}
		s.handleProjectProvers(meta.ProjectID, projectProvers)
		s.pubSubs.Add(meta.ProjectID)
		slog.Info("the project scheduld to this prover", "project_id", meta.ProjectID)
	}
}

func (s *scheduler) getAllProver() []string {
	provers := []string{}
	s.provers.Range(func(key, value any) bool {
		provers = append(provers, key.(string))
		return true
	})
	return provers
}

func Run(epoch uint64, chainEndpoint, proverContractAddress, projectContractAddress, proverPrivateKey string, pubSubs *p2p.PubSubs, handleProjectProvers HandleProjectProvers) error {
	provers := &sync.Map{}
	if err := watchProver(provers, chainEndpoint, proverContractAddress); err != nil {
		return err
	}
	if err := listAndFillProver(provers, chainEndpoint, proverContractAddress); err != nil {
		return err
	}

	projectOffsets := &sync.Map{}
	if err := watchProject(projectOffsets, epoch, chainEndpoint, projectContractAddress); err != nil {
		return err
	}
	if err := listAndFillProject(projectOffsets, epoch, chainEndpoint, projectContractAddress); err != nil {
		return err
	}

	chainHead := make(chan *types.Header)
	if err := watchChainHead(chainHead, chainEndpoint); err != nil {
		return err
	}

	s := &scheduler{
		provers:              provers,
		projectOffsets:       projectOffsets,
		epoch:                epoch,
		pubSubs:              pubSubs,
		chainHead:            chainHead,
		proverPrivateKey:     proverPrivateKey,
		handleProjectProvers: handleProjectProvers,
	}
	go s.schedule()
	return nil
}
