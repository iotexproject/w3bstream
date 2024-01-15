package postgres

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/machinefi/sprout/types"
)

type message struct {
	gorm.Model
	MessageID      string `gorm:"index:message_id,not null"`
	ProjectID      uint64 `gorm:"index:message_fetch,not null"`
	ProjectVersion string `gorm:"not null,default:'0.0'"`
	Data           string `gorm:"size:4096"`
}

type task struct {
	gorm.Model
	TaskID    string          `gorm:"index:task_id,not null"`
	MessageID string          `gorm:"index:message_id,not null"`
	ProjectID uint64          `gorm:"index:task_fetch,not null"`
	State     types.TaskState `gorm:"index:task_fetch,not null"`
}

type taskStateLog struct {
	gorm.Model
	TaskID  string          `gorm:"index:task_id,not null"`
	State   types.TaskState `gorm:"not null"`
	Comment string
}

type Postgres struct {
	db *gorm.DB
}

func (p *Postgres) Save(msg *types.Message) error {
	m := message{
		MessageID:      msg.ID,
		ProjectID:      msg.ProjectID,
		ProjectVersion: msg.ProjectVersion,
		Data:           msg.Data,
	}
	tid := uuid.NewString()
	t := task{
		TaskID:    tid,
		MessageID: msg.ID,
		ProjectID: msg.ProjectID,
		State:     types.TaskStateReceived,
	}
	l := taskStateLog{
		TaskID: tid,
		State:  types.TaskStateReceived,
	}

	return p.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&m).Error; err != nil {
			return errors.Wrap(err, "create message failed")
		}
		if err := tx.Create(&t).Error; err != nil {
			return errors.Wrap(err, "create task failed")
		}
		if err := tx.Create(&l).Error; err != nil {
			return errors.Wrap(err, "create task state log failed")
		}
		return nil
	})
}

func (p *Postgres) Fetch(projectID uint64) (*types.Task, error) {
	t := task{}
	if err := p.db.Where("project_id = ? AND state = ?", projectID, types.TaskStateReceived).First(&t).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "query task failed, projectID %d", projectID)
	}

	m := message{}
	if err := p.db.Where("message_id = ?", t.MessageID).Take(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Errorf("missing message, messageID %s", t.MessageID)
		}
		return nil, errors.Wrapf(err, "query message failed, messageID %s", t.MessageID)
	}
	return &types.Task{
		ID: t.TaskID,
		Messages: []*types.Message{{
			ID:             m.MessageID,
			ProjectID:      m.ProjectID,
			ProjectVersion: m.ProjectVersion,
			Data:           m.Data,
		}},
	}, nil
}

func (p *Postgres) FetchStateLog(messageID string) ([]*types.TaskStateLog, error) {
	ts := []*task{}
	if err := p.db.Where("message_id = ?", messageID).Find(&ts).Error; err != nil {
		return nil, errors.Wrapf(err, "query task by message id failed, messageID %s", messageID)
	}
	tids := []string{}
	for _, t := range ts {
		tids = append(tids, t.TaskID)
	}
	if len(tids) == 0 {
		return nil, nil
	}

	ls := []*taskStateLog{}
	if err := p.db.Order("created_at").Where("task_id IN ?", tids).Find(&ls).Error; err != nil {
		return nil, errors.Wrapf(err, "query task state log failed, taskIDs %v", tids)
	}

	tls := []*types.TaskStateLog{}
	for _, l := range ls {
		tls = append(tls, &types.TaskStateLog{
			TaskID:    l.TaskID,
			State:     l.State,
			Comment:   l.Comment,
			CreatedAt: l.CreatedAt,
		})
	}
	return tls, nil
}

func (p *Postgres) UpdateState(taskID string, state types.TaskState, comment string, createdAt time.Time) error {
	return p.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&task{}).Where("task_id = ?", taskID).Update("state", state).Error; err != nil {
			return errors.Wrapf(err, "update task state failed, task_id %s, target_state %s", taskID, state.String())
		}
		if err := tx.Create(&taskStateLog{
			TaskID:  taskID,
			State:   state,
			Comment: comment,
			Model: gorm.Model{
				CreatedAt: createdAt,
			},
		}).Error; err != nil {
			return errors.Wrap(err, "create task state log failed")
		}
		return nil
	})
}

func NewPostgres(pgEndpoint string) (*Postgres, error) {
	db, err := gorm.Open(postgres.Open(pgEndpoint), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, errors.Wrap(err, "connect postgres failed")
	}
	if err := db.AutoMigrate(&message{}, &task{}, &taskStateLog{}); err != nil {
		return nil, errors.Wrap(err, "migrate model failed")
	}
	return &Postgres{db}, nil
}
