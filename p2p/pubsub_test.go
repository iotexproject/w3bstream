package p2p

import (
	"context"
	"reflect"
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	"github.com/golang/mock/gomock"
	"github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	pubsub_pb "github.com/libp2p/go-libp2p-pubsub/pb"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/testutil"
	"github.com/machinefi/sprout/testutil/mock"
)

func TestNewPubSubs(t *testing.T) {
	require := require.New(t)
	patches := NewPatches()
	defer patches.Reset()

	var handle HandleSubscriptionMessage = nil
	bootNodeMultiaddr := "/dns4/bootnode-0.testnet.iotex.one/tcp/4689/ipfs/12D3KooWFnaTYuLo8Mkbm3wzaWHtUuaxBRe24Uiopu15Wr5EhD3o"
	iotexChainID := 2

	ctrl := gomock.NewController(t)
	host := mock.NewMockHost(ctrl)

	t.Run("NewP2pHostFailed", func(t *testing.T) {
		patches = libp2pNew(patches, nil, errors.New(t.Name()))
		_, err := NewPubSubs(handle, bootNodeMultiaddr, iotexChainID)
		require.ErrorContains(err, t.Name())
	})
	patches = libp2pNew(patches, host, nil)

	t.Run("NewGossipFailed", func(t *testing.T) {
		patches = pubsubNewGossipSub(patches, nil, errors.New(t.Name()))
		_, err := NewPubSubs(handle, bootNodeMultiaddr, iotexChainID)
		require.ErrorContains(err, t.Name())
	})
	patches = pubsubNewGossipSub(patches, &pubsub.PubSub{}, nil)

	t.Run("DiscoveryFailed", func(t *testing.T) {
		patches = p2pDiscoverPeers(patches, errors.New(t.Name()))
		_, err := NewPubSubs(handle, bootNodeMultiaddr, iotexChainID)
		require.ErrorContains(err, t.Name())
	})
	patches = p2pDiscoverPeers(patches, nil)

	t.Run("NewPubSubs", func(t *testing.T) {
		host.EXPECT().ID().Return(peer.ID("ID")).Times(1)
		_, err := NewPubSubs(handle, bootNodeMultiaddr, iotexChainID)
		require.NoError(err)
	})
}

func TestPubSubs_Add(t *testing.T) {
	require := require.New(t)
	patches := NewPatches()
	defer patches.Reset()

	projectID := uint64(0x1)
	p := &PubSubs{pubSubs: make(map[uint64]*pubSub)}

	t.Run("NewPubSubFailed", func(t *testing.T) {
		patches = patches.ApplyFuncReturn(newPubSub, &pubSub{}, errors.New(t.Name()))
		err := p.Add(projectID)
		require.EqualError(err, t.Name())
	})
	patches = patches.ApplyFuncReturn(newPubSub, &pubSub{}, nil)

	t.Run("AddSuccess", func(t *testing.T) {
		patches = patches.ApplyPrivateMethod(&pubSub{}, "run", func() {})
		err := p.Add(projectID)
		require.NoError(err)
	})

	t.Run("AddRepeat", func(t *testing.T) {
		err := p.Add(projectID)
		require.NoError(err)
	})
}

func TestPubSubs_Delete(t *testing.T) {
	require := require.New(t)
	patches := NewPatches()
	defer patches.Reset()

	projectID := uint64(0x1)
	p := &PubSubs{pubSubs: make(map[uint64]*pubSub)}

	t.Run("IDNotExist", func(t *testing.T) {
		patches = patches.ApplyPrivateMethod(&pubSub{}, "release", func() {})
		p.Delete(projectID)
	})

	t.Run("DeleteSuccess", func(t *testing.T) {
		patches = p2pNewPubSub(patches, &pubSub{}, nil)
		patches = patches.ApplyPrivateMethod(&pubSub{}, "run", func() {})
		err := p.Add(projectID)
		require.NoError(err)

		p.Delete(projectID)
	})
}

func TestPubSubs_Publish(t *testing.T) {
	require := require.New(t)
	patches := NewPatches()
	defer patches.Reset()

	projectID := uint64(0x1)
	p := &PubSubs{pubSubs: make(map[uint64]*pubSub)}
	d := &Data{
		Task:         nil,
		TaskStateLog: nil,
	}

	t.Run("NotExist", func(t *testing.T) {
		err := p.Publish(projectID, d)
		require.ErrorContains(err, "project 1 topic not exist")
	})

	patches = p2pNewPubSub(patches, &pubSub{}, nil)
	patches = patches.ApplyPrivateMethod(&pubSub{}, "run", func() {})
	err := p.Add(projectID)
	require.NoError(err)

	t.Run("MarshalFailed", func(t *testing.T) {
		patches = testutil.JsonMarshal(patches, []byte("any"), errors.New(t.Name()))
		err := p.Publish(projectID, d)
		require.ErrorContains(err, t.Name())
	})
	patches = testutil.JsonMarshal(patches, []byte("any"), nil)

	t.Run("PublishDataFailed", func(t *testing.T) {
		patches = pubsubTopicPublish(patches, errors.New(t.Name()))
		err := p.Publish(projectID, d)
		require.ErrorContains(err, t.Name())
	})

	t.Run("PublishDataSuccess", func(t *testing.T) {
		patches = pubsubTopicPublish(patches, nil)
		err := p.Publish(projectID, d)
		require.NoError(err)
	})
}

