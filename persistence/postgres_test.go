package persistence

import (
	"reflect"
	"testing"
	"time"

	. "github.com/agiledragon/gomonkey/v2"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/testutil"
	"github.com/machinefi/sprout/types"
)

func PatchNewPostgres(p *Patches, v *Postgres, err error) *Patches {
	return p.ApplyFunc(
		NewPostgres,
		func(ep string) (*Postgres, error) {
			return v, err
		},
	)
}

func PatchTxAggregateTask(p *Patches, taskID string, err error) *Patches {
	return p.ApplyFunc(
		txAggregateTask,
		func(_ *gorm.DB, _ int, _ *types.Message) (string, error) {
			return taskID, err
		},
	)
}

func PatchTxCreateTaskLog(p *Patches, v *taskStateLog, err error) *Patches {
	return p.ApplyFunc(
		txCreateTaskLog,
		func(_ *gorm.DB, _ string, _ types.TaskState) (*taskStateLog, error) {
			return v, err
		},
	)
}

func PatchTxCreateMessage(p *Patches, msg *message, err error) *Patches {
	return p.ApplyFunc(
		txCreateMessage,
		func(_ *gorm.DB, _ *types.Message) (*message, error) {
			return msg, err
		},
	)
}

var _targetPersistencePostgres = reflect.TypeOf(&Postgres{})

func PatchPostgresFetchByID(p *Patches, task *types.Task, err error) *Patches {
	return p.ApplyMethod(
		_targetPersistencePostgres,
		"FetchByID",
		func(_ *Postgres) (*types.Task, error) {
			return task, err
		},
	)
}

func PatchPostgresFetchTasksByTaskID(p *Patches, tasks []*task, err error) *Patches {
	return p.ApplyMethod(
		_targetPersistencePostgres,
		"FetchTasksByTaskID",
		func(_ *Postgres, _ string) ([]*task, error) {
			return tasks, err
		},
	)
}

func PatchPostgresFetchTasksByMessageID(p *Patches, tasks []*task, err error) *Patches {
	return p.ApplyMethod(
		_targetPersistencePostgres,
		"FetchTasksByMessageID",
		func(_ *Postgres, _ string) ([]*task, error) {
			return tasks, err
		},
	)
}

func PatchPostgresFetchTaskStateLogsByTaskIDs(p *Patches, logs []*taskStateLog, err error) *Patches {
	return p.ApplyMethod(
		_targetPersistencePostgres,
		"FetchTaskStateLogsByTaskIDs",
		func(_ *Postgres, _ ...string) ([]*taskStateLog, error) {
			return logs, err
		},
	)
}

func PatchPostgresFetchMessagesByMessageIDs(p *Patches, messages []*message, err error) *Patches {
	return p.ApplyMethod(
		_targetPersistencePostgres,
		"FetchMessagesByMessageIDs",
		func(_ *Postgres, _ ...string) ([]*message, error) {
			return messages, err
		},
	)
}

func TestTxCreateMessage(t *testing.T) {
	r := require.New(t)
	p := NewPatches()
	defer p.Reset()
	d := &gorm.DB{}

	t.Run("FailedToCreate", func(t *testing.T) {
		p = testutil.GormDBCreate(p, nil, &gorm.DB{Error: errors.New(t.Name())})
		_, err := txCreateMessage(d, &types.Message{})
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p = testutil.GormDBCreate(p, &message{}, &gorm.DB{Error: nil})
		v, err := txCreateMessage(d, &types.Message{})
		r.NotNil(v)
		r.NoError(err)
	})
}

func TestTxCreateTaskLog(t *testing.T) {
	r := require.New(t)
	p := NewPatches()
	defer p.Reset()
	d := &gorm.DB{}

	t.Run("FailedToCreate", func(t *testing.T) {
		p = testutil.GormDBCreate(p, nil, &gorm.DB{Error: errors.New(t.Name())})
		v, err := txCreateTaskLog(d, "any", types.TaskStatePacked)
		r.Nil(v)
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p = testutil.GormDBCreate(p, &taskStateLog{}, &gorm.DB{Error: nil})
		v, err := txCreateTaskLog(d, "any", types.TaskStatePacked)
		r.NotNil(v)
		r.NoError(err)
	})
}

