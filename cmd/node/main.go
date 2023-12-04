package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	"github.com/machinefi/sprout/message"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/types"
	"github.com/machinefi/sprout/vm"
)

func main() {
	initLogger()
	bindEnvConfig()
	if err := migrateDatabase(); err != nil {
		log.Fatal(err)
	}

	vmHandler := vm.NewHandler(
		map[types.VM]string{
			types.VMRisc0:  viper.GetString(Risc0ServerEndpoint),
			types.VMHalo2:  viper.GetString(Halo2ServerEndpoint),
			types.VMZkwasm: viper.GetString(ZkwasmServerEndpoint),
		},
	)
	projectManager, err := project.NewManager(viper.GetString(ChainEndpoint), viper.GetString(ProjectContractAddress), viper.GetString(ProjectFileDirectory))
	if err != nil {
		log.Fatal(err)
	}

	msgProcessor, err := message.NewProcessor(vmHandler, projectManager, viper.GetString(ChainEndpoint), viper.GetString(OperatorPrivateKey))
	if err != nil {
		log.Fatal(err)
	}

	msgProcessor.Run()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
