package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/cmd/apinode/api"
	"github.com/iotexproject/w3bstream/cmd/apinode/config"
	"github.com/iotexproject/w3bstream/cmd/apinode/persistence"
	"github.com/iotexproject/w3bstream/monitor"
	"github.com/iotexproject/w3bstream/p2p"
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

	pubSub, err := p2p.NewPubSub(cfg.BootNodeMultiAddr, cfg.IoTeXChainID, nil)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new p2p pubsub"))
	}

	p, err := persistence.NewPersistence(cfg.DatabaseDSN)
	if err != nil {
		log.Fatal(err)
	}

	if err := monitor.Run(
		&monitor.Handler{
			ScannedBlockNumber:       p.ScannedBlockNumber,
			UpsertScannedBlockNumber: p.UpsertScannedBlockNumber,
			AssignTask:               p.UpsertAssignedTask,
			SettleTask:               p.UpsertSettledTask,
		},
		&monitor.ContractAddr{
			TaskManager: common.HexToAddress(cfg.TaskManagerContractAddr),
		},
		cfg.BeginningBlockNumber,
		cfg.ChainEndpoint,
	); err != nil {
		log.Fatal(errors.Wrap(err, "failed to run contract monitor"))
	}

	go func() {
		if err := api.Run(p, prv, pubSub, cfg.AggregationAmount, cfg.ServiceEndpoint, cfg.ProverServiceEndpoint); err != nil {
			log.Fatal(err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
