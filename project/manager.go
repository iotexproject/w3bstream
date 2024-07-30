package project

import (
	"log/slog"
	"os"
	"path"
	"strconv"
	"sync"

	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/persistence/contract"
)

type ContractProject func(projectID uint64) *contract.Project

type Manager struct {
	contractProject ContractProject
	projects        sync.Map // projectID(uint64) -> *Project for local model
	cache           *cache   // optional
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
	cp := m.contractProject(projectID)
	if cp == nil {
		return nil, errors.Errorf("the project not exist, project_id %v", projectID)
	}

	pm := &Meta{
		ProjectID: projectID,
		Uri:       cp.Uri,
		Hash:      cp.Hash,
	}

	var data []byte
	var err error
	cached := true
	if m.cache != nil {
		data = m.cache.get(projectID, cp.Hash[:])
	}
	if len(data) == 0 {
		cached = false
		data, err = pm.FetchProjectFile()
		if err != nil {
			return nil, errors.Wrapf(err, "failed to fetch project file, project_id %v", projectID)
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

func (m *Manager) watchProject(projectNotification <-chan uint64) {
	for pid := range projectNotification {
		m.projects.Delete(pid)
	}
}

func NewManager(projectCacheDir string, contractProject ContractProject, projectNotification <-chan uint64) (*Manager, error) {
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
		contractProject: contractProject,
		cache:           c,
	}
	go m.watchProject(projectNotification)
	return m, nil
}

func NewLocalManager(projectFileDirectory string) (*Manager, error) {
	m := &Manager{}

	if err := m.loadFromLocal(projectFileDirectory); err != nil {
		return nil, err
	}
	return m, nil
}
