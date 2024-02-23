package task

import (
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/machinefi/sprout/testutil"
	"github.com/machinefi/sprout/types"
	"github.com/machinefi/sprout/vm"
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	"github.com/machinefi/sprout/p2p"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestNewProcessor(t *testing.T) {
	require := require.New(t)
	patches := NewPatches()

	ps := &p2p.PubSubs{}

	t.Run("NewFailed", func(t *testing.T) {
		patches = testutil.P2pNewPubSubs(patches, nil, errors.New(t.Name()))
		_, err := NewProcessor(nil, nil, "", 0)
		require.ErrorContains(err, t.Name())
	})
	patches = testutil.P2pNewPubSubs(patches, ps, nil)

	t.Run("AddProjectFailed", func(t *testing.T) {
		patches = testutil.ProjectManagerGetAllProjectID(patches, append([]uint64{}, 1))
		patches = testutil.P2pPubSubsAdd(patches, errors.New(t.Name()))
		_, err := NewProcessor(nil, nil, "", 0)
		require.ErrorContains(err, t.Name())
	})
}

func TestReportFail(t *testing.T) {
	patches := NewPatches()
	p := &Processor{}

	t.Run("MarshalFailed", func(t *testing.T) {
		patches = testutil.JsonMarshal(patches, []byte("any"), errors.New(t.Name()))
		p.reportFail("taskID", errors.New(t.Name()), nil)
	})
	patches = testutil.JsonMarshal(patches, []byte("any"), nil)

	t.Run("PublishFailed", func(t *testing.T) {
		patches = testutil.TopicPublish(patches, errors.New(t.Name()))
		p.reportFail("taskID", errors.New(t.Name()), nil)
	})
}

func TestReportSuccess(t *testing.T) {
	patches := NewPatches()
	p := &Processor{}

	t.Run("MarshalFailed", func(t *testing.T) {
		patches = testutil.JsonMarshal(patches, []byte("any"), errors.New(t.Name()))
		p.reportSuccess("taskID", types.TaskStatePacked, "", nil)
	})
	patches = testutil.JsonMarshal(patches, []byte("any"), nil)

	t.Run("PublishFailed", func(t *testing.T) {
		patches = testutil.TopicPublish(patches, errors.New(t.Name()))
		p.reportSuccess("taskID", types.TaskStatePacked, "", nil)
	})

}

func TestProcessorHandleP2PData(t *testing.T) {
	patches := NewPatches()
	p := &Processor{
		vmHandler:      &vm.Handler{},
		projectManager: nil,
		ps:             nil,
	}

	t.Run("TaskNil", func(t *testing.T) {
		data := &p2p.Data{
			Task:         nil,
			TaskStateLog: nil,
		}
		p.handleP2PData(data, nil)
	})

	t.Run("GetProjectFailed", func(t *testing.T) {
		data := &p2p.Data{
			Task: &types.Task{
				ID: "",
				Messages: []*types.Message{{
					ID:             "id1",
					ProjectID:      uint64(0x1),
					ProjectVersion: "0.1",
					Data:           "data",
				}},
			},
			TaskStateLog: nil,
		}
		patches = processorReportSuccess(patches)
		patches = testutil.ProjectManagerGet(patches, errors.New(t.Name()))
		patches = processorReportFail(patches)
		p.handleP2PData(data, nil)
	})
}

func processorReportSuccess(p *Patches) *Patches {
	var pro *Processor
	return ApplyPrivateMethod(pro, "reportSuccess", func(taskID string, state types.TaskState, comment string, topic *pubsub.Topic) {})
}

func processorReportFail(p *Patches) *Patches {
	var pro *Processor
	return ApplyPrivateMethod(pro, "reportFail", func(taskID string, err error, topic *pubsub.Topic) {})
}