func TestTxAggregateTask(t *testing.T) {
	r := require.New(t)
	p := NewPatches()
	defer p.Reset()
	d := &gorm.DB{}

	t.Run("FailedToFindMessages", func(t *testing.T) {
		p = testutil.GormDBClauses(p, &gorm.DB{})
		p = testutil.GormDBOrder(p, &gorm.DB{})
		p = testutil.GormDBWhere(p, &gorm.DB{})
		p = testutil.GormDBLimit(p, &gorm.DB{})
		p = testutil.GormDBModel(p, &gorm.DB{Error: nil})

		p = testutil.GormDBFind(p, nil, &gorm.DB{Error: errors.New(t.Name())})
		taskID, err := txAggregateTask(d, 100, &types.Message{})
		r.Equal(len(taskID), 0)
		r.ErrorContains(err, t.Name())
	})

	t.Run("FetchedMessageListEmpty", func(t *testing.T) {
		p = testutil.GormDBFind(p, &([]*message{}), &gorm.DB{Error: nil})
		taskID, err := txAggregateTask(d, 0, &types.Message{})
		r.Equal(len(taskID), 0)
		r.NoError(err)
	})

	t.Run("NoEnoughMessagesForPacking", func(t *testing.T) {
		p = testutil.GormDBFind(p, &([]*message{{}}), &gorm.DB{Error: nil})
		taskID, err := txAggregateTask(d, 100, &types.Message{})
		r.Equal(len(taskID), 0)
		r.NoError(err)
	})

	t.Run("FailedToUpdateMessages", func(t *testing.T) {
		p = testutil.GormDBFind(p, &([]*message{{}, {}}), &gorm.DB{Error: nil})
		p = testutil.GormDBUpdate(p, &gorm.DB{Error: errors.New(t.Name())})

		taskID, err := txAggregateTask(d, 1, &types.Message{})
		r.Equal(len(taskID), 0)
		r.ErrorContains(err, t.Name())
	})

	t.Run("FailedToCreateTasks", func(t *testing.T) {
		p = testutil.GormDBUpdate(p, &gorm.DB{Error: nil})
		p = testutil.GormDBCreate(p, &([]*task{}), &gorm.DB{Error: errors.New(t.Name())})
		taskID, err := txAggregateTask(d, 1, &types.Message{})
		r.Equal(len(taskID), 0)
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p = testutil.GormDBCreate(p, &([]*task{}), &gorm.DB{Error: nil})
		taskID, err := txAggregateTask(d, 1, &types.Message{})
		r.Greater(len(taskID), 0)
		r.NoError(err)
	})
}

func TestPostgres_Save(t *testing.T) {
	r := require.New(t)
	p := NewPatches()
	defer p.Reset()

	db := &gorm.DB{
		Error:     nil,
		Statement: &gorm.Statement{},
	}
	p = PatchNewPostgres(p, &Postgres{db: db}, nil)
	v, err := NewPostgres("any")
	r.NoError(err)
	r.NotNil(v)

	// make sure gorm tx executed ok
	p = testutil.GormDBBegin(p, db)
	p = testutil.GormDBCommit(p, db)
	p = testutil.GormDBClauses(p, db)
	p = testutil.GormDBOrder(p, db)
	p = testutil.GormDBWhere(p, db)
	p = testutil.GormDBLimit(p, db)
	p = testutil.GormDBRollback(p, db)

	t.Run("FailedToCreateMessage", func(t *testing.T) {
		p = PatchTxCreateMessage(p, nil, errors.New(t.Name()))
		err := v.Save(&types.Message{}, &project.Config{})
		r.ErrorContains(err, t.Name())
	})
	p = PatchTxCreateMessage(p, &message{}, nil)

	t.Run("InTx", func(t *testing.T) {
		t.Run("FailedToAggregateTask", func(t *testing.T) {
			p = PatchTxAggregateTask(p, "", errors.New(t.Name()))
			err := v.Save(&types.Message{}, &project.Config{})
			r.ErrorContains(err, t.Name())
		})

		t.Run("FailedToCreateTaskLog", func(t *testing.T) {
			taskID := uuid.NewString()
			p = PatchTxAggregateTask(p, taskID, nil)
			p = PatchTxCreateTaskLog(p, nil, errors.New(t.Name()))
			err := v.Save(&types.Message{}, &project.Config{})
			r.ErrorContains(err, t.Name())
		})

		t.Run("NoNeedCreateTaskLog", func(t *testing.T) {
			p = PatchTxAggregateTask(p, "", nil)
			err := v.Save(&types.Message{}, &project.Config{})
			r.NoError(err)
		})
	})
}

