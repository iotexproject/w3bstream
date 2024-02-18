package clients

import (
	_ "embed" // embed mock clients configuration
	"encoding/json"
	"sync"
)

var (
	//go:embed clients
	mockClientsConfig []byte
)

type Client struct {
	ClientDID string            `json:"clientDID"`
	Projects  []uint64          `json:"projects"`
	Metadata  map[string]string `json:"metadata,omitempty"`
	projects  map[uint64]struct{}
}

var manager *Manager

func NewManager() *Manager {
	if manager != nil {
		return manager
	}
	m := &Manager{
		mux:  sync.Mutex{},
		pool: make(map[string]*Client),
	}
	m.fillByMockClients()
	manager = m
	return manager
}

type Manager struct {
	mux  sync.Mutex
	pool map[string]*Client
}

func (mgr *Manager) GetByClientDID(clientdid string) (*Client, bool) {
	mgr.mux.Lock()
	defer mgr.mux.Unlock()
	c, ok := mgr.pool[clientdid]
	return c, ok
}

func (mgr *Manager) AddClient(c *Client) {
	mgr.mux.Lock()
	defer mgr.mux.Unlock()
	mgr.pool[c.ClientDID] = c
}

// TODO syncFromContract
func (mgr *Manager) syncFromContract() {}

func (mgr *Manager) fillByMockClients() {
	clients := make([]*Client, 0)
	if err := json.Unmarshal(mockClientsConfig, &clients); err != nil {
		panic(err)
	}
	for _, c := range clients {
		c.projects = make(map[uint64]struct{})
		for _, id := range c.Projects {
			c.projects[id] = struct{}{}
		}
		mgr.pool[c.ClientDID] = c
	}
}
