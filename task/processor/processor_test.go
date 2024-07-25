package processor

import (
	"bytes"
	"encoding/binary"
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/iotexproject/w3bstream/output"
	"github.com/iotexproject/w3bstream/p2p"
	"github.com/iotexproject/w3bstream/project"
	"github.com/iotexproject/w3bstream/task"
	"github.com/iotexproject/w3bstream/testutil"
	"github.com/iotexproject/w3bstream/vm"
)

func TestProcessor_ReportFail(t *testing.T) {
	processor := &Processor{}

	t.Run("FailedToMarshal", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		testutil.JsonMarshal(p, []byte("any"), errors.New(t.Name()))
		processor.reportFail(&task.Task{}, errors.New(t.Name()), nil)
	})
	t.Run("FailedToPublish", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		testutil.JsonMarshal(p, []byte("any"), nil)
		testutil.TopicPublish(p, errors.New(t.Name()))
		processor.reportFail(&task.Task{}, errors.New(t.Name()), nil)
	})
}

func TestProcessor_ReportSuccess(t *testing.T) {
	processor := &Processor{}

	t.Run("FailedToMarshal", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		testutil.JsonMarshal(p, []byte("any"), errors.New(t.Name()))
		processor.reportSuccess(&task.Task{}, task.StatePacked, nil, "", nil)
	})
	t.Run("FailedToPublish", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		testutil.JsonMarshal(p, []byte("any"), nil)
		testutil.TopicPublish(p, errors.New(t.Name()))
		processor.reportSuccess(&task.Task{}, task.StatePacked, nil, "", nil)
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
		Task: &task.Task{
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
			VMTypeID:     1,
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
		p.ApplyMethodReturn(&task.Task{}, "VerifySignature", nil)
		p.ApplyMethodReturn(&vm.Handler{}, "Handle", nil, errors.New(t.Name()))
		processorReportFail(p)
		processorReportSuccess(p)

		processor.HandleP2PData(data, nil)
	})
	t.Run("FailedToSignProof", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(&project.Manager{}, "Project", testProject, nil)
		p.ApplyMethodReturn(&task.Task{}, "VerifySignature", nil)
		p.ApplyMethodReturn(&vm.Handler{}, "Handle", []byte("res"), nil)
		p.ApplyPrivateMethod(processor, "signProof", func(*task.Task, []byte) (string, error) { return "", errors.New(t.Name()) })
		processorReportFail(p)
		processorReportSuccess(p)

		processor.HandleP2PData(data, nil)
	})
	t.Run("Success", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyMethodReturn(&project.Manager{}, "Project", testProject, nil)
		p.ApplyMethodReturn(&task.Task{}, "VerifySignature", nil)
		p.ApplyMethodReturn(&vm.Handler{}, "Handle", []byte("res"), nil)
		p.ApplyPrivateMethod(processor, "signProof", func(*task.Task, []byte) (string, error) { return "", nil })
		processorReportSuccess(p)

		processor.HandleP2PData(data, nil)
	})
}

func TestProcessor_signProof(t *testing.T) {
	r := require.New(t)
	t.Run("FailedToWriteBinary", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyFuncReturn(binary.Write, errors.New(t.Name()))

		pr := &Processor{}
		_, err := pr.signProof(&task.Task{}, []byte{})

		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToWriteBinary2", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p.ApplyFuncSeq(binary.Write, []OutputCell{
			{
				Values: Params{nil},
				Times:  1,
			},
			{
				Values: Params{errors.New(t.Name())},
				Times:  1,
			},
		})
		pr := &Processor{}
		_, err := pr.signProof(&task.Task{}, []byte{})

		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToBufWriteString", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		buf := bytes.NewBuffer(nil)
		p.ApplyFuncReturn(binary.Write, nil)
		p.ApplyMethodReturn(buf, "WriteString", int(1), errors.New(t.Name()))

		pr := &Processor{}
		_, err := pr.signProof(&task.Task{}, []byte{})

		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToBufWrite", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		buf := bytes.NewBuffer(nil)
		p.ApplyFuncReturn(hexutil.Decode, nil, nil)
		p.ApplyFuncReturn(binary.Write, nil)
		p.ApplyMethodReturn(buf, "WriteString", int(1), nil)
		p.ApplyMethodReturn(buf, "Write", int(1), errors.New(t.Name()))

		pr := &Processor{}
		_, err := pr.signProof(&task.Task{}, []byte{})

		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToBufWrite2", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		buf := bytes.NewBuffer(nil)
		p.ApplyFuncReturn(hexutil.Decode, nil, nil)
		p.ApplyFuncReturn(binary.Write, nil)
		p.ApplyMethodReturn(buf, "WriteString", int(1), nil)
		p.ApplyMethodSeq(buf, "Write", []OutputCell{
			{
				Values: Params{int(1), nil},
				Times:  1,
			},
			{
				Values: Params{int(1), errors.New(t.Name())},
				Times:  1,
			},
		})
		pr := &Processor{}
		_, err := pr.signProof(&task.Task{}, []byte{})

		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToSign", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		buf := bytes.NewBuffer(nil)
		p.ApplyFuncReturn(binary.Write, nil)
		p.ApplyMethodReturn(buf, "WriteString", int(1), nil)
		p.ApplyMethodReturn(buf, "Write", int(1), nil)
		p.ApplyFuncReturn(crypto.Sign, nil, errors.New(t.Name()))

		pr := &Processor{}
		_, err := pr.signProof(&task.Task{}, []byte{})

		r.ErrorContains(err, t.Name())
	})
	t.Run("Success", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		buf := bytes.NewBuffer(nil)
		p.ApplyFuncReturn(binary.Write, nil)
		p.ApplyMethodReturn(buf, "WriteString", int(1), nil)
		p.ApplyMethodReturn(buf, "Write", int(1), nil)
		p.ApplyFuncReturn(crypto.Sign, nil, nil)

		pr := &Processor{}
		_, err := pr.signProof(&task.Task{}, []byte{})

		r.NoError(err)
	})
}

func TestNewProcessor(t *testing.T) {
	r := require.New(t)
	p := NewProcessor(nil, nil, nil, nil, 0)
	r.NotNil(p)
}

func processorReportSuccess(p *Patches) *Patches {
	var pro *Processor
	return p.ApplyPrivateMethod(pro, "reportSuccess", func(taskID string, state task.State, comment string, topic *pubsub.Topic) {})
}

func processorReportFail(p *Patches) *Patches {
	var pro *Processor
	return p.ApplyPrivateMethod(pro, "reportFail", func(taskID string, err error, topic *pubsub.Topic) {})
}
