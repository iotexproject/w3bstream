package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"strconv"

	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/p2p/muxer/yamux"
	"github.com/pkg/errors"
)

type (
	BootNode struct {
		host host.Host
		dht  *dht.IpfsDHT

		config BootNodeConfig
	}

	BootNodeConfig struct {
		PrivateKey   crypto.PrivKey
		Port         int
		IoTeXChainID int
	}
)

func NewBootNode(config BootNodeConfig) *BootNode {
	h, err := libp2p.New(libp2p.Identity(config.PrivateKey), libp2p.ListenAddrStrings(
		fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", config.Port)), libp2p.Muxer("/yamux/2.0.0", yamux.DefaultTransport))
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to create libp2p host"))
	}

	for _, a := range h.Addrs() {
		slog.Info(fmt.Sprintf("listening on %s/p2p/%s", a, h.ID().String()))
	}

	ctx := context.Background()
	dht, err := dht.New(ctx, h, dht.ProtocolPrefix(protocol.ID("/iotex"+strconv.Itoa(config.IoTeXChainID))), dht.Mode(dht.ModeServer))
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new dht"))
	}

	return &BootNode{
		host:   h,
		dht:    dht,
		config: config,
	}
}

func (b *BootNode) Start() error {
	return b.dht.Bootstrap(context.Background())
}

func (b *BootNode) Stop() error {
	if err := b.dht.Close(); err != nil {
		return err
	}
	return b.host.Close()
}
