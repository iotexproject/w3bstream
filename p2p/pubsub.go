package p2p

import (
	"context"
	"log/slog"
	"strconv"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/pkg/errors"
)

func newSubscriber(projectID uint64, ps *pubsub.PubSub, handle HandleSubscriptionMessage, selfID peer.ID) (*subscriber, error) {
	topic, err := ps.Join("w3bstream-project-" + strconv.FormatUint(projectID, 10))
	if err != nil {
		return nil, errors.Wrapf(err, "join topic %v failed", projectID)
	}
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, errors.Wrapf(err, "topic %v subscription failed", projectID)
	}
	ctx, cancel := context.WithCancel(context.Background())

	_ps := &subscriber{
		selfID:       selfID,
		topic:        topic,
		subscription: sub,
		handle:       handle,
		cancel:       cancel,
	}

	go _ps.run(ctx)

	return _ps, nil
}

type subscriber struct {
	selfID       peer.ID
	topic        *pubsub.Topic
	subscription *pubsub.Subscription
	handle       HandleSubscriptionMessage
	cancel       context.CancelFunc
}

func (p *subscriber) release() {
	p.subscription.Cancel()
	p.cancel()
	if err := p.topic.Close(); err != nil {
		slog.Error("failed to close topic", "error", err, "topic", p.topic.String())
	}
}

func (p *subscriber) run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			slog.With("ctx.Err()", ctx.Err()).Info("pubsub stopped caused by")
			return
		default:
			m, err := p.subscription.Next(ctx)
			if err != nil {
				slog.Error("failed to get p2p data", "error", err)
				continue
			}
			if m.ReceivedFrom == p.selfID {
				continue
			}
			p.handle(m.Message.Data, p.topic)
		}
	}
}
