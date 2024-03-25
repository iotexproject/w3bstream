package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/machinefi/sprout/cmd/prover/config"
	"github.com/machinefi/sprout/persistence"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/task"
	"github.com/machinefi/sprout/vm"
)

func main() {
	conf, err := config.Get()
	if err != nil {
		log.Fatal(err)
	}
	conf.Print()
	slog.Info("prover config loaded")

	if err := migrateDatabase(conf.DatabaseDSN); err != nil {
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

	provers, err := persistence.NewProver(conf.ChainEndpoint, conf.ProverContractAddress)
	if err != nil {
		log.Fatal(err)
	}

	projectManager, err := project.NewManager(conf.ChainEndpoint, conf.ProjectContractAddress, conf.ProjectFileDirectory, conf.ProjectCacheDirectory, conf.IPFSEndpoint, conf.ProverID, provers.GetAll())
	if err != nil {
		log.Fatal(err)
	}

	taskProcessor, err := task.NewProcessor(vmHandler, projectManager, conf.BootNodeMultiAddr, conf.ProverID, conf.IoTeXChainID)
	if err != nil {
		log.Fatal(err)
	}

	taskProcessor.Run()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