func TestPostgres_Fetch(t *testing.T) {
	r := require.New(t)
	p := NewPatches()
	defer p.Reset()

	v := &Postgres{
		db: &gorm.DB{
			Error:     nil,
			Statement: &gorm.Statement{},
		},
	}

	p = testutil.GormDBWhere(p, v.db)

	t.Run("FailedDBFind", func(t *testing.T) {
		p = testutil.GormDBFirst(p, &task{}, &gorm.DB{Error: errors.New(t.Name())})
		_task, err := v.Fetch()
		r.Nil(_task)
		r.ErrorContains(err, t.Name())
	})

	t.Run("NotFoundError", func(t *testing.T) {
		p = testutil.GormDBFirst(p, &task{}, &gorm.DB{Error: gorm.ErrRecordNotFound})
		_task, err := v.Fetch()
		r.Nil(_task)
		r.NoError(err)
	})
	p = testutil.GormDBFirst(p, &task{}, &gorm.DB{Error: nil})

	t.Run("FailedToFetchByID", func(t *testing.T) {
		p = PatchPostgresFetchByID(p, nil, errors.New(t.Name()))
		_task, err := v.Fetch()
		r.Nil(_task)
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p = PatchPostgresFetchByID(p, &types.Task{}, nil)
		_task, err := v.Fetch()
		r.NotNil(_task)
		r.NoError(err)
	})
}

func TestPostgres_FetchTasksByTaskID(t *testing.T) {
	r := require.New(t)
	p := NewPatches()
	defer p.Reset()

	v := &Postgres{db: &gorm.DB{}}
	p = testutil.GormDBWhere(p, v.db)

	t.Run("FailedToFind", func(t *testing.T) {
		p = testutil.GormDBFind(p, &([]*task{}), &gorm.DB{Error: errors.New(t.Name())})
		vs, err := v.FetchTasksByTaskID("any")
		r.Nil(vs)
		r.ErrorContains(err, t.Name())
	})
	t.Run("Success", func(t *testing.T) {
		p = testutil.GormDBFind(p, &([]*task{}), &gorm.DB{Error: nil})
		_, err := v.FetchTasksByTaskID("any")
		r.NoError(err)
	})
}

func TestPostgres_FetchMessagesByMessageIDs(t *testing.T) {
	r := require.New(t)
	p := NewPatches()
	defer p.Reset()

	v := &Postgres{db: &gorm.DB{}}
	p = testutil.GormDBWhere(p, v.db)

	t.Run("EmptyMessageIDList", func(t *testing.T) {
		vs, err := v.FetchMessagesByMessageIDs()
		r.Nil(vs)
		r.NoError(err)
	})

	t.Run("FailedToFind", func(t *testing.T) {
		p = testutil.GormDBFind(p, &([]*message{}), &gorm.DB{Error: errors.New(t.Name())})
		vs, err := v.FetchMessagesByMessageIDs("any")
		r.Nil(vs)
		r.ErrorContains(err, t.Name())
	})
	t.Run("Success", func(t *testing.T) {
		p = testutil.GormDBFind(p, &([]*message{}), &gorm.DB{Error: nil})
		_, err := v.FetchMessagesByMessageIDs("any")
		r.NoError(err)
	})
}

