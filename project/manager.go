package project

import (
	"bytes"
	"log/slog"
	"math/big"
	"os"
	"path"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/persistence/contract"
	"github.com/machinefi/sprout/smartcontracts/go/project"
)

type Manager struct {
	ipfsEndpoint string
	instance     *project.Project
	projects     sync.Map // projectID(uint64) -> *Project
	cache        *cache   // optional
}

func (m *Manager) ProjectIDs() []uint64 {
	var ids []uint64
	m.projects.Range(func(key, value any) bool {
		ids = append(ids, key.(uint64))
		return true
	})
	return ids
}

func (m *Manager) Project(projectID uint64) (*Project, error) {
	var err error
	p, ok := m.projects.Load(projectID)
	if !ok {
		p, err = m.load(projectID)
		if err != nil {
			return nil, err
		}
	}
	return p.(*Project), nil
}

func (m *Manager) load(projectID uint64) (*Project, error) {
	emptyHash := [32]byte{}
	c, err := m.instance.Config(nil, new(big.Int).SetUint64(projectID))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get project meta from chain, project_id %v", projectID)
	}
	if c.Uri == "" || bytes.Equal(c.Hash[:], emptyHash[:]) {
		return nil, errors.Errorf("the project not exist, project_id %v", projectID)
	}

	pm := &Meta{
		ProjectID: projectID,
		Uri:       c.Uri,
		Hash:      c.Hash,
	}

	var data []byte
	cached := true
	if m.cache != nil {
		data = m.cache.get(projectID, c.Hash[:])
	}
	if len(data) == 0 {
		cached = false
		data, err = pm.FetchProjectRawData(m.ipfsEndpoint)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get project raw data, project_id %v", projectID)
		}
	}
	if !cached && m.cache != nil {
		m.cache.set(projectID, data)
	}

	p, err := convertProject(data)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to convert project, project_id %v", projectID)
	}
	m.projects.Store(projectID, p)
	return p, nil
}

func (m *Manager) loadFromLocal(projectFileDir string) error {
	files, err := os.ReadDir(projectFileDir)
	if err != nil {
		return errors.Wrapf(err, "failed to read project directory %s", projectFileDir)
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		data, err := os.ReadFile(path.Join(projectFileDir, f.Name()))
		if err != nil {
			slog.Error("failed to read project file", "filename", f.Name(), "error", err)
			continue
		}

		projectID, err := strconv.ParseUint(f.Name(), 10, 64)
		if err != nil {
			slog.Error("failed to parse filename", "filename", f.Name())
			continue
		}

		p, err := convertProject(data)
		if err != nil {
			slog.Error("failed to convert project", "project_id", projectID, "error", err)
			continue
		}
		m.projects.Store(projectID, p)
	}
	return nil
}

func (m *Manager) watchProjectContract(chainEndpoint, contractAddress string) error {
	projectCh, err := contract.ListAndWatchProject(chainEndpoint, contractAddress, 0)
	if err != nil {
		return err
	}

	go func() {
		for p := range projectCh {
			for id := range p.Projects {
				m.projects.Delete(id)
			}
		}
	}()
	return nil
}

func NewManager(chainEndpoint, contractAddress, projectCacheDir, ipfsEndpoint, projectFileDirectory string) (*Manager, error) {
	var (
		c   *cache
		err error
	)

	if projectCacheDir != "" {
		c, err = newCache(projectCacheDir)
		if err != nil {
			return nil, errors.Wrap(err, "failed to new cache")
		}
	}

	m := &Manager{
		ipfsEndpoint: ipfsEndpoint,
		cache:        c,
	}

	if projectFileDirectory != "" {
		if err := m.loadFromLocal(projectFileDirectory); err != nil {
			return nil, err
		}
		return m, nil
	}

	client, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to dial chain, endpoint %s", chainEndpoint)
	}
	instance, err := project.NewProject(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to new contract instance, endpoint %s, contractAddress %s", chainEndpoint, contractAddress)
	}
	m.instance = instance

	if err := m.watchProjectContract(chainEndpoint, contractAddress); err != nil {
		return nil, err
	}
	return m, nil
}
