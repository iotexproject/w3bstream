package sequencer

import (
	"github.com/machinefi/sprout/enums"
	"github.com/machinefi/sprout/message"
	"github.com/machinefi/sprout/proto"
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
	l := MessageStateLog{
		MessageID: msg.ID,
		State:     proto.MessageState_MESSAGE_STATE_RECEIVED,
	}

	return s.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(&m).Error; err != nil {
			return errors.Wrap(err, "create message failed")
		}
		if err := tx.Create(&l).Error; err != nil {
			return errors.Wrap(err, "create message state log failed")
		}
		return nil
	})
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

func (s *Sequencer) FetchStateLog(messageID string) ([]*MessageStateLog, error) {
	ls := []*MessageStateLog{}

	if err := s.db.Where("message_id = ?", messageID).Find(&ls).Error; err != nil {
		return nil, errors.Wrapf(err, "query message state log failed, messageID %s", messageID)
	}
	return ls, nil
}

func (s *Sequencer) UpdateMessageState(msgIDs []string, state proto.MessageState, comment string) error {
	ls := []*MessageStateLog{}
	for _, id := range msgIDs {
		ls = append(ls, &MessageStateLog{
			MessageID: id,
			State:     state,
			Comment:   comment,
		})
	}

	return s.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Model(&Message{}).Where("message_id IN ?", msgIDs).Update("state", state).Error; err != nil {
			return errors.Wrapf(err, "update message failed, message_ids %v, target_state %v", msgIDs, state)
		}
		if err := tx.Create(ls).Error; err != nil {
			return errors.Wrap(err, "create message state log failed")
		}
		return nil
	})
}

func NewSequencer(pgEndpoint string) (*Sequencer, error) {
	db, err := gorm.Open(postgres.Open(pgEndpoint), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "connect postgres failed")
	}

	if err := db.AutoMigrate(&Message{}, &MessageStateLog{}); err != nil {
		return nil, errors.Wrap(err, "migrate message model failed")
	}

	return &Sequencer{
		db: db,
	}, nil
}