func TestPostgres_FetchByID(t *testing.T) {
	r := require.New(t)
	p := NewPatches()
	defer p.Reset()

	v := &Postgres{db: &gorm.DB{}}

	t.Run("FailedToFetchTasks", func(t *testing.T) {
		p = PatchPostgresFetchTasksByTaskID(p, nil, errors.New(t.Name()))
		_task, err := v.FetchByID("any")
		r.Nil(_task)
		r.ErrorContains(err, t.Name())
	})

	t.Run("FetchTaskListEmpty", func(t *testing.T) {
		p = PatchPostgresFetchTasksByTaskID(p, nil, nil)
		_task, err := v.FetchByID("any")
		r.Nil(_task)
		r.NoError(err)
	})
	p.Reset()
	p = PatchPostgresFetchTasksByTaskID(p, []*task{{}}, nil)

	t.Run("FailedToFetchMessages", func(t *testing.T) {
		p = PatchPostgresFetchMessagesByMessageIDs(p, nil, errors.New(t.Name()))
		_task, err := v.FetchByID("any")
		r.Nil(_task)
		r.ErrorContains(err, t.Name())
	})

	t.Run("FetchMessageListEmpty", func(t *testing.T) {
		p = PatchPostgresFetchMessagesByMessageIDs(p, nil, nil)
		_task, err := v.FetchByID("any")
		r.Nil(_task)
		r.Error(err)
	})

	t.Run("Success", func(t *testing.T) {
		p = PatchPostgresFetchMessagesByMessageIDs(p, []*message{{}}, nil)
		_task, err := v.FetchByID("any")
		r.NotNil(_task)
		r.NoError(err)
	})
}

func TestPostgres_FetchMessage(t *testing.T) {
	r := require.New(t)
	p := NewPatches()
	defer p.Reset()
	v := &Postgres{db: &gorm.DB{}}

	p = testutil.GormDBWhere(p, v.db)

	t.Run("FailedToFind", func(t *testing.T) {
		v.db.Error = errors.New(t.Name())
		p = testutil.GormDBFind(p, nil, v.db)
		_, err := v.FetchMessage("any")
		r.ErrorContains(err, t.Name())
	})

	v.db.Error = nil
	p = testutil.GormDBFind(p, &[]*message{{}, {}}, v.db)
	_, err := v.FetchMessage("any")
	r.NoError(err)
}

func TestPostgres_FetchTasksByMessageID(t *testing.T) {
	r := require.New(t)
	p := NewPatches()
	defer p.Reset()

	v := &Postgres{db: &gorm.DB{}}
	p = testutil.GormDBWhere(p, v.db)

	t.Run("FailedToFind", func(t *testing.T) {
		p = testutil.GormDBFind(p, &([]*task{}), &gorm.DB{Error: errors.New(t.Name())})
		vs, err := v.FetchTasksByMessageID("any")
		r.Nil(vs)
		r.ErrorContains(err, t.Name())
	})
	t.Run("Success", func(t *testing.T) {
		p = testutil.GormDBFind(p, &([]*task{}), &gorm.DB{Error: nil})
		_, err := v.FetchTasksByMessageID("any")
		r.NoError(err)
	})
}

func TestPostgres_FetchTaskStateLogsByTaskIDs(t *testing.T) {
	r := require.New(t)
	p := NewPatches()
	defer p.Reset()

	v := &Postgres{db: &gorm.DB{}}
	p = testutil.GormDBWhere(p, v.db)
	p = testutil.GormDBOrder(p, v.db)

	t.Run("EmptyMessageIDList", func(t *testing.T) {
		vs, err := v.FetchTaskStateLogsByTaskIDs()
		r.Nil(vs)
		r.NoError(err)
	})

	t.Run("FailedToFind", func(t *testing.T) {
		p = testutil.GormDBFind(p, &([]*taskStateLog{}), &gorm.DB{Error: errors.New(t.Name())})
		vs, err := v.FetchTaskStateLogsByTaskIDs("any")
		r.Nil(vs)
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p = testutil.GormDBFind(p, &([]*taskStateLog{}), &gorm.DB{Error: nil})
		_, err := v.FetchTaskStateLogsByTaskIDs("any")
		r.NoError(err)
	})

}

