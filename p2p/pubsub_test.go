package p2p

import (
	"context"
	"encoding/json"
	"log/slog"
	"reflect"
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	. "github.com/bytedance/mockey"
	"github.com/golang/mock/gomock"
	"github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	pubsub_pb "github.com/libp2p/go-libp2p-pubsub/pb"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/machinefi/sprout/testutil/mock"
	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/testutil"
)

func TestNewPubSubs(t *testing.T) {
	require := require.New(t)
	patches := NewPatches()

	var handle HandleSubscriptionMessage = nil
	bootNodeMultiaddr := "/dns4/bootnode-0.testnet.iotex.one/tcp/4689/ipfs/12D3KooWFnaTYuLo8Mkbm3wzaWHtUuaxBRe24Uiopu15Wr5EhD3o"
	iotexChainID := 2

	ctrl := gomock.NewController(t)
	host := mock.NewMockHost(ctrl)

	t.Run("NewP2pHostFailed", func(t *testing.T) {
		PatchConvey("NewP2pHostFailed", t, func() {
			Mock(libp2p.New).Return(nil, errors.New(t.Name())).Build()
			_, err := NewPubSubs(handle, bootNodeMultiaddr, iotexChainID)
			So(err.Error(), ShouldContainSubstring, t.Name())
		})
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

func TestAdd(t *testing.T) {
	projectID := uint64(0x1)
	p := &PubSubs{pubSubs: make(map[uint64]*pubSub)}

	t.Run("NewPubSubFailed", func(t *testing.T) {
		PatchConvey("NewPubSubFailed", t, func() {
			Mock(newPubSub).Return(&pubSub{}, errors.New(t.Name())).Build()
			err := p.Add(projectID)
			So(err.Error(), ShouldEqual, t.Name())
		})
	})

	t.Run("AddOk", func(t *testing.T) {
		PatchConvey("AddOk", t, func() {
			Mock(newPubSub).Return(&pubSub{}, nil).Build()
			Mock((*pubSub).run).Return().Build()
			err := p.Add(projectID)
			So(err, ShouldBeEmpty)
		})
	})

	t.Run("AddRepeat", func(t *testing.T) {
		PatchConvey("AddOk", t, func() {
			err := p.Add(projectID)
			So(err, ShouldBeEmpty)
		})
	})
}

func TestDelete(t *testing.T) {
	projectID := uint64(0x1)
	p := &PubSubs{pubSubs: make(map[uint64]*pubSub)}

	t.Run("IDNotExist", func(t *testing.T) {
		p.Delete(projectID)
	})

	t.Run("DeleteOk", func(t *testing.T) {
		PatchConvey("DeleteOk", t, func() {
			Mock(newPubSub).Return(&pubSub{}, nil).Build()
			Mock((*pubSub).run).Return().Build()
			err := p.Add(projectID)
			So(err, ShouldBeEmpty)

			Mock((*pubSub).release).Return().Build()
			p.Delete(projectID)
		})
	})
}

func TestPublish(t *testing.T) {
	require := require.New(t)
	patches := NewPatches()

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
	Mock((*pubSub).run).Return().Build()
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

	t.Run("PublishDataOk", func(t *testing.T) {
		patches = pubsubTopicPublish(patches, nil)
		err := p.Publish(projectID, d)
		require.NoError(err)
	})
}

func TestRelease(t *testing.T) {
	_, cancel := context.WithCancel(context.Background())
	p := &pubSub{
		topic:     &pubsub.Topic{},
		ctxCancel: cancel,
	}

	t.Run("TopicCloseFailed", func(t *testing.T) {
		PatchConvey("TopicCloseFailed", t, func() {
			Mock((*pubsub.Subscription).Cancel).Return().Build()
			Mock((*pubsub.Topic).Close).Return(errors.New(t.Name())).Build()
			p.release()
		})
	})

	t.Run("TopicCloseOk", func(t *testing.T) {
		PatchConvey("TopicCloseOk", t, func() {
			Mock((*pubsub.Subscription).Cancel).Return().Build()
			Mock((*pubsub.Topic).Close).Return(nil).Build()
			p.release()
		})
	})
}

func TestNextMsg(t *testing.T) {
	p := &pubSub{
		selfID: peer.ID("test01"),
	}

	PatchConvey("GetP2pDataFailed", t, func() {
		Mock((*pubsub.Subscription).Next).Return(nil, errors.New(t.Name())).Build()
		mockerLog := Mock(slog.Error).Return().Build()
		mockerUnmarshal := Mock(json.Unmarshal).Return(nil).Build()
		p.nextMsg()
		So(mockerLog.Times(), ShouldEqual, 1)
		So(mockerUnmarshal.Times(), ShouldEqual, 0)
	})

	PatchConvey("GetP2pDataFromSelf", t, func() {
		Mock((*pubsub.Subscription).Next).Return(&pubsub.Message{ReceivedFrom: p.selfID}, nil).Build()
		mockerLog := Mock(slog.Error).Return().Build()
		mockerUnmarshal := Mock(json.Unmarshal).Return(nil).Build()
		p.nextMsg()
		So(mockerLog.Times(), ShouldEqual, 0)
		So(mockerUnmarshal.Times(), ShouldEqual, 0)
	})

	PatchConvey("UnmarshalP2pDataFailed", t, func() {
		Mock((*pubsub.Subscription).Next).Return(&pubsub.Message{
			ReceivedFrom: peer.ID("test02"),
			Message:      &pubsub_pb.Message{Data: nil},
		}, nil).Build()
		Mock(json.Unmarshal).Return(errors.New(t.Name())).Build()
		mockerLog := Mock(slog.Error).Return().Build()
		p.nextMsg()
		So(mockerLog.Times(), ShouldEqual, 1)
	})
}

func TestNewPubSub(t *testing.T) {

	//t.Run("JoinTopicFailed", func(t *testing.T) {
	//	PatchConvey("JoinTopicFailed", t, func() {
	//		Mock((*pubsub.PubSub).Join).Return(nil, errors.New(t.Name())).Build()
	//		_, err := newPubSub(uint64(0x1), &pubsub.PubSub{}, nil, peer.ID("0"))
	//		So(err.Error(), ShouldContainSubstring, t.Name())
	//	})
	//})
	//
	//t.Run("TopicSubscriptionFailed", func(t *testing.T) {
	//	PatchConvey("TopicSubscriptionFailed", t, func() {
	//		Mock((*pubsub.PubSub).Join).Return(&pubsub.Topic{}, nil).Build()
	//		Mock((*pubsub.Topic).Subscribe).Return(nil, errors.New(t.Name())).Build()
	//		_, err := newPubSub(uint64(0x1), &pubsub.PubSub{}, nil, peer.ID("0"))
	//		So(err.Error(), ShouldContainSubstring, t.Name())
	//	})
	//})

	t.Run("NewPubSubOk", func(t *testing.T) {
		PatchConvey("NewPubSubOk", t, func() {
			Mock((*pubsub.PubSub).Join).Return(&pubsub.Topic{}, nil).Build()
			Mock((*pubsub.Topic).Subscribe).Return(&pubsub.Subscription{}, nil).Build()
			_, err := newPubSub(uint64(0x1), nil, nil, peer.ID("0"))
			So(err, ShouldBeEmpty)
		})
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
