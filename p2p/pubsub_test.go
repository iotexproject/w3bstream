package p2p

import (
	"log/slog"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	pubsub_pb "github.com/libp2p/go-libp2p-pubsub/pb"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/machinefi/sprout/testutil"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

type mockHost struct{ host.Host }

func (h mockHost) ID() peer.ID {
	return ""
}

func TestNewPubSubs(t *testing.T) {
	r := require.New(t)

	var (
		handle = func(data *Data, topic *pubsub.Topic) {}
		_host  = &mockHost{}
	)

	t.Run("FailedToNewP2PHost", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(libp2p.New, nil, errors.New(t.Name()))

		_, err := NewPubSubs(handle, "", 0)
		r.ErrorContains(err, t.Name())
	})
	t.Run("FailedToNewGossip", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(libp2p.New, _host, nil)
		p = p.ApplyFuncReturn(pubsub.NewGossipSub, nil, errors.New(t.Name()))

		_, err := NewPubSubs(handle, "", 0)
		r.ErrorContains(err, t.Name())
	})

	t.Run("FailedToDiscovery", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(libp2p.New, _host, nil)
		p = p.ApplyFuncReturn(pubsub.NewGossipSub, &pubsub.PubSub{}, nil)
		p = p.ApplyFuncReturn(discoverPeers, errors.New(t.Name()))

		_, err := NewPubSubs(handle, "", 0)
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(libp2p.New, _host, nil)
		p = p.ApplyFuncReturn(pubsub.NewGossipSub, nil, nil)
		p = p.ApplyFuncReturn(discoverPeers, nil)

		_, err := NewPubSubs(handle, "any", 0)
		r.NoError(err)
	})
}

func TestPubSubs_Add(t *testing.T) {
	r := require.New(t)

	ps := &PubSubs{pubSubs: map[uint64]*pubSub{1: {}}}

	t.Run("ProjectIDAlreadyExists", func(t *testing.T) {
		r.Nil(ps.Add(1))
	})

	t.Run("FailedToNewPubSub", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyFunc(NewPubSub, func(uint64, *pubsub.PubSub, HandleSubscriptionMessage, peer.ID) (*pubSub, error) {
			slog.Info("mock newpubsub")
			return nil, errors.New(t.Name())
		})

		err := ps.Add(2)
		r.EqualError(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(NewPubSub, &pubSub{}, nil)
		p = p.ApplyPrivateMethod(&pubSub{}, "run", func(_ *pubSub) {})

		err := ps.Add(3)
		r.NoError(err)
	})
}

func TestPubSubs_Delete(t *testing.T) {
	ps := &PubSubs{
		pubSubs: map[uint64]*pubSub{1: {ctxCancel: func() {}}},
	}

	t.Run("IDNotExist", func(t *testing.T) {
		ps.Delete(101)
	})

	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(NewPubSub, &pubSub{}, nil)
		p = p.ApplyPrivateMethod(&pubSub{}, "release", func(_ *pubSub) {})

		ps.Delete(1)
	})
}

func TestPubSubs_Publish(t *testing.T) {
	r := require.New(t)

	projectID := uint64(0x1)
	ps := &PubSubs{pubSubs: map[uint64]*pubSub{1: {}}}
	d := &Data{
		Task:         nil,
		TaskStateLog: nil,
	}

	t.Run("NotExist", func(t *testing.T) {
		r.Error(ps.Publish(102, d))
	})

	t.Run("FailedToMarshalJson", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = testutil.JsonMarshal(p, nil, errors.New(t.Name()))
		r.Error(ps.Publish(1, nil))
	})

	t.Run("FailedToPublishData", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyMethodReturn(&pubsub.Topic{}, "Publish", errors.New(t.Name()))
		r.ErrorContains(ps.Publish(projectID, d), t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyMethodReturn(&pubsub.Topic{}, "Publish", nil)
		r.NoError(ps.Publish(projectID, d))
	})
}

func TestPubSub_Release(t *testing.T) {
	ps := &pubSub{
		topic:        &pubsub.Topic{},
		ctxCancel:    func() {},
		subscription: &pubsub.Subscription{},
	}

	t.Run("FailedToCloseTopic", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyMethod(&pubsub.Subscription{}, "Cancel", func() {})
		p = p.ApplyMethodReturn(&pubsub.Topic{}, "Close", errors.New(t.Name()))

		ps.release()
	})

	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyMethod(&pubsub.Subscription{}, "Cancel", func() {})
		p = p.ApplyMethodReturn(&pubsub.Topic{}, "Close", nil)

		ps.release()
	})
}

func TestPubSub_NextMsg(t *testing.T) {
	r := require.New(t)

	ps := &pubSub{
		selfID:       peer.ID("test01"),
		subscription: &pubsub.Subscription{},
		handle:       func(data *Data, topic *pubsub.Topic) {},
	}

	t.Run("FailedToGetP2PData", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()
		p = p.ApplyMethodReturn(&pubsub.Subscription{}, "Next", nil, errors.New(t.Name()))
		defer p.Reset()

		err := ps.nextMsg()
		r.ErrorContains(err, t.Name())
	})

	t.Run("GetP2pDataFromSelf", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()
		p = p.ApplyMethodReturn(&pubsub.Subscription{}, "Next", &pubsub.Message{ReceivedFrom: ps.selfID}, nil)

		err := ps.nextMsg()
		r.NoError(err)
	})

	t.Run("FailedToUnmarshalP2PData", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyMethodReturn(&pubsub.Subscription{}, "Next", &pubsub.Message{
			ReceivedFrom: peer.ID("test02"),
			Message:      &pubsub_pb.Message{Data: nil},
		}, nil)

		err := ps.nextMsg()
		r.ErrorContains(err, "failed to json unmarshal p2p data")
	})

	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyMethodReturn(&pubsub.Subscription{}, "Next", &pubsub.Message{
			ReceivedFrom: peer.ID("test02"),
			Message:      &pubsub_pb.Message{Data: []byte("{}")},
		}, nil)
		r.Nil(ps.nextMsg())
	})
}

func TestNewPubSub(t *testing.T) {
	r := require.New(t)

	t.Run("FailedToJoinTopic", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyMethodReturn(&pubsub.PubSub{}, "Join", nil, errors.New(t.Name()))

		_, err := NewPubSub(uint64(0x1), &pubsub.PubSub{}, nil, peer.ID("0"))
		r.ErrorContains(err, t.Name())
	})

	t.Run("FailedToSubscribeTopic", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyMethodReturn(&pubsub.PubSub{}, "Join", &pubsub.Topic{}, nil)
		p = p.ApplyMethodReturn(&pubsub.Topic{}, "Subscribe", nil, errors.New(t.Name()))

		_, err := NewPubSub(uint64(0x1), &pubsub.PubSub{}, nil, peer.ID("0"))
		r.ErrorContains(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyMethodReturn(&pubsub.PubSub{}, "Join", &pubsub.Topic{}, nil)
		p = p.ApplyMethodReturn(&pubsub.Topic{}, "Subscribe", &pubsub.Subscription{}, nil)

		_, err := NewPubSub(uint64(0x1), &pubsub.PubSub{}, nil, peer.ID("0"))
		r.NoError(err)
	})
}
