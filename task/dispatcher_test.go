package task

import (
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/types"
)

type mockPersistence struct{}

func (m *mockPersistence) Create(t *types.Task, tl *types.TaskStateLog) error {
	return nil
}

type mockDatasourceNil struct{}

func (m *mockDatasourceNil) Retrieve(nextTaskID uint64) (*types.Task, error) {
	return nil, nil
}

type mockDatasourceErr struct {
	err error
}

func (m *mockDatasourceErr) Retrieve(nextTaskID uint64) (*types.Task, error) {
	return nil, m.err
}

type mockDatasourceSuccess struct {
	task *types.Task
}

func (m *mockDatasourceSuccess) Retrieve(nextTaskID uint64) (*types.Task, error) {
	return m.task, nil
}

func TestNewDispatcher(t *testing.T) {
	r := require.New(t)

	t.Run("NewFailed", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(p2p.NewPubSubs, nil, errors.New(t.Name()))
		err := RunDispatcher(nil, nil, nil, "", "", "", "", "", 0)
		r.ErrorContains(err, t.Name())
	})

	t.Run("New", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()
		p = p.ApplyFuncReturn(p2p.NewPubSubs, nil, nil)

		err := RunDispatcher(nil, nil, nil, "", "", "", "", "", 0)
		r.NoError(err)
	})
}

// func TestDispatcher_HandleP2PData(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	op := mock.NewMockOutput(ctrl)

// 	d := &Dispatcher{
// 		pubSubs:                   nil,
// 		persistence:               &mockPersistence{},
// 		projectConfigManager:      &project.ConfigManager{},
// 		operatorPrivateKeyECDSA:   "",
// 		operatorPrivateKeyED25519: "",
// 	}

// 	t.Run("TaskStateLogNil", func(t *testing.T) {
// 		d.handleP2PData(&p2p.Data{
// 			Task:         nil,
// 			TaskStateLog: nil,
// 		}, nil)
// 	})

// 	data := &p2p.Data{
// 		Task: nil,
// 		TaskStateLog: &types.TaskStateLog{
// 			TaskID:    1,
// 			State:     types.TaskStatePacked,
// 			Comment:   "Comment",
// 			CreatedAt: time.Now(),
// 		},
// 	}

// 	t.Run("FailedToCreateTaskStateLog", func(t *testing.T) {
// 		p := NewPatches()
// 		defer p.Reset()
// 		p = p.ApplyMethodReturn(&mockPersistence{}, "Create", errors.New(t.Name()))
// 		d.handleP2PData(data, nil)
// 	})

// 	t.Run("NotTaskStateProved", func(t *testing.T) {
// 		p := NewPatches()
// 		defer p.Reset()
// 		p = p.ApplyMethodReturn(&mockPersistence{}, "Create", nil)

// 		d.handleP2PData(data, nil)
// 	})

// 	data = &p2p.Data{
// 		Task: nil,
// 		TaskStateLog: &types.TaskStateLog{
// 			TaskID:    1,
// 			State:     types.TaskStateProved,
// 			Comment:   "Comment",
// 			CreatedAt: time.Now(),
// 		},
// 	}

// 	t.Run("FailedToGetProject", func(t *testing.T) {
// 		p := NewPatches()
// 		defer p.Reset()
// 		p = p.ApplyMethodReturn(&mockPersistence{}, "Create", nil)
// 		p = p.ApplyMethodReturn(&types.TaskStateLog{}, "VerifySignature", nil)
// 		p = testproject.ProjectConfigManagerGet(p, nil, errors.New(t.Name()))
// 		d.handleP2PData(data, nil)
// 	})

// 	t.Run("FailedToInitOutputAndFailedToCreateTaskStateLog", func(t *testing.T) {
// 		p := NewPatches()
// 		defer p.Reset()

