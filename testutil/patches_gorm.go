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

func GormDBFind(p *Patches, inputmut any, ret *gorm.DB) *Patches {
	return p.ApplyMethod(
		_targetGormDatabase,
		"Find",
		func(_ *gorm.DB, v any) *gorm.DB {
			v = inputmut
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
