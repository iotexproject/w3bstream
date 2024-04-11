package scheduler

import (
	"context"
	"log/slog"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/utils/contract"
	"github.com/machinefi/sprout/utils/distance"
	"github.com/machinefi/sprout/utils/hash"
)

type HandleProjectProvers func(projectID uint64, provers []string)

type scheduler struct {
	provers              *sync.Map // proverID(string) -> Prover(*contract.Prover)
	projectOffsets       *sync.Map // project offset in interval(uint64) -> Project(*contract.Project)
	epoch                uint64
	pubSubs              *p2p.PubSubs // TODO define interface
	chainHead            chan uint64
	proverID             string
	handleProjectProvers HandleProjectProvers
}

func (s *scheduler) schedule() {
	for head := range s.chainHead {
		offset := s.epoch - (head % s.epoch)
		cp, ok := s.projectOffsets.Load(offset)
		if !ok {
			continue
		}
		projectID := cp.(*contract.Project).ID
		slog.Info("a new epoch has arrived", "head_number", head, "project_id", projectID)

		provers := s.getAllProver()

		amount := uint64(1) // TODO fetch amount from project attr
		if amount > uint64(len(provers)) {
			slog.Error("no enough resource for the project", "require prover amount", amount, "current prover", len(provers), "project_id", projectID)
			continue
		}

		projectProvers := distance.GetMinNLocation(provers, projectID, amount)

		isMy := false
		for _, p := range projectProvers {
			if p == s.proverID {
				isMy = true
			}
		}
		if !isMy {
			slog.Info("the project not scheduld to this prover", "project_id", projectID)
			s.pubSubs.Delete(projectID)
			continue
		}
		s.handleProjectProvers(projectID, projectProvers)
		s.pubSubs.Add(projectID)
		slog.Info("the project scheduld to this prover", "project_id", projectID)
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

func Run(epoch uint64, chainEndpoint, proverContractAddress, projectContractAddress, proverID string, pubSubs *p2p.PubSubs, handleProjectProvers HandleProjectProvers) error {
	provers := &sync.Map{}
	proverCh, err := contract.ListAndWatchProver(chainEndpoint, proverContractAddress)
	if err != nil {
		return err
	}
	go func() {
		for p := range proverCh {
			slog.Info("get a new prover", "prover_operator", p.OperatorAddress)
			e, ok := provers.Load(p.OperatorAddress)
			if ok {
				if ep := e.(*contract.Prover); ep.BlockNumber > p.BlockNumber {
					p = ep
				}
			}
			provers.Store(p.OperatorAddress, p)
		}
	}()

	projectOffsets := &sync.Map{}
	projectCh, err := contract.ListAndWatchProject(chainEndpoint, projectContractAddress)
	if err != nil {
		return err
	}
	go func() {
		for p := range projectCh {
			slog.Info("get a new project", "project_id", p.ID)
			projectIDHash := hash.Sum256Uint64(p.ID)
			offset := new(big.Int).SetBytes(projectIDHash[:]).Uint64() % epoch

			e, ok := projectOffsets.Load(offset)
			if ok {
				if ep := e.(*contract.Project); ep.BlockNumber > p.BlockNumber {
					p = ep
				}
			}
			projectOffsets.Store(offset, p) // TODO different project may have same offset
		}
	}()

	chainHead := make(chan uint64)
	if err := watchChainHead(chainHead, chainEndpoint); err != nil {
		return err
	}

	s := &scheduler{
		provers:              provers,
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
