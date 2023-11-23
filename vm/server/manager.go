package server

import (
	"context"
	"sync"
)

type Mgr struct {
	mux  sync.Mutex
	idle map[uint64]*Instance
}

func (m *Mgr) Acquire(projectID uint64, endpoint string, code string, expParam string) (*Instance, error) {
	m.mux.Lock()
	defer m.mux.Unlock()

	if i, ok := m.idle[projectID]; ok {
		return i, nil
	}

	return NewInstance(context.Background(), endpoint, projectID, code, expParam)
}

func (m *Mgr) Release(projectID uint64, i *Instance) {
	m.mux.Lock()
	defer m.mux.Unlock()

	m.idle[projectID] = i
}

func NewMgr() *Mgr {
	return &Mgr{
		idle: make(map[uint64]*Instance),
	}
}
