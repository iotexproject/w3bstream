package task

import (
	"testing"
	"time"

	. "github.com/agiledragon/gomonkey/v2"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/persistence"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/testutil/mock"
	testp2p "github.com/machinefi/sprout/testutil/p2p"
	testproject "github.com/machinefi/sprout/testutil/project"
	"github.com/machinefi/sprout/types"
)

func TestNewDispatcher(t *testing.T) {
	r := require.New(t)

	t.Run("NewFailed", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		p = testp2p.P2pNewPubSubs(p, nil, errors.New(t.Name()))
		_, err := NewDispatcher(nil, nil, nil, "", "", "", 0)
		r.ErrorContains(err, t.Name())
	})

	t.Run("New", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()
		p = testp2p.P2pNewPubSubs(p, nil, nil)

		_, err := NewDispatcher(nil, nil, nil, "", "", "", 0)
		r.NoError(err)
	})
}

func TestDispatcher_HandleP2PData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	op := mock.NewMockOutput(ctrl)

	d := &Dispatcher{
		pubSubs:                   nil,
		persistence:               &persistence.Postgres{},
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
			Comment:   "Comment",
			CreatedAt: time.Now(),
		},
	}

	t.Run("FailedToCreateTaskStateLog", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()
		p = p.ApplyMethodReturn(&persistence.Postgres{}, "Create", errors.New(t.Name()))
		d.handleP2PData(data, nil)
	})

	t.Run("NotTaskStateProved", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()
		p = p.ApplyMethodReturn(&persistence.Postgres{}, "Create", nil)

		d.handleP2PData(data, nil)
	})

	data.TaskStateLog.State = types.TaskStateProved

	t.Run("FailedToGetProject", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()
		p = p.ApplyMethodReturn(&persistence.Postgres{}, "Create", nil)

		p = testproject.ProjectManagerGet(p, nil, errors.New(t.Name()))
		d.handleP2PData(data, nil)
	})

	t.Run("FailedToInitOutputAndFailedToCreateTaskStateLog", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		outputCell := []OutputCell{
			{Values: Params{nil}},
			{Values: Params{errors.New(t.Name())}},
		}
		p = p.ApplyMethodSeq(&persistence.Postgres{}, "Create", outputCell)
		p = testproject.ProjectManagerGet(p, &project.Config{}, nil)

		p = p.ApplyMethodReturn(&project.Config{}, "GetOutput", nil, errors.New(t.Name()))
		d.handleP2PData(data, nil)
	})

	t.Run("FailedToInitOutput", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		outputCell := []OutputCell{
			{Values: Params{nil}},
			{Values: Params{nil}},
		}
		p = p.ApplyMethodSeq(&persistence.Postgres{}, "Create", outputCell)
		p = testproject.ProjectManagerGet(p, &project.Config{}, nil)

		p = p.ApplyMethodReturn(&project.Config{}, "GetOutput", nil, errors.New(t.Name()))
		d.handleP2PData(data, nil)
	})

	t.Run("FailedToOutputAndFailedToCreateTaskStateLog", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()
		outputCell := []OutputCell{
			{Values: Params{nil}},
			{Values: Params{errors.New(t.Name())}},
		}
		p = p.ApplyMethodSeq(&persistence.Postgres{}, "Create", outputCell)
		p = testproject.ProjectManagerGet(p, &project.Config{}, nil)
		p = p.ApplyMethodReturn(&project.Config{}, "GetOutput", op, nil)

		op.EXPECT().Output(gomock.Any(), gomock.Any()).Return("", errors.New(t.Name())).Times(1)
		d.handleP2PData(data, nil)
	})

	t.Run("FailedToOutput", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()
		outputCell := []OutputCell{
			{Values: Params{nil}},
			{Values: Params{nil}},
		}
		p = p.ApplyMethodSeq(&persistence.Postgres{}, "Create", outputCell)
		p = testproject.ProjectManagerGet(p, &project.Config{}, nil)
		p = p.ApplyMethodReturn(&project.Config{}, "GetOutput", op, nil)

		op.EXPECT().Output(gomock.Any(), gomock.Any()).Return("", errors.New(t.Name())).Times(1)
		d.handleP2PData(data, nil)
	})

	t.Run("FailedToCreateOutputtedTaskState", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()
		outputCell := []OutputCell{
			{Values: Params{nil}},
			{Values: Params{errors.New(t.Name())}},
		}
		p = p.ApplyMethodSeq(&persistence.Postgres{}, "Create", outputCell)
		p = testproject.ProjectManagerGet(p, &project.Config{}, nil)
		p = p.ApplyMethodReturn(&project.Config{}, "GetOutput", op, nil)
		op.EXPECT().Output(gomock.Any(), gomock.Any()).Return("", nil).Times(1)

		d.handleP2PData(data, nil)
	})

	t.Run("HandleP2PDataSuccess", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()
		outputCell := []OutputCell{
			{Values: Params{nil}},
			{Values: Params{nil}},
		}
		p = p.ApplyMethodSeq(&persistence.Postgres{}, "Create", outputCell)
		p = testproject.ProjectManagerGet(p, &project.Config{}, nil)
		p = p.ApplyMethodReturn(&project.Config{}, "GetOutput", op, nil)
		op.EXPECT().Output(gomock.Any(), gomock.Any()).Return("", nil).Times(1)

		d.handleP2PData(data, nil)
	})
}

