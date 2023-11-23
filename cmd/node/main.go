package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	"github.com/machinefi/sprout/cmd/node/apis"
	"github.com/machinefi/sprout/message"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/vm"
)

func main() {
	initLogger()
	bindEnvConfig()
	if err := migrateDatabase(); err != nil {
		log.Fatal(err)
	}

	vmHandler := vm.NewHandler(
		map[vm.Type]string{
			vm.Risc0: viper.GetString(Risc0ServerEndpoint),
			vm.Halo2: viper.GetString(Halo2ServerEndpoint),
		},
	)
	projectManager, err := project.NewManager(viper.GetString(ChainEndpoint), viper.GetString(ProjectContractAddress), viper.GetString(ProjectFileDirectory))
	if err != nil {
		log.Fatal(err)
	}

	msgHandler, err := message.NewHandler(
		vmHandler,
		projectManager,
		viper.GetString(ChainEndpoint),
		viper.GetString(SequencerServerEndpoint),
		viper.GetString(OperatorPrivateKey),
		viper.GetUint64(ProjectID),
	)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := apis.NewServer(viper.GetString(ServiceEndpoint), msgHandler).Run(); err != nil {
			log.Fatal(err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
