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

func (p *Postgres) Retrieve(projectID uint64, taskID common.Hash) (*task.Task, error) {
	t := persistence.Task{}
	if err := p.db.Where("task_id = ? AND project_id = ?", taskID, projectID).First(&t).Error; err != nil {
		return nil, errors.Wrap(err, "failed to query task")
	}
	ps := [][]byte{}
	if err := json.Unmarshal(t.Payloads, &ps); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal task payloads, task_id %v", t.TaskID)
	}

	return &task.Task{
		ID:             t.TaskID,
		ProjectID:      t.ProjectID,
		ProjectVersion: t.ProjectVersion,
		Payloads:       ps,
		DeviceID:       t.DeviceID,
		Signature:      t.Signature,
	}, nil
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
