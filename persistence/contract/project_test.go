package contract

import (
	"math/big"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/smartcontracts/go/project"
	"github.com/machinefi/sprout/util/hash"
)

func TestNewProject(t *testing.T) {
	r := require.New(t)
	p := newProject()
	r.Equal(p.Paused, true)
}

func TestProject_merge(t *testing.T) {
	r := require.New(t)

	np := &Project{Attributes: map[common.Hash][]byte{}}

	paused := true
	hash := hash.Keccak256Uint64(1)
	attr := map[common.Hash][]byte{}
	attr[hash] = []byte("1")
	diff := &projectDiff{
		id:         1,
		paused:     &paused,
		uri:        "uri",
		hash:       &hash,
		attributes: attr,
	}
	np.merge(diff)
	r.Equal(np.ID, diff.id)
	r.Equal(np.Uri, diff.uri)
}

func TestBlockProject_merge(t *testing.T) {
	r := require.New(t)

	np := &blockProject{Projects: map[uint64]*Project{}}

	paused := true
	hash := hash.Keccak256Uint64(1)
	attr := map[common.Hash][]byte{}
	attr[hash] = []byte("1")
	diff := &blockProjectDiff{
		diffs: map[uint64]*projectDiff{
			1: {
				id:         1,
				paused:     &paused,
				uri:        "uri",
				hash:       &hash,
				attributes: attr,
			},
		},
	}
	np.merge(diff)
	r.Equal(len(np.Projects), 1)
}

func TestContract_processProjectLogs(t *testing.T) {
	r := require.New(t)
	id := new(big.Int).SetUint64(1)
	filterer := &project.ProjectFilterer{}
	c := &Contract{projectInstance: &project.Project{ProjectFilterer: *filterer}}

	t.Run("FailedToParseAttributeSetEvent", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(filterer, "ParseAttributeSet", &project.ProjectAttributeSet{ProjectId: id}, errors.New(t.Name()))

		logs := []types.Log{
			{
				Topics:      []common.Hash{attributeSetTopic},
				BlockNumber: 100,
				TxIndex:     1,
			},
		}

		_, err := c.processProjectLogs(logs)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToParseProjectPausedEvent", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(filterer, "ParseProjectPaused", &project.ProjectProjectPaused{ProjectId: id}, errors.New(t.Name()))

		logs := []types.Log{
			{
				Topics:      []common.Hash{projectPausedTopic},
				BlockNumber: 100,
				TxIndex:     1,
			},
		}

		_, err := c.processProjectLogs(logs)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToParseProjectResumedEvent", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(filterer, "ParseProjectResumed", &project.ProjectProjectResumed{ProjectId: id}, errors.New(t.Name()))

		logs := []types.Log{
			{
				Topics:      []common.Hash{projectResumedTopic},
				BlockNumber: 100,
				TxIndex:     1,
			},
		}

		_, err := c.processProjectLogs(logs)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToParseProjectConfigUpdatedEvent", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(filterer, "ParseProjectConfigUpdated", &project.ProjectProjectConfigUpdated{ProjectId: id}, errors.New(t.Name()))

		logs := []types.Log{
			{
				Topics:      []common.Hash{projectConfigUpdatedTopic},
				BlockNumber: 100,
				TxIndex:     1,
			},
		}

		_, err := c.processProjectLogs(logs)
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(filterer, "ParseAttributeSet", &project.ProjectAttributeSet{ProjectId: id}, nil)
		p.ApplyMethodReturn(filterer, "ParseProjectPaused", &project.ProjectProjectPaused{ProjectId: id}, nil)
		p.ApplyMethodReturn(filterer, "ParseProjectResumed", &project.ProjectProjectResumed{ProjectId: id}, nil)
		p.ApplyMethodReturn(filterer, "ParseProjectConfigUpdated", &project.ProjectProjectConfigUpdated{ProjectId: id}, nil)

		logs := []types.Log{
			{
				Topics:      []common.Hash{attributeSetTopic},
				BlockNumber: 100,
				TxIndex:     1,
			},
			{
				Topics:      []common.Hash{projectPausedTopic},
				BlockNumber: 99,
				TxIndex:     1,
			},
			{
				Topics:      []common.Hash{projectResumedTopic},
				BlockNumber: 100,
				TxIndex:     2,
			},
			{
				Topics:      []common.Hash{projectConfigUpdatedTopic},
				BlockNumber: 101,
				TxIndex:     1,
			},
			{
				Topics:      []common.Hash{projectConfigUpdatedTopic},
				BlockNumber: 101,
				TxIndex:     2,
			},
			{
				Topics:      []common.Hash{projectConfigUpdatedTopic},
				BlockNumber: 98,
				TxIndex:     2,
			},
			{
				Topics:      []common.Hash{projectConfigUpdatedTopic},
				BlockNumber: 98,
				TxIndex:     1,
			},
		}
		diffs, err := c.processProjectLogs(logs)
		r.NoError(err)
		r.Equal(len(diffs), 4)
	})
}
