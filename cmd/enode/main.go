package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	"github.com/machinefi/sprout/cmd/enode/api"
	"github.com/machinefi/sprout/persistence"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/task"
)

func main() {
	initLogger()
	bindEnvConfig()

	projectManager, err := project.NewManager(viper.GetString(ChainEndpoint), viper.GetString(ProjectContractAddress), viper.GetString(ProjectFileDirectory))
	if err != nil {
		log.Fatal(err)
	}
	projectIDs := projectManager.GetAllProjectID()

	pg, err := persistence.NewPostgres(viper.GetString(DatabaseDSN))
	if err != nil {
		log.Fatal(err)
	}

	dispatcher, err := task.NewDispatcher(projectIDs, pg, viper.GetString(BootNodeMultiaddr), viper.GetInt(IotexChainID))
	if err != nil {
		log.Fatal(err)
	}
	go dispatcher.Dispatch()

	go func() {
		if err := api.NewHttpServer(pg).Run(viper.GetString(HttpServiceEndpoint)); err != nil {
			log.Fatal(err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
