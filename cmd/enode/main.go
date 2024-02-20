package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	"github.com/machinefi/sprout/clients"
	"github.com/machinefi/sprout/cmd/enode/api"
	"github.com/machinefi/sprout/cmd/enode/constant"
	"github.com/machinefi/sprout/persistence"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/task"
)

func main() {
	initLogger()
	initConfig()

	pg, err := persistence.NewPostgres(viper.GetString(constant.DatabaseDSN))
	if err != nil {
		log.Fatal(err)
	}

	_ = clients.NewManager()

	projectManager, err := project.NewManager(viper.GetString(constant.ChainEndpoint), viper.GetString(constant.ProjectContractAddress), viper.GetString(constant.IPFSEndpoint))
	if err != nil {
		log.Fatal(err)
	}

	dispatcher, err := task.NewDispatcher(pg, projectManager, viper.GetString(constant.BootNodeMultiaddr), viper.GetString(constant.OperatorPrivateKey), viper.GetString(constant.OperatorPrivateKeyED25519), viper.GetInt(constant.IotexChainID))
	if err != nil {
		log.Fatal(err)
	}
	go dispatcher.Dispatch()

	go func() {
		if err := api.NewHttpServer(pg, viper.GetString(constant.DIDAuthServerEndpoint), projectManager).Run(viper.GetString(constant.HttpServiceEndpoint)); err != nil {
			log.Fatal(err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
