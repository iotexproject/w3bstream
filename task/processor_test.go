package task

import (
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	"github.com/ethereum/go-ethereum/crypto"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/output"
	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/testutil"
	"github.com/machinefi/sprout/types"
	"github.com/machinefi/sprout/vm"
)

func TestProcessor_ReportFail(t *testing.T) {
	processor := &Processor{}

	t.Run("FailedToMarshal", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		testutil.JsonMarshal(p, []byte("any"), errors.New(t.Name()))
		processor.reportFail(&types.Task{}, errors.New(t.Name()), nil)
	})
	t.Run("FailedToPublish", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		testutil.JsonMarshal(p, []byte("any"), nil)
		testutil.TopicPublish(p, errors.New(t.Name()))
		processor.reportFail(&types.Task{}, errors.New(t.Name()), nil)
	})
}

func TestProcessor_ReportSuccess(t *testing.T) {
	processor := &Processor{}

	t.Run("FailedToMarshal", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		testutil.JsonMarshal(p, []byte("any"), errors.New(t.Name()))
		processor.reportSuccess(&types.Task{}, types.TaskStatePacked, nil, "", nil)
	})
	t.Run("FailedToPublish", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		testutil.JsonMarshal(p, []byte("any"), nil)
		testutil.TopicPublish(p, errors.New(t.Name()))
		processor.reportSuccess(&types.Task{}, types.TaskStatePacked, nil, "", nil)
	})
}

func TestProcessor_HandleProjectProvers(t *testing.T) {
	r := require.New(t)
	p := &Processor{}
	p.HandleProjectProvers(1, []uint64{1})

	v, ok := p.projectProvers.Load(uint64(1))
	r.True(ok)
	r.Equal(v.([]uint64), []uint64{1})
}

func TestProcessor_HandleP2PData(t *testing.T) {
	m := &project.Manager{}
	processor := &Processor{
		vmHandler: &vm.Handler{},
		project:   m.Project,
	}
	projectID := uint64(1)

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

	t.Run("FailedToGetProject", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		processorReportSuccess(p)
		p.ApplyMethodReturn(&project.Manager{}, "Project", nil, errors.New(t.Name()))
		processorReportFail(p)
		processor.HandleP2PData(data, nil)
	})

	testProject := &project.Project{
		DefaultVersion: "0.1",
		Versions: []*project.Config{{
			Code:         "code",
			CodeExpParam: "codeExpParam",
			VMType:       vm.Risc0,
			Output:       output.Config{},
			Version:      "0.1",
		}},
	}

	t.Run("FailedToGetProjectDefaultConfig", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(&project.Manager{}, "Project", testProject, nil)
		p.ApplyMethodReturn(&project.Project{}, "DefaultConfig", nil, errors.New(t.Name()))
		processorReportSuccess(p)
		processorReportFail(p)

		processor.HandleP2PData(data, nil)
	})
	t.Run("FailedToGetProjectDefaultConfig", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		processor.projectProvers.Store(projectID, []uint64{1, 2, 3})
		defer processor.projectProvers.Delete(projectID)

		p.ApplyMethodReturn(&project.Manager{}, "Project", testProject, nil)

		processor.HandleP2PData(data, nil)
	})
	t.Run("FailedToProof", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(&project.Manager{}, "Project", testProject, nil)
		processorReportSuccess(p)
		p.ApplyMethodReturn(&vm.Handler{}, "Handle", nil, errors.New(t.Name()))
		processorReportFail(p)

		processor.HandleP2PData(data, nil)
	})
	t.Run("FailedToVerifySignature", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(&project.Manager{}, "Project", testProject, nil)
		p.ApplyMethodReturn(&vm.Handler{}, "Handle", []byte("res"), nil)
		processorReportSuccess(p)

		processor.HandleP2PData(data, nil)
	})
	t.Run("FailedToCallVMHandle", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(&project.Manager{}, "Project", testProject, nil)
		p.ApplyMethodReturn(&types.Task{}, "VerifySignature", nil)
		p.ApplyMethodReturn(&vm.Handler{}, "Handle", nil, errors.New(t.Name()))
		processorReportFail(p)
		processorReportSuccess(p)

		processor.HandleP2PData(data, nil)
	})
	t.Run("FailedToSignProof", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(&project.Manager{}, "Project", testProject, nil)
		p.ApplyMethodReturn(&types.Task{}, "VerifySignature", nil)
		p.ApplyMethodReturn(&vm.Handler{}, "Handle", []byte("res"), nil)
		p.ApplyPrivateMethod(processor, "signProof", func(*types.Task, []byte) (string, error) { return "", errors.New(t.Name()) })
		processorReportFail(p)
		processorReportSuccess(p)

		processor.HandleP2PData(data, nil)
	})
	t.Run("Success", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(&project.Manager{}, "Project", testProject, nil)
		p.ApplyMethodReturn(&types.Task{}, "VerifySignature", nil)
		p.ApplyMethodReturn(&vm.Handler{}, "Handle", []byte("res"), nil)
		p.ApplyPrivateMethod(processor, "signProof", func(*types.Task, []byte) (string, error) { return "", nil })
		processorReportSuccess(p)

		processor.HandleP2PData(data, nil)
	})
}

func TestProcessor_signProof(t *testing.T) {
	r := require.New(t)
	sk, err := crypto.GenerateKey()
	r.NoError(err)
	p := &Processor{
		proverPrivateKey: sk,
	}
	_, err = p.signProof(&types.Task{}, []byte{})
	r.NoError(err)
}

func TestNewProcessor(t *testing.T) {
	r := require.New(t)
	p := NewProcessor(nil, nil, nil, nil, 0)
	r.NotNil(p)
}

func processorReportSuccess(p *Patches) *Patches {
	var pro *Processor
	return p.ApplyPrivateMethod(pro, "reportSuccess", func(taskID string, state types.TaskState, comment string, topic *pubsub.Topic) {})
}

func processorReportFail(p *Patches) *Patches {
	var pro *Processor
	return p.ApplyPrivateMethod(pro, "reportFail", func(taskID string, err error, topic *pubsub.Topic) {})
}
