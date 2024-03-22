package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"

	"github.com/machinefi/sprout/cmd/znode/config"
	"github.com/machinefi/sprout/persistence"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/task"
	"github.com/machinefi/sprout/vm"
)

var conf *config.Config

func main() {
	var err error
	conf, err = config.Get()
	if err != nil {
		panic(errors.Wrap(err, "failed to init enode config"))
	}
	conf.Print()
	slog.Info("znode config loaded")

	if err := migrateDatabase(); err != nil {
		log.Fatal(err)
	}

	vmHandler := vm.NewHandler(
		map[vm.Type]string{
			vm.Risc0:  conf.Risc0ServerEndpoint,
			vm.Halo2:  conf.Halo2ServerEndpoint,
			vm.ZKwasm: conf.ZKWasmServerEndpoint,
			vm.Wasm:   conf.WasmServerEndpoint,
		},
	)

	znodes, err := persistence.NewZNode(conf.ChainEndpoint, conf.ZnodeContractAddress)
	if err != nil {
		log.Fatal(err)
	}

	projectManager, err := project.NewManager(conf.ChainEndpoint, conf.ProjectContractAddress, conf.ProjectFileDirectory, conf.ProjectCacheDirectory, conf.IPFSEndpoint, conf.IoID, znodes.GetAll())
	if err != nil {
		log.Fatal(err)
	}

	taskProcessor, err := task.NewProcessor(vmHandler, projectManager, conf.BootNodeMultiAddr, conf.IoID, conf.IoTeXChainID)
	if err != nil {
		log.Fatal(err)
	}

	taskProcessor.Run()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
