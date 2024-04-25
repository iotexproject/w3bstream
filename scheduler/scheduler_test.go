package scheduler

import (
	"testing"

	"github.com/agiledragon/gomonkey/v2"

	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/project"
)

// func TestScheduler_schedule(t *testing.T) {
// 	r := require.New(t)
// 	p := gomonkey.NewPatches()
// 	defer p.Reset()

// 	p.ApplyMethodReturn(&p2p.PubSubs{}, "Add", nil)

// 	scheduledProverID := atomic.Uint64{}

// 	s := &scheduler{
// 		contractProver: &contractProver{
// 			epoch:  1,
// 			blocks: list.New(),
// 		},
// 		contractProject: &contractProject{
// 			epoch:  10,
// 			blocks: list.New(),
// 		},
// 		projectOffsets: &sync.Map{},
// 		epoch:          1,
// 		chainHead:      make(chan uint64, 10),
// 		proverID:       1,
// 		handleProjectProvers: func(projectID uint64, proverIDs []uint64) {
// 			scheduledProverID.Store(proverIDs[0])
// 		},
// 	}

// 	s.chainHead <- 100
// 	s.contractProver.add(&contract.BlockProver{
// 		BlockNumber: 100,
// 		Provers: map[uint64]*contract.Prover{
// 			1: {
// 				ID: 1,
// 			},
// 		},
// 	})
// 	pf := &projectOffset{}
// 	projectID := uint64(1)
// 	pf.projectIDs.Store(projectID, true)
// 	s.projectOffsets.Store(hash.Keccak256Uint64(projectID).Big().Uint64()%s.epoch, pf)
// 	go s.schedule()
// 	for scheduledProverID.Load() == 0 {
// 	}
// 	close(s.chainHead)
// 	r.Equal(scheduledProverID.Load(), uint64(1))
// }

// func TestRun(t *testing.T) {
// 	r := require.New(t)
// 	p := gomonkey.NewPatches()
// 	defer p.Reset()

// 	p.ApplyFuncReturn(contract.ListAndWatchProver, make(chan *contract.BlockProver), nil)
// 	p.ApplyFuncReturn(contract.ListAndWatchProject, make(chan *contract.BlockProject), nil)
// 	p.ApplyFuncReturn(watchChainHead, nil)

// 	err := Run(1, "", "", "", "", 1, nil, nil, nil)
// 	r.NoError(err)
// }

func TestDummySchedule(t *testing.T) {
	p := gomonkey.NewPatches()
	defer p.Reset()

	ps := &p2p.PubSubs{}
	f := func(uint64, []uint64) {
	}
	pm := &project.Manager{}

	p.ApplyMethodReturn(pm, "ProjectIDs", []uint64{1})
	p.ApplyMethodReturn(&p2p.PubSubs{}, "Add", nil)

	RunLocal(ps, f, pm.ProjectIDs)
}
