package p2p

import (
	"context"
	"reflect"
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestDiscoverPeers(t *testing.T) {
	require := require.New(t)
	patches := NewPatches()

	ctx := context.Background()
	bootNodeMultiaddr := "/dns4/bootnode-0.testnet.iotex.one/tcp/4689/ipfs/12D3KooWFnaTYuLo8Mkbm3wzaWHtUuaxBRe24Uiopu15Wr5EhD3o"
	iotexChainID := 2

	t.Run("NewDhtFailed", func(t *testing.T) {
		patches = newDht(patches, errors.New(t.Name()))
		err := discoverPeers(ctx, nil, bootNodeMultiaddr, iotexChainID)
		require.ErrorContains(err, t.Name())
	})
	patches = newDht(patches, nil)

	t.Run("DhtBootstrapFailed", func(t *testing.T) {
		patches = ipfsDHTBootstrap(patches, errors.New(t.Name()))
		err := discoverPeers(ctx, nil, bootNodeMultiaddr, iotexChainID)
		require.ErrorContains(err, t.Name())
	})
	patches = ipfsDHTBootstrap(patches, nil)

	t.Run("ParseBootNodeMultiaddrFailed", func(t *testing.T) {
		patches = multiaddrNewMultiaddr(patches, errors.New(t.Name()))
		err := discoverPeers(ctx, nil, bootNodeMultiaddr, iotexChainID)
		require.ErrorContains(err, t.Name())
	})
	patches = multiaddrNewMultiaddr(patches, nil)

	t.Run("GetBootnodeFailed", func(t *testing.T) {
		patches = peerAddrInfoFromP2pAddr(patches, errors.New(t.Name()))
		err := discoverPeers(ctx, nil, bootNodeMultiaddr, iotexChainID)
		require.ErrorContains(err, t.Name())
	})
	patches = peerAddrInfoFromP2pAddr(patches, nil)

	//t.Run("ConnectBootnodeFailed", func(t *testing.T) {
	//	patches = hostConnect(patches, errors.New(t.Name()))
	//	err := discoverPeers(ctx, nil, bootNodeMultiaddr, iotexChainID)
	//	require.ErrorContains(err, t.Name())
	//})

	//t.Run("DiscoverOK", func(t *testing.T) {
	//	err := discoverPeers(ctx, nil, bootNodeMultiaddr, iotexChainID)
	//	require.NoError(err)
	//})
}

func newDht(p *Patches, err error) *Patches {
	return p.ApplyFunc(
		dht.New,
		func(ctx context.Context, h host.Host, options ...dht.Option) (*dht.IpfsDHT, error) {
			return nil, err
		},
	)
}

func ipfsDHTBootstrap(p *Patches, err error) *Patches {
	ipfsDHT := &dht.IpfsDHT{}
	return p.ApplyMethodFunc(
		reflect.TypeOf(ipfsDHT),
		"Bootstrap",
		func(context.Context) error {
			return err
		},
	)
}

func multiaddrNewMultiaddr(p *Patches, err error) *Patches {
	return p.ApplyFunc(
		multiaddr.NewMultiaddr,
		func(s string) (multiaddr.Multiaddr, error) {
			return nil, err
		},
	)
}

func peerAddrInfoFromP2pAddr(p *Patches, err error) *Patches {
	return p.ApplyFunc(
		peer.AddrInfoFromP2pAddr,
		func(m multiaddr.Multiaddr) (*peer.AddrInfo, error) {
			return nil, err
		},
	)
}

func hostConnect(p *Patches, err error) *Patches {
	var h *host.Host
	return p.ApplyMethodFunc(
		reflect.TypeOf(h),
		"Connect",
		func(ctx context.Context, pi peer.AddrInfo) error {
			return err
		},
	)
}