func TestDispatcher_DispatchTask(t *testing.T) {
	r := require.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ds := mock.NewMockDatasource(ctrl)

	d := &Dispatcher{
		datasource:                ds,
		pubSubs:                   nil,
		persistence:               nil,
		projectManager:            nil,
		operatorPrivateKeyECDSA:   "",
		operatorPrivateKeyED25519: "",
	}

	t.Run("FailedToRetrieve", func(t *testing.T) {
		ds.EXPECT().Retrieve(gomock.Any()).Return(nil, errors.New(t.Name())).Times(1)
		_, err := d.dispatchTask(uint64(0x1))
		r.ErrorContains(err, t.Name())
	})

	t.Run("NilTask", func(t *testing.T) {
		ds.EXPECT().Retrieve(gomock.Any()).Return(nil, nil).Times(1)
		taskId, err := d.dispatchTask(uint64(0x1))
		r.NoError(err)
		r.Equal(uint64(0x1), taskId)
	})

	t.Run("FailedToAdd", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		ds.EXPECT().Retrieve(gomock.Any()).Return(&types.Task{ProjectID: uint64(0x1)}, nil).Times(1)
		p = p.ApplyMethodReturn(&p2p.PubSubs{}, "Add", errors.New(t.Name()))
		_, err := d.dispatchTask(uint64(0x1))
		r.ErrorContains(err, t.Name())
	})

	t.Run("FailedToPublish", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		ds.EXPECT().Retrieve(gomock.Any()).Return(&types.Task{ProjectID: uint64(0x1)}, nil).Times(1)
		p = p.ApplyMethodReturn(&p2p.PubSubs{}, "Add", nil)
		p = p.ApplyMethodReturn(&p2p.PubSubs{}, "Publish", errors.New(t.Name()))
		_, err := d.dispatchTask(uint64(0x1))
		r.ErrorContains(err, t.Name())
	})

	t.Run("DispatchTaskSuccess", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		ds.EXPECT().Retrieve(gomock.Any()).Return(&types.Task{ID: uint64(0x1), ProjectID: uint64(0x1)}, nil).Times(1)
		p = p.ApplyMethodReturn(&p2p.PubSubs{}, "Add", nil)
		p = p.ApplyMethodReturn(&p2p.PubSubs{}, "Publish", nil)
		taskId, err := d.dispatchTask(uint64(0x1))
		r.NoError(err)
		r.Equal(uint64(0x1)+1, taskId)
	})
}

func TestDispatcher_Dispatch(t *testing.T) {
	p := NewPatches()
	defer p.Reset()

	d := &Dispatcher{}

	t.Run("FailedToDispatchTask", func(t *testing.T) {
		ch := make(chan time.Time, 1)
		ticker := &time.Timer{C: ch}
		go func() { ch <- time.Now() }()
		p = p.ApplyFuncReturn(time.NewTimer, ticker)
		p = p.ApplyPrivateMethod(d, "dispatchTask", func(nextTaskID uint64) (uint64, error) {
			return 0, errors.New(t.Name())
		})
		go d.Dispatch(uint64(0x1))
		time.Sleep(1 * time.Second)
		close(ch)
	})

	t.Run("DispatchTaskSuccess", func(t *testing.T) {
		ch := make(chan time.Time, 1)
		ticker := &time.Timer{C: ch}
		go func() { ch <- time.Now() }()
		p = p.ApplyFuncReturn(time.NewTimer, ticker)
		p = p.ApplyPrivateMethod(d, "dispatchTask", func(nextTaskID uint64) (uint64, error) {
			return 0, nil
		})
		go d.Dispatch(uint64(0x1))
		time.Sleep(1 * time.Second)
		close(ch)
	})
}
