package persistence

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/binary"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/datatypes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type Message struct {
	gorm.Model
	MessageID      string `gorm:"index:message_id,not null"`
	ClientID       string `gorm:"index:message_fetch,not null,default:''"`
	ProjectID      uint64 `gorm:"index:message_fetch,not null"`
	ProjectVersion string `gorm:"index:message_fetch,not null,default:'0.0'"`
	Data           []byte `gorm:"size:4096"`
	InternalTaskID string `gorm:"index:internal_task_id,not null,default:''"`
}

type Task struct {
	gorm.Model
	ProjectID      uint64         `gorm:"index:task_fetch,not null"`
	InternalTaskID string         `gorm:"index:internal_task_id,not null"`
	MessageIDs     datatypes.JSON `gorm:"not null"`
	Signature      string         `gorm:"not null,default:''"`
}

func (t *Task) sign(sk *ecdsa.PrivateKey, projectID uint64, clientID string, messages ...[]byte) (string, error) {
	buf := bytes.NewBuffer(nil)

	if err := binary.Write(buf, binary.BigEndian, uint64(t.ID)); err != nil {
		return "", err
	}
	if err := binary.Write(buf, binary.BigEndian, projectID); err != nil {
		return "", err
	}
	if _, err := buf.WriteString(clientID); err != nil {
		return "", err
	}
	if _, err := buf.Write(crypto.Keccak256Hash(messages...).Bytes()); err != nil {
		return "", err
	}

	h := crypto.Keccak256Hash(buf.Bytes())
	sig, err := crypto.Sign(h.Bytes(), sk)
	if err != nil {
		return "", err
	}
	return hexutil.Encode(sig), nil
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

func (p *Persistence) aggregateTaskTx(tx *gorm.DB, amount int, m *Message, sk *ecdsa.PrivateKey) error {
	messages := make([]*Message, 0)
	if amount == 0 {
		amount = 1
	}

	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Order("created_at").
		Where(
			"project_id = ? AND project_version = ? AND client_id = ? AND internal_task_id = ?",
			m.ProjectID, m.ProjectVersion, m.ClientID, "",
		).Limit(amount).Find(&messages).Error; err != nil {
		return errors.Wrap(err, "failed to fetch unpacked messages")
	}

	// no enough message for pack Task
	if len(messages) < amount {
		return nil
	}

	taskID := uuid.NewString()
	messageIDs := make([]string, 0, amount)
	for _, v := range messages {
		messageIDs = append(messageIDs, v.MessageID)
	}
	if err := tx.Model(&Message{}).Where("message_id IN ?", messageIDs).Update("internal_task_id", taskID).Error; err != nil {
		return errors.Wrap(err, "failed to update message internal Task id")
	}
	messageIDsJson, err := json.Marshal(messageIDs)
	if err != nil {
		return errors.Wrap(err, "failed to marshal message id array")
	}

	t := &Task{
		InternalTaskID: taskID,
		ProjectID:      m.ProjectID,
		MessageIDs:     messageIDsJson,
	}

	if err := tx.Create(t).Error; err != nil {
		return errors.Wrap(err, "failed to create Task")
	}
	data := make([][]byte, 0, len(messages))
	for _, v := range messages {
		data = append(data, v.Data)
	}

	sig, err := t.sign(sk, m.ProjectID, m.ClientID, data...)
	if err != nil {
		return errors.Wrap(err, "failed to sign Task")
	}

	if err := tx.Model(t).Update("signature", sig).Where("id = ?", t.ID).Error; err != nil {
		return errors.Wrap(err, "failed to update Task sign")
	}

	return nil
}

func (p *Persistence) Save(msg *Message, aggregationAmount uint, sk *ecdsa.PrivateKey) error {
	return p.db.Transaction(func(tx *gorm.DB) error {
		if err := p.createMessageTx(tx, msg); err != nil {
			return err
		}
		if err := p.aggregateTaskTx(tx, int(aggregationAmount), msg, sk); err != nil {
			return err
		}
		return nil
	})
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
