package main

import (
	"encoding/json"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/cmd/prover/config"
	"github.com/iotexproject/w3bstream/cmd/prover/db"
	"github.com/iotexproject/w3bstream/datasource"
	"github.com/iotexproject/w3bstream/monitor"
	"github.com/iotexproject/w3bstream/project"
	"github.com/iotexproject/w3bstream/task/processor"
	"github.com/iotexproject/w3bstream/vm"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to get config"))
	}
	cfg.Print()
	slog.Info("prover config loaded")

	if err := migrateDatabase(cfg.DatasourceDSN); err != nil {
		log.Fatal(errors.Wrap(err, "failed to migrate database"))
	}

	prv, err := crypto.HexToECDSA(cfg.ProverOperatorPriKey)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to parse prover private key"))
	}
	proverOperatorAddress := crypto.PubkeyToAddress(prv.PublicKey)
	slog.Info("my prover", "address", proverOperatorAddress.String())

	vmEndpoints := map[uint64]string{}
	if err := json.Unmarshal([]byte(cfg.VMEndpoints), &vmEndpoints); err != nil {
		log.Fatal(errors.Wrap(err, "failed to unmarshal vm endpoints"))
	}

	vmHandler, err := vm.NewHandler(vmEndpoints)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new vm handler"))
	}

	db, err := db.New(cfg.LocalDBDir, proverOperatorAddress)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new db"))
	}

	if err := monitor.Run(
		&monitor.Handler{
			ScannedBlockNumber:       db.ScannedBlockNumber,
			UpsertScannedBlockNumber: db.UpsertScannedBlockNumber,
			AssignTask:               db.CreateTask,
			UpsertProject:            db.UpsertProject,
			DeleteTask:               db.DeleteTask,
		},
		&monitor.ContractAddr{
			Project:     common.HexToAddress(cfg.ProjectContractAddr),
			TaskManager: common.HexToAddress(cfg.TaskManagerContractAddr),
		},
		cfg.BeginningBlockNumber,
		cfg.ChainEndpoint,
	); err != nil {
		log.Fatal(errors.Wrap(err, "failed to run contract monitor"))
	}

	projectManager := project.NewManager(db.Project, db.ProjectFile, db.UpsertProjectFile)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new project manager"))
	}

	datasource, err := datasource.NewPostgres(cfg.DatasourceDSN)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to connect datasource"))
	}

	if err := processor.Run(vmHandler.Handle, projectManager.Project, db, datasource.Retrieve, prv, cfg.ChainEndpoint, common.HexToAddress(cfg.RouterContractAddr)); err != nil {
		log.Fatal(errors.Wrap(err, "failed to run task processor"))
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
