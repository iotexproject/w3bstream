package task

import (
	"reflect"
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/testutil"
	testp2p "github.com/machinefi/sprout/testutil/p2p"
	testproject "github.com/machinefi/sprout/testutil/project"
	"github.com/machinefi/sprout/types"
	"github.com/machinefi/sprout/vm"
)

func TestNewProcessor(t *testing.T) {
	require := require.New(t)
	patches := NewPatches()
	defer patches.Reset()

	ps := &p2p.PubSubs{}

	t.Run("NewFailed", func(t *testing.T) {
		patches = testp2p.P2pNewPubSubs(patches, nil, errors.New(t.Name()))
		_, err := NewProcessor(nil, nil, "", 0)
		require.ErrorContains(err, t.Name())
	})
	patches = testp2p.P2pNewPubSubs(patches, ps, nil)

	t.Run("AddProjectFailed", func(t *testing.T) {
		patches = testproject.ProjectManagerGetAllProjectID(patches, append([]uint64{}, 1))
		patches = testp2p.P2pPubSubsAdd(patches, errors.New(t.Name()))
		_, err := NewProcessor(nil, nil, "", 0)
		require.ErrorContains(err, t.Name())
	})
}

func TestProcessor_ReportFail(t *testing.T) {
	patches := NewPatches()
	defer patches.Reset()
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

func TestProcessor_ReportSuccess(t *testing.T) {
	patches := NewPatches()
	defer patches.Reset()
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

func TestProcessor_HandleP2PData(t *testing.T) {
	patches := NewPatches()
	defer patches.Reset()
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

	data := &p2p.Data{
		Task: &types.Task{
			ID:             "",
			ProjectID:      uint64(0x1),
			ProjectVersion: "0.1",
			Data:           [][]byte{[]byte("data")},
		},
		TaskStateLog: nil,
	}

	t.Run("GetProjectFailed", func(t *testing.T) {
		patches = processorReportSuccess(patches)
		patches = testproject.ProjectManagerGet(patches, nil, errors.New(t.Name()))
		patches = processorReportFail(patches)
		p.handleP2PData(data, nil)
	})
	conf := &project.Config{
		Code:         "code",
		CodeExpParam: "codeExpParam",
		VMType:       "vmType",
		Output:       project.OutputConfig{},
		Aggregation:  project.AggregationConfig{},
		Version:      "",
	}
	patches = testproject.ProjectManagerGet(patches, conf, nil)

	t.Run("ProofFailed", func(t *testing.T) {
		patches = processorReportSuccess(patches)
		patches = vmHandlerHandle(patches, nil, errors.New(t.Name()))
		patches = processorReportFail(patches)
		p.handleP2PData(data, nil)
	})
	patches = vmHandlerHandle(patches, []byte("res"), nil)

	t.Run("HandleSuccess", func(t *testing.T) {
		patches = processorReportSuccess(patches)
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

func vmHandlerHandle(p *Patches, res []byte, err error) *Patches {
	var handler *vm.Handler
	return p.ApplyMethodFunc(
		reflect.TypeOf(handler),
		"Handle",
		func(msgs []*types.Message, vmtype types.VM, code string, expParam string) ([]byte, error) {
			return res, err
		},
	)
}
