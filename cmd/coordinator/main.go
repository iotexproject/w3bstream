package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/cmd/coordinator/api"
	"github.com/machinefi/sprout/cmd/coordinator/config"
	"github.com/machinefi/sprout/datasource"
	"github.com/machinefi/sprout/persistence/contract"
	"github.com/machinefi/sprout/persistence/postgres"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/task"
)

func main() {
	conf, err := config.Get()
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to get config"))
	}
	conf.Print()
	slog.Info("coordinator config loaded")

	persistence, err := postgres.New(conf.DatabaseDSN)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new postgres persistence"))
	}

	projectManagerNotification := make(chan *contract.Project, 10)
	dispatcherNotification := make(chan *contract.Project, 10)

	projectNotifications := []chan<- *contract.Project{projectManagerNotification, dispatcherNotification}

	local := conf.ProjectFileDirectory != ""

	var contractPersistence *contract.Contract
	if !local {
		contractPersistence, err = contract.New(conf.SchedulerEpoch, conf.ChainEndpoint, common.HexToAddress(conf.ProverContractAddress),
			common.HexToAddress(conf.ProjectContractAddress), common.HexToAddress(conf.BlockNumberContractAddress),
			common.HexToAddress(conf.MultiCallContractAddress), nil, projectNotifications)
		if err != nil {
			log.Fatal(errors.Wrap(err, "failed to new contract persistence"))
		}
	}

	var projectManager *project.Manager
	if local {
		projectManager, err = project.NewLocalManager(conf.ProjectFileDirectory)
	} else {
		projectManager, err = project.NewManager(conf.ProjectCacheDirectory, conf.IPFSEndpoint, contractPersistence.LatestProject, projectManagerNotification)
	}
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new project manager"))
	}

	if local {
		err = task.RunLocalDispatcher(persistence, datasource.NewPostgres, projectManager.ProjectIDs, projectManager.Project,
			conf.OperatorPrivateKey, conf.OperatorPrivateKeyED25519, conf.BootNodeMultiAddr, conf.IoTeXChainID)
	} else {
		err = task.RunDispatcher(persistence, datasource.NewPostgres, projectManager.Project, conf.BootNodeMultiAddr,
			conf.OperatorPrivateKey, conf.OperatorPrivateKeyED25519, conf.IoTeXChainID,
			dispatcherNotification, contractPersistence.LatestProjects)
	}
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to run dispatcher"))
	}

	// TODO verify sig
	// sequencerPubKey, err := hexutil.Decode(conf.SequencerPubKey)
	// if err != nil {
	// 	log.Fatal(errors.Wrap(err, "failed to decode sequencer pubkey"))
	// }
	//go dispatcher.Dispatch(nextTaskID, sequencerPubKey)

	go func() {
		if err := api.NewHttpServer(persistence, conf).Run(conf.ServiceEndpoint); err != nil {
			log.Fatal(errors.Wrap(err, "failed to run http server"))
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
