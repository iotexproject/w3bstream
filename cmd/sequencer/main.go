package main

import (
	"flag"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/clients"
	"github.com/machinefi/sprout/cmd/sequencer/api"
	"github.com/machinefi/sprout/cmd/sequencer/persistence"
)

var (
	logLevel           int
	aggregationAmount  uint
	address            string
	coordinatorAddress string
	databaseDSN        string
	didAuthServer      string
	privateKey         string
	did                bool
)

func init() {
	flag.IntVar(&logLevel, "logLevel", int(slog.LevelDebug), "golang slog level")
	flag.UintVar(&aggregationAmount, "aggregationAmount", 1, "the amount for pack how many messages into one task")
	flag.StringVar(&address, "address", ":9000", "http listen address")
	flag.StringVar(&coordinatorAddress, "coordinatorAddress", "localhost:9001", "coordinator address")
	flag.StringVar(&databaseDSN, "databaseDSN", "postgres://test_user:test_passwd@localhost:5432/test?sslmode=disable", "database dsn")
	flag.StringVar(&didAuthServer, "didAuthServer", "srv-did-vc:9999", "did auth server endpoint")
	flag.StringVar(&privateKey, "privateKey", "dbfe03b0406549232b8dccc04be8224fcc0afa300a33d4f335dcfdfead861c85", "sequencer private key")
	flag.BoolVar(&did, "did", true, "did flag")
}

func main() {
	flag.Parse()

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.Level(logLevel)})))

	sk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed parse private key"))
	}

	slog.Info("sequencer public key", "public_key", hexutil.Encode(crypto.FromECDSAPub(&sk.PublicKey)))

	_ = clients.NewManager()

	p, err := persistence.NewPersistence(databaseDSN)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := api.NewHttpServer(p, aggregationAmount, coordinatorAddress, didAuthServer, sk, did).Run(address); err != nil {
			log.Fatal(err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
