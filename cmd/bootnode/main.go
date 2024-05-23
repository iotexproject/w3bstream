package main

import (
	"context"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/p2p/muxer/yamux"
	"github.com/pkg/errors"
)

var (
	privateKey   string
	ioTeXChainID int
)

func init() {
	flag.StringVar(&privateKey, "privateKey", "975da455cedd5c64a151594358e66b445d703d08cd560ff452cd8c9f02d94b1284c6f964e8e4497876c811f8590d4a4f7e9fac9f77252d37f5cb8ac5478ea5a7", "bootnode private key")
	flag.IntVar(&ioTeXChainID, "ioTeXChainID", 2, "iotex chain id")
}

func main() {
	flag.Parse()

	priBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to decode private key"))
	}
	priKey, err := crypto.UnmarshalEd25519PrivateKey(priBytes)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to unmarshal private key"))
	}
	h, err := libp2p.New(libp2p.Identity(priKey), libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/8000"), libp2p.Muxer("/yamux/2.0.0", yamux.DefaultTransport))
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to create libp2p host"))
	}

	for _, a := range h.Addrs() {
		slog.Info(fmt.Sprintf("listening on %s/p2p/%s", a, h.ID().String()))
	}

	ctx := context.Background()
	dht, err := dht.New(ctx, h, dht.ProtocolPrefix(protocol.ID("/iotex"+strconv.Itoa(ioTeXChainID))), dht.Mode(dht.ModeServer))
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new dht"))
	}
	if err = dht.Bootstrap(ctx); err != nil {
		log.Fatal(errors.Wrap(err, "failed to bootstrap dht"))
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
