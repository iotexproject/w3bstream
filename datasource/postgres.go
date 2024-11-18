package datasource

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/iotexproject/w3bstream/service/apinode/persistence"
	"github.com/iotexproject/w3bstream/task"
)

type Postgres struct {
	db *gorm.DB
}

func (p *Postgres) Retrieve(taskIDs []common.Hash) ([]*task.Task, error) {
	if len(taskIDs) == 0 {
		return nil, errors.New("empty query task ids")
	}
	ts := []*persistence.Task{}
	if err := p.db.Where("task_id IN ?", taskIDs).First(&ts).Error; err != nil {
		return nil, errors.Wrap(err, "failed to query tasks")
	}

	res := []*task.Task{}
	for _, t := range ts {
		ps := [][]byte{}
		if err := json.Unmarshal(t.Payloads, &ps); err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal task payloads, task_id %v", t.TaskID)
		}
		res = append(res, &task.Task{
			ID:             t.TaskID,
			ProjectID:      t.ProjectID,
			ProjectVersion: t.ProjectVersion,
			Payloads:       ps,
			DeviceID:       t.DeviceID,
			Signature:      t.Signature,
		})
	}
	if len(res) != len(taskIDs) {
		return nil, errors.Errorf("cannot find all tasks, task_ids %v", taskIDs)
	}
	return res, nil
}

func NewPostgres(dsn string) (*Postgres, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to connect postgres")
	}
	return &Postgres{db}, nil
}