func TestPubSub_Release(t *testing.T) {
	patches := NewPatches()
	defer patches.Reset()

	_, cancel := context.WithCancel(context.Background())
	p := &pubSub{
		topic:     &pubsub.Topic{},
		ctxCancel: cancel,
	}

	t.Run("TopicCloseFailed", func(t *testing.T) {
		patches = patches.ApplyMethodReturn(&pubsub.Subscription{}, "Cancel")
		patches = patches.ApplyMethodReturn(&pubsub.Topic{}, "Close", errors.New(t.Name()))
		p.release()
	})
	patches = patches.ApplyMethodReturn(&pubsub.Topic{}, "Close", nil)

	t.Run("TopicCloseSuccess", func(t *testing.T) {
		p.release()
	})
}

func TestPubSub_NextMsg(t *testing.T) {
	require := require.New(t)
	patches := NewPatches()
	defer patches.Reset()

	p := &pubSub{
		selfID: peer.ID("test01"),
	}

	t.Run("GetP2pDataFailed", func(t *testing.T) {
		patches = patches.ApplyMethodReturn(&pubsub.Subscription{}, "Next", nil, errors.New(t.Name()))
		err := p.nextMsg()
		require.ErrorContains(err, t.Name())
	})

	t.Run("GetP2pDataFromSelf", func(t *testing.T) {
		patches = patches.ApplyMethodReturn(&pubsub.Subscription{}, "Next", &pubsub.Message{ReceivedFrom: p.selfID}, nil)
		err := p.nextMsg()
		require.NoError(err)
	})

	t.Run("UnmarshalP2pDataFailed", func(t *testing.T) {
		patches = patches.ApplyMethodReturn(&pubsub.Subscription{}, "Next", &pubsub.Message{
			ReceivedFrom: peer.ID("test02"),
			Message:      &pubsub_pb.Message{Data: nil},
		}, nil)
		err := p.nextMsg()
		require.ErrorContains(err, "failed to json unmarshal p2p data")
	})
}

func TestNewPubSub(t *testing.T) {
	require := require.New(t)
	patches := NewPatches()
	defer patches.Reset()

	t.Run("JoinTopicFailed", func(t *testing.T) {
		patches = patches.ApplyMethodReturn(&pubsub.PubSub{}, "Join", nil, errors.New(t.Name()))
		_, err := newPubSub(uint64(0x1), &pubsub.PubSub{}, nil, peer.ID("0"))
		require.ErrorContains(err, t.Name())
	})

	t.Run("TopicSubscriptionFailed", func(t *testing.T) {
		patches = patches.ApplyMethodReturn(&pubsub.PubSub{}, "Join", nil, nil)
		patches = patches.ApplyMethodReturn(&pubsub.Topic{}, "Subscribe", nil, errors.New(t.Name()))
		_, err := newPubSub(uint64(0x1), &pubsub.PubSub{}, nil, peer.ID("0"))
		require.ErrorContains(err, t.Name())
	})

	t.Run("NewPubSubSuccess", func(t *testing.T) {
		patches = patches.ApplyMethodReturn(&pubsub.PubSub{}, "Join", nil, nil)
		patches = patches.ApplyMethodReturn(&pubsub.Topic{}, "Subscribe", nil, nil)
		_, err := newPubSub(uint64(0x1), &pubsub.PubSub{}, nil, peer.ID("0"))
		require.NoError(err)
	})
}

func libp2pNew(p *Patches, h host.Host, err error) *Patches {
	return p.ApplyFunc(
		libp2p.New,
		func(opts ...libp2p.Option) (host.Host, error) {
			return h, err
		},
	)
}

func pubsubNewGossipSub(p *Patches, ps *pubsub.PubSub, err error) *Patches {
	return p.ApplyFunc(
		pubsub.NewGossipSub,
		func(ctx context.Context, h host.Host, opts ...pubsub.Option) (*pubsub.PubSub, error) {
			return ps, err
		},
	)
}

func p2pDiscoverPeers(p *Patches, err error) *Patches {
	return p.ApplyFunc(
		discoverPeers,
		func(ctx context.Context, h host.Host, bootNodeMultiaddr string, iotexChainID int) error {
			return err
		},
	)
}

func p2pNewPubSub(p *Patches, pub *pubSub, err error) *Patches {
	return p.ApplyFunc(
		newPubSub,
		func(projectID uint64, ps *pubsub.PubSub, handle HandleSubscriptionMessage, selfID peer.ID) (*pubSub, error) {
			return pub, err
		},
	)
}

func pubsubTopicPublish(p *Patches, err error) *Patches {
	var topic *pubsub.Topic
	return p.ApplyMethodFunc(
		reflect.TypeOf(topic),
		"Publish",
		func(ctx context.Context, data []byte, opts ...pubsub.PubOpt) error {
			return err
		},
	)
}
