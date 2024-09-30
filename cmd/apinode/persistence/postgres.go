package persistence

import (
	"crypto/ecdsa"
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/iotexproject/w3bstream/p2p"
	"github.com/iotexproject/w3bstream/task"
	"github.com/pkg/errors"
	"gorm.io/datatypes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type Message struct {
	gorm.Model
	DeviceID       common.Address `gorm:"index:message_fetch,not null"`
	ProjectID      uint64         `gorm:"index:message_fetch,not null"`
	ProjectVersion string         `gorm:"index:message_fetch,not null"`
	TaskID         common.Hash    `gorm:"index:task_id,not null"`
	Data           []byte         `gorm:"size:4096"`
}

type Task struct {
	gorm.Model
	TaskID     common.Hash    `gorm:"uniqueIndex:task,not null"`
	ProjectID  uint64         `gorm:"uniqueIndex:task,not null"`
	MessageIDs datatypes.JSON `gorm:"not null"`
	Signature  []byte         `gorm:"not null"`
}

type Persistence struct {
	db *gorm.DB
}

func (p *Persistence) createMessageTx(tx *gorm.DB, m *Message) error {
	if err := tx.Create(m).Error; err != nil {
		return errors.Wrap(err, "failed to create message")
	}
	return nil
}

func (p *Persistence) aggregateTaskTx(tx *gorm.DB, pubSub *p2p.PubSub, amount int, m *Message, prv *ecdsa.PrivateKey) (common.Hash, error) {
	messages := make([]*Message, 0)
	if amount <= 0 {
		amount = 1
	}

	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Order("created_at").
		Where(
			"project_id = ? AND project_version = ? AND device_id = ? AND task_id = ?",
			m.ProjectID, m.ProjectVersion, m.DeviceID, common.Hash{},
		).Limit(amount).Find(&messages).Error; err != nil {
		return common.Hash{}, errors.Wrap(err, "failed to fetch unpacked messages")
	}

	// no enough message for pack task
	if len(messages) < amount {
		return common.Hash{}, nil
	}

	taskID := crypto.Keccak256Hash([]byte(uuid.NewString()))
	messageIDs := make([]uint, 0, amount)
	for _, v := range messages {
		messageIDs = append(messageIDs, v.ID)
	}
	if err := tx.Model(&Message{}).Where("id IN ?", messageIDs).Update("task_id", taskID).Error; err != nil {
		return common.Hash{}, errors.Wrap(err, "failed to update message task id")
	}
	messageIDsJson, err := json.Marshal(messageIDs)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "failed to marshal message id array")
	}
	data := make([][]byte, 0, len(messages))
	for _, v := range messages {
		data = append(data, v.Data)
	}

	t := &task.Task{
		ID:             taskID,
		ProjectID:      m.ProjectID,
		ProjectVersion: m.ProjectVersion,
		DeviceID:       m.DeviceID,
		Payloads:       data,
	}
	sig, err := t.Sign(prv)
	if err != nil {
		return common.Hash{}, err
	}

	mt := &Task{
		TaskID:     taskID,
		ProjectID:  m.ProjectID,
		MessageIDs: messageIDsJson,
		Signature:  sig,
	}

	if err := tx.Create(mt).Error; err != nil {
		return common.Hash{}, errors.Wrap(err, "failed to create Task")
	}
	return taskID, pubSub.Publish(m.ProjectID, taskID)
}

func (p *Persistence) Save(pubSub *p2p.PubSub, msg *Message, aggregationAmount int, prv *ecdsa.PrivateKey) (common.Hash, error) {
	taskID := common.Hash{}
	err := p.db.Transaction(func(tx *gorm.DB) error {
		if err := p.createMessageTx(tx, msg); err != nil {
			return err
		}
		id, err := p.aggregateTaskTx(tx, pubSub, aggregationAmount, msg, prv)
		if err != nil {
			return err
		}
		taskID = id
		return nil
	})
	return taskID, err
}

func (p *Persistence) FetchMessage(messageID string) ([]*Message, error) {
	ms := []*Message{}
	if err := p.db.Where("message_id = ?", messageID).Find(&ms).Error; err != nil {
		return nil, errors.Wrapf(err, "query message by messageID failed, messageID %s", messageID)
	}

	return ms, nil
}

func (p *Persistence) FetchTask(internalTaskID string) ([]*Task, error) {
	ts := []*Task{}
	if err := p.db.Where("internal_task_id = ?", internalTaskID).Find(&ts).Error; err != nil {
		return nil, errors.Wrapf(err, "query Task by internal Task id failed, internal_task_id %s", internalTaskID)
	}

	return ts, nil
}

func NewPersistence(pgEndpoint string) (*Persistence, error) {
	db, err := gorm.Open(postgres.Open(pgEndpoint), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect postgres")
	}
	if err := db.AutoMigrate(&Message{}, &Task{}); err != nil {
		return nil, errors.Wrap(err, "failed to migrate model")
	}
	return &Persistence{db}, nil
}
