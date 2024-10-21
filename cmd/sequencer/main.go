package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/cmd/sequencer/config"
	"github.com/iotexproject/w3bstream/cmd/sequencer/db"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to get config"))
	}
	cfg.Print()
	slog.Info("sequencer config loaded")

	prv, err := crypto.HexToECDSA(cfg.OperatorPrvKey)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to parse private key"))
	}

	db, err := db.New(cfg.LocalDBDir)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new db"))
	}

	s := NewSequencer(cfg, db, prv)
	if err := s.Start(); err != nil {
		log.Fatal(errors.Wrap(err, "failed to start sequencer"))
	}
	defer func() {
		if err := s.Stop(); err != nil {
			log.Fatal(errors.Wrap(err, "failed	to stop sequencer"))
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
