package scheduler

import (
	"container/list"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/utils/contract"
)

type contractProjects struct {
	mu    sync.Mutex
	epoch uint64
	datas *list.List
}

func (c *contractProjects) get(projectID, expectedBlockNumber uint64) (*contract.Project, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	np := &contract.Project{Attributes: map[common.Hash][]byte{}}

	for e := c.datas.Front(); e != nil; e = e.Next() {
		ep := e.Value.(*contract.Projects)
		if expectedBlockNumber < ep.BlockNumber {
			break
		}
		p, ok := ep.Projects[projectID]
		if ok {
			np.Merge(p)
		}
	}
	if np.ID == 0 {
		return nil, errors.Errorf("failed to find project contract data at the block number, project_id %v, expected_block_number %v", projectID, expectedBlockNumber)
	}
	return np, nil
}

func (c *contractProjects) set(diff *contract.Projects) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.datas.PushBack(diff)

	if c.datas.Len() > int(c.epoch) {
		h := c.datas.Front()
		h.Value.(*contract.Projects).Merge(h.Next().Value.(*contract.Projects))
		c.datas.Remove(h.Next())
	}
}

type contractProvers struct {
	mu    sync.Mutex
	epoch uint64
	datas *list.List
}

func (c *contractProvers) get(expectedBlockNumber uint64) *contract.Provers {
	c.mu.Lock()
	defer c.mu.Unlock()

	np := &contract.Provers{Provers: map[uint64]*contract.Prover{}}

	for e := c.datas.Front(); e != nil; e = e.Next() {
		ep := e.Value.(*contract.Provers)
		if expectedBlockNumber < ep.BlockNumber {
			break
		}
		np.Merge(ep)
	}
	return np
}

func (c *contractProvers) set(diff *contract.Provers) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.datas.PushBack(diff)
	if c.datas.Len() > int(c.epoch) {
		h := c.datas.Front()
		h.Value.(*contract.Provers).Merge(h.Next().Value.(*contract.Provers))
		c.datas.Remove(h.Next())
	}
}
