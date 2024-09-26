package main

import (
	"encoding/json"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"

	"github.com/iotexproject/w3bstream/cmd/prover/config"
	"github.com/iotexproject/w3bstream/p2p"
	"github.com/iotexproject/w3bstream/project"
	"github.com/iotexproject/w3bstream/scheduler"
	"github.com/iotexproject/w3bstream/task/processor"
	"github.com/iotexproject/w3bstream/vm"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to get config"))
	}
	cfg.Print()
	slog.Info("prover config loaded")

	if err := migrateDatabase(cfg.DatabaseDSN); err != nil {
		log.Fatal(errors.Wrap(err, "failed to migrate database"))
	}

	sk, err := crypto.HexToECDSA(cfg.ProverOperatorPriKey)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to parse prover private key"))
	}
	proverOperatorAddress := crypto.PubkeyToAddress(sk.PublicKey)
	slog.Info("my prover", "address", proverOperatorAddress.String())

	vmEndpoints := map[uint64]string{}
	if err := json.Unmarshal([]byte(cfg.VMEndpoints), &vmEndpoints); err != nil {
		log.Fatal(errors.Wrap(err, "failed to unmarshal vm endpoints"))
	}

	vmHandler, err := vm.NewHandler(vmEndpoints)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new vm handler"))
	}

	projectManager := project.NewManager(kvDB, contractPersistence.LatestProject)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new project manager"))
	}

	taskProcessor := processor.NewProcessor(vmHandler, projectManager.Project, sk, defaultDatasourcePubKey, proverID)

	pubSubs, err := p2p.NewPubSub(taskProcessor.HandleP2PData, cfg.BootNodeMultiAddr, cfg.IoTeXChainID)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new pubsubs"))
	}

	if local {
		scheduler.RunLocal(pubSubs, taskProcessor.HandleProjectProvers, projectManager)
	} else {
		projectOffsets := scheduler.NewProjectEpochOffsets(cfg.SchedulerEpoch, contractPersistence.LatestProjects, schedulerNotification)

		if err := scheduler.Run(cfg.SchedulerEpoch, proverID, pubSubs, taskProcessor.HandleProjectProvers,
			chainHeadNotification, contractPersistence.Project, contractPersistence.Provers, projectOffsets, projectManager); err != nil {
			log.Fatal(errors.Wrap(err, "failed to run scheduler"))
		}
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
