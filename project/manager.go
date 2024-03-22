package project

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"
	"os"
	"path"
	"sort"
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
	pool         map[key]*Project
	projectIDs   map[uint64]bool
	instance     *contracts.Contracts
	ipfsEndpoint string
	notify       chan uint64
	cache        *cache   // optional
	znodes       []string // optional
	ioID         string   // optional
}

type key string

func getKey(projectID uint64, version string) key {
	return key(fmt.Sprintf("%d_%s", projectID, version))
}

func (m *Manager) Get(projectID uint64, version string) (*Project, error) {
	m.mux.Lock()
	defer m.mux.Unlock()

	if p, ok := m.pool[getKey(projectID, version)]; ok {
		return p, nil
	}
	return nil, errors.Errorf("project config not exist, projectID %d, version %s", projectID, version)
}

func (m *Manager) set(projectID uint64, version string, c *Config, provers []string) {
	m.mux.Lock()
	defer m.mux.Unlock()

	m.pool[getKey(projectID, version)] = &Project{Config: c, Provers: provers}
	m.projectIDs[projectID] = true
}

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
		data, err := pm.GetConfigData(m.ipfsEndpoint)
		if err != nil {
			slog.Error("failed to fetch project", "error", err, "project_id", ev.ProjectId)
			return
		}
		cs, err := convertConfigs(data)
		if err != nil {
			slog.Error("failed to convert project configs", "error", err, "project_id", ev.ProjectId)
			return
		}
		for _, c := range cs {
			slog.Info("monitor project", "project_id", pm.ProjectID, "version", c.Version, "vm_type", c.VMType, "code_size", len(c.Code))
			m.set(pm.ProjectID, c.Version, c, nil)
		}
		if m.cache != nil {
			m.cache.set(ev.ProjectId, data)
		}

		select {
		case m.notify <- pm.ProjectID:
		default:
			slog.Info("project notify channel full", "project_id", pm.ProjectID)
		}
	}
}

func (m *Manager) watchProjectRegistrar(logs <-chan *types.Log, subs event.Subscription) {
	for {
		m.doProjectRegistrarWatch(logs, subs)
	}
}

type distance struct {
	distance *big.Int
	hash     [sha256.Size]byte
}

func (m *Manager) fillProjectPoolFromContract() {
	for projectID := uint64(1); ; projectID++ {
		emptyHash := [32]byte{}
		mp, err := m.instance.Projects(nil, projectID)
		if err != nil {
			slog.Error("failed to get project meta from chain ", "project_id", projectID, "error", err)
			continue
		}
		// query empty, means reached the maximum projectID value
		if mp.Uri == "" || bytes.Equal(mp.Hash[:], emptyHash[:]) {
			slog.Info("load project from contract completed", "max project_id", projectID-1)
			return
		}

		pm := &ProjectMeta{
			ProjectID: projectID,
			Uri:       mp.Uri,
			Hash:      mp.Hash,
			Paused:    mp.Paused,
		}

		var data []byte
		cached := true
		if m.cache != nil {
			data = m.cache.get(projectID, mp.Hash[:])
		}
		if len(data) == 0 {
			cached = false
			data, err = pm.GetConfigData(m.ipfsEndpoint)
			if err != nil {
				slog.Error("failed to fetch project", "error", err, "project_id", projectID)
				continue
			}
		}
		if !cached && m.cache != nil {
			m.cache.set(projectID, data)
		}

		cs, err := convertConfigs(data)
		if err != nil {
			slog.Error("failed to convert project config", "error", err, "project_id", projectID)
			continue
		}

		var provers []string

		if m.ioID != "" {
			c := cs[0]
			if c.ResourceRequest.ProverAmount > uint(len(m.znodes)) {
				slog.Error("no enough resource for the project", "require prover amount", c.ResourceRequest.ProverAmount, "current prover", len(m.znodes), "project_id", projectID)
				continue
			}
			znodeMap := map[[sha256.Size]byte]string{}
			for _, n := range m.znodes {
				znodeMap[sha256.Sum256([]byte(n))] = n
			}

			b := make([]byte, 8)
			binary.LittleEndian.PutUint64(b, projectID)
			projectIDHash := sha256.Sum256(b)

			ds := make([]distance, 0, len(m.znodes))

			for h := range znodeMap {
				n := new(big.Int).Xor(new(big.Int).SetBytes(h[:]), new(big.Int).SetBytes(projectIDHash[:]))
				ds = append(ds, distance{
					distance: n,
					hash:     h,
				})
			}

			sort.SliceStable(ds, func(i, j int) bool {
				return ds[i].distance.Cmp(ds[j].distance) < 0
			})

			amount := c.ResourceRequest.ProverAmount
			if amount == 0 {
				amount = 1
			}

			ds = ds[:amount]
			for _, d := range ds {
				provers = append(provers, znodeMap[d.hash])
			}
			isMe := false
			for _, p := range provers {
				if p == m.ioID {
					isMe = true
				}
			}
			if !isMe {
				slog.Info("the project not scheduld to this znode", "project_id", projectID)
				continue
			}
		}

		for _, c := range cs {
			slog.Debug("contract project loaded", "project_id", pm.ProjectID, "version", c.Version, "vm_type", c.VMType, "code_size", len(c.Code))
			m.set(projectID, c.Version, c, provers)
		}
	}
}

