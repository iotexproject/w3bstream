package project

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"log/slog"
	"os"
	"path"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
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

			m.mux.Lock()

			for _, c := range cs {
				m.pool[getKey(pm.ProjectID, c.Version)] = c
			}
			m.projectIDs[pm.ProjectID] = true

			m.mux.Unlock()
		}
	}
}

func (m *Manager) fillProjectPoolFromLocal(projectFileDirectory string) error {
	files, err := os.ReadDir(projectFileDirectory)
	if err != nil {
		if !os.IsNotExist(err) {
			return errors.Wrap(err, "read project file directory failed")
		}
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		data, err := os.ReadFile(path.Join(projectFileDirectory, f.Name()))
		if err != nil {
			return errors.Wrapf(err, "read project file %s failed", f.Name())
		}

		projectID, err := strconv.ParseUint(f.Name(), 10, 64)
		if err != nil {
			return errors.Wrapf(err, "parse file name %s to projectID failed", f.Name())
		}

		cs := []*Config{}
		if err := json.Unmarshal(data, &cs); err != nil {
			return errors.Wrapf(err, "unmarshal config file %s failed", f.Name())
		}

		for _, c := range cs {
			m.pool[getKey(projectID, c.Version)] = c
		}
		m.projectIDs[projectID] = true
	}
	return nil
}

func (m *Manager) fillProjectPoolFromChain() error {
	emptyHash := [32]byte{}
	for i := uint64(1); ; i++ {
		mp, err := m.instance.Projects(nil, i)
		if err != nil {
			return errors.Wrapf(err, "get project meta from chain failed, projectID %d", i)
		}
		// query empty, means reached the maximum projectID value
		if mp.Uri == "" || bytes.Equal(mp.Hash[:], emptyHash[:]) {
			return nil
		}
		slog.Debug("queried project", "project_id", i, "uri", mp.Uri)
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
			m.pool[getKey(pm.ProjectID, c.Version)] = c
		}
		m.projectIDs[pm.ProjectID] = true
	}
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

	if err = m.fillProjectPoolFromChain(); err != nil {
		return nil, errors.Wrap(err, "read project file from chain failed")
	}
	if err = m.fillProjectPoolFromLocal(projectFileDirectory); err != nil {
		return nil, errors.Wrap(err, "read project file from local failed")
	}

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
