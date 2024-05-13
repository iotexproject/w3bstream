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
	"github.com/machinefi/ioconnect-go/pkg/ioconnect"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/clients"
	"github.com/machinefi/sprout/cmd/sequencer/api"
	"github.com/machinefi/sprout/cmd/sequencer/persistence"
)

var (
	logLevel             int
	aggregationAmount    uint
	address              string
	coordinatorAddress   string
	databaseDSN          string
	didAuthServer        string
	privateKey           string
	jwkSecret            string
	jwk                  *ioconnect.JWK
	ioIDContractAddress  string
	chainEndpoint        string
	ioIDRegistryEndpoint string
)

func init() {
	flag.IntVar(&logLevel, "logLevel", int(slog.LevelDebug), "golang slog level")
	flag.UintVar(&aggregationAmount, "aggregationAmount", 1, "the amount for pack how many messages into one task")
	flag.StringVar(&address, "address", ":9000", "http listen address")
	flag.StringVar(&coordinatorAddress, "coordinatorAddress", "localhost:9001", "coordinator address")
	flag.StringVar(&databaseDSN, "databaseDSN", "postgres://test_user:test_passwd@localhost:5432/test?sslmode=disable", "database dsn")
	flag.StringVar(&didAuthServer, "didAuthServer", "srv-did-vc:9999", "did auth server endpoint")
	flag.StringVar(&privateKey, "privateKey", "dbfe03b0406549232b8dccc04be8224fcc0afa300a33d4f335dcfdfead861c85", "sequencer private key")
	flag.StringVar(&jwkSecret, "jwkSecret", "R3QNJihYLjtcaxALSTsKe1cYSX0pS28wZitFVXE4Y2klf2hxVCczYHw2dVg4fXJdSgdCcnM4PgV1aTo9DwYqEw==", "jwk secret base64 string")
	flag.StringVar(&ioIDContractAddress, "ioIDContract", "0xB63e6034361283dc8516480a46EcB9C122c983Bb", "ioIDRegistry contract address")
	flag.StringVar(&chainEndpoint, "chainEndpoint", "http://iotex.chainendpoint.io", "chain endpoint")
	flag.StringVar(&ioIDRegistryEndpoint, "ioIDRegistryEndpoint", "did.iotex.me", "ioID registry endpoint")

	// initialize jwk context from secrets
	if jwkSecret != "" {
		var (
			secrets = ioconnect.JWKSecrets{}
			err     error
		)
		if err = secrets.UnmarshalText([]byte(jwkSecret)); err != nil {
			panic(errors.Wrap(err, "invalid jwk secrets from flag"))
		}
		if jwk, err = ioconnect.NewJWKBySecret(secrets); err != nil {
			panic(errors.Wrap(err, "failed to new jwk from secrets"))
		}
		return
	}
}

func main() {
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: slog.Level(logLevel)}))
	slog.SetDefault(logger)

	sk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed parse private key"))
	}

	slog.Info("sequencer public key", "public_key", hexutil.Encode(crypto.FromECDSAPub(&sk.PublicKey)))

	_, err = clients.NewManager(ioIDContractAddress, chainEndpoint, ioIDRegistryEndpoint)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new clients manager"))
	}

	p, err := persistence.NewPersistence(databaseDSN)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := api.NewHttpServer(p, aggregationAmount, coordinatorAddress, didAuthServer, sk, jwk).Run(address); err != nil {
			log.Fatal(err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