func (m *Manager) fillProjectPoolFromLocal(projectFileDir string) {
	if projectFileDir == "" {
		return
	}
	files, err := os.ReadDir(projectFileDir)
	if err != nil {
		if !os.IsNotExist(err) {
			slog.Error("read project directory failed", "path", projectFileDir, "error", err)
			return
		}
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		data, err := os.ReadFile(path.Join(projectFileDir, f.Name()))
		if err != nil {
			slog.Error("read project config failed", "filename", f.Name(), "error", err)
			continue
		}

		projectID, err := strconv.ParseUint(f.Name(), 10, 64)
		if err != nil {
			slog.Error("parse filename failed", "filename", f.Name())
			continue
		}
		cs := []*Config{}
		if err := json.Unmarshal(data, &cs); err != nil {
			slog.Error("parse project config failed", "filename", f.Name())
			continue
		}
		var provers []string
		if m.ioID != "" {
			provers = append(provers, m.ioID)
		}

		for _, c := range cs {
			slog.Info("local project loaded", "project_id", projectID, "version", c.Version, "vm_type", c.VMType, "code_size", len(c.Code))
			m.set(projectID, c.Version, c, provers)
		}
	}
}

func NewManager(chainEndpoint, contractAddress, projectFileDir, projectCacheDir, ipfsEndpoint, ioID string, znodes []string) (*Manager, error) {
	var c *cache
	var err error
	if projectCacheDir != "" {
		c, err = newCache(projectCacheDir)
		if err != nil {
			return nil, errors.Wrap(err, "failed to new cache")
		}
	}
	m := &Manager{
		pool:         make(map[key]*Project),
		projectIDs:   make(map[uint64]bool),
		ipfsEndpoint: ipfsEndpoint,
		notify:       make(chan uint64, 32),
		cache:        c,
		ioID:         ioID,
		znodes:       znodes,
	}

	if contractAddress != "" {
		client, err := ethclient.Dial(chainEndpoint)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to dial chain, endpoint %s", chainEndpoint)
		}
		m.instance, err = contracts.NewContracts(common.HexToAddress(contractAddress), client)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to new contract instance, endpoint %s, contractAddress %s", chainEndpoint, contractAddress)
		}

		m.fillProjectPoolFromContract()

		topic := "ProjectUpserted(uint64,string,bytes32)"
		monitor, err := newDefaultMonitor(
			chainEndpoint,
			[]string{contractAddress},
			[]string{topic},
		)
		if err != nil {
			return nil, errors.Wrap(err, "failed to new contract monitor")
		}
		go monitor.run()
		go m.watchProjectRegistrar(monitor.mustEvents(topic), monitor)
	}

	m.fillProjectPoolFromLocal(projectFileDir)

	return m, nil
}
