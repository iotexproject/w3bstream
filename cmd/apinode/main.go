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

	"github.com/iotexproject/w3bstream/cmd/apinode/api"
	"github.com/iotexproject/w3bstream/cmd/apinode/config"
	"github.com/iotexproject/w3bstream/cmd/apinode/persistence"
)

func main() {
	conf, err := config.Get()
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to get config"))
	}
	conf.Print()
	slog.Info("apinode config loaded")

	priKey, err := crypto.HexToECDSA(conf.PriKey)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed parse private key"))
	}

	slog.Info("sequencer public key", "public_key", hexutil.Encode(crypto.FromECDSAPub(&priKey.PublicKey)))

	p, err := persistence.NewPersistence(conf.DatabaseDSN)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := api.NewHttpServer(p, conf.AggregationAmount, priKey).Run(conf.ServiceEndpoint); err != nil {
			log.Fatal(err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
