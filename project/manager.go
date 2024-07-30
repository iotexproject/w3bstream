package project

import (
	"bytes"
	"log/slog"
	"os"
	"path"
	"strconv"
	"sync"

	"github.com/cockroachdb/pebble"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/persistence/contract"
)

type ContractProject func(projectID uint64) *contract.Project

type Manager struct {
	local           bool
	contractProject ContractProject // optional
	db              *pebble.DB      // optional
	localProjects   sync.Map        // projectID(uint64) -> *Project for local model
}

func (m *Manager) dbKey(projectID uint64) []byte {
	return []byte("project_file_" + strconv.FormatUint(projectID, 10))
}

func (m *Manager) dbHashKey(projectID uint64) []byte {
	return []byte("project_file_hash_" + strconv.FormatUint(projectID, 10))
}

func (m *Manager) ProjectIDs() ([]uint64, error) {
	if !m.local {
		return nil, errors.New("get project ids not supported")
	}
	var ids []uint64
	m.localProjects.Range(func(key, value any) bool {
		ids = append(ids, key.(uint64))
		return true
	})
	return ids, nil
}

func (m *Manager) Project(projectID uint64) (*Project, error) {
	if m.local {
		return m.loadFromLocal(projectID)
	}
	return m.loadFromContract(projectID)
}

func (m *Manager) loadFromLocal(projectID uint64) (*Project, error) {
	p, ok := m.localProjects.Load(projectID)
	if !ok {
		return nil, errors.Errorf("project not exist, project_id %v", projectID)
	}
	return p.(*Project), nil
}

func (m *Manager) loadFromContract(projectID uint64) (*Project, error) {
	cp := m.contractProject(projectID)
	if cp == nil {
		return nil, errors.Errorf("project not exist, project_id %v", projectID)
	}
	dataBytes, closer, err := m.db.Get(m.dbHashKey(projectID))
	if err != nil && err != pebble.ErrNotFound {
		return nil, errors.Wrapf(err, "failed to get db project file hash data, project_id %v", projectID)
	}
	hash := make([]byte, len(dataBytes))
	copy(hash, dataBytes)
	if err := closer.Close(); err != nil {
		return nil, errors.Wrapf(err, "failed to close result of project file hash data, project_id %v", projectID)
	}

	if bytes.Equal(cp.Hash[:], hash) {
		dataBytes, closer, err := m.db.Get(m.dbKey(projectID))
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get db project file data, project_id %v", projectID)
		}
		p, err := convertProject(dataBytes)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal project file, project_id %v", projectID)
		}
		if err := closer.Close(); err != nil {
			return nil, errors.Wrapf(err, "failed to close result of project file data, project_id %v", projectID)
		}
		return p, nil
	} else {
		pm := &Meta{ProjectID: projectID, Uri: cp.Uri, Hash: cp.Hash}
		data, err := pm.FetchProjectFile()
		if err != nil {
			return nil, errors.Wrapf(err, "failed to fetch project file, project_id %v", projectID)
		}
		p, err := convertProject(data)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal project file, project_id %v", projectID)
		}
		if err := m.setDB(projectID, data, pm.Hash[:]); err != nil {
			return nil, err
		}
		return p, nil
	}
}

func (m *Manager) setDB(projectID uint64, file, hash []byte) error {
	batch := m.db.NewBatch()
	defer batch.Close()

	if err := batch.Set(m.dbHashKey(projectID), hash, nil); err != nil {
		return errors.Wrapf(err, "failed to set project file hash, project_id %v", projectID)
	}
	if err := batch.Set(m.dbKey(projectID), file, nil); err != nil {
		return errors.Wrapf(err, "failed to set project file, project_id %v", projectID)
	}
	if err := batch.Commit(pebble.Sync); err != nil {
		return errors.Wrapf(err, "failed to commit batch, project_id %v", projectID)
	}
	return nil
}

func (m *Manager) fillLocal(projectFileDir string) error {
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
		m.localProjects.Store(projectID, p)
	}
	return nil
}

func NewManager(db *pebble.DB, contractProject ContractProject) *Manager {
	return &Manager{
		local:           false,
		db:              db,
		contractProject: contractProject,
	}
}

func NewLocalManager(projectFileDirectory string) (*Manager, error) {
	m := &Manager{
		local: true,
	}

	if err := m.fillLocal(projectFileDirectory); err != nil {
		return nil, err
	}
	return m, nil
}
