package datasource

import (
	"encoding/json"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	pgdriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/iotexproject/w3bstream/cmd/apinode/persistence"
	"github.com/iotexproject/w3bstream/task"
)

type postgres struct {
	db *gorm.DB
}

func (p *postgres) Retrieve(projectID uint64, taskID common.Hash) (*task.Task, error) {
	t := persistence.Task{}
	if err := p.db.Where("task_id = ? AND project_id = ?", taskID, projectID).First(&t).Error; err != nil {
		return nil, errors.Wrap(err, "failed to query task")
	}
	messageIDs := []string{}
	if err := json.Unmarshal(t.MessageIDs, &messageIDs); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal task message ids, task_id %v", t.TaskID)
	}

	ms := []*persistence.Message{}
	if err := p.db.Where("id IN ?", messageIDs).Find(&ms).Error; err != nil {
		return nil, errors.Wrapf(err, "failed to query task messages, task_id %v", t.TaskID)
	}
	if len(ms) == 0 {
		return nil, errors.Errorf("invalid task, task_id %v", t.TaskID)
	}

	ds := [][]byte{}
	for _, m := range ms {
		ds = append(ds, m.Data)
	}

	return &task.Task{
		ID:             t.TaskID,
		ProjectID:      ms[0].ProjectID,
		ProjectVersion: ms[0].ProjectVersion,
		Payloads:       ds,
		DeviceID:       ms[0].DeviceID,
		Signature:      t.Signature,
	}, nil
}

type Postgres struct {
	mux sync.Mutex
	ps  map[string]*postgres
}

func (p *Postgres) New(dsn string) (Datasource, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	d, ok := p.ps[dsn]
	if ok {
		return d, nil
	}

	db, err := gorm.Open(pgdriver.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to connect postgres, dsn %s", dsn)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get sql db")
	}
	sqlDB.SetMaxOpenConns(500)

	d = &postgres{db}
	p.ps[dsn] = d
	return d, nil
}

func NewPostgres() *Postgres {
	return &Postgres{
		ps: map[string]*postgres{},
	}
}
