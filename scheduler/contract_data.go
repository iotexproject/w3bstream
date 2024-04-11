package scheduler

import (
	"container/list"
	"sync"

	"github.com/machinefi/sprout/utils/contract"
	"github.com/pkg/errors"
)

type contractProjects struct {
	mu            sync.Mutex
	epoch         uint64
	projects      map[uint64]*contract.Project
	projectEvents *list.List
}

func (c *contractProjects) get(projectID, askedBlockNumber uint64) (*contract.Project, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	p, ok := c.projects[projectID]
	if !ok {
		return nil, errors.New("project not exist in contract")
	}
	if askedBlockNumber < p.BlockNumber {
		return nil, errors.New("illegal blockNumber")
	}

	np := &contract.Project{}
	np.Merge(p)
	if np.BlockNumber == askedBlockNumber {
		return np, nil
	}

	for e := c.projectEvents.Front(); e != nil; e = e.Next() {
		ep := e.Value.(*contract.Project)
		np.Merge(ep)
		if np.BlockNumber == askedBlockNumber {
			return np, nil
		}
	}
	return nil, errors.Errorf("failed to find project contract data at the block number, project_id %v, asked_block_number %v, max_block_number %v", projectID, askedBlockNumber, np.BlockNumber)
}

func (c *contractProjects) set(diff *contract.Project) {
	c.mu.Lock()
	defer c.mu.Unlock()

	p, ok := c.projects[diff.ID]
	if !ok {
		c.projects[diff.ID] = diff
		return
	}
	c.projectEvents.PushBack(diff)
	if c.projectEvents.Len() > (int(c.epoch) - 1) {
		h := c.projectEvents.Front()
		p.Merge(h.Value.(*contract.Project))
		c.projectEvents.Remove(h)
	}
}

type contractProvers struct {
	mu           sync.Mutex
	epoch        uint64
	proverEvents *list.List
}

func (c *contractProvers) get(askedBlockNumber uint64) (*contract.Prover, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	p := &contract.Prover{}
	for e := c.proverEvents.Front(); e != nil; e = e.Next() {
		ep := e.Value.(*contract.Prover)
		p.Merge(ep)
		if p.BlockNumber == askedBlockNumber {
			return p, nil
		}
	}
	return nil, errors.Errorf("failed to find prover contract data at the block number, asked_block_number %v, max_block_number %v", askedBlockNumber, p.BlockNumber)
}

func (c *contractProvers) set(diff *contract.Prover) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.proverEvents.PushBack(diff)
	if c.proverEvents.Len() > (int(c.epoch) - 1) {
		h := c.proverEvents.Front()
		h.Value.(*contract.Project).Merge(h.Next().Value.(*contract.Project))
		c.proverEvents.Remove(h.Next())
	}
}
