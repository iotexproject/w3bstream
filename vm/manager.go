package vm

import (
	"context"
	"sync"
)

type manager struct {
	mux  sync.Mutex
	idle map[uint64]*instance
}

func (m *manager) acquire(projectID uint64, endpoint string, code string, expParam string) (*instance, error) {
	m.mux.Lock()
	defer m.mux.Unlock()

	if i, ok := m.idle[projectID]; ok {
		return i, nil
	}

	return newInstance(context.Background(), endpoint, projectID, code, expParam)
}

func (m *manager) release(projectID uint64, i *instance) {
	m.mux.Lock()
	defer m.mux.Unlock()

	m.idle[projectID] = i
}

func newManager() *manager {
	return &manager{
		idle: make(map[uint64]*instance),
	}
}