func TestPostgres_FetchStateLog(t *testing.T) {
	r := require.New(t)
	p := NewPatches()
	defer p.Reset()
	v := &Postgres{db: &gorm.DB{}}

	t.Run("FailedToFetchTasks", func(t *testing.T) {
		p = PatchPostgresFetchTasksByMessageID(p, nil, errors.New(t.Name()))
		_, err := v.FetchStateLog("any")
		r.ErrorContains(err, t.Name())
	})

	t.Run("FetchTasksListEmpty", func(t *testing.T) {
		p = PatchPostgresFetchTasksByMessageID(p, []*task{}, nil)
		vs, err := v.FetchStateLog("any")
		r.Nil(vs)
		r.NoError(err)
	})

	t.Run("FailedToFetchTaskStateLogs", func(t *testing.T) {
		p = PatchPostgresFetchTasksByMessageID(p, []*task{{}}, nil)
		p = PatchPostgresFetchTaskStateLogsByTaskIDs(p, nil, errors.New(t.Name()))
		vs, err := v.FetchStateLog("any")
		r.Nil(vs)
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p = PatchPostgresFetchTaskStateLogsByTaskIDs(p, []*taskStateLog{{}}, nil)
		vs, err := v.FetchStateLog("any")
		r.NotNil(vs)
		r.NoError(err)
	})
}

func TestPostgres_UpdateState(t *testing.T) {
	r := require.New(t)
	p := NewPatches()
	defer p.Reset()
	d := &gorm.DB{Statement: &gorm.Statement{}}
	v := &Postgres{db: d}

	// make sure gorm tx executed ok
	p = testutil.GormDBBegin(p, d)
	p = testutil.GormDBCommit(p, d)
	p = testutil.GormDBClauses(p, d)
	p = testutil.GormDBOrder(p, d)
	p = testutil.GormDBWhere(p, d)
	p = testutil.GormDBLimit(p, d)
	p = testutil.GormDBRollback(p, d)

	t.Run("InTx", func(t *testing.T) {
		t.Run("FailedToUpdateState", func(t *testing.T) {
			p = testutil.GormDBModel(p, d)
			p = testutil.GormDBUpdate(p, &gorm.DB{Error: errors.New(t.Name())})
			err := v.UpdateState("any", types.TaskStatePacked, "any", time.Now())
			r.ErrorContains(err, t.Name())
		})
		t.Run("FailedToCreateStateLog", func(t *testing.T) {
			p = testutil.GormDBUpdate(p, &gorm.DB{Error: nil})
			p = testutil.GormDBCreate(p, nil, &gorm.DB{Error: errors.New(t.Name())})
			err := v.UpdateState("any", types.TaskStatePacked, "any", time.Now())
			r.ErrorContains(err, t.Name())
		})
		t.Run("Success", func(t *testing.T) {
			p = testutil.GormDBCreate(p, nil, &gorm.DB{Error: nil})
			err := v.UpdateState("any", types.TaskStatePacked, "any", time.Now())
			r.NoError(err)
		})
	})
}

func TestNewPostgres(t *testing.T) {
	r := require.New(t)
	p := NewPatches()
	defer p.Reset()
	d := &gorm.DB{}

	t.Run("FailedToOpenDSN", func(t *testing.T) {
		p = testutil.GormOpen(p, nil, errors.New(t.Name()))
		v, err := NewPostgres("any")
		r.Nil(v)
		r.ErrorContains(err, t.Name())
	})

	t.Run("FailedToMigrate", func(t *testing.T) {
		p = testutil.GormOpen(p, d, nil)
		p = testutil.GormDBAutoMigrate(p, errors.New(t.Name()))
		v, err := NewPostgres("any")
		r.Nil(v)
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p = testutil.GormOpen(p, d, nil)
		p = testutil.GormDBAutoMigrate(p, nil)
		v, err := NewPostgres("any")
		r.NotNil(v)
		r.NoError(err)
	})
}
