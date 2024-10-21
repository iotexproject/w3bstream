package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/cmd/prover/config"
	"github.com/iotexproject/w3bstream/cmd/prover/db"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to get config"))
	}
	cfg.Print()
	slog.Info("prover config loaded")

	prv, err := crypto.HexToECDSA(cfg.ProverOperatorPrvKey)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to parse prover private key"))
	}
	proverOperatorAddress := crypto.PubkeyToAddress(prv.PublicKey)
	slog.Info("my prover", "address", proverOperatorAddress.String())

	db, err := db.New(cfg.LocalDBDir, proverOperatorAddress)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new db"))
	}

	prover := NewProver(cfg, prv, db)
	if err := prover.Start(); err != nil {
		log.Fatal(errors.Wrap(err, "failed to start prover"))
	}
	defer func() {
		if err := prover.Stop(); err != nil {
			log.Fatal(errors.Wrap(err, "failed to stop prover"))
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
