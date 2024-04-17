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

func (c *Client) init() {
	if c.Metadata == nil {
		c.Metadata = make(map[string]string)
	}
	if c.projects == nil {
		c.projects = make(map[uint64]struct{})
	}
	for _, v := range c.Projects {
		c.projects[v] = struct{}{}
	}
}

func (c *Client) HasProjectPermission(projectID uint64) bool {
	_, ok := c.projects[projectID]
	return ok
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
	m.syncFromContract()
	m.fillByMockClients()
	manager = m
	return manager
}

type Manager struct {
	mux  sync.Mutex
	pool map[string]*Client
}

func (mgr *Manager) ClientByDID(clientdid string) (*Client, bool) {
	mgr.mux.Lock()
	defer mgr.mux.Unlock()
	c, ok := mgr.pool[clientdid]
	return c, ok
}

func (mgr *Manager) AddClient(c *Client) {
	c.init()
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
		c.init()
		mgr.pool[c.ClientDID] = c
	}
}
