package contract

import (
	"container/list"
	"errors"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/machinefi/sprout/smartcontracts/go/project"
	"github.com/machinefi/sprout/smartcontracts/go/prover"
	"github.com/stretchr/testify/require"
)

func TestContract_Project(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	bp := &blockProjects{}
	p.ApplyPrivateMethod(bp, "project", func(uint64, uint64) *Project { return nil })

	c := &Contract{
		blockProjects: bp,
	}
	r.Nil(c.Project(0, 0))
}

func TestContract_LatestProject(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	bp := &blockProjects{}
	p.ApplyPrivateMethod(bp, "project", func(uint64, uint64) *Project { return nil })

	c := &Contract{
		blockProjects: bp,
	}
	r.Nil(c.LatestProject(0))
}

func TestContract_LatestProjects(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	bp := &blockProjects{}
	p.ApplyPrivateMethod(bp, "projects", func() *blockProject {
		return &blockProject{
			Projects: map[uint64]*Project{1: {}},
		}
	})

	c := &Contract{
		blockProjects: bp,
	}
	r.Equal(len(c.LatestProjects()), 1)
}

func TestContract_Provers(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	bp := &blockProvers{}
	p.ApplyPrivateMethod(bp, "provers", func(blockNumber uint64) *blockProver {
		return &blockProver{
			Provers: map[uint64]*Prover{1: {}},
		}
	})

	c := &Contract{
		blockProvers: bp,
	}
	r.Equal(len(c.Provers(0)), 1)
}

func TestContract_LatestProvers(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	bp := &blockProvers{}
	p.ApplyPrivateMethod(bp, "provers", func(blockNumber uint64) *blockProver {
		return &blockProver{
			Provers: map[uint64]*Prover{1: {}},
		}
	})

	c := &Contract{
		blockProvers: bp,
	}
	r.Equal(len(c.LatestProvers()), 1)
}

func TestContract_Prover(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	bp := &blockProvers{}
	p.ApplyPrivateMethod(bp, "prover", func(common.Address) *Prover { return nil })

	c := &Contract{
		blockProvers: bp,
	}
	r.Nil(c.Prover(common.Address{}))
}

func TestContract_notifyProject(t *testing.T) {
	r := require.New(t)

	n := make(chan *Project, 10)
	c := &Contract{
		projectNotifications: []chan<- *Project{n},
	}
	c.notifyProject(&blockProject{
		Projects: map[uint64]*Project{1: {}},
	})
	p := <-n
	r.Equal(uint64(0), p.ID)
}

func TestContract_notifyChainHead(t *testing.T) {
	r := require.New(t)

	n := make(chan uint64, 10)
	c := &Contract{
		chainHeadNotifications: []chan<- uint64{n},
	}
	c.notifyChainHead(1)
	p := <-n
	r.Equal(uint64(1), p)
}

func TestContract_addBlockProject(t *testing.T) {
	p := gomonkey.NewPatches()
	defer p.Reset()

	bp := &blockProjects{}
	c := &Contract{
		blockProjects: bp,
	}
	p.ApplyPrivateMethod(bp, "add", func(*blockProject) {})
	p.ApplyPrivateMethod(c, "notifyProject", func(*blockProject) {})
	c.addBlockProject(nil)
}

func TestContract_list(t *testing.T) {
	r := require.New(t)
	t.Run("FailedToListProject", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		c := &Contract{}
		p.ApplyFuncReturn(listProject, nil, uint64(0), uint64(0), errors.New(t.Name()))

		_, err := c.list()
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToListProver", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		c := &Contract{}
		p.ApplyFuncReturn(listProject, nil, uint64(0), uint64(0), nil)
		p.ApplyFuncReturn(listProver, nil, uint64(0), uint64(0), errors.New(t.Name()))

		_, err := c.list()
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToFilterContractLogs", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		cli := &ethclient.Client{}
		c := &Contract{client: cli}
		p.ApplyFuncReturn(listProject, nil, uint64(100), uint64(200), nil)
		p.ApplyFuncReturn(listProver, nil, uint64(100), uint64(200), nil)
		p.ApplyMethodReturn(cli, "FilterLogs", nil, errors.New(t.Name()))

		_, err := c.list()
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToProcessProjectLogs", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		cli := &ethclient.Client{}
		c := &Contract{client: cli}
		p.ApplyFuncReturn(listProject, nil, uint64(100), uint64(200), nil)
		p.ApplyFuncReturn(listProver, nil, uint64(100), uint64(200), nil)
		p.ApplyMethodReturn(cli, "FilterLogs", nil, nil)
		p.ApplyFuncReturn(processProjectLogs, errors.New(t.Name()))

		_, err := c.list()
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToProcessProverLogs", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		cli := &ethclient.Client{}
		c := &Contract{client: cli}
		p.ApplyFuncReturn(listProject, nil, uint64(100), uint64(200), nil)
		p.ApplyFuncReturn(listProver, nil, uint64(100), uint64(200), nil)
		p.ApplyMethodReturn(cli, "FilterLogs", nil, nil)
		p.ApplyFuncReturn(processProjectLogs, nil)
		p.ApplyFuncReturn(processProverLogs, errors.New(t.Name()))

		_, err := c.list()
		r.ErrorContains(err, t.Name())
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		cli := &ethclient.Client{}
		c := &Contract{
			client: cli,
			blockProjects: &blockProjects{
				capacity: 10,
				blocks:   list.New(),
			},
			blockProvers: &blockProvers{
				capacity: 10,
				blocks:   list.New(),
			},
		}
		p.ApplyFuncReturn(listProject, nil, uint64(100), uint64(200), nil)
		p.ApplyFuncReturn(listProver, nil, uint64(100), uint64(200), nil)
		p.ApplyMethodReturn(cli, "FilterLogs", nil, nil)

		_, err := c.list()
		r.NoError(err)
	})
}

