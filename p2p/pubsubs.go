package p2p

import (
	"context"
	"log/slog"

	"github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/muxer/yamux"
	"github.com/pkg/errors"
)

type HandleSubscription func([]byte)

type PubSub struct {
	topic *pubsub.Topic
}

func run(ctx context.Context, sub *pubsub.Subscription, selfID peer.ID, handle HandleSubscription) {
	for {
		m, err := sub.Next(ctx)
		if err != nil {
			slog.Error("failed to get p2p message", "error", err)
			continue
		}
		if m.ReceivedFrom == selfID {
			continue
		}
		handle(m.Message.Data)
	}
}

func (p *PubSub) Publish(data []byte) error {
	if err := p.topic.Publish(context.Background(), data); err != nil {
		return errors.Wrap(err, "failed to publish data to p2p network")
	}
	return nil
}

func NewPubSub(handle HandleSubscription, bootNodeMultiaddr string, iotexChainID int) (*PubSub, error) {
	ctx := context.Background()
	h, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"), libp2p.Muxer("/yamux/2.0.0", yamux.DefaultTransport))
	if err != nil {
		return nil, errors.Wrap(err, "failed to new libp2p host")
	}

	ps, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		return nil, errors.Wrap(err, "failed to new gossip subscription")
	}
	if err := discoverPeers(ctx, h, bootNodeMultiaddr, iotexChainID); err != nil {
		return nil, err
	}
	topic, err := ps.Join("w3bstream-task")
	if err != nil {
		return nil, errors.Wrapf(err, "failed to join topic w3bstream-task")
	}
	if handle != nil {
		sub, err := topic.Subscribe()
		if err != nil {
			return nil, errors.Wrapf(err, "failed to subscribe topic w3bstream-task")
		}
		go run(context.Background(), sub, h.ID(), handle)
	}

	return &PubSub{
		topic: topic,
	}, nil
}
