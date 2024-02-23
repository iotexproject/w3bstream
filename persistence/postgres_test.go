package persistence

import (
	. "github.com/agiledragon/gomonkey/v2"
	"github.com/google/uuid"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/testutil"
	"github.com/machinefi/sprout/types"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"reflect"
	"testing"
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

func PatchTxCreateTaskLog(p *Patches, err error) *Patches {
	return p.ApplyFunc(
		txCreateTaskLog,
		func(_ *gorm.DB, _ string) error {
			return err
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

func PatchPostgresFetchMessagesByMessageIDs(p *Patches, messages []*message, err error) *Patches {
	return p.ApplyMethod(
		_targetPersistencePostgres,
		"FetchMessagesByMessageIDs",
		func(_ *Postgres, _ ...string) ([]*message, error) {
			return messages, err
		},
	)
}

func TestPostgres_Save(t *testing.T) {
	r := require.New(t)
	p := NewPatches()

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
			p = PatchTxCreateTaskLog(p, errors.New(t.Name()))
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

func TestPostgres_FetchByID(t *testing.T) {
	t.SkipNow()
	r := require.New(t)
	p := NewPatches()

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
	v := &Postgres{db: &gorm.DB{}}

	p = testutil.GormDBWhere(p, v.db)

	t.Run("FailedToFind", func(t *testing.T) {
		v.db.Error = errors.New(t.Name())
		p = testutil.GormDBFind(p, nil, v.db)
		_, err := v.FetchMessage("any")
		r.ErrorContains(err, t.Name())
	})

	v.db.Error = nil
	p = testutil.GormDBFind(p, []*message{}, v.db)
	_, err := v.FetchMessage("any")
	r.NoError(err)
}

func TestPostgres_FetchStateLog(t *testing.T) {
	// r := require.New(t)
	// p := NewPatches()
	// v := &Postgres{db: &gorm.DB{}}

	// t.Run("FailedToFetchTasks", func(t *testing.T) {
	// 	p = PatchFetchTasksByMessageID()
	// })

	// t.Run("FetchTasksListEmpty", func(t *testing.T) {
	// })

	// t.Run("FailedToFetchTaskStateLogs", func(t *testing.T) {
	// })

	// t.Run("Success", func(t *testing.T) {
	// })

}

func TestPostgres_UpdateState(t *testing.T) {

}

func TestNewPostgres(t *testing.T) {

}
