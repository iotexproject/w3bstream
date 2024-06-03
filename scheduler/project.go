package scheduler

import (
	"log/slog"
	"sync"

	"github.com/machinefi/sprout/persistence/contract"
	"github.com/machinefi/sprout/util/hash"
)

type LatestProjects func() []*contract.Project

type ScheduledProject struct {
	ID                   uint64
	ScheduledBlockNumber uint64
}

type projectEpochOffset struct {
	projectIDs sync.Map // projectID(uint64)->scheduledBlockNumber(uint64)
}

type ProjectEpochOffsets struct {
	epoch    uint64
	projects sync.Map // project offset in epoch (uint64) -> *projectEpochOffset
}

func (pe *ProjectEpochOffsets) storeProject(pid uint64) {
	offset := hash.Keccak256Uint64(pid).Big().Uint64() % pe.epoch

	projects, _ := pe.projects.LoadOrStore(offset, &projectEpochOffset{})
	projects.(*projectEpochOffset).projectIDs.LoadOrStore(pid, uint64(0))
}

func (pe *ProjectEpochOffsets) offset(blockNumber uint64) uint64 {
	return blockNumber % pe.epoch
}

func (pe *ProjectEpochOffsets) setBlockNumber(projectID, blockNumber uint64) {
	projects, ok := pe.projects.Load(pe.offset(blockNumber))
	if !ok {
		return
	}
	projects.(*projectEpochOffset).projectIDs.Store(projectID, blockNumber)
}

func (pe *ProjectEpochOffsets) Projects(blockNumber uint64) []*ScheduledProject {
	projects, ok := pe.projects.Load(pe.offset(blockNumber))
	if !ok {
		return nil
	}
	ps := []*ScheduledProject{}
	projects.(*projectEpochOffset).projectIDs.Range(func(key, value any) bool {
		ps = append(ps, &ScheduledProject{
			ID:                   key.(uint64),
			ScheduledBlockNumber: value.(uint64),
		})
		return true
	})
	return ps
}

func NewProjectEpochOffsets(epoch uint64, latestProjects LatestProjects, projectNotification <-chan uint64) *ProjectEpochOffsets {
	pe := &ProjectEpochOffsets{
		epoch: epoch,
	}

	ps := latestProjects()
	for _, p := range ps {
		pe.storeProject(p.ID)
	}
	go func() {
		for pid := range projectNotification {
			slog.Info("get new project contract events", "project_id", pid)
			pe.storeProject(pid)
		}
	}()
	return pe
}
