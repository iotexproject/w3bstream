package manager

import (
	"context"
	"sync"

	"github.com/machinefi/w3bstream-mainnet/msg"
	"github.com/machinefi/w3bstream-mainnet/vm/instance"
	"github.com/machinefi/w3bstream-mainnet/vm/risc0"
)

type Config struct {
	Risc0ServerAddr       string
	ProjectConfigFilePath string
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

	// TODO get project bin data by real project info
	testdata := getTestData(m.conf.ProjectConfigFilePath)
	return risc0.NewInstance(context.Background(), m.conf.Risc0ServerAddr, msg.Key(), testdata.Content, testdata.ExpParam)
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
