package project

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path"
	"strconv"
	"sync"

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
			return nil, errors.Wrapf(err, "unmarshal config file %s failed", f.Name())
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

	return &Manager{
		pool:            pool,
		chainEndpoint:   chainEndpoint,
		contractAddress: contractAddress,
	}, nil
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
