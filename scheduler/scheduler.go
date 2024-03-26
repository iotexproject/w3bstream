package scheduler

import (
	"crypto/sha256"
	"encoding/binary"
	"log/slog"
	"math/big"
	"sort"
	"sync"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/project"
)

type HandleProjectProvers func(projectID uint64, provers []string)

type scheduler struct {
	provers              *sync.Map // proverID(string) -> true(bool)
	projectOffsets       *sync.Map // project offset in interval(uint64) -> projectID(uint64)
	interval             uint64
	pubSubs              *p2p.PubSubs
	chainHead            chan *types.Header
	projectManager       *project.Manager // TODO define interface
	proverID             string
	handleProjectProvers HandleProjectProvers
}

type distance struct {
	distance *big.Int
	hash     [sha256.Size]byte
}

func (s *scheduler) schedule() {
	for head := range s.chainHead {
		offset := s.interval - (head.Number.Uint64() % s.interval)
		projectIDValue, ok := s.projectOffsets.Load(offset)
		if !ok {
			continue
		}
		projectID := projectIDValue.(uint64)
		provers := s.getAllProver()

		projectConfig, err := s.projectManager.Get(projectID, "0.1") // TODO change project version
		if err != nil {
			slog.Error("failed to get project config", "error", err, "project_id", projectID)
			continue
		}
		if a := projectConfig.Config.ResourceRequest.ProverAmount; a > uint(len(provers)) {
			slog.Error("no enough resource for the project", "require prover amount", a, "current prover", len(provers), "project_id", projectID)
			continue
		}
		proverMap := map[[sha256.Size]byte]string{}
		for _, n := range provers {
			proverMap[sha256.Sum256([]byte(n))] = n
		}

		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, projectID)
		projectIDHash := sha256.Sum256(b)

		ds := make([]distance, 0, len(provers))

		for h := range proverMap {
			n := new(big.Int).Xor(new(big.Int).SetBytes(h[:]), new(big.Int).SetBytes(projectIDHash[:]))
			ds = append(ds, distance{
				distance: n,
				hash:     h,
			})
		}

		sort.Slice(ds, func(i, j int) bool {
			return ds[i].distance.Cmp(ds[j].distance) < 0
		})

		amount := projectConfig.Config.ResourceRequest.ProverAmount
		if amount == 0 {
			amount = 1
		}

		projectProvers := []string{}

		ds = ds[:amount]
		for _, d := range ds {
			projectProvers = append(projectProvers, proverMap[d.hash])
		}
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

func Run(interval uint64, chainEndpoint, proverContractAddress, projectContractAddress, proverID string, pubSubs *p2p.PubSubs, projectManager *project.Manager, handleProjectProvers HandleProjectProvers) error {
	provers := &sync.Map{}
	if err := watchProver(provers, chainEndpoint, proverContractAddress); err != nil {
		return err
	}
	if err := fillProver(provers, chainEndpoint, proverContractAddress); err != nil {
		return err
	}

	projectOffsets := &sync.Map{}
	if err := watchProject(projectOffsets, interval, chainEndpoint, projectContractAddress); err != nil {
		return err
	}
	if err := fillProject(projectOffsets, interval, chainEndpoint, projectContractAddress); err != nil {
		return err
	}

	chainHead := make(chan *types.Header)
	if err := watchChainHead(chainHead, chainEndpoint); err != nil {
		return err
	}

	s := &scheduler{
		provers:              provers,
		projectOffsets:       projectOffsets,
		interval:             interval,
		pubSubs:              pubSubs,
		chainHead:            chainHead,
		proverID:             proverID,
		projectManager:       projectManager,
		handleProjectProvers: handleProjectProvers,
	}
	go s.schedule()
	return nil
}
