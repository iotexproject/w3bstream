package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/cmd/prover/config"
	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/scheduler"
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

	projectConfigManager, err := project.NewConfigManager(conf.ChainEndpoint, conf.ProjectContractAddress, conf.ProjectCacheDirectory, conf.IPFSEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	sequencerPubKey, err := hexutil.Decode(conf.SequencerPubKey)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to decode sequencer pubkey"))
	}
	proverPubKey, err := hexutil.Decode(conf.ProverPubKey)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to decode prover pubkey"))
	}
	taskProcessor := task.NewProcessor(vmHandler, projectConfigManager, conf.ProverPrivateKey, sequencerPubKey, proverPubKey)

	pubSubs, err := p2p.NewPubSubs(taskProcessor.HandleP2PData, conf.BootNodeMultiAddr, conf.IoTeXChainID)
	if err != nil {
		log.Fatal(err)
	}

	if err := scheduler.Run(conf.SchedulerEpoch, conf.ChainEndpoint, conf.ProverContractAddress, conf.ProjectContractAddress, conf.ProverPrivateKey, pubSubs, taskProcessor.HandleProjectProvers); err != nil {
		log.Fatal(err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
