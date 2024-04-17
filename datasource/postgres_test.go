package datasource

import (
	"encoding/json"
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/machinefi/sprout/testutil"
)

func TestNewPostgres(t *testing.T) {
	r := require.New(t)

	t.Run("FailedToOpenDSN", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		testutil.GormOpen(p, nil, errors.New(t.Name()))
		datasource, err := NewPostgres("any")
		r.Nil(datasource)
		r.ErrorContains(err, t.Name())
	})
	t.Run("Success", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		d := &gorm.DB{}

		testutil.GormOpen(p, d, nil)
		datasource, err := NewPostgres("any")
		r.NotNil(datasource)
		r.NoError(err)
	})
}

func TestPostgres_Retrieve(t *testing.T) {
	r := require.New(t)

	d := &postgres{
		db: &gorm.DB{
			Error:     nil,
			Statement: &gorm.Statement{},
		},
	}

	t.Run("FailedToQueryTask", func(t *testing.T) {
		t.Run("NotExist", func(t *testing.T) {
			p := NewPatches()
			defer p.Reset()

			testutil.GormDBWhere(p, d.db)
			p.ApplyMethodReturn(&gorm.DB{}, "First", &gorm.DB{Error: gorm.ErrRecordNotFound})

			task, err := d.Retrieve(uint64(1), uint64(1))
			r.NoError(err)
			r.Empty(task)
		})

		t.Run("FailedToQuery", func(t *testing.T) {
			p := NewPatches()
			defer p.Reset()

			testutil.GormDBWhere(p, d.db)
			p.ApplyMethodReturn(&gorm.DB{}, "First", &gorm.DB{Error: errors.New(t.Name())})

			_, err := d.Retrieve(uint64(1), uint64(1))
			r.ErrorContains(err, t.Name())
		})
	})

	t.Run("FailedToUnmarshal", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		testutil.GormDBWhere(p, d.db)
		testutil.GormDBFirst(p, &task{}, d.db)
		p.ApplyFuncReturn(json.Unmarshal, errors.New(t.Name()))

		_, err := d.Retrieve(uint64(1), uint64(1))
		r.ErrorContains(err, t.Name())
	})

	t.Run("FailedToQueryMessages", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		testutil.GormDBWhere(p, d.db)
		testutil.GormDBFirst(p, &task{}, d.db)
		p.ApplyFuncReturn(json.Unmarshal, nil)
		p.ApplyMethodReturn(&gorm.DB{}, "Find", &gorm.DB{Error: errors.New(t.Name())})

		_, err := d.Retrieve(uint64(1), uint64(1))
		r.ErrorContains(err, t.Name())
	})

	t.Run("InvalidTask", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		testutil.GormDBWhere(p, d.db)
		testutil.GormDBFirst(p, &task{}, d.db)
		p.ApplyFuncReturn(json.Unmarshal, nil)
		testutil.GormDBFind(p, &([]*message{}), d.db)

		_, err := d.Retrieve(uint64(1), uint64(1))
		r.ErrorContains(err, "invalid task,")
	})

	t.Run("Success", func(t *testing.T) {
		p := NewPatches()
		defer p.Reset()

		testutil.GormDBWhere(p, d.db)
		testutil.GormDBFirst(p, &task{}, d.db)
		p.ApplyFuncReturn(json.Unmarshal, nil)
		testutil.GormDBFind(p, &([]*message{{}}), d.db)

		task, err := d.Retrieve(uint64(1), uint64(1))
		r.NoError(err)
		r.NotNil(task)
	})
}
