package p2p

import (
	"context"
	"reflect"
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	"github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/machinefi/sprout/testutil"
)

func TestNewPubSubs(t *testing.T) {
	require := require.New(t)
	patches := NewPatches()

	var handle HandleSubscriptionMessage = nil
	bootNodeMultiaddr := "/dns4/bootnode-0.testnet.iotex.one/tcp/4689/ipfs/12D3KooWFnaTYuLo8Mkbm3wzaWHtUuaxBRe24Uiopu15Wr5EhD3o"
	iotexChainID := 2

	t.Run("NewP2pHostFailed", func(t *testing.T) {
		patches = libp2pNew(patches, errors.New(t.Name()))
		_, err := NewPubSubs(handle, bootNodeMultiaddr, iotexChainID)
		require.ErrorContains(err, t.Name())
	})
	patches = libp2pNew(patches, nil)

	t.Run("NewGossipFailed", func(t *testing.T) {
		patches = pubsubNewGossipSub(patches, errors.New(t.Name()))
		_, err := NewPubSubs(handle, bootNodeMultiaddr, iotexChainID)
		require.ErrorContains(err, t.Name())
	})
	patches = pubsubNewGossipSub(patches, nil)

	t.Run("DiscoveryFailed", func(t *testing.T) {
		patches = p2pDiscoverPeers(patches, errors.New(t.Name()))
		_, err := NewPubSubs(handle, bootNodeMultiaddr, iotexChainID)
		require.ErrorContains(err, t.Name())
	})
	patches = p2pDiscoverPeers(patches, nil)

	t.Run("NewPubSubs", func(t *testing.T) {
		_, err := NewPubSubs(handle, bootNodeMultiaddr, iotexChainID)
		require.NoError(err)
	})
}

func TestAdd(t *testing.T) {
	require := require.New(t)
	patches := NewPatches()

	projectID := uint64(0x1)
	p := &PubSubs{pubSubs: make(map[uint64]*pubSub)}

	t.Run("NewPubSubFailed", func(t *testing.T) {
		patches = p2pNewPubSub(patches, &pubSub{}, errors.New(t.Name()))
		err := p.Add(projectID)
		require.ErrorContains(err, t.Name())
	})
	patches = p2pNewPubSub(patches, &pubSub{}, nil)

	t.Run("Add", func(t *testing.T) {
		err := p.Add(projectID)
		require.NoError(err)
	})

	t.Run("AddRepeat", func(t *testing.T) {
		err := p.Add(projectID)
		require.NoError(err)
	})

}

func TestDelete(t *testing.T) {
	//require := require.New(t)
	//patches := NewPatches()

	projectID := uint64(0x1)
	p := &PubSubs{pubSubs: make(map[uint64]*pubSub)}

	t.Run("IDNotExist", func(t *testing.T) {
		p.Delete(projectID)
	})

	//t.Run("Delete", func(t *testing.T) {
	//	patches = p2pNewPubSub(patches, &pubSub{}, nil)
	//	patches = pubSubRun(patches)
	//	err := p.Add(projectID)
	//	require.NoError(err)
	//	patches = pubSubRelease(patches)
	//	p.Delete(projectID)
	//})
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
	//patches = pubSubRun(patches)
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
}

func libp2pNew(p *Patches, err error) *Patches {
	return p.ApplyFunc(
		libp2p.New,
		func(opts ...libp2p.Option) (host.Host, error) {
			return nil, err
		},
	)
}

func pubsubNewGossipSub(p *Patches, err error) *Patches {
	return p.ApplyFunc(
		pubsub.NewGossipSub,
		func(ctx context.Context, h host.Host, opts ...pubsub.Option) (*pubsub.PubSub, error) {
			return nil, err
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

func pubSubRelease(p *Patches) *Patches {
	var ps *pubSub
	return p.ApplyMethodFunc(
		reflect.TypeOf(ps),
		"Release",
		func() {
		},
	)
}

func pubSubRun(p *Patches) *Patches {
	var ps *pubSub
	return p.ApplyMethodFunc(
		reflect.TypeOf(ps),
		"Run",
		func() {
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