func TestContract_watch(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	ct := make(chan time.Time, 10)
	ticker := &time.Ticker{
		C: ct,
	}

	cli := &ethclient.Client{}
	head := make(chan uint64, 10)
	c := &Contract{
		client: cli,
		blockProjects: &blockProjects{
			capacity: 10,
			blocks:   list.New(),
		},
		blockProvers: &blockProvers{
			capacity: 10,
			blocks:   list.New(),
		},
		chainHeadNotifications: []chan<- uint64{head},
	}
	p.ApplyFuncReturn(time.NewTicker, ticker)
	p.ApplyMethodSeq(cli, "FilterLogs", []gomonkey.OutputCell{
		{
			Values: gomonkey.Params{nil, errors.New(t.Name())},
			Times:  1,
		},
		{
			Values: gomonkey.Params{nil, nil},
			Times:  3,
		},
	})
	p.ApplyFuncSeq(processProjectLogs, []gomonkey.OutputCell{
		{
			Values: gomonkey.Params{errors.New(t.Name())},
			Times:  1,
		},
		{
			Values: gomonkey.Params{nil},
			Times:  2,
		},
	})
	p.ApplyFuncSeq(processProverLogs, []gomonkey.OutputCell{
		{
			Values: gomonkey.Params{errors.New(t.Name())},
			Times:  1,
		},
		{
			Values: gomonkey.Params{nil},
			Times:  1,
		},
	})
	for i := 0; i < 4; i++ {
		ct <- time.Now()
	}
	close(ct)
	c.watch(0)
	res := <-head
	r.Equal(uint64(1), res)
}

func TestNew(t *testing.T) {
	r := require.New(t)
	t.Run("FailedToDialChainEndpoint", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(ethclient.Dial, nil, errors.New(t.Name()))

		addr := common.Address{}
		_, err := New(1, "", addr, addr, addr, addr, nil, nil)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToNewProjectContractInstance", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(ethclient.Dial, nil, nil)
		p.ApplyFuncReturn(project.NewProject, nil, errors.New(t.Name()))

		addr := common.Address{}
		_, err := New(1, "", addr, addr, addr, addr, nil, nil)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToNewProverContractInstance", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(ethclient.Dial, nil, nil)
		p.ApplyFuncReturn(project.NewProject, nil, nil)
		p.ApplyFuncReturn(prover.NewProver, nil, errors.New(t.Name()))

		addr := common.Address{}
		_, err := New(1, "", addr, addr, addr, addr, nil, nil)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToList", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(ethclient.Dial, nil, nil)
		p.ApplyFuncReturn(project.NewProject, nil, nil)
		p.ApplyFuncReturn(prover.NewProver, nil, nil)
		p.ApplyPrivateMethod(&Contract{}, "list", func() (uint64, error) { return 0, errors.New(t.Name()) })

		addr := common.Address{}
		_, err := New(1, "", addr, addr, addr, addr, nil, nil)
		r.ErrorContains(err, t.Name())
	})
	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(ethclient.Dial, nil, nil)
		p.ApplyFuncReturn(project.NewProject, nil, nil)
		p.ApplyFuncReturn(prover.NewProver, nil, nil)
		p.ApplyPrivateMethod(&Contract{}, "list", func() (uint64, error) { return 0, nil })
		p.ApplyPrivateMethod(&Contract{}, "watch", func(uint64) {})

		addr := common.Address{}
		_, err := New(1, "", addr, addr, addr, addr, nil, nil)
		time.Sleep(10 * time.Millisecond)

		r.NoError(err)
	})
}
