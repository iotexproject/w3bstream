package persistence

import (
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/machinefi/sprout/testutil"
	"github.com/machinefi/sprout/types"
)

func TestPostgres_Create(t *testing.T) {
	r := require.New(t)
	p := NewPatches()
	defer p.Reset()

	db := &gorm.DB{
		Error:     nil,
		Statement: &gorm.Statement{},
	}
	p = p.ApplyFuncReturn(NewPostgres, &Postgres{db: db}, nil)
	v, err := NewPostgres("any")
	r.NoError(err)
	r.NotNil(v)

	t.Run("FailedToCreateTaskStateLog", func(t *testing.T) {
		ndb := *db
		ndb.Error = errors.New(t.Name())
		p = testutil.GormDBCreate(p, nil, &ndb)

		err := v.Create(&types.TaskStateLog{})
		r.ErrorContains(err, t.Name())
	})
	p = testutil.GormDBCreate(p, nil, db)

	t.Run("Success", func(t *testing.T) {
		err := v.Create(&types.TaskStateLog{})
		r.NoError(err)
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
	p = testutil.GormDBOrder(p, v.db)

	t.Run("FailedToFindDB", func(t *testing.T) {
		p = p.ApplyMethodReturn(&gorm.DB{}, "Find", &gorm.DB{Error: errors.New(t.Name())})
		_, err := v.Fetch(1, 1)
		r.ErrorContains(err, t.Name())
	})

	p = testutil.GormDBFind(p, &([]*taskStateLog{{}, {}, {}}), v.db)

	t.Run("Success", func(t *testing.T) {
		_task, err := v.Fetch(1, 1)
		r.NotNil(_task)
		r.NoError(err)
	})
}

func TestPostgres_FetchNextTaskID(t *testing.T) {
	r := require.New(t)
	p := NewPatches()
	defer p.Reset()

	v := &Postgres{
		db: &gorm.DB{
			Error:     nil,
			Statement: &gorm.Statement{},
		},
	}

	p = testutil.GormDBModel(p, v.db)
	p = testutil.GormDBSelect(p, v.db)

	t.Run("FailedToTakeMaxTaskID", func(t *testing.T) {
		p = testutil.GormDBTake(p, nil, &gorm.DB{Error: errors.New(t.Name())})

		taskID, err := v.FetchNextTaskID()

		r.Equal(taskID, uint64(0))
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		_taskID := uint64(100)
		p = testutil.GormDBTake(p, &_taskID, v.db)

		taskID, err := v.FetchNextTaskID()

		r.Equal(taskID, _taskID+1)
		r.NoError(err)
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
