package db

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"

	"gorm.io/gorm"
)

// TODO: remove this combinator
type DB struct {
	sqlite *gorm.DB
	ch     driver.Conn
}

func New(localDBDir, dsn string) (*DB, error) {
	sqlite, err := newSqlite(localDBDir)
	if err != nil {
		return nil, err
	}
	ch, err := newCH(dsn)
	if err != nil {
		return nil, err
	}
	return &DB{sqlite: sqlite, ch: ch}, nil
}
