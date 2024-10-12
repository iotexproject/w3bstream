package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/cmd/sequencer/api"
	"github.com/iotexproject/w3bstream/cmd/sequencer/config"
	"github.com/iotexproject/w3bstream/cmd/sequencer/db"
	"github.com/iotexproject/w3bstream/monitor"
	"github.com/iotexproject/w3bstream/p2p"
	"github.com/iotexproject/w3bstream/task/assigner"
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

	if err := monitor.Run(
		&monitor.Handler{
			ScannedBlockNumber:       db.ScannedBlockNumber,
			UpsertScannedBlockNumber: db.UpsertScannedBlockNumber,
			UpsertNBits:              db.UpsertNBits,
			UpsertBlockHead:          db.UpsertBlockHead,
			UpsertProver:             db.UpsertProver,
			SettleTask:               db.DeleteTask,
		},
		&monitor.ContractAddr{
			Prover:      common.HexToAddress(cfg.ProverContractAddr),
			Dao:         common.HexToAddress(cfg.DaoContractAddr),
			Minter:      common.HexToAddress(cfg.MinterContractAddr),
			TaskManager: common.HexToAddress(cfg.TaskManagerContractAddr),
		},
		cfg.BeginningBlockNumber,
		cfg.ChainEndpoint,
	); err != nil {
		log.Fatal(errors.Wrap(err, "failed to run contract monitor"))
	}

	if _, err := p2p.NewPubSub(cfg.BootNodeMultiAddr, cfg.IoTeXChainID, db.CreateTask); err != nil {
		log.Fatal(errors.Wrap(err, "failed to new p2p pubsub"))
	}

	if err := assigner.Run(db, prv, cfg.ChainEndpoint, common.HexToAddress(cfg.MinterContractAddr)); err != nil {
		log.Fatal(errors.Wrap(err, "failed to run task assigner"))
	}

	go func() {
		if err := api.Run(db, cfg, prv); err != nil {
			log.Fatal(errors.Wrap(err, "failed to run http server"))
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
