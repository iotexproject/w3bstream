package project

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

type ContractProject func(projectID uint64) (string, common.Hash, error)
type ProjectFile func(projectID uint64) ([]byte, common.Hash, error)
type UpsertProjectFile func(projectID uint64, file []byte, hash common.Hash) error

type Manager struct {
	contractProject   ContractProject
	projectFile       ProjectFile
	upsertProjectFile UpsertProjectFile
}

func (m *Manager) Project(projectID uint64) (*Project, error) {
	uri, hash, err := m.contractProject(projectID)
	if err != nil {
		return nil, errors.Errorf("failed to get project metadata, project_id %v", projectID)
	}
	pf, fileHash, err := m.projectFile(projectID)
	if err != nil {
		return nil, errors.Errorf("failed to get project file, project_id %v", projectID)
	}

	if bytes.Equal(fileHash[:], hash[:]) {
		p, err := convertProject(pf)
		return p, errors.Wrapf(err, "failed to unmarshal project file, project_id %v", projectID)
	}

	pm := &Meta{ProjectID: projectID, Uri: uri, Hash: hash}
	data, err := pm.FetchProjectFile()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to fetch project file, project_id %v", projectID)
	}
	p, err := convertProject(data)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal project file, project_id %v", projectID)
	}
	err = m.upsertProjectFile(projectID, data, hash)
	return p, err
}

func NewManager(cp ContractProject, pf ProjectFile, upf UpsertProjectFile) *Manager {
	return &Manager{
		contractProject:   cp,
		projectFile:       pf,
		upsertProjectFile: upf,
	}
}
