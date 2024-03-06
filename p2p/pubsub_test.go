package p2p

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/golang/mock/gomock"
	"github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	pubsub_pb "github.com/libp2p/go-libp2p-pubsub/pb"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/testutil/mock"
)

func TestNewPubSubs(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	host := mock.NewMockHost(ctrl)

	var handle HandleSubscriptionMessage = nil

	t.Run("FailedToNewP2PHost", func(t *testing.T) {
		p = p.ApplyFuncReturn(libp2p.New, nil, errors.New(t.Name()))

		_, err := NewPubSubs(handle, "", 0)
		r.ErrorContains(err, t.Name())
	})
	p = p.ApplyFuncReturn(libp2p.New, host, nil)
	t.Run("FailedToNewGossip", func(t *testing.T) {
		p = p.ApplyFuncReturn(pubsub.NewGossipSub, nil, errors.New(t.Name()))

		_, err := NewPubSubs(handle, "", 0)
		r.ErrorContains(err, t.Name())
	})
	p = p.ApplyFuncReturn(pubsub.NewGossipSub, &pubsub.PubSub{}, nil)

	t.Run("FailedToDiscovery", func(t *testing.T) {
		p = p.ApplyFuncReturn(discoverPeers, errors.New(t.Name()))

		_, err := NewPubSubs(handle, "", 0)
		r.ErrorContains(err, t.Name())
	})
	p = p.ApplyFuncReturn(discoverPeers, nil)

	t.Run("Success", func(t *testing.T) {
		host.EXPECT().ID().Return(peer.ID("ID")).Times(1)
		_, err := NewPubSubs(handle, "", 0)
		r.NoError(err)
	})
}

func TestPubSubs_Add(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	projectID := uint64(0x1)
	pubSubs := &PubSubs{pubSubs: make(map[uint64]*pubSub)}

	t.Run("FailedToNewPubSub", func(t *testing.T) {
		p = p.ApplyFuncReturn(newPubSub, &pubSub{}, errors.New(t.Name()))

		err := pubSubs.Add(projectID)
		r.EqualError(err, t.Name())
	})
	p = p.ApplyFuncReturn(newPubSub, &pubSub{}, nil)

	t.Run("Success", func(t *testing.T) {
		p = p.ApplyPrivateMethod(&pubSub{}, "run", func() {})
		err := pubSubs.Add(projectID)
		r.NoError(err)
	})
	t.Run("Repeat", func(t *testing.T) {
		err := pubSubs.Add(projectID)
		r.NoError(err)
	})
}

func TestPubSubs_Delete(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	projectID := uint64(0x1)
	pubSubs := &PubSubs{pubSubs: make(map[uint64]*pubSub)}

	t.Run("IDNotExist", func(t *testing.T) {
		pubSubs.Delete(projectID)
	})
	t.Run("Success", func(t *testing.T) {
		p = p.ApplyFuncReturn(newPubSub, &pubSub{}, nil)
		p = p.ApplyPrivateMethod(&pubSub{}, "run", func() {})
		p = p.ApplyPrivateMethod(&pubSub{}, "release", func() {})

		err := pubSubs.Add(projectID)
		r.NoError(err)

		pubSubs.Delete(projectID)
	})
}

