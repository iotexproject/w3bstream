package p2p

import (
	"context"
	"log/slog"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
)

type discoveryNotifee struct {
	h host.Host
}

func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	slog.Info("discovered new peer", "id", pi.ID)
	if err := n.h.Connect(context.Background(), pi); err != nil {
		slog.Error("connecting to peer failed", "id", pi.ID, "error", err)
	}
}

// SetupDiscovery creates an mDNS discovery service and attaches it to the libp2p Host.
// This lets us automatically discover peers on the same LAN and connect to them.
func SetupDiscovery(h host.Host) error {
	s := mdns.NewMdnsService(h, "w3bstream-mdns", &discoveryNotifee{h: h})
	return s.Start()
}
