package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/cmd/apinode/config"
	"github.com/iotexproject/w3bstream/cmd/apinode/persistence"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to get config"))
	}
	cfg.Print()
	slog.Info("apinode config loaded")

	prv, err := crypto.HexToECDSA(cfg.PrvKey)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to parse private key"))
	}

	slog.Info("sequencer public key", "public_key", hexutil.Encode(crypto.FromECDSAPub(&prv.PublicKey)))

	p, err := persistence.NewPersistence(cfg.DatabaseDSN)
	if err != nil {
		log.Fatal(err)
	}

	apinode := NewAPINode(cfg, p, prv)

	if err := apinode.Start(); err != nil {
		log.Fatal(err)
	}
	defer apinode.Stop()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
