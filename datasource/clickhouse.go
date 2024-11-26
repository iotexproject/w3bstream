package datasource

import (
	"context"
	"encoding/json"
	"math/big"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/service/apinode/db"
	"github.com/iotexproject/w3bstream/task"
)

type Clickhouse struct {
	db driver.Conn
}

func (p *Clickhouse) Retrieve(taskIDs []common.Hash) ([]*task.Task, error) {
	if len(taskIDs) == 0 {
		return nil, errors.New("empty query task ids")
	}
	tids := make([][]byte, 0, len(taskIDs))
	for _, t := range taskIDs {
		tids = append(tids, t.Bytes())
	}
	var ts []db.Task
	if err := p.db.Select(context.Background(), &ts, "SELECT * FROM w3bstream_tasks WHERE task_id IN ?", tids); err != nil {
		return nil, errors.Wrap(err, "failed to query tasks")
	}

	res := []*task.Task{}
	for i := range ts {
		ps := [][]byte{}
		if err := json.Unmarshal(ts[i].Payloads, &ps); err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal task payloads, task_id %v", ts[i].TaskID)
		}
		pid := new(big.Int)
		if _, ok := pid.SetString(ts[i].ProjectID, 10); !ok {
			return nil, errors.New("failed to decode project id string")
		}
		res = append(res, &task.Task{
			ID:             common.BytesToHash(ts[i].TaskID),
			ProjectID:      pid,
			ProjectVersion: ts[i].ProjectVersion,
			Payloads:       ps,
			DeviceID:       common.BytesToAddress(ts[i].DeviceID),
			Signature:      ts[i].Signature,
		})
	}
	if len(res) != len(taskIDs) {
		return nil, errors.Errorf("cannot find all tasks, task_ids %v", taskIDs)
	}
	return res, nil
}

func NewClickhouse(dsn string) (*Clickhouse, error) {
	op, err := clickhouse.ParseDSN(dsn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse clickhouse dsn")
	}
	conn, err := clickhouse.Open(op)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect clickhouse")
	}
	return &Clickhouse{db: conn}, nil
}
