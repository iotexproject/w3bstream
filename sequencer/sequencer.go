package sequencer

import (
	"github.com/machinefi/sprout/proto"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Sequencer struct {
	db *gorm.DB
}

func (s *Sequencer) Save(msg *proto.Message) error {
	m := Message{
		MessageID: msg.MessageID,
		ProjectID: msg.ProjectID,
		Data:      msg.Data,
		State:     proto.MessageState_RECEIVED,
	}
	l := MessageStateLog{
		MessageID: msg.MessageID,
		State:     proto.MessageState_RECEIVED,
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

func (s *Sequencer) Fetch(projectID uint64) (*proto.Message, error) {
	m := Message{}
	if err := s.db.Where("project_id = ? AND state = ?", projectID, proto.MessageState_RECEIVED).First(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "query message failed, projectID %d", projectID)
	}

	return &proto.Message{
		MessageID: m.MessageID,
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
	db, err := gorm.Open(postgres.Open(pgEndpoint), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
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
