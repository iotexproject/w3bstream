package coordinator

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/types"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Coordinator struct {
	db    *gorm.DB
	topic *pubsub.Topic
	sub   *pubsub.Subscription
}

func (s *Coordinator) Save(msg *types.Message) error {
	m := Message{
		MessageID: msg.MessageID,
		ProjectID: msg.ProjectID,
		Data:      msg.Data,
		State:     types.MessageStateReceived,
	}
	l := MessageStateLog{
		MessageID: msg.MessageID,
		State:     types.MessageStateReceived,
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

func (s *Coordinator) fetch(projectID uint64) (*types.Message, error) {
	m := Message{}
	if err := s.db.Where("project_id = ? AND state = ?", projectID, types.MessageStateReceived).First(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "query message failed, projectID %d", projectID)
	}

	return &types.Message{
		MessageID: m.MessageID,
		ProjectID: m.ProjectID,
		Data:      m.Data,
	}, nil
}

func (s *Coordinator) FetchStateLog(messageID string) ([]*MessageStateLog, error) {
	ls := []*MessageStateLog{}

	if err := s.db.Where("message_id = ?", messageID).Find(&ls).Error; err != nil {
		return nil, errors.Wrapf(err, "query message state log failed, messageID %s", messageID)
	}
	return ls, nil
}

func (s *Coordinator) updateMessageState(msgIDs []string, state types.MessageState, comment string) error {
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

func (r *Coordinator) Run() {
	for {
		m, err := r.sub.Next(context.Background())
		if err != nil {
			slog.Error("get p2p data failed", "error", err)
			continue
		}
		d := p2p.Data{}
		if err := json.Unmarshal(m.Message.Data, &d); err != nil {
			slog.Error("json unmarshal p2p data failed", "error", err)
			continue
		}
		switch {
		case d.Request != nil:
			r.handleRequest(d.Request.ProjectID)
		case d.Response != nil:
			r.handleResponse(d.Response.MessageIDs, d.Response.State, d.Response.Comment)
		}
	}
}

func (r *Coordinator) handleRequest(projectID uint64) {
	m, err := r.fetch(projectID)
	if err != nil {
		slog.Error("fetch message failed", "error", err)
		return
	}
	if m == nil {
		return
	}
	d := p2p.Data{
		Message: &p2p.MessageData{Messages: []*types.Message{m}},
	}
	j, err := json.Marshal(&d)
	if err != nil {
		slog.Error("json marshal p2p data failed", "error", err)
		return
	}

	if err := r.topic.Publish(context.Background(), j); err != nil {
		slog.Error("publish data to p2p network failed", "error", err)
	}
}

func (r *Coordinator) handleResponse(messageIDs []string, state types.MessageState, comment string) {
	if len(messageIDs) == 0 {
		return
	}
	if err := r.updateMessageState(messageIDs, state, comment); err != nil {
		slog.Error("update message state failed", "error", err, "messageIDs", messageIDs)
		return
	}
}

func NewCoordinator(pgEndpoint, p2pMultiaddr string) (*Coordinator, error) {
	db, err := gorm.Open(postgres.Open(pgEndpoint), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, errors.Wrap(err, "connect postgres failed")
	}
	if err := db.AutoMigrate(&Message{}, &MessageStateLog{}); err != nil {
		return nil, errors.Wrap(err, "migrate message model failed")
	}

	h, err := libp2p.New(libp2p.ListenAddrStrings(p2pMultiaddr))
	if err != nil {
		return nil, errors.Wrap(err, "new libp2p host failed")
	}
	ctx := context.Background()
	go p2p.DiscoverPeers(ctx, h, p2p.Topic)

	ps, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		return nil, errors.Wrap(err, "new gossip subscription failed")
	}
	topic, err := ps.Join(p2p.Topic)
	if err != nil {
		return nil, errors.Wrap(err, "join topic failed")
	}

	sub, err := topic.Subscribe()
	if err != nil {
		return nil, errors.Wrap(err, "topic subscription failed")
	}

	return &Coordinator{
		db:    db,
		topic: topic,
		sub:   sub,
	}, nil
}
