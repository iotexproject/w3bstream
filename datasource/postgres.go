package datasource

import (
	"encoding/json"
	"sync"

	"github.com/pkg/errors"
	"gorm.io/datatypes"
	pgdriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	tasktype "github.com/machinefi/sprout/task"
)

type message struct {
	gorm.Model
	MessageID      string `gorm:"index:message_id,not null"`
	ClientID       string `gorm:"index:message_fetch,not null,default:''"`
	ProjectID      uint64 `gorm:"index:message_fetch,not null"`
	ProjectVersion string `gorm:"index:message_fetch,not null,default:'0.0'"`
	Data           []byte `gorm:"size:4096"`
	InternalTaskID string `gorm:"index:internal_task_id,not null,default:''"`
}

type task struct {
	gorm.Model
	ProjectID      uint64         `gorm:"index:task_fetch,not null"`
	InternalTaskID string         `gorm:"index:internal_task_id,not null"`
	MessageIDs     datatypes.JSON `gorm:"not null"`
	Signature      string         `gorm:"not null,default:''"`
}

type postgres struct {
	db *gorm.DB
}

func (p *postgres) Retrieve(projectID, nextTaskID uint64) (*tasktype.Task, error) {
	t := task{}
	if err := p.db.Where("id >= ? AND project_id = ?", nextTaskID, projectID).First(&t).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "failed to query task, next_task_id %v", nextTaskID)
	}

	messageIDs := []string{}
	if err := json.Unmarshal(t.MessageIDs, &messageIDs); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal task message ids, task_id %v", t.ID)
	}

	ms := []*message{}
	if err := p.db.Where("message_id IN ?", messageIDs).Find(&ms).Error; err != nil {
		return nil, errors.Wrapf(err, "failed to query task messages, task_id %v", t.ID)
	}
	if len(ms) == 0 {
		return nil, errors.Errorf("invalid task, task_id %v", t.ID)
	}

	ds := [][]byte{}
	for _, m := range ms {
		ds = append(ds, m.Data)
	}

	return &tasktype.Task{
		ID:             uint64(t.ID),
		ProjectID:      ms[0].ProjectID,
		ProjectVersion: ms[0].ProjectVersion,
		Data:           ds,
		ClientID:       ms[0].ClientID,
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
