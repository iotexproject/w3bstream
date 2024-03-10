package main

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"

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
	TaskID    string `gorm:"index:task_id,not null"`
	MessageID string `gorm:"index:message_id,not null"`
}

type persistence struct {
	db *gorm.DB
}

func (p *persistence) createMessageTx(tx *gorm.DB, msg *types.Message) error {
	m := &message{
		MessageID:      msg.ID,
		ClientDID:      msg.ClientDID,
		ProjectID:      msg.ProjectID,
		ProjectVersion: msg.ProjectVersion,
		Data:           msg.Data,
	}
	if err := tx.Create(m).Error; err != nil {
		return errors.Wrap(err, "failed to create message")
	}
	return nil
}

func (p *persistence) aggregateTaskTx(tx *gorm.DB, amount int, m *types.Message) error {
	messages := make([]*message, 0)
	if amount == 0 {
		amount = 1
	}

	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Order("created_at").
		Where(
			"project_id = ? AND project_version = ? AND client_did = ? AND task_id = ?",
			m.ProjectID, m.ProjectVersion, m.ClientDID, "",
		).Limit(amount).Find(&messages).Error; err != nil {
		return errors.Wrap(err, "failed to fetch unpacked messages")
	}

	// no enough message for pack task
	if len(messages) < amount {
		return nil
	}

	taskID := uuid.NewString()
	tasks := make([]*task, 0, amount)
	messageIDs := make([]string, 0, amount)
	for _, v := range messages {
		messageIDs = append(messageIDs, v.MessageID)
		tasks = append(tasks, &task{
			TaskID:    taskID,
			MessageID: v.MessageID,
		})
	}
	if err := tx.Model(&message{}).Where("message_id IN ?", messageIDs).Update("task_id", taskID).Error; err != nil {
		return errors.Wrap(err, "failed to update message task id")
	}
	if err := tx.Create(&tasks).Error; err != nil {
		return errors.Wrap(err, "failed to create tasks")
	}
	return nil
}

func (p *persistence) save(msg *types.Message, aggregationAmount uint) error {
	return p.db.Transaction(func(tx *gorm.DB) error {
		if err := p.createMessageTx(tx, msg); err != nil {
			return err
		}
		if err := p.aggregateTaskTx(tx, int(aggregationAmount), msg); err != nil {
			return err
		}
		return nil
	})
}

func (p *persistence) fetchMessage(messageID string) ([]*types.MessageWithTime, error) {
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

func newPersistence(pgEndpoint string) (*persistence, error) {
	db, err := gorm.Open(postgres.Open(pgEndpoint), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect postgres")
	}
	if err := db.AutoMigrate(&message{}, &task{}); err != nil {
		return nil, errors.Wrap(err, "failed to migrate model")
	}
	return &persistence{db}, nil
}
