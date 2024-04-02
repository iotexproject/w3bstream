package task

import (
	"reflect"
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/output"
	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/testutil"
	testproject "github.com/machinefi/sprout/testutil/project"
	"github.com/machinefi/sprout/types"
	"github.com/machinefi/sprout/vm"
)

func TestProcessor_ReportFail(t *testing.T) {
	processor := &Processor{}

	t.Run("MarshalFailed", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()
		p = testutil.JsonMarshal(p, []byte("any"), errors.New(t.Name()))
		processor.reportFail(&types.Task{}, errors.New(t.Name()), nil)
	})

	t.Run("PublishFailed", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()
		p = testutil.JsonMarshal(p, []byte("any"), nil)

		p = testutil.TopicPublish(p, errors.New(t.Name()))
		processor.reportFail(&types.Task{}, errors.New(t.Name()), nil)
	})
}

func TestProcessor_ReportSuccess(t *testing.T) {
	processor := &Processor{}

	t.Run("MarshalFailed", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()
		p = testutil.JsonMarshal(p, []byte("any"), errors.New(t.Name()))
		processor.reportSuccess(&types.Task{}, types.TaskStatePacked, nil, "", nil)
	})

	t.Run("PublishFailed", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()
		p = testutil.JsonMarshal(p, []byte("any"), nil)

		p = testutil.TopicPublish(p, errors.New(t.Name()))
		processor.reportSuccess(&types.Task{}, types.TaskStatePacked, nil, "", nil)
	})

}

func TestProcessor_HandleP2PData(t *testing.T) {
	processor := &Processor{
		vmHandler:            &vm.Handler{},
		projectConfigManager: &project.ConfigManager{},
	}

	t.Run("TaskNil", func(t *testing.T) {
		processor.HandleP2PData(&p2p.Data{
			Task:         nil,
			TaskStateLog: nil,
		}, nil)
	})

	data := &p2p.Data{
		Task: &types.Task{
			ID:             1,
			ProjectID:      uint64(0x1),
			ProjectVersion: "0.1",
			Data:           [][]byte{[]byte("data")},
		},
		TaskStateLog: nil,
	}

	t.Run("GetProjectFailed", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p = processorReportSuccess(p)
		p = testproject.ProjectConfigManagerGet(p, nil, errors.New(t.Name()))
		p = processorReportFail(p)
		processor.HandleP2PData(data, nil)
	})

	conf := &project.Config{
		Code:         "code",
		CodeExpParam: "codeExpParam",
		VMType:       "vmType",
		Output:       output.Config{},
		Aggregation:  project.AggregationConfig{},
		Version:      "",
	}

	t.Run("ProofFailed", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()
		p = testproject.ProjectConfigManagerGet(p, conf, nil)

		p = processorReportSuccess(p)
		p = vmHandlerHandle(p, nil, errors.New(t.Name()))
		p = processorReportFail(p)
		processor.HandleP2PData(data, nil)
	})

	t.Run("HandleSuccess", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()
		p = testproject.ProjectConfigManagerGet(p, conf, nil)
		p = vmHandlerHandle(p, []byte("res"), nil)

		p = processorReportSuccess(p)
		processor.HandleP2PData(data, nil)
	})

}

func processorReportSuccess(p *Patches) *Patches {
	var pro *Processor
	return p.ApplyPrivateMethod(pro, "reportSuccess", func(taskID string, state types.TaskState, comment string, topic *pubsub.Topic) {})
}

func processorReportFail(p *Patches) *Patches {
	var pro *Processor
	return p.ApplyPrivateMethod(pro, "reportFail", func(taskID string, err error, topic *pubsub.Topic) {})
}

func vmHandlerHandle(p *Patches, res []byte, err error) *Patches {
	var handler *vm.Handler
	return p.ApplyMethodFunc(
		reflect.TypeOf(handler),
		"Handle",
		func(taskID, projectID uint64, clientDID, sign string, vmtype vm.Type, code string, expParam string, data [][]byte) ([]byte, error) {
			return res, err
		},
	)
}
