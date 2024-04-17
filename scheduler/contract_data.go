package scheduler

import (
	"container/list"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/persistence/contract"
)

type contractProject struct {
	mu     sync.Mutex
	epoch  uint64
	blocks *list.List
}

func (c *contractProject) project(projectID, blockNumber uint64) (*contract.Project, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	np := &contract.Project{Attributes: map[common.Hash][]byte{}}

	for e := c.blocks.Front(); e != nil; e = e.Next() {
		ep := e.Value.(*contract.BlockProject)
		if blockNumber < ep.BlockNumber {
			break
		}
		p, ok := ep.Projects[projectID]
		if ok {
			np.Merge(p)
		}
	}
	if np.ID == 0 {
		return nil, errors.Errorf("failed to find project contract data at the block number, project_id %v, expected_block_number %v", projectID, blockNumber)
	}
	return np, nil
}

func (c *contractProject) add(diff *contract.BlockProject) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.blocks.PushBack(diff)

	if uint64(c.blocks.Len()) > c.epoch {
		h := c.blocks.Front()
		np := &contract.BlockProject{Projects: map[uint64]*contract.Project{}}
		np.Merge(h.Value.(*contract.BlockProject))
		np.Merge(h.Next().Value.(*contract.BlockProject))
		c.blocks.Remove(h.Next())
		c.blocks.Remove(h)
		c.blocks.PushFront(np)
	}
}

type contractProver struct {
	mu     sync.Mutex
	epoch  uint64
	blocks *list.List
}

func (c *contractProver) blockProver(blockNumber uint64) *contract.BlockProver {
	c.mu.Lock()
	defer c.mu.Unlock()

	np := &contract.BlockProver{Provers: map[uint64]*contract.Prover{}}

	for e := c.blocks.Front(); e != nil; e = e.Next() {
		ep := e.Value.(*contract.BlockProver)
		if blockNumber < ep.BlockNumber {
			break
		}
		np.Merge(ep)
	}
	return np
}

func (c *contractProver) add(diff *contract.BlockProver) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.blocks.PushBack(diff)

	if uint64(c.blocks.Len()) > c.epoch {
		h := c.blocks.Front()
		np := &contract.BlockProver{Provers: map[uint64]*contract.Prover{}}
		np.Merge(h.Value.(*contract.BlockProver))
		np.Merge(h.Next().Value.(*contract.BlockProver))
		c.blocks.Remove(h.Next())
		c.blocks.Remove(h)
		c.blocks.PushFront(np)
	}
}
