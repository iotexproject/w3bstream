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
	TaskID         []byte    `ch:"task_id"`
	DeviceID       []byte    `ch:"device_id"`
	Nonce          uint64    `ch:"nonce"`
	ProjectID      string    `ch:"project_id"`
	ProjectVersion string    `ch:"project_version"`
	Payload        []byte    `ch:"payload"`
	Signature      []byte    `ch:"signature"`
	Algorithm      string    `ch:"algorithm"`
	CreatedAt      time.Time `ch:"create_at"`
}

func (p *DB) CreateTask(m *Task) error {
	batch, err := p.ch.PrepareBatch(context.Background(), "INSERT INTO w3bstream_tasks")
	if err != nil {
		return errors.Wrap(err, "failed to prepare batch")
	}
	if err := batch.AppendStruct(m); err != nil {
		return errors.Wrap(err, "failed to append struct")
	}
	err = batch.Send()
	return errors.Wrap(err, "failed to create task")
}

func (p *DB) FetchTask(taskID common.Hash) (*Task, error) {
	t := Task{}
	if err := p.ch.QueryRow(context.Background(), "SELECT * FROM w3bstream_tasks WHERE task_id = ?", taskID.Bytes()).ScanStruct(&t); err != nil {
		return nil, errors.Wrap(err, "failed to query task")
	}
	return &t, nil
}

func migrateCH(conn driver.Conn) error {
	err := conn.Exec(context.Background(), `
        CREATE TABLE IF NOT EXISTS w3bstream_tasks
        (
            task_id Array(UInt8) NOT NULL,
            device_id Array(UInt8) NOT NULL,
            nonce UInt64 NOT NULL,
            project_id String NOT NULL,
            project_version String NOT NULL,
            payload Array(UInt8) NOT NULL,
            signature Array(UInt8) NOT NULL,
            algorithm String NOT NULL,
            create_at DateTime NOT NULL
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
