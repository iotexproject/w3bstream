package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/iotexproject/w3bstream/cmd/sequencer/api"
	"github.com/iotexproject/w3bstream/cmd/sequencer/config"
	"github.com/iotexproject/w3bstream/cmd/sequencer/db"
	"github.com/iotexproject/w3bstream/monitor"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to get config"))
	}
	cfg.Print()
	slog.Info("sequencer config loaded")

	sqliteDB, err := gorm.Open(sqlite.Open(cfg.LocalDBDir), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to connect sqlite"))
	}
	db, err := db.New(sqliteDB)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new db"))
	}

	if err := monitor.Run(
		&monitor.Handler{
			ScannedBlockNumber:       db.ScannedBlockNumber,
			UpsertScannedBlockNumber: db.UpsertScannedBlockNumber,
			UpsertNBits:              db.UpsertNBits,
			UpsertBlockHead:          db.UpsertBlockHead,
			DeleteTask:               db.DeleteTask,
		},
		&monitor.ContractAddr{
			Prover:      common.HexToAddress(cfg.ProverContractAddr),
			Project:     common.HexToAddress(cfg.ProjectContractAddr),
			Dao:         common.HexToAddress(cfg.DaoContractAddr),
			Minter:      common.HexToAddress(cfg.MinterContractAddr),
			TaskManager: common.HexToAddress(cfg.TaskManagerContractAddr),
		},
		cfg.BeginningBlockNumber,
		cfg.ChainEndpoint,
	); err != nil {
		log.Fatal(errors.Wrap(err, "failed to run contract monitor"))
	}

	go func() {
		if err := api.Run(db, cfg); err != nil {
			log.Fatal(errors.Wrap(err, "failed to run http server"))
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
