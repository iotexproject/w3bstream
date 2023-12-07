package p2p

import (
	"context"
	"strconv"

	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/p2p/discovery/routing"
	"github.com/libp2p/go-libp2p/p2p/discovery/util"
	"github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
)

func discoverPeers(ctx context.Context, h host.Host, bootNodeMultiaddr string, iotexChainID int) error {
	dht, err := dht.New(ctx, h, dht.ProtocolPrefix(protocol.ID("/iotex"+strconv.Itoa(iotexChainID))), dht.Mode(dht.ModeServer))
	if err != nil {
		return errors.Wrap(err, "new dht failed")
	}
	if err = dht.Bootstrap(ctx); err != nil {
		return errors.Wrap(err, "dht bootstrap failed")
	}

	bm, err := multiaddr.NewMultiaddr(bootNodeMultiaddr)
	if err != nil {
		return errors.Wrap(err, "parse boot node multiaddr failed")
	}
	bi, err := peer.AddrInfoFromP2pAddr(bm)
	if err != nil {
		return errors.Wrap(err, "get boot node multiaddr addr info failed")
	}
	if err := h.Connect(context.Background(), *bi); err != nil {
		return errors.Wrap(err, "connect boot node failed")
	}

	routingDiscovery := routing.NewRoutingDiscovery(dht)
	util.Advertise(ctx, routingDiscovery, "w3bstream")
	return nil
}
