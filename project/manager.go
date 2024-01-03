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

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/project/contracts"
)

type Manager struct {
	mux             sync.Mutex
	pool            map[key]*Config
	chainEndpoint   string
	contractAddress string
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
	return nil
}

func (m *Manager) watchProjectRegistrar(events chan *contracts.ContractsProjectUpserted, subs event.Subscription) {
	for {
		select {
		case err := <-subs.Err():
			slog.Error("project upserted event subscription failed", "err", err)
		case ev := <-events:
			if ev.ProjectId == 0 {
				continue
			}
			pm := &ProjectMeta{
				ProjectID: ev.ProjectId,
				Uri:       ev.Uri,
				Hash:      ev.Hash,
			}
			cs, err := pm.GetConfigs()
			if err != nil {
				slog.Error("fetch project failed", "err", err)
				continue
			}

			m.mux.Lock()

			for _, c := range cs {
				m.pool[getKey(pm.ProjectID, c.Version)] = c
			}

			m.mux.Unlock()
		}
	}
}

func fillProjectPoolFromLocal(pool map[key]*Config, projectFileDirectory string) error {
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
			pool[getKey(projectID, c.Version)] = c
		}
	}
	return nil
}

func fillProjectPoolFromChain(pool map[key]*Config, instance *contracts.Contracts) error {
	emptyHash := [32]byte{}
	for i := uint64(1); ; i++ {
		mp, err := instance.Projects(nil, i)
		if err != nil {
			return errors.Wrapf(err, "get project meta from chain failed, projectID %d", i)
		}
		// query empty, means reached the maximum projectID value
		if mp.Uri == "" || bytes.Equal(mp.Hash[:], emptyHash[:]) {
			return nil
		}
		m := &ProjectMeta{
			ProjectID: i,
			Uri:       mp.Uri,
			Hash:      mp.Hash,
			Paused:    mp.Paused,
		}
		cs, err := m.GetConfigs()
		if err != nil {
			slog.Error("fetch project failed", "err", err)
			continue
		}

		for _, c := range cs {
			pool[getKey(m.ProjectID, c.Version)] = c
		}
	}
}

func NewManager(chainEndpoint, contractAddress, projectFileDirectory string) (*Manager, error) {
	pool := make(map[key]*Config)

	client, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		return nil, errors.Wrapf(err, "dial chain endpoint failed, endpoint %s", chainEndpoint)
	}
	instance, err := contracts.NewContracts(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, errors.Wrapf(err, "new contract instance failed, endpoint %s, contractAddress %s", chainEndpoint, contractAddress)
	}

	if err := fillProjectPoolFromChain(pool, instance); err != nil {
		return nil, errors.Wrap(err, "read project file from chain failed")
	}
	if err := fillProjectPoolFromLocal(pool, projectFileDirectory); err != nil {
		return nil, errors.Wrap(err, "read project file from local failed")
	}

	m := &Manager{
		pool:            pool,
		chainEndpoint:   chainEndpoint,
		contractAddress: contractAddress,
	}

	events := make(chan *contracts.ContractsProjectUpserted)
	subs, err := instance.WatchProjectUpserted(&bind.WatchOpts{}, events, nil)
	if err != nil {
		return nil, errors.Wrap(err, "watch project upserted event failed")
	}
	go m.watchProjectRegistrar(events, subs)

	return m, nil
}
