package task

import (
	"testing"
	"time"

	. "github.com/agiledragon/gomonkey/v2"
	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/persistence"
	"github.com/machinefi/sprout/testutil"
	"github.com/machinefi/sprout/types"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestNewDispatcher(t *testing.T) {
	require := require.New(t)
	patches := NewPatches()

	t.Run("NewFailed", func(t *testing.T) {
		patches = testutil.P2pNewPubSubs(patches, nil, errors.New(t.Name()))
		_, err := NewDispatcher(nil, nil, "", "", "", 0)
		require.ErrorContains(err, t.Name())
	})
	patches = testutil.P2pNewPubSubs(patches, nil, nil)

	t.Run("New", func(t *testing.T) {
		_, err := NewDispatcher(nil, nil, "", "", "", 0)
		require.NoError(err)
	})
}

func TestHandleP2PData(t *testing.T) {
	patches := NewPatches()

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
			TaskID:    "TaskID",
			State:     types.TaskStatePacked,
			Comment:   "Comment",
			CreatedAt: time.Now(),
		},
	}

	t.Run("UpdateStateFailed", func(t *testing.T) {
		patches = testutil.PersistencePostgresUpdateState(patches, errors.New(t.Name()))
		d.handleP2PData(data, nil)
	})
	patches = testutil.PersistencePostgresUpdateState(patches, nil)

	t.Run("TaskStateProved", func(t *testing.T) {
		d.handleP2PData(data, nil)
	})

	t.Run("FetchTaskFailed", func(t *testing.T) {
		data.TaskStateLog.State = types.TaskStateProved
		patches = testutil.PersistencePostgresFetchByID(patches, nil, errors.New(t.Name()))
		d.handleP2PData(data, nil)
	})

	task := &types.Task{
		ID: "",
		Messages: []*types.Message{{
			ID:             "id1",
			ProjectID:      uint64(0x1),
			ProjectVersion: "0.1",
			Data:           "data",
		}},
	}
	patches = testutil.PersistencePostgresFetchByID(patches, task, nil)

	t.Run("GetProjectFailed", func(t *testing.T) {
		patches = testutil.ProjectManagerGet(patches, errors.New(t.Name()))
		d.handleP2PData(data, nil)
	})
	patches = testutil.ProjectManagerGet(patches, nil)

	t.Run("InitOutputFailed", func(t *testing.T) {
		patches = testutil.ProjectConfigGetOutput(patches, errors.New(t.Name()))
		d.handleP2PData(data, nil)
	})
}
