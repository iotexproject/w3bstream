package scheduler

import (
	"container/list"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/persistence/contract"
	"github.com/machinefi/sprout/util/hash"
)

func TestScheduler_schedule(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	p.ApplyMethodReturn(&p2p.PubSubs{}, "Add", nil)

	scheduledProverID := atomic.Uint64{}

	s := &scheduler{
		contractProver: &contractProver{
			epoch:  1,
			blocks: list.New(),
		},
		contractProject: &contractProject{
			epoch:  10,
			blocks: list.New(),
		},
		projectOffsets: &sync.Map{},
		epoch:          1,
		chainHead:      make(chan uint64, 10),
		proverID:       1,
		handleProjectProvers: func(projectID uint64, proverIDs []uint64) {
			scheduledProverID.Store(proverIDs[0])
		},
	}

	s.chainHead <- 100
	s.contractProver.add(&contract.BlockProver{
		BlockNumber: 100,
		Provers: map[uint64]*contract.Prover{
			1: {
				ID: 1,
			},
		},
	})
	pf := &projectOffset{}
	projectID := uint64(1)
	pf.projectIDs.Store(projectID, true)
	s.projectOffsets.Store(hash.Keccak256Uint64(projectID).Big().Uint64()%s.epoch, pf)
	go s.schedule()
	for scheduledProverID.Load() == 0 {
	}
	close(s.chainHead)
	r.Equal(scheduledProverID.Load(), uint64(1))
}

func TestWatchChainHead(t *testing.T) {
	r := require.New(t)
	t.Run("FailedToDialChain", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(ethclient.Dial, nil, errors.New(t.Name()))

		err := watchChainHead(make(chan<- uint64), "")
		r.ErrorContains(err, t.Name())
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		c := make(chan time.Time, 10)
		p.ApplyFuncReturn(ethclient.Dial, &ethclient.Client{}, nil)
		p.ApplyFuncReturn(time.NewTicker, &time.Ticker{C: c})
		p.ApplyMethodReturn(&ethclient.Client{}, "BlockNumber", uint64(1), nil)

		c <- time.Now()
		h := make(chan uint64, 10)
		err := watchChainHead(h, "")
		r.NoError(err)
		d := <-h
		r.Equal(d, uint64(1))
		close(c)
	})
}

func TestRun(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	p.ApplyFuncReturn(contract.ListAndWatchProver, make(chan *contract.BlockProver), nil)
	p.ApplyFuncReturn(contract.ListAndWatchProject, make(chan *contract.BlockProject), nil)
	p.ApplyFuncReturn(watchChainHead, nil)

	err := Run(1, "", "", "", "", 1, nil, nil, nil)
	r.NoError(err)
}
