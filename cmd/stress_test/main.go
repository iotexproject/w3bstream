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
)

var (
	projectMinterPrivateKey string
)

func init() {
	flag.StringVar(&projectMinterPrivateKey, "projectMinterPrivateKey", "1232997017ad09d7e3af812a22d61b46ec0d640a7d45cc26908960e9471f7024", "project minter private key")
}

func main() {
	flag.Parse()

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})))

	sk, err := crypto.HexToECDSA(projectMinterPrivateKey)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed parse project minter private key"))
	}

	slog.Info("project minter public key", "public_key", hexutil.Encode(crypto.FromECDSAPub(&sk.PublicKey)))

	// TODO test logic

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
