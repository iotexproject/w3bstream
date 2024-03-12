package task

import (
	"testing"
	"time"

	. "github.com/agiledragon/gomonkey/v2"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/persistence"
	testp2p "github.com/machinefi/sprout/testutil/p2p"
	testproject "github.com/machinefi/sprout/testutil/project"
	"github.com/machinefi/sprout/types"
)

func TestNewDispatcher(t *testing.T) {
	require := require.New(t)
	patches := NewPatches()
	defer patches.Reset()

	t.Run("NewFailed", func(t *testing.T) {
		patches = testp2p.P2pNewPubSubs(patches, nil, errors.New(t.Name()))
		_, err := NewDispatcher(nil, nil, "", "", "", 0, nil)
		require.ErrorContains(err, t.Name())
	})
	patches = testp2p.P2pNewPubSubs(patches, nil, nil)

	t.Run("New", func(t *testing.T) {
		_, err := NewDispatcher(nil, nil, "", "", "", 0, nil)
		require.NoError(err)
	})
}

func TestDispatcher_HandleP2PData(t *testing.T) {
	patches := NewPatches()
	defer patches.Reset()

	d := &Dispatcher{
		pubSubs:                   nil,
		pg:                        &persistence.Postgres{},
		projectManager:            nil,
		operatorPrivateKeyECDSA:   "",
		operatorPrivateKeyED25519: "",
	}

	t.Run("TaskStateLogNil", func(t *testing.T) {
		data := &p2p.Data{
			Task:         nil,
			TaskStateLog: nil,
		}
		d.handleP2PData(data, nil)
	})

	data := &p2p.Data{
		Task: nil,
		TaskStateLog: &types.TaskStateLog{
			Task:      types.Task{ID: 1},
			State:     types.TaskStatePacked,
			Comment:   []byte("Comment"),
			CreatedAt: time.Now(),
		},
	}

	t.Run("UpdateStateFailed", func(t *testing.T) {
		patches = patches.ApplyMethodReturn(&persistence.Postgres{}, "Create", errors.New(t.Name()))
		d.handleP2PData(data, nil)
	})
	patches = patches.ApplyMethodReturn(&persistence.Postgres{}, "Create", nil)

	t.Run("TaskStateProved", func(t *testing.T) {
		d.handleP2PData(data, nil)
	})

	t.Run("GetProjectFailed", func(t *testing.T) {
		patches = testproject.ProjectManagerGet(patches, nil, errors.New(t.Name()))
		d.handleP2PData(data, nil)
	})
	patches = testproject.ProjectManagerGet(patches, nil, nil)

	t.Run("InitOutputFailed", func(t *testing.T) {
		patches = testproject.ProjectConfigGetOutput(patches, nil, errors.New(t.Name()))
		d.handleP2PData(data, nil)
	})
}
