package main

import (
	"encoding/hex"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/libp2p/go-libp2p/core/crypto"
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

	bootnode := NewBootNode(BootNodeConfig{
		PrivateKey:   priKey,
		Port:         8000,
		IoTeXChainID: ioTeXChainID,
	})

	if err := bootnode.Start(); err != nil {
		log.Fatal(errors.Wrap(err, "failed to start bootnode"))
	}
	defer func() {
		if err := bootnode.Stop(); err != nil {
			log.Fatal(errors.Wrap(err, "failed to stop bootnode"))
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
