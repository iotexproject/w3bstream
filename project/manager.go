package project

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/project/contracts"
)

type Manager struct {
	mux          sync.Mutex
	pool         map[key]*Config
	projectIDs   map[uint64]bool
	instance     *contracts.Contracts
	ipfsEndpoint string
}

type key string

func getKey(projectID uint64, version string) key {
	return key(fmt.Sprintf("%d_%s", projectID, version))
}

func (m *Manager) Get(projectID uint64, version string) (*Config, error) {
	m.mux.Lock()
	defer m.mux.Unlock()

	if p, ok := m.pool[getKey(projectID, version)]; ok {
		return p, nil
	}
	return nil, errors.Errorf("project config not exist, projectID %d, version %s", projectID, version)
}

func (m *Manager) Set(projectID uint64, version string, c *Config) {
	m.mux.Lock()
	defer m.mux.Unlock()

	key := getKey(projectID, version)
	old, ok := m.pool[key]
	m.pool[key] = c
	if ok {
		slog.Warn("project was overwritten", "project_id", projectID, "version", version, "old", old, "new", c)
	}
	m.projectIDs[projectID] = true
}

// TODO will delete when node konw how to fetch message
func (m *Manager) GetAllProjectID() []uint64 {
	m.mux.Lock()
	defer m.mux.Unlock()

	ids := []uint64{}
	for id := range m.projectIDs {
		ids = append(ids, id)
	}
	return ids
}

func (m *Manager) watchProjectRegistrar(logs <-chan *types.Log, subs event.Subscription) {
	for {
		select {
		case err := <-subs.Err():
			slog.Error("project upserted event subscription failed", "err", err)
		case l := <-logs:
			ev, err := m.instance.ParseProjectUpserted(*l)
			if err != nil {
				slog.Error("failed to parse target event", "msg", err)
				continue
			}
			if ev.ProjectId == 0 {
				continue
			}
			pm := &ProjectMeta{
				ProjectID: ev.ProjectId,
				Uri:       ev.Uri,
				Hash:      ev.Hash,
			}
			cs, err := pm.GetConfigs(m.ipfsEndpoint)
			if err != nil {
				slog.Error("fetch project failed", "err", err)
				continue
			}

			for _, c := range cs {
				slog.Info("monitor project", "project_id", pm.ProjectID, "version", c.Version, "vm_type", c.VMType, "code_size", len(c.Code))
				m.Set(pm.ProjectID, c.Version, c)
			}
		}
	}
}

func (m *Manager) fillProjectPoolFromLocal(projectFileDirectory string) {
	files, err := os.ReadDir(projectFileDirectory)
	if err != nil {
		if !os.IsNotExist(err) {
			slog.Warn("failed to read project directory", "path", projectFileDirectory, "msg", err)
			return
		}
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		data, err := os.ReadFile(path.Join(projectFileDirectory, f.Name()))
		if err != nil {
			slog.Error("failed to read project config", "filename", f.Name(), "msg", err)
			continue
		}

		projectID, err := strconv.ParseUint(f.Name(), 10, 64)
		if err != nil {
			slog.Error("failed to parse filename", "filename", f.Name())
			continue
		}

		cs := []*Config{}
		if err := json.Unmarshal(data, &cs); err != nil {
			slog.Error("failed to parse project config", "filename", f.Name())
			continue
		}

		for _, c := range cs {
			slog.Info("local project loaded", "project_id", projectID, "version", c.Version, "vm_type", c.VMType, "code_size", len(c.Code))
			m.Set(projectID, c.Version, c)
		}
	}
}

func (m *Manager) fillProjectPoolFromChain() {
	emptyHash := [32]byte{}
	for i := uint64(1); ; i++ {
		mp, err := m.instance.Projects(nil, i)
		if err != nil {
			slog.Error("failed to get project meta from chain", "project_id", i, "msg", err)
			continue
		}
		// query empty, means reached the maximum projectID value
		if mp.Uri == "" || bytes.Equal(mp.Hash[:], emptyHash[:]) {
			slog.Info("project from contract read completed", "max", i-1)
			return
		}

		pm := &ProjectMeta{
			ProjectID: i,
			Uri:       mp.Uri,
			Hash:      mp.Hash,
			Paused:    mp.Paused,
		}
		cs, err := pm.GetConfigs(m.ipfsEndpoint)
		if err != nil {
			slog.Error("fetch project failed", "err", err)
			continue
		}

		for _, c := range cs {
			slog.Info("contract project loaded", "project_id", pm.ProjectID, "version", c.Version, "vm_type", c.VMType, "code_size", len(c.Code))
			m.Set(i, c.Version, c)
		}
	}
}

func (m *Manager) loadProjects(projectFileLocalDir string) {
	m.fillProjectPoolFromLocal(projectFileLocalDir)
	m.fillProjectPoolFromChain()
}

func NewManager(chainEndpoint, contractAddress, projectFileDirectory, ipfsEndpoint string) (*Manager, error) {
	m := &Manager{
		mux:          sync.Mutex{},
		pool:         make(map[key]*Config),
		projectIDs:   make(map[uint64]bool),
		ipfsEndpoint: ipfsEndpoint,
	}

	client, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		return nil, errors.Wrapf(err, "dial chain endpoint failed, endpoint %s", chainEndpoint)
	}
	m.instance, err = contracts.NewContracts(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, errors.Wrapf(err, "new contract instance failed, endpoint %s, contractAddress %s", chainEndpoint, contractAddress)
	}

	go m.loadProjects(projectFileDirectory)

	topic := "ProjectUpserted(uint64,string,bytes32)"
	monitor, err := NewDefaultMonitor(
		chainEndpoint,
		[]string{contractAddress},
		[]string{topic},
	)
	if err != nil {
		slog.Error("failed to new contract monitor", "msg", err)
		return nil, err
	}
	go monitor.run()
	go m.watchProjectRegistrar(monitor.MustEvents(topic), monitor)

	return m, nil
}
