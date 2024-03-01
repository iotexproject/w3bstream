package project

import (
	"bytes"
	"fmt"
	"log/slog"
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
	notify       chan uint64
	// znodes       []string
	// ioID         string
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
		slog.Warn("project was overwritten", "project_id", projectID, "version", version, "old_version", old.Version)
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

func (m *Manager) GetNotify() <-chan uint64 {
	return m.notify
}

func (m *Manager) doProjectRegistrarWatch(logs <-chan *types.Log, subs event.Subscription) {
	select {
	case err := <-subs.Err():
		slog.Error("project upserted event subscription failed", "err", err)
	case l := <-logs:
		ev, err := m.instance.ParseProjectUpserted(*l)
		if err != nil {
			slog.Error("failed to parse target event", "msg", err)
			return
		}
		if ev.ProjectId == 0 {
			return
		}
		pm := &ProjectMeta{
			ProjectID: ev.ProjectId,
			Uri:       ev.Uri,
			Hash:      ev.Hash,
		}
		cs, err := pm.GetConfigs(m.ipfsEndpoint)
		if err != nil {
			slog.Error("fetch project failed", "err", err)
			return
		}

		for _, c := range cs {
			slog.Info("monitor project", "projectID", pm.ProjectID, "version", c.Version, "vm_type", c.VMType, "code_size", len(c.Code))
			m.Set(pm.ProjectID, c.Version, c)
		}
		select {
		case m.notify <- pm.ProjectID:
		default:
			slog.Info("project notify channel full", "projectID", pm.ProjectID)
		}
	}
}

func (m *Manager) watchProjectRegistrar(logs <-chan *types.Log, subs event.Subscription) {
	for {
		m.doProjectRegistrarWatch(logs, subs)
	}
}

func (m *Manager) doProjectPoolFill(projectID uint64) bool {
	emptyHash := [32]byte{}
	mp, err := m.instance.Projects(nil, projectID)
	if err != nil {
		slog.Error("get project meta from chain failed", "project_id", projectID, "error", err)
		return false
	}
	// query empty, means reached the maximum projectID value
	if mp.Uri == "" || bytes.Equal(mp.Hash[:], emptyHash[:]) {
		slog.Info("project from contract read completed", "max", projectID-1)
		return true
	}

	// znodeMap := map[[sha256.Size]byte]string{}
	// for _, n := range znodes {
	// 	znodeMap[sha256.Sum256([]byte(n))] = n
	// }

	// max := new(big.Int).SetUint64(0)
	// maxZnode := ioID
	// for h, id := range znodeMap {
	// 	n := new(big.Int).Xor(new(big.Int).SetBytes(h[:]), new(big.Int).SetUint64(i))
	// 	if n.Cmp(max) > 0 {
	// 		max = n
	// 		maxZnode = id
	// 	}
	// }
	// if maxZnode != ioID {
	// 	slog.Info("the project not scheduld to this znode", "projectID", i)
	// 	continue
	// }

	pm := &ProjectMeta{
		ProjectID: projectID,
		Uri:       mp.Uri,
		Hash:      mp.Hash,
		Paused:    mp.Paused,
	}
	cs, err := pm.GetConfigs(m.ipfsEndpoint)
	if err != nil {
		slog.Error("fetch project failed", "err", err)
		return false
	}

	for _, c := range cs {
		slog.Debug("contract project loaded", "project_id", pm.ProjectID, "version", c.Version, "vm_type", c.VMType, "code_size", len(c.Code))
		m.Set(projectID, c.Version, c)
	}
	return false
}

func (m *Manager) fillProjectPool() {
	for i := uint64(1); ; i++ {
		if finished := m.doProjectPoolFill(i); finished {
			return
		}
	}
}

func NewManager(chainEndpoint, contractAddress, ipfsEndpoint string) (*Manager, error) {
	m := &Manager{
		pool:         make(map[key]*Config),
		projectIDs:   make(map[uint64]bool),
		ipfsEndpoint: ipfsEndpoint,
		notify:       make(chan uint64, 32),
	}

	client, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		return nil, errors.Wrapf(err, "dial chain endpoint failed, endpoint %s", chainEndpoint)
	}
	m.instance, err = contracts.NewContracts(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, errors.Wrapf(err, "new contract instance failed, endpoint %s, contractAddress %s", chainEndpoint, contractAddress)
	}

	m.fillProjectPool()

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
