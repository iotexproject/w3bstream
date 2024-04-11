package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"

	"github.com/machinefi/sprout/cmd/coordinator/api"
	"github.com/machinefi/sprout/cmd/coordinator/config"
	"github.com/machinefi/sprout/datasource"
	"github.com/machinefi/sprout/persistence"
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

	persistence, err := persistence.NewPostgres(conf.DatabaseDSN)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new postgres persistence"))
	}

	projectConfigManager, err := project.NewManager(conf.ChainEndpoint, conf.ProjectContractAddress,
		conf.ProjectCacheDirectory, conf.IPFSEndpoint, conf.ProjectFileDirectory)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new project config manager"))
	}

	if err := task.RunDispatcher(persistence, datasource.NewPostgres, projectConfigManager.GetAllCacheProjectIDs, projectConfigManager.Get, conf.BootNodeMultiAddr, conf.OperatorPrivateKey, conf.OperatorPrivateKeyED25519, conf.ChainEndpoint, conf.ProjectContractAddress, conf.ProjectFileDirectory, conf.IoTeXChainID); err != nil {
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
