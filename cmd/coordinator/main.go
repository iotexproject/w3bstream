package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/cmd/coordinator/api"
	"github.com/machinefi/sprout/cmd/coordinator/config"
	"github.com/machinefi/sprout/datasource"
	"github.com/machinefi/sprout/persistence/contract"
	"github.com/machinefi/sprout/persistence/postgres"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/scheduler"
	"github.com/machinefi/sprout/task/dispatcher"
)

func main() {
	conf, err := config.Get()
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to get config"))
	}
	conf.Print()
	slog.Info("coordinator config loaded")

	sequencerPubKey, err := hexutil.Decode(conf.SequencerPubKey)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to decode sequencer pubkey"))
	}

	persistence, err := postgres.New(conf.DatabaseDSN)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new postgres persistence"))
	}

	projectManagerNotification := make(chan uint64, 10)
	schedulerNotification := make(chan uint64, 10)
	dispatcherNotification := make(chan uint64, 10)
	chainHeadNotification := make(chan uint64, 10)

	projectNotifications := []chan<- uint64{projectManagerNotification, dispatcherNotification, schedulerNotification}
	chainHeadNotifications := []chan<- uint64{chainHeadNotification}

	local := conf.ProjectFileDir != ""

	var contractPersistence *contract.Contract
	if !local {
		contractPersistence, err = contract.New(conf.SchedulerEpoch, conf.BeginningBlockNumber, conf.LocalDBDir,
			conf.ChainEndpoint, common.HexToAddress(conf.ProverContractAddr),
			common.HexToAddress(conf.ProjectContractAddr), chainHeadNotifications, projectNotifications)
		if err != nil {
			log.Fatal(errors.Wrap(err, "failed to new contract persistence"))
		}
		defer contractPersistence.Release()
	}

	var projectManager *project.Manager
	if local {
		projectManager, err = project.NewLocalManager(conf.ProjectFileDir)
	} else {
		projectManager, err = project.NewManager(conf.ProjectCacheDir, conf.IPFSEndpoint, contractPersistence.LatestProject, projectManagerNotification)
	}
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new project manager"))
	}

	datasourcePG := datasource.NewPostgres()
	var taskDispatcher *dispatcher.Dispatcher
	if local {
		taskDispatcher, err = dispatcher.NewLocal(persistence, datasourcePG.New, projectManager, conf.DefaultDatasourceURI,
			conf.OperatorPriKey, conf.OperatorPriKeyED25519, conf.BootNodeMultiAddr, sequencerPubKey, conf.IoTeXChainID)
	} else {
		projectOffsets := scheduler.NewProjectEpochOffsets(conf.SchedulerEpoch, contractPersistence.LatestProjects, schedulerNotification)

		taskDispatcher, err = dispatcher.New(persistence, datasourcePG.New, projectManager, conf.DefaultDatasourceURI, conf.BootNodeMultiAddr,
			conf.OperatorPriKey, conf.OperatorPriKeyED25519, sequencerPubKey, conf.IoTeXChainID,
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
