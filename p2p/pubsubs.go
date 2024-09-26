package p2p

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/ethereum/go-ethereum/common"
	"github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/muxer/yamux"
	"github.com/pkg/errors"
)

type HandleSubscription func(projectID uint64, taskID common.Hash) error

type PubSub struct {
	topic *pubsub.Topic
}

type data struct {
	ProjectID uint64
	TaskID    common.Hash
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
		d := &data{}
		if err := json.Unmarshal(m.Message.Data, d); err != nil {
			slog.Error("failed to unmarshal p2p message", "error", err)
			continue
		}
		if err := handle(d.ProjectID, d.TaskID); err != nil {
			slog.Error("failed to handle p2p message", "error", err)
		}
	}
}

func (p *PubSub) Publish(projectID uint64, taskID common.Hash) error {
	j, err := json.Marshal(&data{ProjectID: projectID, TaskID: taskID})
	if err != nil {
		return errors.Wrap(err, "failed to marshal data")
	}
	err = p.topic.Publish(context.Background(), j)
	return errors.Wrap(err, "failed to publish data to p2p network")
}

func NewPubSub(bootNodeMultiaddr string, iotexChainID int, handle HandleSubscription) (*PubSub, error) {
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
