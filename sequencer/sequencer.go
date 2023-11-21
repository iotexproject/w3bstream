package sequencer

import (
	"github.com/machinefi/sprout/enums"
	"github.com/machinefi/sprout/message"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Sequencer struct {
	db *gorm.DB
}

func (s *Sequencer) Save(msg *message.Message) error {
	m := Message{
		MessageID: msg.ID,
		ProjectID: msg.ProjectID,
		Data:      msg.Data,
	}
	result := s.db.Create(&m)
	if result.Error != nil {
		return errors.Wrap(result.Error, "save message failed")
	}
	return nil
}

func (s *Sequencer) Fetch(projectID uint64) (*message.Message, error) {
	m := Message{}
	result := s.db.Where("project_id = ? AND state = ?", projectID, enums.MessageReceived).First(&m)
	if result.Error != nil {
		return nil, errors.Wrapf(result.Error, "query message failed, projectID %d", projectID)
	}

	return &message.Message{
		ID:        m.MessageID,
		ProjectID: m.ProjectID,
		Data:      m.Data,
	}, nil
}

func (s *Sequencer) UpdateMessageState(msgID string, state enums.MessageState) error {
	result := s.db.Model(&Message{}).Where("message_id = ?", msgID).Update("state", state)
	if result.Error != nil {
		return errors.Wrapf(result.Error, "update message failed, message_id %s, target_state %v", msgID, state)
	}

	return nil
}

func NewSequencer(pgEndpoint string) (*Sequencer, error) {
	db, err := gorm.Open(postgres.Open(pgEndpoint), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "connect postgres failed")
	}

	if err := db.AutoMigrate(&Message{}); err != nil {
		return nil, errors.Wrap(err, "migrate message model failed")
	}

	return &Sequencer{
		db: db,
	}, nil
}
