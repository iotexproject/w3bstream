package server

import (
	"context"
	"sync"

	"github.com/machinefi/w3bstream-mainnet/msg"
)

type Mgr struct {
	mux  sync.Mutex
	idle map[msg.MsgKey]*Instance
}

func (m *Mgr) Acquire(msg *msg.Msg, serverAddr string, code []byte, expParam string) (*Instance, error) {
	m.mux.Lock()
	defer m.mux.Unlock()

	if i, ok := m.idle[msg.Key()]; ok {
		return i, nil
	}

	return NewInstance(context.Background(), serverAddr, msg.Key(), code, expParam)
}

func (m *Mgr) Release(key msg.MsgKey, i *Instance) {
	m.mux.Lock()
	defer m.mux.Unlock()

	m.idle[key] = i
}

func NewMgr() *Mgr {
	return &Mgr{
		idle: make(map[msg.MsgKey]*Instance),
	}
}
