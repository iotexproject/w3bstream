package project

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/iotexproject/w3bstream/util/filefetcher"
	"github.com/pkg/errors"
)

type ContractProject func(projectID *big.Int) (string, common.Hash, error)
type ProjectFile func(projectID *big.Int) ([]byte, common.Hash, error)
type UpsertProjectFile func(projectID *big.Int, file []byte, hash common.Hash) error

type Manager struct {
	contractProject   ContractProject
	projectFile       ProjectFile
	upsertProjectFile UpsertProjectFile
}

func (m *Manager) Project(projectID *big.Int) (*Project, error) {
	uri, hash, err := m.contractProject(projectID)
	if err != nil {
		return nil, errors.Errorf("failed to get project metadata, project_id %v", projectID)
	}
	pf, fileHash, err := m.projectFile(projectID)
	if err != nil {
		return nil, errors.Errorf("failed to get project file, project_id %v", projectID)
	}

	p := &Project{}
	if bytes.Equal(fileHash[:], hash[:]) {
		err := p.Unmarshal(pf)
		return p, errors.Wrapf(err, "failed to unmarshal project file, project_id %v", projectID)
	}

	// project file hash mismatch, fetch new project file from uri
	fetcher := &filefetcher.Filedescriptor{Uri: uri, Hash: hash}
	data, err := fetcher.FetchFile()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to fetch project file, project_id %v", projectID)
	}
	err = p.Unmarshal(data)
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
