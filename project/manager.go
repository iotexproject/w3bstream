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

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/project/contracts"
)

type Manager struct {
	mux             sync.Mutex
	pool            map[uint64]*Project
	chainEndpoint   string
	contractAddress string
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
			p, err := pm.GetProject()
			if err != nil {
				slog.Error("fetch project failed", "err", err)
				continue
			}

			m.mux.Lock()
			m.pool[p.ID] = p
			m.mux.Unlock()
		}
	}
}

func fillProjectPoolFromLocal(pool map[uint64]*Project, projectFileDirectory string) error {
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

		c := Config{}
		if err := json.Unmarshal(data, &c); err != nil {
			return errors.Wrapf(err, "unmarshal config file %s failed", f.Name())
		}

		p := Project{
			ID:     projectID,
			Config: c,
		}

		pool[p.ID] = &p
	}
	return nil
}

func fillProjectPoolFromChain(pool map[uint64]*Project, instance *contracts.Contracts) error {
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
		p, err := m.GetProject()
		if err != nil {
			slog.Error("fetch project failed", "err", err)
			continue
		}
		pool[p.ID] = p
	}
}

func NewManager(chainEndpoint, contractAddress, projectFileDirectory string) (*Manager, error) {
	pool := make(map[uint64]*Project)

	// client, err := ethclient.Dial(chainEndpoint)
	// if err != nil {
	// 	return nil, errors.Wrapf(err, "dial chain endpoint failed, endpoint %s", chainEndpoint)
	// }
	// instance, err := contracts.NewContracts(common.HexToAddress(contractAddress), client)
	// if err != nil {
	// 	return nil, errors.Wrapf(err, "new contract instance failed, endpoint %s, contractAddress %s", chainEndpoint, contractAddress)
	// }

	// if err := fillProjectPoolFromChain(pool, instance); err != nil {
	// 	return nil, errors.Wrap(err, "read project file from chain failed")
	// }
	if err := fillProjectPoolFromLocal(pool, projectFileDirectory); err != nil {
		return nil, errors.Wrap(err, "read project file from local failed")
	}

	m := &Manager{
		pool:            pool,
		chainEndpoint:   chainEndpoint,
		contractAddress: contractAddress,
	}

	// events := make(chan *contracts.ContractsProjectUpserted)
	// subs, err := instance.WatchProjectUpserted(&bind.WatchOpts{}, events, nil)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "watch project upserted event failed")
	// }
	// go m.watchProjectRegistrar(events, subs)

	return m, nil
}
