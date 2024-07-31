package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/cockroachdb/pebble"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/cmd/coordinator/api"
	"github.com/iotexproject/w3bstream/cmd/coordinator/config"
	"github.com/iotexproject/w3bstream/datasource"
	"github.com/iotexproject/w3bstream/persistence/contract"
	"github.com/iotexproject/w3bstream/persistence/postgres"
	"github.com/iotexproject/w3bstream/project"
	"github.com/iotexproject/w3bstream/scheduler"
	"github.com/iotexproject/w3bstream/task/dispatcher"
)

func main() {
	conf, err := config.Get()
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to get config"))
	}
	conf.Print()
	slog.Info("coordinator config loaded")

	defaultDatasourcePubKey, err := hexutil.Decode(conf.DefaultDatasourcePubKey)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to decode default datasource public key"))
	}

	persistence, err := postgres.New(conf.DatabaseDSN)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new postgres persistence"))
	}

	schedulerNotification := make(chan uint64, 10)
	dispatcherNotification := make(chan uint64, 10)
	chainHeadNotification := make(chan uint64, 10)

	projectNotifications := []chan<- uint64{dispatcherNotification, schedulerNotification}
	chainHeadNotifications := []chan<- uint64{chainHeadNotification}

	local := conf.ProjectFileDir != ""

	var contractPersistence *contract.Contract
	var kvDB *pebble.DB
	if !local {
		kvDB, err = pebble.Open(conf.LocalDBDir, &pebble.Options{})
		if err != nil {
			log.Fatal(errors.Wrap(err, "failed to open pebble db"))
		}
		defer kvDB.Close()

		contractPersistence, err = contract.New(kvDB, conf.SchedulerEpoch, conf.BeginningBlockNumber,
			conf.ChainEndpoint, common.HexToAddress(conf.ProverContractAddr),
			common.HexToAddress(conf.ProjectContractAddr), chainHeadNotifications, projectNotifications)
		if err != nil {
			log.Fatal(errors.Wrap(err, "failed to new contract persistence"))
		}
	}

	var projectManager *project.Manager
	if local {
		projectManager, err = project.NewLocalManager(conf.ProjectFileDir)
	} else {
		projectManager = project.NewManager(kvDB, contractPersistence.LatestProject)
	}
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new project manager"))
	}

	datasourcePG := datasource.NewPostgres()
	var taskDispatcher *dispatcher.Dispatcher
	if local {
		taskDispatcher, err = dispatcher.NewLocal(persistence, datasourcePG.New, projectManager, conf.DefaultDatasourceURI,
			conf.OperatorPriKey, conf.OperatorPriKeyED25519, conf.BootNodeMultiAddr, conf.ContractWhitelist, defaultDatasourcePubKey, conf.IoTeXChainID)
	} else {
		projectOffsets := scheduler.NewProjectEpochOffsets(conf.SchedulerEpoch, contractPersistence.LatestProjects, schedulerNotification)

		taskDispatcher, err = dispatcher.New(persistence, datasourcePG.New, projectManager, conf.DefaultDatasourceURI, conf.BootNodeMultiAddr,
			conf.OperatorPriKey, conf.OperatorPriKeyED25519, conf.ContractWhitelist, defaultDatasourcePubKey, conf.IoTeXChainID,
			dispatcherNotification, chainHeadNotification, contractPersistence, projectOffsets)
	}
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new dispatcher"))
	}
	taskDispatcher.Run()

	go func() {
		if err := api.NewHttpServer(persistence, conf).Run(conf.ServiceEndpoint); err != nil {
			log.Fatal(errors.Wrap(err, "failed to run http server"))
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
