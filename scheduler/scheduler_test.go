package scheduler

import (
	"sync/atomic"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/iotexproject/w3bstream/p2p"
	"github.com/iotexproject/w3bstream/persistence/contract"
	"github.com/iotexproject/w3bstream/project"
	"github.com/iotexproject/w3bstream/util/distance"
)

func TestScheduler_schedule(t *testing.T) {
	r := require.New(t)
	scheduledProverID := atomic.Uint64{}
	handleProjectProvers := func(projectID uint64, proverIDs []uint64) {
		scheduledProverID.Store(proverIDs[0])
	}
	t.Run("NoProject", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		pes := &ProjectEpochOffsets{}
		p.ApplyMethodReturn(pes, "Projects", nil)

		chainHead := make(chan uint64, 10)
		chainHead <- 1
		close(chainHead)
		s := &scheduler{
			chainHead:            chainHead,
			projectOffsets:       pes,
			epoch:                1,
			handleProjectProvers: handleProjectProvers,
		}
		s.schedule()
		r.Equal(uint64(0), scheduledProverID.Load())
	})
	t.Run("ProjectAlreadyScheduled", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		pes := &ProjectEpochOffsets{}
		p.ApplyMethodReturn(pes, "Projects", []*ScheduledProject{{1, 1}})

		chainHead := make(chan uint64, 10)
		chainHead <- 1
		close(chainHead)
		s := &scheduler{
			chainHead:            chainHead,
			projectOffsets:       pes,
			epoch:                1,
			handleProjectProvers: handleProjectProvers,
		}
		pe := &projectEpochOffset{}
		pe.projectIDs.Store(uint64(1), uint64(1))
		s.projectOffsets.projects.Store(uint64(0), pe)
		s.schedule()
		r.Equal(uint64(0), scheduledProverID.Load())
	})
	t.Run("ContractProjectNotExist", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		paused := false
		pm := &contract.Contract{}
		pes := &ProjectEpochOffsets{}
		p.ApplyMethodReturn(pes, "Projects", []*ScheduledProject{{1, 0}})
		p.ApplyMethodReturn(pm, "Provers", []*contract.Prover{{ID: 1, Paused: paused}})
		p.ApplyMethodReturn(pm, "Project", nil)

		chainHead := make(chan uint64, 10)
		chainHead <- 1
		close(chainHead)
		s := &scheduler{
			contractProvers:      pm.Provers,
			contractProject:      pm.Project,
			chainHead:            chainHead,
			projectOffsets:       pes,
			epoch:                1,
			handleProjectProvers: handleProjectProvers,
		}
		pe := &projectEpochOffset{}
		pe.projectIDs.Store(uint64(1), uint64(1))
		s.projectOffsets.projects.Store(uint64(0), pe)
		s.schedule()
		r.Equal(uint64(0), scheduledProverID.Load())
	})
	t.Run("FailedToParseProjectRequiredProverAmount", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		paused := false
		pm := &contract.Contract{}
		pes := &ProjectEpochOffsets{}
		pf := &project.Manager{}
		p.ApplyMethodReturn(pes, "Projects", []*ScheduledProject{{1, 0}})
		p.ApplyMethodReturn(pm, "Provers", []*contract.Prover{{ID: 1, Paused: paused}})
		p.ApplyMethodReturn(pm, "Project", &contract.Project{
			Attributes: map[common.Hash][]byte{contract.RequiredProverAmount: []byte("err")},
		})
		p.ApplyMethodReturn(pf, "Project", &project.Project{}, nil)
		p.ApplyMethodReturn(&project.Project{}, "DefaultConfig", &project.Config{}, nil)

		chainHead := make(chan uint64, 10)
		chainHead <- 1
		close(chainHead)
		s := &scheduler{
			contractProvers:      pm.Provers,
			contractProject:      pm.Project,
			chainHead:            chainHead,
			projectOffsets:       pes,
			epoch:                1,
			handleProjectProvers: handleProjectProvers,
			projectManager:       pf,
		}
		pe := &projectEpochOffset{}
		pe.projectIDs.Store(uint64(1), uint64(1))
		s.projectOffsets.projects.Store(uint64(0), pe)
		s.schedule()
		r.Equal(uint64(0), scheduledProverID.Load())
	})
	t.Run("NoEnoughResource", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		paused := false
		pm := &contract.Contract{}
		pes := &ProjectEpochOffsets{epoch: 1}
		pf := &project.Manager{}
		p.ApplyMethodReturn(pes, "Projects", []*ScheduledProject{{1, 0}})
		p.ApplyMethodReturn(pm, "Provers", []*contract.Prover{{ID: 1, Paused: paused}})
		p.ApplyMethodReturn(pm, "Project", &contract.Project{
			Attributes: map[common.Hash][]byte{contract.RequiredProverAmount: []byte("10")},
		})
		p.ApplyMethodReturn(&p2p.PubSub{}, "Delete")
		p.ApplyMethodReturn(pf, "Project", &project.Project{}, nil)
		p.ApplyMethodReturn(&project.Project{}, "DefaultConfig", &project.Config{}, nil)

		chainHead := make(chan uint64, 10)
		chainHead <- 1
		close(chainHead)
		s := &scheduler{
			contractProvers:      pm.Provers,
			contractProject:      pm.Project,
			chainHead:            chainHead,
			projectOffsets:       pes,
			epoch:                1,
			handleProjectProvers: handleProjectProvers,
			projectManager:       pf,
		}
		pe := &projectEpochOffset{}
		pe.projectIDs.Store(uint64(1), uint64(1))
		s.projectOffsets.projects.Store(uint64(0), pe)
		s.schedule()
		r.Equal(uint64(0), scheduledProverID.Load())
	})
	t.Run("TheProjectNotScheduledToThisProver", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		paused := false
		pm := &contract.Contract{}
		pes := &ProjectEpochOffsets{epoch: 1}
		pf := &project.Manager{}
		p.ApplyMethodReturn(pes, "Projects", []*ScheduledProject{{1, 0}})
		p.ApplyMethodReturn(pm, "Provers", []*contract.Prover{{ID: 1, Paused: paused}})
		p.ApplyMethodReturn(pm, "Project", &contract.Project{ID: 1})
		p.ApplyFuncReturn(distance.Sort, []uint64{100})
		p.ApplyMethodReturn(&p2p.PubSub{}, "Delete")
		p.ApplyMethodReturn(pf, "Project", &project.Project{}, nil)
		p.ApplyMethodReturn(&project.Project{}, "DefaultConfig", &project.Config{}, nil)

		chainHead := make(chan uint64, 10)
		chainHead <- 1
		close(chainHead)
		s := &scheduler{
			contractProvers:      pm.Provers,
			contractProject:      pm.Project,
			chainHead:            chainHead,
			projectOffsets:       pes,
			epoch:                1,
			handleProjectProvers: handleProjectProvers,
			projectManager:       pf,
		}
		pe := &projectEpochOffset{}
		pe.projectIDs.Store(uint64(1), uint64(1))
		s.projectOffsets.projects.Store(uint64(0), pe)
		s.schedule()
		r.Equal(uint64(0), scheduledProverID.Load())
	})
	t.Run("FailedToAddPubSubs", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		paused := false
		pm := &contract.Contract{}
		pes := &ProjectEpochOffsets{}
		pf := &project.Manager{}
		p.ApplyMethodReturn(pes, "Projects", []*ScheduledProject{{1, 0}})
		p.ApplyMethodReturn(pm, "Provers", []*contract.Prover{{ID: 1, Paused: paused, VMTypes: map[uint64]bool{0: true}}})
		p.ApplyMethodReturn(pm, "Project", &contract.Project{ID: 1})
		p.ApplyMethodReturn(&p2p.PubSub{}, "Add", errors.New(t.Name()))
		p.ApplyMethodReturn(pf, "Project", &project.Project{}, nil)
		p.ApplyMethodReturn(&project.Project{}, "DefaultConfig", &project.Config{}, nil)

		chainHead := make(chan uint64, 10)
		chainHead <- 1
		close(chainHead)
		s := &scheduler{
			proverID:             1,
			contractProvers:      pm.Provers,
			contractProject:      pm.Project,
			chainHead:            chainHead,
			projectOffsets:       pes,
			epoch:                1,
			handleProjectProvers: handleProjectProvers,
			projectManager:       pf,
		}
		pe := &projectEpochOffset{}
		pe.projectIDs.Store(uint64(1), uint64(1))
		s.projectOffsets.projects.Store(uint64(0), pe)
		s.schedule()
		r.Equal(uint64(1), scheduledProverID.Load())
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		paused := false
		pm := &contract.Contract{}
		pes := &ProjectEpochOffsets{epoch: 1}
		pf := &project.Manager{}
		p.ApplyMethodReturn(pes, "Projects", []*ScheduledProject{{1, 0}})
		p.ApplyMethodReturn(pm, "Provers", []*contract.Prover{{ID: 1, Paused: paused, VMTypes: map[uint64]bool{0: true}}})
		p.ApplyMethodReturn(pm, "Project", &contract.Project{ID: 1})
		p.ApplyMethodReturn(&p2p.PubSub{}, "Add", nil)
		p.ApplyMethodReturn(pf, "Project", &project.Project{}, nil)
		p.ApplyMethodReturn(&project.Project{}, "DefaultConfig", &project.Config{}, nil)

		chainHead := make(chan uint64, 10)
		chainHead <- 1
		close(chainHead)
		s := &scheduler{
			proverID:             1,
			contractProvers:      pm.Provers,
			contractProject:      pm.Project,
			chainHead:            chainHead,
			projectOffsets:       pes,
			epoch:                1,
			handleProjectProvers: handleProjectProvers,
			projectManager:       pf,
		}
		pe := &projectEpochOffset{}
		pe.projectIDs.Store(uint64(1), uint64(1))
		s.projectOffsets.projects.Store(uint64(0), pe)
		s.schedule()
		r.Equal(uint64(1), scheduledProverID.Load())
		k, _ := pe.projectIDs.Load(uint64(1))
		r.Equal(k.(uint64), uint64(1))
	})
}

func TestRun(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	p.ApplyPrivateMethod(&scheduler{}, "schedule", func() {})

	err := Run(10, 1, nil, nil, nil, nil, nil, &ProjectEpochOffsets{}, nil)
	r.NoError(err)
}

func TestRunLocal(t *testing.T) {
	t.Run("FailedToAddPubsubs", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		ps := &p2p.PubSub{}
		f := func(uint64, []uint64) {
		}
		pm := &project.Manager{}

		p.ApplyMethodReturn(pm, "ProjectIDs", []uint64{1}, nil)
		p.ApplyMethodReturn(&p2p.PubSub{}, "Add", errors.New(t.Name()))

		RunLocal(ps, f, pm)
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		ps := &p2p.PubSub{}
		f := func(uint64, []uint64) {
		}
		pm := &project.Manager{}

		p.ApplyMethodReturn(pm, "ProjectIDs", []uint64{1}, nil)
		p.ApplyMethodReturn(&p2p.PubSub{}, "Add", nil)

		RunLocal(ps, f, pm)
	})
}
