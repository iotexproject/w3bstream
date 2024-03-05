package p2p

import (
	"context"
	"encoding/json"
	"log/slog"
	"strconv"
	"sync"

	"github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/muxer/yamux"
	"github.com/pkg/errors"
)

type HandleSubscriptionMessage func(*Data, *pubsub.Topic)

type PubSubs struct {
	mux     sync.RWMutex
	pubSubs map[uint64]*pubSub
	ps      *pubsub.PubSub
	selfID  peer.ID
	handle  HandleSubscriptionMessage
}

func (p *PubSubs) Add(projectID uint64) error {
	p.mux.Lock()
	defer p.mux.Unlock()

	if _, ok := p.pubSubs[projectID]; ok {
		return nil
	}

	nps, err := newPubSub(projectID, p.ps, p.handle, p.selfID)
	if err != nil {
		return err
	}
	go nps.run()

	p.pubSubs[projectID] = nps
	return nil
}

// TODO delete not used currently
func (p *PubSubs) Delete(projectID uint64) {
	p.mux.Lock()
	defer p.mux.Unlock()

	pubSub, ok := p.pubSubs[projectID]
	if !ok {
		return
	}
	pubSub.release()
	delete(p.pubSubs, projectID)
}

func (p *PubSubs) Publish(projectID uint64, d *Data) error {
	p.mux.RLock()
	defer p.mux.RUnlock()

	s, ok := p.pubSubs[projectID]
	if !ok {
		return errors.Errorf("project %v topic not exist", projectID)
	}
	j, err := json.Marshal(d)
	if err != nil {
		return errors.Wrap(err, "json marshal p2p data failed")
	}
	if err := s.topic.Publish(context.Background(), j); err != nil {
		return errors.Wrap(err, "publish data to p2p network failed")
	}
	return nil
}

func NewPubSubs(handle HandleSubscriptionMessage, bootNodeMultiaddr string, iotexChainID int) (*PubSubs, error) {
	ctx := context.Background()
	h, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"), libp2p.Muxer("/yamux/2.0.0", yamux.DefaultTransport))
	if err != nil {
		return nil, errors.Wrap(err, "new libp2p host failed")
	}

	ps, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		return nil, errors.Wrap(err, "new gossip subscription failed")
	}
	if err := discoverPeers(ctx, h, bootNodeMultiaddr, iotexChainID); err != nil {
		return nil, err
	}

	return &PubSubs{
		ps:      ps,
		pubSubs: make(map[uint64]*pubSub),
		selfID:  h.ID(),
		handle:  handle,
	}, nil
}

type pubSub struct {
	selfID       peer.ID
	topic        *pubsub.Topic
	subscription *pubsub.Subscription
	handle       HandleSubscriptionMessage
	ctx          context.Context
	ctxCancel    context.CancelFunc
}

func (p *pubSub) release() {
	p.ctxCancel()
	p.subscription.Cancel()
	if err := p.topic.Close(); err != nil {
		slog.Error("close topic failed", "error", err, "topic", p.topic.String())
	}
}

func (p *pubSub) run() {
	for {
		if err := p.nextMsg(); err != nil {
			slog.Error("failed to pubSub get msg", "error", err)
		}
	}
}

func (p *pubSub) nextMsg() error {
	m, err := p.subscription.Next(p.ctx)
	if err != nil {
		return errors.Wrapf(err, "failed to get p2p data")
	}
	if m.ReceivedFrom == p.selfID {
		return nil
	}
	d := &Data{}
	if err := json.Unmarshal(m.Message.Data, d); err != nil {
		return errors.Wrapf(err, "failed to json unmarshal p2p data")
	}
	p.handle(d, p.topic)
	return nil
}

func newPubSub(projectID uint64, ps *pubsub.PubSub, handle HandleSubscriptionMessage, selfID peer.ID) (*pubSub, error) {
	topic, err := ps.Join("w3bstream-project-" + strconv.FormatUint(projectID, 10))
	if err != nil {
		return nil, errors.Wrapf(err, "join topic %v failed", projectID)
	}
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, errors.Wrapf(err, "topic %v subscription failed", projectID)
	}
	ctx, cancel := context.WithCancel(context.Background())

	return &pubSub{
		selfID:       selfID,
		topic:        topic,
		subscription: sub,
		handle:       handle,
		ctx:          ctx,
		ctxCancel:    cancel,
	}, nil
}
