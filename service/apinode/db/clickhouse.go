package db

import (
	"context"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/ethereum/go-ethereum/common"

	"github.com/pkg/errors"
)

type Task struct {
	TaskID             string    `ch:"task_id"             gorm:"primarykey"`
	DevicePubKey       string    `ch:"device_public_key"   gorm:"not null"`
	Nonce              uint64    `ch:"nonce"               gorm:"not null"`
	ProjectID          string    `ch:"project_id"          gorm:"not null"`
	ProjectVersion     string    `ch:"project_version"     gorm:"not null"`
	Payload            string    `ch:"payload"             gorm:"not null"`
	Signature          string    `ch:"signature"           gorm:"not null"`
	SignatureAlgorithm string    `ch:"signature_algorithm" gorm:"not null"`
	HashAlgorithm      string    `ch:"hash_algorithm"      gorm:"not null"`
	CreatedAt          time.Time `ch:"created_at"          gorm:"not null"`
	PrevTaskID         string    `ch:"previous_task_id"    gorm:"not null"`
}

func (p *DB) CreateTasks(ts []*Task) error {
	batch, err := p.ch.PrepareBatch(context.Background(), "INSERT INTO w3bstream_tasks")
	if err != nil {
		return errors.Wrap(err, "failed to prepare batch")
	}
	for _, t := range ts {
		if err := batch.AppendStruct(t); err != nil {
			return errors.Wrap(err, "failed to append struct")
		}
	}
	err = batch.Send()
	return errors.Wrap(err, "failed to create tasks")
}

func (p *DB) FetchTask(taskID common.Hash) (*Task, error) {
	t := Task{}
	if err := p.ch.QueryRow(context.Background(), "SELECT * FROM w3bstream_tasks WHERE task_id = ?", taskID.Hex()).ScanStruct(&t); err != nil {
		return nil, errors.Wrap(err, "failed to query task")
	}
	return &t, nil
}

func migrateCH(conn driver.Conn) error {
	err := conn.Exec(context.Background(), `
        CREATE TABLE IF NOT EXISTS w3bstream_tasks
        (
            task_id String NOT NULL,
            device_public_key String NOT NULL,
            nonce UInt64 NOT NULL,
            project_id String NOT NULL,
            project_version String NOT NULL,
            payload String NOT NULL,
            signature String NOT NULL,
            signature_algorithm String NOT NULL,
            hash_algorithm String NOT NULL,
            created_at DateTime NOT NULL,
			previous_task_id String NOT NULL
        )
        ENGINE = ReplacingMergeTree()
        PRIMARY KEY task_id
        ORDER BY task_id`,
	)
	return errors.Wrap(err, "failed to create clickhouse table")
}

func newCH(dsn string) (driver.Conn, error) {
	op, err := clickhouse.ParseDSN(dsn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse clickhouse dsn")
	}
	conn, err := clickhouse.Open(op)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect clickhouse")
	}
	if err := migrateCH(conn); err != nil {
		return nil, err
	}
	return conn, nil
}
