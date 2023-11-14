package main

import (
	"github.com/machinefi/w3bstream-mainnet/cmd/node/apis"
	"github.com/machinefi/w3bstream-mainnet/enums"
	"github.com/machinefi/w3bstream-mainnet/msg/handler"
	"github.com/machinefi/w3bstream-mainnet/vm"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	vmHandler := vm.NewHandler(
		map[vm.Type]string{
			vm.Risc0: viper.GetString(enums.EnvKeyRisc0ServerEndpoint),
			vm.Halo2: viper.GetString(enums.EnvKeyHalo2ServerEndpoint),
		},
	)

	msgHandler := handler.New(
		vmHandler,
		viper.GetString(enums.EnvKeyChainEndpoint),
		viper.GetString(enums.EnvKeyOperatorPrivateKey),
		viper.GetString(enums.EnvKeyProjectConfigPath),
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