// 		outputCell := []OutputCell{
// 			{Values: Params{nil}},
// 			{Values: Params{errors.New(t.Name())}},
// 		}
// 		p = p.ApplyMethodSeq(&mockPersistence{}, "Create", outputCell)
// 		p = p.ApplyMethodReturn(&types.TaskStateLog{}, "VerifySignature", nil)
// 		p = testproject.ProjectConfigManagerGet(p, &project.Config{}, nil)
// 		p = p.ApplyFuncReturn(output.New, nil, errors.New(t.Name()))

// 		d.handleP2PData(data, nil)
// 	})

// 	t.Run("FailedToInitOutput", func(t *testing.T) {
// 		p := NewPatches()
// 		defer p.Reset()

// 		outputCell := []OutputCell{
// 			{Values: Params{nil}},
// 			{Values: Params{nil}},
// 		}
// 		p = p.ApplyMethodSeq(&mockPersistence{}, "Create", outputCell)
// 		p = p.ApplyMethodReturn(&types.TaskStateLog{}, "VerifySignature", nil)
// 		p = testproject.ProjectConfigManagerGet(p, &project.Config{}, nil)
// 		p = p.ApplyFuncReturn(output.New, nil, errors.New(t.Name()))

// 		d.handleP2PData(data, nil)
// 	})

// 	t.Run("FailedToOutputAndFailedToCreateTaskStateLog", func(t *testing.T) {
// 		p := NewPatches()
// 		defer p.Reset()
// 		outputCell := []OutputCell{
// 			{Values: Params{nil}},
// 			{Values: Params{errors.New(t.Name())}},
// 		}
// 		p = p.ApplyMethodSeq(&mockPersistence{}, "Create", outputCell)
// 		p = p.ApplyMethodReturn(&types.TaskStateLog{}, "VerifySignature", nil)
// 		p = testproject.ProjectConfigManagerGet(p, &project.Config{}, nil)
// 		p = p.ApplyFuncReturn(output.New, op, nil)

// 		op.EXPECT().Output(gomock.Any(), gomock.Any(), gomock.Any()).Return("", errors.New(t.Name())).Times(1)
// 		d.handleP2PData(data, nil)
// 	})

// 	t.Run("FailedToOutput", func(t *testing.T) {
// 		p := NewPatches()
// 		defer p.Reset()
// 		outputCell := []OutputCell{
// 			{Values: Params{nil}},
// 			{Values: Params{nil}},
// 		}
// 		p = p.ApplyMethodSeq(&mockPersistence{}, "Create", outputCell)
// 		p = p.ApplyMethodReturn(&types.TaskStateLog{}, "VerifySignature", nil)
// 		p = testproject.ProjectConfigManagerGet(p, &project.Config{}, nil)
// 		p = p.ApplyFuncReturn(output.New, op, nil)

// 		op.EXPECT().Output(gomock.Any(), gomock.Any(), gomock.Any()).Return("", errors.New(t.Name())).Times(1)
// 		d.handleP2PData(data, nil)
// 	})

// 	t.Run("FailedToCreateOutputtedTaskState", func(t *testing.T) {
// 		p := NewPatches()
// 		defer p.Reset()
// 		outputCell := []OutputCell{
// 			{Values: Params{nil}},
// 			{Values: Params{errors.New(t.Name())}},
// 		}
// 		p = p.ApplyMethodSeq(&mockPersistence{}, "Create", outputCell)
// 		p = p.ApplyMethodReturn(&types.TaskStateLog{}, "VerifySignature", nil)
// 		p = testproject.ProjectConfigManagerGet(p, &project.Config{}, nil)
// 		p = p.ApplyFuncReturn(output.New, op, nil)
// 		op.EXPECT().Output(gomock.Any(), gomock.Any(), gomock.Any()).Return("", nil).Times(1)

// 		d.handleP2PData(data, nil)
// 	})

