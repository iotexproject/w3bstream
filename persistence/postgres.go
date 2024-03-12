package persistence

import (
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/machinefi/sprout/types"
)

type taskStateLog struct {
	gorm.Model
	TaskID         uint64          `gorm:"index:state_fetch,not null"`
	ProjectID      uint64          `gorm:"index:state_fetch,not null"`
	ProjectVersion string          `gorm:"index:state_fetch,not null"`
	State          types.TaskState `gorm:"not null"`
	Comment        string
	Result         []byte
}

type Postgres struct {
	db *gorm.DB
}

func (p *Postgres) Create(tl *types.TaskStateLog) error {
	l := &taskStateLog{
		TaskID:         tl.Task.ID,
		ProjectID:      tl.Task.ProjectID,
		ProjectVersion: tl.Task.ProjectVersion,
		State:          tl.State,
		Comment:        tl.Comment,
		Result:         tl.Result,
		Model: gorm.Model{
			CreatedAt: tl.CreatedAt,
		},
	}
	if err := p.db.Create(l).Error; err != nil {
		return errors.Wrap(err, "failed to create task state log")
	}
	return nil
}

func (p *Postgres) Fetch(taskID, projectID uint64) ([]*types.TaskStateLog, error) {
	ls := []*taskStateLog{}
	if err := p.db.Order("created_at").Where("task_id = ? AND project_id = ?", taskID, projectID).Find(&ls).Error; err != nil {
		return nil, errors.Wrapf(err, "failed to query task state log, task_id %v, project_id %v", taskID, projectID)
	}
	tls := []*types.TaskStateLog{}
	for _, l := range ls {
		tls = append(tls, &types.TaskStateLog{
			Task:      types.Task{ID: taskID},
			State:     l.State,
			Comment:   l.Comment,
			Result:    l.Result,
			CreatedAt: l.CreatedAt,
		})
	}
	return tls, nil
}

func NewPostgres(pgEndpoint string) (*Postgres, error) {
	db, err := gorm.Open(postgres.Open(pgEndpoint), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect postgres")
	}
	if err := db.AutoMigrate(&taskStateLog{}); err != nil {
		return nil, errors.Wrap(err, "failed to migrate model")
	}
	return &Postgres{db}, nil
}
