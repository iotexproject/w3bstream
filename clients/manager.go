package clients

import (
	"encoding/json"
	"log/slog"
	"os"
	"sync"

	"github.com/pkg/errors"
)

type Client struct {
	ClientDID string            `json:"clientDID"`
	Projects  []uint64          `json:"projects"`
	Metadata  map[string]string `json:"metadata,omitempty"`
	projects  map[uint64]struct{}
}

var manager *Manager

func NewManager(confPath string) *Manager {
	if manager != nil {
		return manager
	}
	m := &Manager{
		mux:  sync.Mutex{},
		pool: make(map[string]*Client),
	}
	if err := m.syncFromLocal(confPath); err != nil {
		slog.Error("failed to sync clients from local", "msg", err)
	}
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

func (mgr *Manager) syncFromLocal(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(errors.Wrap(err, "failed to read local config"))
	}
	clients := make([]*Client, 0)
	if err = json.Unmarshal(content, &clients); err != nil {
		return errors.Wrap(err, "failed to parse local config")
	}
	for _, c := range clients {
		c.projects = make(map[uint64]struct{})
		for _, id := range c.Projects {
			c.projects[id] = struct{}{}
		}
		mgr.pool[c.ClientDID] = c
	}
	return nil
}
