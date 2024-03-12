package datasource

import (
	"github.com/pkg/errors"
	pgdriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/machinefi/sprout/types"
)

type message struct {
	gorm.Model
	MessageID      string `gorm:"index:message_id,not null"`
	ClientDID      string `gorm:"column:client_did;index:message_fetch,not null,default:''"`
	ProjectID      uint64 `gorm:"index:message_fetch,not null"`
	ProjectVersion string `gorm:"index:message_fetch,not null,default:'0.0'"`
	Data           []byte `gorm:"size:4096"`
	TaskID         string `gorm:"index:task_id,not null,default:''"`
}

type task struct {
	gorm.Model
	TaskID     string   `gorm:"index:task_id,not null"`
	MessageIDs []string `gorm:"index:message_id,not null"`
}

type postgres struct {
	db *gorm.DB
}

func (p *postgres) Retrieve(nextTaskID uint64) (*types.Task, error) {
	t := task{}
	if err := p.db.Where("id >= ?", nextTaskID).First(&t).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "failed to query task, next_task_id %v", nextTaskID)
	}

	ms := []*message{}
	if err := p.db.Where("message_id IN ?", t.MessageIDs).Find(&ms).Error; err != nil {
		return nil, errors.Wrapf(err, "failed to query task messages, task_id %v", t.ID)
	}
	if len(ms) == 0 {
		return nil, errors.Errorf("invalid task, task_id %v", t.ID)
	}

	ds := [][]byte{}
	for _, m := range ms {
		ds = append(ds, m.Data)
	}

	return &types.Task{
		ID:             uint64(t.ID),
		ProjectID:      ms[0].ProjectID,
		ProjectVersion: ms[0].ProjectVersion,
		Data:           ds,
	}, nil
}

func NewPostgres(dsn string) (Datasource, error) {
	db, err := gorm.Open(pgdriver.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect postgres")
	}
	return &postgres{db}, nil
}
