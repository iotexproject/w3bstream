package postgres

import (
	"github.com/machinefi/sprout/msg"
	"github.com/machinefi/sprout/sequencer"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type pgSequencer struct {
	db *gorm.DB
}

func (s *pgSequencer) Save(msg *msg.Msg) (msgID uint64, err error) {
	m := Message{
		ProjectID: msg.ProjectID,
		Data:      msg.Data,
	}
	result := s.db.Create(&m)
	if result.Error != nil {
		return 0, result.Error
	}
	return uint64(m.ID), nil
}

func (s *pgSequencer) Fetch(projectID, afterMsgID uint64, strategy msg.FetchStrategy) ([]*msg.Msg, error) {
	// TODO FetchStrategy support

	ms := []*Message{}
	result := s.db.Find(&ms, "project_id = ? AND id > ?", projectID, afterMsgID)
	if result.Error != nil {
		return nil, result.Error
	}

	mms := []*msg.Msg{}
	for _, m := range ms {
		mms = append(mms, &msg.Msg{
			ProjectID: m.ProjectID,
			Data:      m.Data,
		})
	}
	return mms, nil
}

func NewSequencer(pgEndpoint string) (sequencer.Sequencer, error) {
	db, err := gorm.Open(postgres.Open(pgEndpoint), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "connect postgres failed")
	}

	if err := db.AutoMigrate(&Message{}); err != nil {
		return nil, errors.Wrap(err, "migrate message model failed")
	}

	return &pgSequencer{
		db: db,
	}, nil
}
