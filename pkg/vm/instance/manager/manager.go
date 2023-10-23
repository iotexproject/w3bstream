package manager

import (
	"sync"

	"github.com/machinefi/w3bstream-mainnet/pkg/msg"
	"github.com/machinefi/w3bstream-mainnet/pkg/vm/instance"
)

type Config struct {
	Risc0ServerAddr string
}

type Mgr struct {
	mux  sync.Mutex
	runs map[msg.MsgKey]instance.Instance
	conf *Config
}

func (m *Mgr) Acquire(msg *msg.Msg) (instance.Instance, error) {
	m.mux.Lock()
	defer m.mux.Unlock()

	if i, ok := m.runs[msg.Key()]; ok {
		return i, nil
	}

	return nil, nil

}

func (m *Mgr) Release(key msg.MsgKey, i instance.Instance) {
	m.mux.Lock()
	defer m.mux.Unlock()

	m.runs[key] = i
}

func NewMgr(conf *Config) *Mgr {
	return &Mgr{
		runs: make(map[msg.MsgKey]instance.Instance),
		conf: conf,
	}
}
