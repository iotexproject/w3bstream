package clients

import (
	"sync"
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
	clients := []*Client{
		{
			ClientDID: "did:ethr:0x9d9250fb4e08ba7a858fe7196a6ba946c6083ff0",
			Projects: []uint64{
				1,
				2,
				3,
				4,
				5,
				6,
				7,
				8,
				9,
				10,
				11,
				12,
				13,
				14,
				15,
				16,
				17,
				18,
				19,
				20,
			},
		},
		{
			ClientDID: "did:key:z6MkeeChrUs1EoKkNNzoy9FwJJb9gNQ92UT8kcXZHMbwj67B",
			Projects: []uint64{
				1,
				2,
				3,
				4,
				5,
				6,
				7,
				8,
				9,
				10,
				11,
				12,
				13,
				14,
				15,
				16,
				17,
				18,
				19,
				20,
			},
		},
		{
			ClientDID: "did:example:d23dd687a7dc6787646f2eb98d0",
			Projects: []uint64{
				1,
				2,
				3,
				4,
				5,
				6,
				7,
				8,
				9,
				10,
				11,
				12,
				13,
				14,
				15,
				16,
				17,
				18,
				19,
				20,
			},
		},
	}
	for _, c := range clients {
		c.projects = make(map[uint64]struct{})
		for _, id := range c.Projects {
			c.projects[id] = struct{}{}
		}
		mgr.pool[c.ClientDID] = c
	}
}
