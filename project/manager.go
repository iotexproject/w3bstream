package project

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/project/contracts"
)

type Manager struct {
	mux             sync.Mutex
	pool            map[uint64]*Project
	chainEndpoint   string
	contractAddress string
}

func NewManager(chainEndpoint, contractAddress, projectFileDirectory string) (*Manager, error) {
	pool := make(map[uint64]*Project)
	files, err := os.ReadDir(projectFileDirectory)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, errors.Wrap(err, "read project file directory failed")
		}
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		data, err := os.ReadFile(path.Join(projectFileDirectory, f.Name()))
		if err != nil {
			return nil, errors.Wrapf(err, "read project file %s failed", f.Name())
		}

		projectID, err := strconv.ParseUint(f.Name(), 10, 64)
		if err != nil {
			return nil, errors.Wrapf(err, "parse file name %s to projectID failed", f.Name())
		}

		c := Config{}
		if err := json.Unmarshal(data, &c); err != nil {
			return nil, errors.Wrapf(err, "unmarshal config file %s failed", f.Name())
		}

		p := Project{
			ID:     projectID,
			Config: c,
		}

		pool[p.ID] = &p
	}
	m := &Manager{
		pool:            pool,
		chainEndpoint:   chainEndpoint,
		contractAddress: contractAddress,
	}

	metas, err := m.syncProjects()
	if err != nil {
		return nil, errors.Wrap(err, "sync project configs from contract failed")
	}

	for _, meta := range metas {
		if meta.Paused {
			slog.Debug("project paused", "project_id", meta.ProjectID)
			continue
		}
		c, err := meta.ProjectConfig()
		if err != nil {
			slog.Error(err.Error())
			continue
		}
		pool[meta.ProjectID] = &Project{
			ID:     meta.ProjectID,
			Config: *c,
		}
	}

	// start monitor contract
	go func() {
		err := m.monitorContract()
		slog.Error(err.Error())
		os.Exit(1)
	}()

	return m, nil
}

func (m *Manager) Get(projectID uint64) (*Project, error) {
	m.mux.Lock()
	defer m.mux.Unlock()

	if p, ok := m.pool[projectID]; ok {
		return p, nil
	}

	client, err := ethclient.Dial(m.chainEndpoint)
	if err != nil {
		return nil, errors.Wrap(err, "dial chain endpoint failed")
	}
	address := common.HexToAddress(m.contractAddress)
	instance, err := contracts.NewContracts(address, client)
	if err != nil {
		return nil, errors.Wrap(err, "new contracts instance failed")
	}
	p, err := instance.Projects(nil, projectID)
	if err != nil {
		return nil, errors.Wrap(err, "get project from contracts failed")
	}

	if p.Uri == "" {
		return nil, errors.New("project not exist")
	}

	slog.Debug("get project file uri", "projectID", projectID, "uri", p.Uri)

	resp, err := http.Get(p.Uri)
	if err != nil {
		return nil, errors.Wrap(err, "get project config file failed")
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read project config file failed")
	}

	// TODO hash check

	np := Project{ID: projectID}
	if err := json.Unmarshal(data, &np.Config); err != nil {
		return nil, errors.Wrap(err, "unmarshal project config file failed")
	}
	m.pool[projectID] = &np
	return &np, nil
}

// TODO will delete when node konw how to fetch message
func (m *Manager) GetAllProjectID() []uint64 {
	m.mux.Lock()
	defer m.mux.Unlock()

	ids := []uint64{}
	for id := range m.pool {
		ids = append(ids, id)
	}
	return ids
}

func (m *Manager) syncProjects() ([]*ProjectMeta, error) {
	client, err := ethclient.Dial(m.chainEndpoint)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial chain endpoint")
	}

	address := common.HexToAddress(m.contractAddress)
	instance, err := contracts.NewContracts(address, client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to bind contract")
	}

	emptyHash := [32]byte{}
	projects := make([]*ProjectMeta, 0)
	for i := uint64(1); ; i++ {
		p, err := instance.Projects(nil, i)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get project meta: %d", i)
		}
		// query empty, sync completed
		if p.Uri == "" && bytes.Equal(p.Hash[:], emptyHash[:]) && !p.Paused {
			break
		}
		projects = append(projects, &ProjectMeta{
			ProjectID: i,
			Uri:       p.Uri,
			Hash:      p.Hash,
			Paused:    p.Paused,
		})
	}

	return projects, nil
}

func (m *Manager) monitorContract() error {
	client, err := ethclient.Dial(m.chainEndpoint)
	if err != nil {
		return errors.Wrap(err, "failed to dial chain endpoint")
	}

	instance, err := contracts.NewContracts(common.HexToAddress(m.contractAddress), client)
	if err != nil {
		return errors.Wrap(err, "failed to create contract instance")
	}

	sink := make(chan *contracts.ContractsProjectUpserted)
	subs, err := instance.WatchProjectUpserted(&bind.WatchOpts{}, sink, nil)
	if err != nil {
		return errors.Wrap(err, "failed to watch event")
	}
	defer subs.Unsubscribe()

	for {
		select {
		case <-subs.Err():
			return errors.Wrap(err, "subscription canceled")
		case ev := <-sink:
			if ev.ProjectId == 0 {
				slog.Debug("invalid project id")
				continue
			}
			conf, err := (&ProjectMeta{
				ProjectID: ev.ProjectId,
				Uri:       ev.Uri,
				Hash:      ev.Hash,
				Paused:    true, // TODO if project can be upsert. how to confirm the state of current project?
			}).ProjectConfig()
			if err != nil {
				slog.Error("fetch project config failed", "err", err)
				continue
			}
			m.Upsert(ev.ProjectId, conf)
		}
	}
}

func (m *Manager) Upsert(projectID uint64, c *Config) {
	m.mux.Lock()
	defer m.mux.Unlock()
	m.pool[projectID] = &Project{
		ID:     projectID,
		Config: *c,
	}
}
