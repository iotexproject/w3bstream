package scheduler

import (
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/persistence/contract"
)

func TestProjectEpochOffsets_storeProject(t *testing.T) {
	r := require.New(t)

	pes := &ProjectEpochOffsets{epoch: 1}
	pes.storeProject(1)

	ps := pes.Projects(1)
	r.True(len(ps) > 0)
}

func TestProjectEpochOffsets_setBlockNumber(t *testing.T) {
	r := require.New(t)

	pes := &ProjectEpochOffsets{epoch: 1}
	pes.storeProject(1)
	pes.setBlockNumber(1, 100)

	ps := pes.Projects(1)
	r.True(len(ps) > 0)
	r.Equal(ps[0].ScheduledBlockNumber, uint64(100))
}

func TestProjectEpochOffsets_Projects(t *testing.T) {
	r := require.New(t)

	pes := &ProjectEpochOffsets{epoch: 1}
	ps := pes.Projects(1)
	r.True(len(ps) == 0)

	pes.storeProject(1)
	ps = pes.Projects(1)
	r.True(len(ps) > 0)
	r.Equal(ps[0].ScheduledBlockNumber, uint64(0))
}

func TestNewProjectEpochOffsets(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	pm := &contract.Contract{}
	p.ApplyMethodReturn(pm, "LatestProjects", []*contract.Project{{ID: 1}})

	notification := make(chan uint64, 10)
	notification <- 2
	close(notification)

	pe := NewProjectEpochOffsets(1, pm.LatestProjects, notification)
	time.Sleep(10 * time.Millisecond)
	ps := pe.Projects(1)
	r.True(len(ps) == 2)
}
