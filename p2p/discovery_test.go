package p2p

import (
	"context"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/routing"
	"github.com/libp2p/go-libp2p/p2p/discovery/util"
	"github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/iotexproject/w3bstream/testutil/mock"
)

func TestDiscoverPeers(t *testing.T) {
	r := require.New(t)
	p := gomonkey.NewPatches()
	defer p.Reset()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	host := mock.NewMockHost(ctrl)

	ctx := context.Background()
	iotexChainID := 2
	addr, err := multiaddr.NewMultiaddr("/dns4/bootnode-0.testnet.iotex.one/tcp/4689/ipfs/12D3KooWFnaTYuLo8Mkbm3wzaWHtUuaxBRe24Uiopu15Wr5EhD3o")
	r.NoError(err)

	t.Run("FailedToNewDht", func(t *testing.T) {
		p = p.ApplyFuncReturn(dht.New, nil, errors.New(t.Name()))

		err := discoverPeers(ctx, nil, "", iotexChainID)
		r.ErrorContains(err, t.Name())
	})
	p = p.ApplyFuncReturn(dht.New, &dht.IpfsDHT{}, nil)

	t.Run("FailedToBootstrapDht", func(t *testing.T) {
		p = p.ApplyMethodReturn(&dht.IpfsDHT{}, "Bootstrap", errors.New(t.Name()))

		err := discoverPeers(ctx, nil, "", iotexChainID)
		r.ErrorContains(err, t.Name())
	})
	p = p.ApplyMethodReturn(&dht.IpfsDHT{}, "Bootstrap", nil)

	t.Run("FailedToParseBootNodeMultiaddr", func(t *testing.T) {
		p = p.ApplyFuncReturn(multiaddr.NewMultiaddr, nil, errors.New(t.Name()))

		err := discoverPeers(ctx, nil, "", iotexChainID)
		r.ErrorContains(err, t.Name())
	})
	p = p.ApplyFuncReturn(multiaddr.NewMultiaddr, addr, nil)

	t.Run("FailedToGetBootnode", func(t *testing.T) {
		p = p.ApplyFuncReturn(peer.AddrInfoFromP2pAddr, nil, errors.New(t.Name()))

		err := discoverPeers(ctx, nil, "", iotexChainID)
		r.ErrorContains(err, t.Name())
	})
	p = p.ApplyFuncReturn(peer.AddrInfoFromP2pAddr, &peer.AddrInfo{}, nil)

	t.Run("FailedToConnectBootNode", func(t *testing.T) {
		host.EXPECT().Connect(gomock.Any(), gomock.Any()).Return(errors.New(t.Name())).Times(1)
		err := discoverPeers(ctx, host, "", iotexChainID)
		r.ErrorContains(err, t.Name())
	})
	t.Run("Success", func(t *testing.T) {
		host.EXPECT().Connect(gomock.Any(), gomock.Any()).Return(nil).Times(1)
		p = p.ApplyFuncReturn(routing.NewRoutingDiscovery, nil)
		p = p.ApplyFuncReturn(util.Advertise)

		err := discoverPeers(ctx, host, "", iotexChainID)
		r.NoError(err)
	})
}
