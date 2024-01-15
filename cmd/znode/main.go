package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/task"
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
	// TODO: remove ChainEndpoint, use ChainConfig instead
	projectManager, err := project.NewManager(
		viper.GetString(ChainEndpoint),
		viper.GetString(ProjectContractAddress),
		viper.GetString(ProjectFileDirectory),
		viper.GetString(IPFSEndpoint),
	)
	if err != nil {
		log.Fatal(err)
	}

	taskProcessor, err := task.NewProcessor(vmHandler, projectManager, viper.GetString(OperatorPrivateKey),
		viper.GetString(OperatorPrivateKeyED25519), viper.GetString(BootNodeMultiaddr), viper.GetInt(IotexChainID))
	if err != nil {
		log.Fatal(err)
	}

	taskProcessor.Run()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