// 	t.Run("HandleP2PDataSuccess", func(t *testing.T) {
// 		p := NewPatches()
// 		defer p.Reset()
// 		outputCell := []OutputCell{
// 			{Values: Params{nil}},
// 			{Values: Params{nil}},
// 		}
// 		p = p.ApplyMethodSeq(&mockPersistence{}, "Create", outputCell)
// 		p = p.ApplyMethodReturn(&types.TaskStateLog{}, "VerifySignature", nil)
// 		p = testproject.ProjectConfigManagerGet(p, &project.Config{}, nil)
// 		p = p.ApplyFuncReturn(output.New, op, nil)
// 		op.EXPECT().Output(gomock.Any(), gomock.Any(), gomock.Any()).Return("", nil).Times(1)

// 		d.handleP2PData(data, nil)
// 	})
// }

// func TestDispatcher_DispatchTask(t *testing.T) {
// 	r := require.New(t)

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	d := &Dispatcher{
// 		pubSubs:                   nil,
// 		persistence:               nil,
// 		projectConfigManager:      nil,
// 		operatorPrivateKeyECDSA:   "",
// 		operatorPrivateKeyED25519: "",
// 	}

// 	t.Run("FailedToRetrieve", func(t *testing.T) {
// 		d.datasource = &mockDatasourceErr{errors.New(t.Name())}
// 		_, err := d.dispatchTask(uint64(0x1), []byte("any"))
// 		r.ErrorContains(err, t.Name())
// 	})

// 	t.Run("NilTask", func(t *testing.T) {
// 		d.datasource = &mockDatasourceNil{}
// 		taskId, err := d.dispatchTask(uint64(0x1), []byte("any"))
// 		r.NoError(err)
// 		r.Equal(uint64(0x1), taskId)
// 	})

// 	t.Run("FailedToVerifyTaskSign", func(t *testing.T) {
// 		d.datasource = &mockDatasourceSuccess{&types.Task{}}

// 		p := NewPatches()
// 		defer p.Reset()

// 		p = p.ApplyMethodReturn(&types.Task{}, "VerifySignature", errors.New(t.Name()))

// 		_, err := d.dispatchTask(uint64(0x1), []byte("any"))
// 		r.ErrorContains(err, t.Name())
// 	})

// 	t.Run("FailedToAdd", func(t *testing.T) {
// 		p := NewPatches()
// 		defer p.Reset()

// 		d.datasource = &mockDatasourceSuccess{&types.Task{ID: uint64(1)}}

// 		p = p.ApplyMethodReturn(&types.Task{}, "VerifySignature", errors.New(t.Name()))
// 		p = p.ApplyMethodReturn(&p2p.PubSubs{}, "Add", errors.New(t.Name()))
// 		_, err := d.dispatchTask(uint64(0x1), []byte("any"))
// 		r.ErrorContains(err, t.Name())
// 	})

// 	t.Run("FailedToPublish", func(t *testing.T) {
// 		p := NewPatches()
// 		defer p.Reset()

// 		p = p.ApplyMethodReturn(&types.Task{}, "VerifySignature", errors.New(t.Name()))
// 		p = p.ApplyMethodReturn(&p2p.PubSubs{}, "Add", nil)
// 		p = p.ApplyMethodReturn(&p2p.PubSubs{}, "Publish", errors.New(t.Name()))
// 		_, err := d.dispatchTask(uint64(0x1), []byte("any"))
// 		r.ErrorContains(err, t.Name())
// 	})

// 	t.Run("DispatchTaskSuccess", func(t *testing.T) {
// 		p := NewPatches()
// 		defer p.Reset()

// 		p = p.ApplyMethodReturn(&types.Task{}, "VerifySignature", errors.New(t.Name()))
// 		p = p.ApplyMethodReturn(&p2p.PubSubs{}, "Add", nil)
// 		p = p.ApplyMethodReturn(&p2p.PubSubs{}, "Publish", nil)
// 		taskId, err := d.dispatchTask(uint64(0x1), []byte("any"))
// 		r.NoError(err)
// 		r.Equal(uint64(1)+1, taskId)
// 	})
// }
