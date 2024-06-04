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
	logLevel                        int
	aggregationAmount               uint
	address                         string
	coordinatorAddr                 string
	databaseDSN                     string
	privateKey                      string
	jwkSecret                       string
	jwk                             *ioconnect.JWK
	ioIDRegistryEndpoint            string
	ioIDRegistryContractAddress     string
	projectClientContractAddress    string
	w3bstreamProjectContractAddress string
	chainEndpoint                   string
)

func init() {
	flag.IntVar(&logLevel, "logLevel", int(slog.LevelDebug), "golang slog level")
	flag.UintVar(&aggregationAmount, "aggregationAmount", 1, "the amount for pack how many messages into one task")
	flag.StringVar(&address, "address", ":9000", "http listen address")
	flag.StringVar(&coordinatorAddr, "coordinatorAddress", "localhost:9001", "coordinator address")
	flag.StringVar(&databaseDSN, "databaseDSN", "postgres://test_user:test_passwd@localhost:5432/test?sslmode=disable", "database dsn")
	flag.StringVar(&privateKey, "privateKey", "dbfe03b0406549232b8dccc04be8224fcc0afa300a33d4f335dcfdfead861c85", "sequencer private key")
	flag.StringVar(&jwkSecret, "jwkSecret", "R3QNJihYLjtcaxALSTsKe1cYSX0pS28wZitFVXE4Y2klf2hxVCczYHw2dVg4fXJdSgdCcnM4PgV1aTo9DwYqEw==", "jwk secret base64 string")
	flag.StringVar(&ioIDRegistryContractAddress, "ioIDRegistryContract", "0x06b3Fcda51e01EE96e8E8873F0302381c955Fddd", "ioIDRegistry contract address")
	flag.StringVar(&projectClientContractAddress, "projectClientContract", "0xF4d6282C5dDD474663eF9e70c927c0d4926d1CEb", "projectClient contract address")
	flag.StringVar(&w3bstreamProjectContractAddress, "w3bstreamProjectContract", "0x6AfCB0EB71B7246A68Bb9c0bFbe5cD7c11c4839f", "w3bstream project contract address")
	flag.StringVar(&chainEndpoint, "chainEndpoint", "https://babel-api.testnet.iotex.io", "chain endpoint")
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

	clientMgr, err := clients.NewManager(projectClientContractAddress, ioIDRegistryContractAddress, w3bstreamProjectContractAddress, ioIDRegistryEndpoint, chainEndpoint)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new clients manager"))
	}

	p, err := persistence.NewPersistence(databaseDSN)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := api.NewHttpServer(p, aggregationAmount, coordinatorAddr, sk, jwk, clientMgr).Run(address); err != nil {
			log.Fatal(err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
