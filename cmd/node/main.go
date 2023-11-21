package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	"github.com/machinefi/sprout/cmd/node/apis"
	"github.com/machinefi/sprout/enums"
	"github.com/machinefi/sprout/message/handler"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/vm"
)

func main() {
	vmHandler := vm.NewHandler(
		map[vm.Type]string{
			vm.Risc0: viper.GetString(enums.EnvKeyRisc0ServerEndpoint),
			vm.Halo2: viper.GetString(enums.EnvKeyHalo2ServerEndpoint),
		},
	)
	projectManager, err := project.NewManager(viper.GetString(enums.EnvKeyChainEndpoint), viper.GetString(enums.EnvKeyProjectContractAddress), viper.GetString(enums.EnvKeyProjectFileDirectory))
	if err != nil {
		log.Fatal(err)
	}

	msgHandler := handler.New(
		vmHandler,
		projectManager,
		viper.GetString(enums.EnvKeyChainEndpoint),
		viper.GetString(enums.EnvKeyOperatorPrivateKey),
	)

	go func() {
		if err := apis.NewServer(viper.GetString(enums.EnvKeyServiceEndpoint), msgHandler).Run(); err != nil {
			log.Fatal(err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
