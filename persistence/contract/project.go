package contract

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

var (
	RequiredProverAmount         = crypto.Keccak256Hash([]byte("RequiredProverAmount"))
	VmType                       = crypto.Keccak256Hash([]byte("VmType"))
	ClientManagementContractAddr = crypto.Keccak256Hash([]byte("ClientManagementContractAddress"))

	attributeSetTopic         = crypto.Keccak256Hash([]byte("AttributeSet(uint256,bytes32,bytes)"))
	projectPausedTopic        = crypto.Keccak256Hash([]byte("ProjectPaused(uint256)"))
	projectResumedTopic       = crypto.Keccak256Hash([]byte("ProjectResumed(uint256)"))
	projectConfigUpdatedTopic = crypto.Keccak256Hash([]byte("ProjectConfigUpdated(uint256,string,bytes32)"))
)

type Project struct {
	ID         uint64
	Paused     bool
	Uri        string
	Hash       common.Hash
	Attributes map[common.Hash][]byte
}

type projectDiff struct {
	id         uint64
	paused     *bool
	uri        string
	hash       *common.Hash
	attributes map[common.Hash][]byte
}

type blockProject struct {
	Projects map[uint64]*Project
}

type blockProjectDiff struct {
	diffs map[uint64]*projectDiff
}

func newProject() *Project {
	return &Project{
		Paused:     true,
		Attributes: map[common.Hash][]byte{},
	}
}

func (p *Project) merge(diff *projectDiff) {
	if diff.id != 0 {
		p.ID = diff.id
	}
	if diff.paused != nil {
		p.Paused = *diff.paused
	}
	if diff.uri != "" {
		p.Uri = diff.uri
	}
	if diff.hash != nil {
		p.Hash = *diff.hash
	}
	for h, d := range diff.attributes {
		p.Attributes[h] = d
	}
}

func (ps *blockProject) merge(diff *blockProjectDiff) {
	for id, p := range ps.Projects {
		diffP, ok := diff.diffs[id]
		if ok {
			p.merge(diffP)
		}
	}
	for id, p := range diff.diffs {
		if _, ok := ps.Projects[id]; !ok {
			np := newProject()
			np.merge(p)
			ps.Projects[id] = np
		}
	}
}

// return blockNumber -> *blockProjectDiff
func (c *Contract) processProjectLogs(logs []types.Log) (map[uint64]*blockProjectDiff, error) {
	psMap := map[uint64]*blockProjectDiff{}

	for _, l := range logs {
		ps, ok := psMap[l.BlockNumber]
		if !ok {
			ps = &blockProjectDiff{
				diffs: map[uint64]*projectDiff{},
			}
		}
		switch l.Topics[0] {
		case attributeSetTopic:
			e, err := c.projectInstance.ParseAttributeSet(l)
			if err != nil {
				return nil, errors.Wrap(err, "failed to parse project attribute set event")
			}

			p, ok := ps.diffs[e.ProjectId.Uint64()]
			if !ok {
				p = &projectDiff{
					id:         e.ProjectId.Uint64(),
					attributes: map[common.Hash][]byte{},
				}
			}
			p.attributes[e.Key] = e.Value
			ps.diffs[e.ProjectId.Uint64()] = p

		case projectPausedTopic:
			e, err := c.projectInstance.ParseProjectPaused(l)
			if err != nil {
				return nil, errors.Wrap(err, "failed to parse project paused event")
			}

			p, ok := ps.diffs[e.ProjectId.Uint64()]
			if !ok {
				p = &projectDiff{
					id:         e.ProjectId.Uint64(),
					attributes: map[common.Hash][]byte{},
				}
			}
			paused := true
			p.paused = &paused
			ps.diffs[e.ProjectId.Uint64()] = p

		case projectResumedTopic:
			e, err := c.projectInstance.ParseProjectResumed(l)
			if err != nil {
				return nil, errors.Wrap(err, "failed to parse project resumed event")
			}

			p, ok := ps.diffs[e.ProjectId.Uint64()]
			if !ok {
				p = &projectDiff{
					id:         e.ProjectId.Uint64(),
					attributes: map[common.Hash][]byte{},
				}
			}
			paused := false
			p.paused = &paused
			ps.diffs[e.ProjectId.Uint64()] = p

		case projectConfigUpdatedTopic:
			e, err := c.projectInstance.ParseProjectConfigUpdated(l)
			if err != nil {
				return nil, errors.Wrap(err, "failed to parse project config updated event")
			}

			p, ok := ps.diffs[e.ProjectId.Uint64()]
			if !ok {
				p = &projectDiff{
					id:         e.ProjectId.Uint64(),
					attributes: map[common.Hash][]byte{},
				}
			}
			h := common.Hash(e.Hash)
			p.uri = e.Uri
			p.hash = &h
			ps.diffs[e.ProjectId.Uint64()] = p
		}
		psMap[l.BlockNumber] = ps
	}

	return psMap, nil
}
