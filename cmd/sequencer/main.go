package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/cockroachdb/pebble"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/cmd/sequencer/api"
	"github.com/iotexproject/w3bstream/cmd/sequencer/config"
	"github.com/iotexproject/w3bstream/persistence/contract"
	"github.com/iotexproject/w3bstream/persistence/postgres"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to get config"))
	}
	cfg.Print()
	slog.Info("coordinator config loaded")

	// defaultDatasourcePubKey, err := hexutil.Decode(cfg.DefaultDatasourcePubKey)
	// if err != nil {
	// 	log.Fatal(errors.Wrap(err, "failed to decode default datasource public key"))
	// }

	persistence, err := postgres.New(cfg.DatabaseDSN)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new postgres persistence"))
	}

	schedulerNotification := make(chan uint64, 10)
	dispatcherNotification := make(chan uint64, 10)
	chainHeadNotification := make(chan uint64, 10)

	projectNotifications := []chan<- uint64{dispatcherNotification, schedulerNotification}
	chainHeadNotifications := []chan<- uint64{chainHeadNotification}

	local := cfg.ProjectFileDir != ""

	var kvDB *pebble.DB
	if !local {
		kvDB, err = pebble.Open(cfg.LocalDBDir, &pebble.Options{})
		if err != nil {
			log.Fatal(errors.Wrap(err, "failed to open pebble db"))
		}
		defer kvDB.Close()

		_, err = contract.New(kvDB, persistence, cfg.SchedulerEpoch, cfg.BeginningBlockNumber,
			cfg.ChainEndpoint, common.HexToAddress(cfg.MinterContractAddr),
			common.HexToAddress(cfg.DaoContractAddr), chainHeadNotifications, projectNotifications)
		if err != nil {
			log.Fatal(errors.Wrap(err, "failed to new contract persistence"))
		}
	}

	// var projectManager *project.Manager
	// if local {
	// 	projectManager, err = project.NewLocalManager(cfg.ProjectFileDir)
	// } else {
	// 	projectManager = project.NewManager(kvDB, contractPersistence.LatestProject)
	// }
	// if err != nil {
	// 	log.Fatal(errors.Wrap(err, "failed to new project manager"))
	// }

	// datasourcePG := datasource.NewPostgres()
	// var taskDispatcher *dispatcher.Dispatcher
	// if local {
	// 	taskDispatcher, err = dispatcher.NewLocal(persistence, datasourcePG.New, projectManager, cfg.DefaultDatasourceURI,
	// 		cfg.OperatorPriKey, cfg.OperatorPriKeyED25519, cfg.BootNodeMultiAddr, cfg.ContractWhitelist, defaultDatasourcePubKey, cfg.IoTeXChainID)
	// } else {
	// 	projectOffsets := scheduler.NewProjectEpochOffsets(cfg.SchedulerEpoch, contractPersistence.LatestProjects, schedulerNotification)

	// 	taskDispatcher, err = dispatcher.New(persistence, datasourcePG.New, projectManager, cfg.DefaultDatasourceURI, cfg.BootNodeMultiAddr,
	// 		cfg.OperatorPriKey, cfg.OperatorPriKeyED25519, cfg.ContractWhitelist, defaultDatasourcePubKey, cfg.IoTeXChainID,
	// 		dispatcherNotification, chainHeadNotification, contractPersistence, projectOffsets)
	// }
	// if err != nil {
	// 	log.Fatal(errors.Wrap(err, "failed to new dispatcher"))
	// }
	// taskDispatcher.Run()

	go func() {
		if err := api.NewHttpServer(persistence, cfg).Run(cfg.ServiceEndpoint); err != nil {
			log.Fatal(errors.Wrap(err, "failed to run http server"))
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
