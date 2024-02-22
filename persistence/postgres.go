package persistence

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"

	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/types"
)

type message struct {
	gorm.Model
	MessageID      string `gorm:"index:message_id,not null"`
	ClientDID      string `gorm:"column:client_did;index:message_fetch,not null,default:''"`
	ProjectID      uint64 `gorm:"index:message_fetch,not null"`
	ProjectVersion string `gorm:"index:message_fetch,not null,default:'0.0'"`
	Data           string `gorm:"size:4096"`
	TaskID         string `gorm:"index:task_id,not null,default:''"`
}

type task struct {
	gorm.Model
	TaskID    string          `gorm:"index:task_id,not null"`
	MessageID string          `gorm:"index:message_id,not null"`
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

func (p *Postgres) CreateMessage(msg *types.Message) (*message, error) {
	m := &message{
		MessageID:      msg.ID,
		ClientDID:      msg.ClientDID,
		ProjectID:      msg.ProjectID,
		ProjectVersion: msg.ProjectVersion,
		Data:           msg.Data,
	}
	if err := p.db.Create(m).Error; err != nil {
		return nil, errors.Wrap(err, "failed to create message")
	}
	return m, nil
}

func txAggregateTask(tx *gorm.DB, amount int, m *types.Message) (string, error) {
	messages := make([]*message, 0)

	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Order("created_at").
		Where(
			"project_id = ? AND project_version = ? AND client_did = ? AND task_id = ?",
			m.ProjectID, m.ProjectVersion, m.ClientDID, "",
		).Limit(amount).Find(&messages).Error; err != nil {
		return "", errors.Wrap(err, "failed to fetch unpacked messages")
	}

	if len(messages) == 0 {
		return "", nil
	}

	// not enough message for pack task
	if amount > 1 && len(messages) < amount {
		return "", nil
	}

	// generate task id, batch update messages
	taskID := uuid.NewString()
	tasks := make([]*task, 0, amount)
	messageIDs := make([]string, 0, amount)
	for _, v := range messages {
		messageIDs = append(messageIDs, v.MessageID)
		tasks = append(tasks, &task{
			TaskID:    taskID,
			MessageID: v.MessageID,
			State:     types.TaskStatePacked,
		})
	}

	if err := tx.Model(&message{}).Where("message_id IN ?", messageIDs).Update("task_id", taskID).Error; err != nil {
		return "", errors.Wrap(err, "failed to update message task ID")
	}

	if err := tx.Create(&tasks).Error; err != nil {
		return "", errors.Wrap(err, "failed to create tasks")
	}
	return taskID, nil
}

func txCreateTaskLog(tx *gorm.DB, taskID string) error {
	if err := tx.Create(&taskStateLog{
		TaskID: taskID,
		State:  types.TaskStatePacked,
	}).Error; err != nil {
		return errors.Wrap(err, "create task state log failed")
	}
	return nil
}

func (p *Postgres) Save(msg *types.Message, config *project.Config) error {
	_, err := p.CreateMessage(msg)
	if err != nil {
		return err
	}

	return p.db.Transaction(func(tx *gorm.DB) error {
		taskID, err := txAggregateTask(tx, int(config.Aggregation.Amount), msg)
		if err != nil {
			return err
		}

		if taskID != "" {
			return txCreateTaskLog(tx, taskID)
		}
		return nil
	})
}

func (p *Postgres) Fetch() (*types.Task, error) {
	t := task{}
	if err := p.db.Where("state = ?", types.TaskStatePacked).First(&t).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "query task failed")
	}
	return p.FetchByID(t.TaskID)
}

func (p *Postgres) FetchTasksByTaskID(taskID string) ([]*task, error) {
	tasks := []*task{}
	if err := p.db.Where("task_id = ?", taskID).Find(&tasks).Error; err != nil {
		return nil, errors.Wrap(err, "query task failed")
	}
	return tasks, nil
}

func (p *Postgres) FetchMessagesByMessageIDs(messageIDs ...string) ([]*message, error) {
	messages := make([]*message, 0)
	if err := p.db.Where("message_id IN ?", messageIDs).Find(&messages).Error; err != nil {
		return nil, errors.Wrapf(err, "query message failed")
	}
	return messages, nil
}

func (p *Postgres) FetchByID(taskID string) (*types.Task, error) {
	tasks, err := p.FetchTasksByTaskID(taskID)
	if err != nil {
		return nil, err
	}
	if len(tasks) == 0 {
		return nil, nil
	}

	mids := make([]string, 0, len(tasks))
	for _, t := range tasks {
		mids = append(mids, t.MessageID)
	}

	messages, err := p.FetchMessagesByMessageIDs(mids...)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to fetch messages rel with task: %s", taskID)
	}

	if len(messages) == 0 {
		return nil, errors.Errorf("missing message, taskID %s", taskID)
	}
	tms := []*types.Message{}
	for _, m := range messages {
		tms = append(tms, &types.Message{
			ID:             m.MessageID,
			ClientDID:      m.ClientDID,
			ProjectID:      m.ProjectID,
			ProjectVersion: m.ProjectVersion,
			Data:           m.Data,
		})
	}
	return &types.Task{
		ID:       taskID,
		Messages: tms,
	}, nil
}

func (p *Postgres) FetchMessage(messageID string) ([]*types.MessageWithTime, error) {
	ms := []*message{}
	if err := p.db.Where("message_id = ?", messageID).Find(&ms).Error; err != nil {
		return nil, errors.Wrapf(err, "query message by messageID failed, messageID %s", messageID)
	}

	tms := []*types.MessageWithTime{}
	for _, m := range ms {
		tms = append(tms, &types.MessageWithTime{
			Message: types.Message{
				ID:             m.MessageID,
				ClientDID:      m.ClientDID,
				ProjectID:      m.ProjectID,
				ProjectVersion: m.ProjectVersion,
				Data:           m.Data,
			},
			CreatedAt: m.CreatedAt,
		})
	}
	return tms, nil
}

func (p *Postgres) FetchStateLog(messageID string) ([]*types.TaskStateLog, error) {
	ts := []*task{}
	if err := p.db.Where("message_id = ?", messageID).Find(&ts).Error; err != nil {
		return nil, errors.Wrapf(err, "query task by messageID failed, messageID %s", messageID)
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