func TestPubSubs_Publish(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	projectID := uint64(0x1)
	pubSubs := &PubSubs{pubSubs: make(map[uint64]*pubSub)}
	d := &Data{
		Task:         nil,
		TaskStateLog: nil,
	}

	t.Run("NotExist", func(t *testing.T) {
		err := pubSubs.Publish(projectID, d)
		r.ErrorContains(err, "project 1 topic not exist")
	})
	p = p.ApplyFuncReturn(newPubSub, &pubSub{}, nil)
	p = p.ApplyPrivateMethod(&pubSub{}, "run", func() {})

	err := pubSubs.Add(projectID)
	r.NoError(err)

	t.Run("FailedToMarshalJson", func(t *testing.T) {
		p = p.ApplyFuncReturn(json.Marshal, nil, errors.New(t.Name()))
		err := pubSubs.Publish(projectID, d)
		r.ErrorContains(err, t.Name())
	})
	p = p.ApplyFuncReturn(json.Marshal, []byte("any"), nil)

	t.Run("FailedToPublishData", func(t *testing.T) {
		p = p.ApplyMethodFunc(&pubsub.Topic{}, "Publish", errors.New(t.Name()))
		err := pubSubs.Publish(projectID, d)
		r.ErrorContains(err, t.Name())
	})
	p = p.ApplyMethodFunc(&pubsub.Topic{}, "Publish", nil)

	t.Run("Success", func(t *testing.T) {
		err := pubSubs.Publish(projectID, d)
		r.NoError(err)
	})
}

func TestPubSub_Release(t *testing.T) {
	p := gomonkey.NewPatches()
	defer p.Reset()

	_, cancel := context.WithCancel(context.Background())
	pubSubs := &pubSub{
		topic:     &pubsub.Topic{},
		ctxCancel: cancel,
	}

	t.Run("FailedToCloseTopic", func(t *testing.T) {
		p = p.ApplyMethodReturn(&pubsub.Subscription{}, "Cancel")
		p = p.ApplyMethodReturn(&pubsub.Topic{}, "Close", errors.New(t.Name()))

		pubSubs.release()
	})
	p = p.ApplyMethodReturn(&pubsub.Topic{}, "Close", nil)

	t.Run("Success", func(t *testing.T) {
		pubSubs.release()
	})
}

func TestPubSub_NextMsg(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	pubSub := &pubSub{
		selfID: peer.ID("test01"),
	}

	t.Run("FailedToGetP2PData", func(t *testing.T) {
		p = p.ApplyMethodReturn(&pubsub.Subscription{}, "Next", nil, errors.New(t.Name()))

		err := pubSub.nextMsg()
		r.ErrorContains(err, t.Name())
	})
	t.Run("GetP2pDataFromSelf", func(t *testing.T) {
		p = p.ApplyMethodReturn(&pubsub.Subscription{}, "Next", &pubsub.Message{ReceivedFrom: pubSub.selfID}, nil)

		err := pubSub.nextMsg()
		r.NoError(err)
	})
	t.Run("FailedToUnmarshalP2PData", func(t *testing.T) {
		p = p.ApplyMethodReturn(&pubsub.Subscription{}, "Next", &pubsub.Message{
			ReceivedFrom: peer.ID("test02"),
			Message:      &pubsub_pb.Message{Data: nil},
		}, nil)

		err := pubSub.nextMsg()
		r.ErrorContains(err, "failed to json unmarshal p2p data")
	})
}

func TestNewPubSub(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	t.Run("FailedToJoinTopic", func(t *testing.T) {
		p = p.ApplyMethodReturn(&pubsub.PubSub{}, "Join", nil, errors.New(t.Name()))

		_, err := newPubSub(uint64(0x1), &pubsub.PubSub{}, nil, peer.ID("0"))
		r.ErrorContains(err, t.Name())
	})
	p = p.ApplyMethodReturn(&pubsub.PubSub{}, "Join", nil, nil)

	t.Run("FailedToSubscribeTopic", func(t *testing.T) {
		p = p.ApplyMethodReturn(&pubsub.Topic{}, "Subscribe", nil, errors.New(t.Name()))
		_, err := newPubSub(uint64(0x1), &pubsub.PubSub{}, nil, peer.ID("0"))
		r.ErrorContains(err, t.Name())
	})
	p = p.ApplyMethodReturn(&pubsub.Topic{}, "Subscribe", nil, nil)

	t.Run("Success", func(t *testing.T) {
		_, err := newPubSub(uint64(0x1), &pubsub.PubSub{}, nil, peer.ID("0"))
		r.NoError(err)
	})
}
