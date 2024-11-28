package datasource

import (
	"context"
	"math/big"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
	tids := make([]string, 0, len(taskIDs))
	for _, t := range taskIDs {
		tids = append(tids, t.Hex())
	}
	var ts []db.Task
	if err := p.db.Select(context.Background(), &ts, "SELECT * FROM w3bstream_tasks WHERE task_id IN ?", tids); err != nil {
		return nil, errors.Wrap(err, "failed to query tasks")
	}

	res := []*task.Task{}
	for i := range ts {
		pid, ok := new(big.Int).SetString(ts[i].ProjectID, 10)
		if !ok {
			return nil, errors.New("failed to decode project id string")
		}
		payload, err := hexutil.Decode(ts[i].Payload)
		if err != nil {
			return nil, errors.Wrap(err, "failed to decode payload from hex format")
		}
		sig, err := hexutil.Decode(ts[i].Signature)
		if err != nil {
			return nil, errors.Wrap(err, "failed to decode signature from hex format")
		}
		res = append(res, &task.Task{
			ID:             common.HexToHash(ts[i].TaskID),
			ProjectID:      pid,
			ProjectVersion: ts[i].ProjectVersion,
			Payload:        payload,
			DeviceID:       common.HexToAddress(ts[i].DeviceID),
			Signature:      sig,
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
