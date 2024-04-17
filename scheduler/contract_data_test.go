package scheduler

import (
	"container/list"
	"testing"

	"github.com/machinefi/sprout/persistence/contract"
	"github.com/stretchr/testify/require"
)

func TestContractProject(t *testing.T) {
	r := require.New(t)

	cp := &contractProject{
		epoch:  2,
		blocks: list.New(),
	}

	cp.add(&contract.BlockProject{
		BlockNumber: 100,
		Projects: map[uint64]*contract.Project{
			1: {
				ID:  1,
				Uri: "uri",
			},
		},
	})
	_, err := cp.project(1, 99)
	r.Error(err)
	p, err := cp.project(1, 100)
	r.NoError(err)
	r.Equal(p.ID, uint64(1))

	cp.add(&contract.BlockProject{
		BlockNumber: 101,
		Projects: map[uint64]*contract.Project{
			1: {
				ID:  1,
				Uri: "uri1",
			},
		},
	})
	p, err = cp.project(1, 101)
	r.NoError(err)
	r.Equal(p.ID, uint64(1))
	r.Equal(p.Uri, "uri1")

	p, err = cp.project(1, 100)
	r.NoError(err)
	r.Equal(p.ID, uint64(1))
	r.Equal(p.Uri, "uri")

	cp.add(&contract.BlockProject{
		BlockNumber: 102,
		Projects: map[uint64]*contract.Project{
			2: {
				ID:  2,
				Uri: "uri2",
			},
			1: {
				ID:  1,
				Uri: "uri2",
			},
		},
	})
	p, err = cp.project(2, 102)
	r.NoError(err)
	r.Equal(p.ID, uint64(2))
	r.Equal(p.Uri, "uri2")
	r.Equal(uint64(cp.blocks.Len()), cp.epoch)

	p, err = cp.project(1, 102)
	r.NoError(err)
	r.Equal(p.ID, uint64(1))
	r.Equal(p.Uri, "uri2")

	p, err = cp.project(1, 101)
	r.NoError(err)
	r.Equal(p.ID, uint64(1))
	r.Equal(p.Uri, "uri1")

	_, err = cp.project(1, 100)
	r.Error(err)

	cp.add(&contract.BlockProject{
		BlockNumber: 105,
		Projects: map[uint64]*contract.Project{
			1: {
				ID:  1,
				Uri: "uri1",
			},
		},
	})
	r.Equal(uint64(cp.blocks.Len()), cp.epoch)
}

func TestContractProver(t *testing.T) {
	r := require.New(t)

	cp := &contractProver{
		epoch:  2,
		blocks: list.New(),
	}

	cp.add(&contract.BlockProver{
		BlockNumber: 100,
		Provers: map[uint64]*contract.Prover{
			1: {
				ID:        1,
				NodeTypes: 1,
			},
		},
	})
	ps := cp.blockProver(99)
	r.Equal(len(ps.Provers), 0)
	ps = cp.blockProver(100)
	r.Equal(len(ps.Provers), 1)
	r.Equal(ps.Provers[1].NodeTypes, uint64(1))

	cp.add(&contract.BlockProver{
		BlockNumber: 101,
		Provers: map[uint64]*contract.Prover{
			1: {
				ID:        1,
				NodeTypes: 2,
			},
		},
	})
	ps = cp.blockProver(100)
	r.Equal(len(ps.Provers), 1)
	r.Equal(ps.Provers[1].NodeTypes, uint64(1))
	ps = cp.blockProver(101)
	r.Equal(len(ps.Provers), 1)
	r.Equal(ps.Provers[1].NodeTypes, uint64(2))

	cp.add(&contract.BlockProver{
		BlockNumber: 102,
		Provers: map[uint64]*contract.Prover{
			1: {
				ID:        1,
				NodeTypes: 3,
			},
			2: {
				ID:        2,
				NodeTypes: 1,
			},
		},
	})
	ps = cp.blockProver(100)
	r.Equal(len(ps.Provers), 0)
	ps = cp.blockProver(101)
	r.Equal(len(ps.Provers), 1)
	ps = cp.blockProver(102)
	r.Equal(len(ps.Provers), 2)
	r.Equal(ps.Provers[1].NodeTypes, uint64(3))
	r.Equal(ps.Provers[2].NodeTypes, uint64(1))
	r.Equal(uint64(cp.blocks.Len()), cp.epoch)

	cp.add(&contract.BlockProver{
		BlockNumber: 103,
		Provers: map[uint64]*contract.Prover{
			1: {
				ID:        1,
				NodeTypes: 2,
			},
		},
	})
	r.Equal(uint64(cp.blocks.Len()), cp.epoch)
}
