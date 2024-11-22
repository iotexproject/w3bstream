package db

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"

	"gorm.io/gorm"
)

type DB struct {
	sqlite *gorm.DB
	ch     driver.Conn
}

func New(localDBDir, chEndpoint, chPasswd string, isChTLS bool) (*DB, error) {
	sqlite, err := newSqlite(localDBDir)
	if err != nil {
		return nil, err
	}
	ch, err := newCH(chEndpoint, chPasswd, isChTLS)
	if err != nil {
		return nil, err
	}
	return &DB{sqlite: sqlite, ch: ch}, nil
}
