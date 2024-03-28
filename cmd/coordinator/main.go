package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/machinefi/sprout/clients"
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
		log.Fatal(err)
	}
	conf.Print()
	slog.Info("coordinator config loaded")

	persistence, err := persistence.NewPostgres(conf.DatabaseDSN)
	if err != nil {
		log.Fatal(err)
	}

	_ = clients.NewManager()

	projectConfigManager, err := project.NewConfigManager(conf.ChainEndpoint, conf.ProjectContractAddress, conf.ProjectCacheDirectory, conf.IPFSEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	datasource, err := datasource.NewPostgres(conf.DatasourceDSN)
	if err != nil {
		log.Fatal(err)
	}

	nextTaskID, err := persistence.FetchNextTaskID()
	if err != nil {
		log.Fatal(err)
	}

	dispatcher, err := task.NewDispatcher(persistence, projectConfigManager, datasource, conf.BootNodeMultiAddr, conf.OperatorPrivateKey, conf.OperatorPrivateKeyED25519, conf.IoTeXChainID)
	if err != nil {
		log.Fatal(err)
	}
	go dispatcher.Dispatch(nextTaskID)

	go func() {
		if err := api.NewHttpServer(persistence, conf).Run(conf.ServiceEndpoint); err != nil {
			log.Fatal(err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
