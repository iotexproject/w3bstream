package p2p

import (
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
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

	ps := &PubSubs{pubSubs: map[uint64]*subscriber{1: {}}}

	t.Run("ProjectIDAlreadyExists", func(t *testing.T) {
		r.Nil(ps.Add(1))
	})

	t.Run("FailedToNewSubscriber", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(newSubscriber, nil, errors.New(t.Name()))

		err := ps.Add(2)
		r.EqualError(err, t.Name())
	})

	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(newSubscriber, &subscriber{}, nil)

		err := ps.Add(3)
		r.NoError(err)
	})
}

func TestPubSubs_Delete(t *testing.T) {
	ps := &PubSubs{
		pubSubs: map[uint64]*subscriber{1: {}},
	}

	t.Run("IDNotExist", func(t *testing.T) {
		ps.Delete(101)
	})

	t.Run("Success", func(t *testing.T) {
		p := gomonkey.NewPatches()
		defer p.Reset()

		p = p.ApplyFuncReturn(newSubscriber, &subscriber{}, nil)
		p = p.ApplyPrivateMethod(&subscriber{}, "release", func(_ *subscriber) {})

		ps.Delete(1)
	})
}

func TestPubSubs_Publish(t *testing.T) {
	r := require.New(t)

	projectID := uint64(0x1)
	ps := &PubSubs{pubSubs: map[uint64]*subscriber{1: {}}}
	d := &Data{}

	t.Run("NotExist", func(t *testing.T) {
		r.Error(ps.Publish(102, d))
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
