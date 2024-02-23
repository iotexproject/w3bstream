package testutil

import (
	"database/sql"
	. "github.com/agiledragon/gomonkey/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"reflect"
)

var _targetGormDatabase = reflect.TypeOf(&gorm.DB{})

func GormDBTransaction(p *Patches, err error) *Patches {
	return p.ApplyMethod(
		_targetGormDatabase,
		"Transaction",
		func(*gorm.DB, func(*gorm.DB) error, ...*sql.TxOptions) error {
			return err
		},
	)
}

func GormDBBegin(p *Patches, ret *gorm.DB) *Patches {
	return p.ApplyMethod(
		_targetGormDatabase,
		"Begin",
		func(*gorm.DB, ...*sql.TxOptions) *gorm.DB {
			return ret
		},
	)
}
func GormDBCommit(p *Patches, ret *gorm.DB) *Patches {
	return p.ApplyMethod(
		_targetGormDatabase,
		"Commit",
		func(*gorm.DB) *gorm.DB {
			return ret
		},
	)
}

func GormDBClauses(p *Patches, ret *gorm.DB) *Patches {
	return p.ApplyMethod(
		_targetGormDatabase,
		"Clauses",
		func(_ *gorm.DB, _ ...clause.Expression) *gorm.DB {
			return ret
		},
	)
}

func GormDBOrder(p *Patches, ret *gorm.DB) *Patches {
	return p.ApplyMethod(
		_targetGormDatabase,
		"Order",
		func(_ *gorm.DB, _ interface{}) *gorm.DB {
			return ret
		},
	)
}

func GormDBWhere(p *Patches, ret *gorm.DB) *Patches {
	return p.ApplyMethod(
		_targetGormDatabase,
		"Where",
		func(_ *gorm.DB, _ interface{}, _ ...interface{}) *gorm.DB {
			return ret
		},
	)
}

func GormDBLimit(p *Patches, ret *gorm.DB) *Patches {
	return p.ApplyMethod(
		_targetGormDatabase,
		"Limit",
		func(_ *gorm.DB, _ int) *gorm.DB {
			return ret
		},
	)
}

func GormDBRollback(p *Patches, ret *gorm.DB) *Patches {
	return p.ApplyMethod(
		_targetGormDatabase,
		"Rollback",
		func(_ *gorm.DB) *gorm.DB {
			return ret
		},
	)
}

func GormDBModel(p *Patches, ret *gorm.DB) *Patches {
	return p.ApplyMethod(
		_targetGormDatabase,
		"Model",
		func(_ *gorm.DB, _ any) *gorm.DB {
			return ret
		},
	)
}

func GormDBUpdate(p *Patches, ret *gorm.DB) *Patches {
	return p.ApplyMethod(
		_targetGormDatabase,
		"Update",
		func(_ *gorm.DB, _string, _ any) *gorm.DB {
			return ret
		},
	)
}

func GormDBFind(p *Patches, inputmut any, ret *gorm.DB) *Patches {
	return p.ApplyMethod(
		_targetGormDatabase,
		"Find",
		func(_ *gorm.DB, v any) *gorm.DB {
			vi := reflect.ValueOf(inputmut)
			vo := reflect.ValueOf(v)
			if vi.IsValid() && vo.IsValid() {
				vo.Elem().Set(vi.Elem())
			}
			return ret
		},
	)
}

func GormDBFirst(p *Patches, inputmut any, ret *gorm.DB) *Patches {
	return p.ApplyMethod(
		_targetGormDatabase,
		"First",
		func(_ *gorm.DB, dst any, _ ...any) *gorm.DB {
			dst = inputmut
			return ret
		},
	)
}

func GormDBCreate(p *Patches, inputmut any, ret *gorm.DB) *Patches {
	return p.ApplyMethod(
		_targetGormDatabase,
		"Create",
		func(_ *gorm.DB, dst any) *gorm.DB {
			dst = inputmut
			return ret
		},
	)
}

func GormDBAutoMigrate(p *Patches, err error) *Patches {
	return p.ApplyMethod(
		_targetGormDatabase,
		"AutoMigrate",
		func(_ *gorm.DB, _ ...any) error {
			return err
		},
	)
}

func GormOpen(p *Patches, db *gorm.DB, err error) *Patches {
	return p.ApplyFunc(
		gorm.Open,
		func(_ gorm.Dialector, _ ...gorm.Option) (*gorm.DB, error) {
			return db, err
		},
	)
}
