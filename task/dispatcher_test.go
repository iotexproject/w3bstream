package task

import (
	"log/slog"
	"reflect"
	"testing"
	"time"

	. "github.com/agiledragon/gomonkey/v2"
	. "github.com/bytedance/mockey"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/persistence"
	"github.com/machinefi/sprout/testutil/mock"
	testp2p "github.com/machinefi/sprout/testutil/p2p"
	testpersistence "github.com/machinefi/sprout/testutil/persistence"
	testproject "github.com/machinefi/sprout/testutil/project"
	"github.com/machinefi/sprout/types"
)

func TestPubTask(t *testing.T) {

	d := &Dispatcher{}
	task := &types.Task{
		ID: "",
		Messages: []*types.Message{{
			ID:             "id1",
			ProjectID:      uint64(0x1),
			ProjectVersion: "0.1",
			Data:           "data",
		}},
	}

	PatchConvey("GetTaskFailed", t, func() {
		Mock((*persistence.Postgres).Fetch).Return(nil, errors.New(t.Name())).Build()
		mockerAdd := Mock((*p2p.PubSubs).Add).Return(nil).Build()
		mockerLog := Mock(slog.Error).Return().Build()
		d.pubTask()
		So(mockerAdd.Times(), ShouldEqual, 0)
		So(mockerLog.Times(), ShouldEqual, 1)
	})

	PatchConvey("TaskNil", t, func() {
		Mock((*persistence.Postgres).Fetch).Return(nil, nil).Build()
		mockerLog := Mock(slog.Error).Return().Build()
		mockerAdd := Mock((*p2p.PubSubs).Add).Return(nil).Build()
		d.pubTask()
		So(mockerLog.Times(), ShouldEqual, 0)
		So(mockerAdd.Times(), ShouldEqual, 0)
	})

	PatchConvey("AddPubsubFailed", t, func() {
		Mock((*persistence.Postgres).Fetch).Return(task, nil).Build()
		Mock((*p2p.PubSubs).Add).Return(errors.New(t.Name())).Build()
		mockerLog := Mock(slog.Error).Return().Build()
		mockerPub := Mock((*p2p.PubSubs).Publish).Return(nil).Build()
		d.pubTask()
		So(mockerLog.Times(), ShouldEqual, 1)
		So(mockerPub.Times(), ShouldEqual, 0)
	})

	PatchConvey("PublishFailed", t, func() {
		Mock((*persistence.Postgres).Fetch).Return(task, nil).Build()
		Mock((*p2p.PubSubs).Add).Return(nil).Build()
		Mock((*p2p.PubSubs).Publish).Return(errors.New(t.Name())).Build()
		mockerLog := Mock(slog.Error).Return().Build()
		d.pubTask()
		So(mockerLog.Times(), ShouldEqual, 1)
	})
}

func TestNewDispatcher(t *testing.T) {
	require := require.New(t)
	patches := NewPatches()

	t.Run("NewFailed", func(t *testing.T) {
		patches = testp2p.P2pNewPubSubs(patches, nil, errors.New(t.Name()))
		_, err := NewDispatcher(nil, nil, "", "", "", 0)
		require.ErrorContains(err, t.Name())
	})
	patches = testp2p.P2pNewPubSubs(patches, nil, nil)

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
		patches = testpersistence.PersistencePostgresUpdateState(patches, errors.New(t.Name()))
		d.handleP2PData(data, nil)
	})
	patches = testpersistence.PersistencePostgresUpdateState(patches, nil)

	t.Run("TaskStateProved", func(t *testing.T) {
		d.handleP2PData(data, nil)
	})

	t.Run("FetchTaskFailed", func(t *testing.T) {
		data.TaskStateLog.State = types.TaskStateProved
		patches = testpersistence.PersistencePostgresFetchByID(patches, nil, errors.New(t.Name()))
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
	patches = testpersistence.PersistencePostgresFetchByID(patches, task, nil)

	t.Run("GetProjectFailed", func(t *testing.T) {
		patches = testproject.ProjectManagerGet(patches, nil, errors.New(t.Name()))
		d.handleP2PData(data, nil)
	})
	patches = testproject.ProjectManagerGet(patches, nil, nil)

	t.Run("InitOutputFailed", func(t *testing.T) {
		patches = testproject.ProjectConfigGetOutput(patches, nil, errors.New(t.Name()))
		d.handleP2PData(data, nil)
	})

	t.Run("InitOutputFailedAndUpdateStateFailed", func(t *testing.T) {
		patches = testproject.ProjectConfigGetOutput(patches, nil, errors.New(t.Name()))
		patches = persistencePostgresUpdateStateSeq(patches, []OutputCell{
			{Values: Params{nil}},
			{Values: Params{errors.New(t.Name())}},
		})

		d.handleP2PData(data, nil)
	})
	ctrl := gomock.NewController(t)
	ot := mock.NewMockOutput(ctrl)
	patches = testproject.ProjectConfigGetOutput(patches, ot, nil)

	t.Run("OutputFailed", func(t *testing.T) {
		ot.EXPECT().Output(gomock.Any(), gomock.Any()).Return("", errors.New("output failed")).Times(1)
		patches = testpersistence.PersistencePostgresUpdateState(patches, nil)
		d.handleP2PData(data, nil)
	})

	t.Run("OutputFailedAndUpdateStateFailed", func(t *testing.T) {
		ot.EXPECT().Output(gomock.Any(), gomock.Any()).Return("", errors.New("output failed")).Times(1)
		patches = persistencePostgresUpdateStateSeq(patches, []OutputCell{
			{Values: Params{nil}},
			{Values: Params{errors.New(t.Name())}},
		})
		d.handleP2PData(data, nil)
	})

	t.Run("StateOutputtedFailed", func(t *testing.T) {
		ot.EXPECT().Output(gomock.Any(), gomock.Any()).Return("outRes", nil).Times(1)
		patches = persistencePostgresUpdateStateSeq(patches, []OutputCell{
			{Values: Params{nil}},
			{Values: Params{errors.New(t.Name())}},
		})
		d.handleP2PData(data, nil)
	})

	t.Run("HandleOK", func(t *testing.T) {
		ot.EXPECT().Output(gomock.Any(), gomock.Any()).Return("outRes", nil).Times(1)
		patches = testpersistence.PersistencePostgresUpdateState(patches, nil)
		d.handleP2PData(data, nil)
	})
}

func persistencePostgresUpdateStateSeq(p *Patches, outCell []OutputCell) *Patches {
	var pg *persistence.Postgres
	return p.ApplyMethodSeq(
		reflect.TypeOf(pg),
		"UpdateState",
		outCell,
	)
}
