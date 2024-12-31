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
	taskIDsHex := make([]string, 0, len(taskIDs))
	for _, t := range taskIDs {
		taskIDsHex = append(taskIDsHex, t.Hex())
	}
	var ts []db.Task
	if err := p.db.Select(context.Background(), &ts, "SELECT * FROM w3bstream_tasks WHERE task_id IN ?", taskIDsHex); err != nil {
		return nil, errors.Wrap(err, "failed to query tasks")
	}

	// filter out prev task that has been fetched
	prevTasksPool := map[string]db.Task{}
	for i := range ts {
		prevTasksPool[ts[i].TaskID] = ts[i]
	}
	fetchPrevTaskIDs := make([]string, 0)
	for i := range ts {
		if ts[i].PrevTaskID != "" {
			if _, exist := prevTasksPool[ts[i].PrevTaskID]; !exist {
				fetchPrevTaskIDs = append(fetchPrevTaskIDs, ts[i].PrevTaskID)
			}
		}
	}

	if len(fetchPrevTaskIDs) != 0 {
		var pdts []db.Task
		if err := p.db.Select(context.Background(), &pdts, "SELECT * FROM w3bstream_tasks WHERE task_id IN ?", fetchPrevTaskIDs); err != nil {
			return nil, errors.Wrap(err, "failed to query previous tasks")
		}
		for i := range pdts {
			prevTasksPool[pdts[i].TaskID] = pdts[i]
		}
	}

	res := []*task.Task{}
	for i := range ts {
		t, err := p.conv(&ts[i])
		if err != nil {
			return nil, err
		}
		if ts[i].PrevTaskID != "" {
			pdt, ok := prevTasksPool[ts[i].PrevTaskID]
			if !ok {
				return nil, errors.New("failed to get previous task")
			}
			pt, err := p.conv(&pdt)
			if err != nil {
				return nil, err
			}
			t.PrevTask = pt
		}
		res = append(res, t)
	}
	if len(res) != len(taskIDs) {
		return nil, errors.Errorf("cannot find all tasks, task_ids %v", taskIDs)
	}

	return res, nil
}

func (p *Clickhouse) conv(dt *db.Task) (*task.Task, error) {
	pid, ok := new(big.Int).SetString(dt.ProjectID, 10)
	if !ok {
		return nil, errors.New("failed to decode project id string")
	}
	sig, err := hexutil.Decode(dt.Signature)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode signature from hex format")
	}
	pubkey, err := hexutil.Decode(dt.DevicePubKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode device public key from hex format")
	}
	return &task.Task{
		ID:             common.HexToHash(dt.TaskID),
		Nonce:          dt.Nonce,
		ProjectID:      pid,
		ProjectVersion: dt.ProjectVersion,
		DevicePubKey:   pubkey,
		Payload:        []byte(dt.Payload),
		Signature:      sig,
	}, nil
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
