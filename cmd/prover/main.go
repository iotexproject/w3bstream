package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/cmd/prover/config"
	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/persistence/contract"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/scheduler"
	"github.com/machinefi/sprout/task/processor"
	"github.com/machinefi/sprout/vm"
)

func main() {
	conf, err := config.Get()
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to get config"))
	}
	conf.Print()
	slog.Info("prover config loaded")

	if err := migrateDatabase(conf.DatabaseDSN); err != nil {
		log.Fatal(errors.Wrap(err, "failed to migrate database"))
	}

	sk, err := crypto.HexToECDSA(conf.ProverOperatorPrivateKey)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to parse prover private key"))
	}
	proverOperatorAddress := crypto.PubkeyToAddress(sk.PublicKey)

	sequencerPubKey, err := hexutil.Decode(conf.SequencerPubKey)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to decode sequencer pubkey"))
	}

	vmHandler := vm.NewHandler(
		map[vm.Type]string{
			vm.Risc0:  conf.Risc0ServerEndpoint,
			vm.Halo2:  conf.Halo2ServerEndpoint,
			vm.ZKwasm: conf.ZKWasmServerEndpoint,
			vm.Wasm:   conf.WasmServerEndpoint,
		},
	)

	projectManagerNotification := make(chan *contract.Project, 10)
	schedulerNotification := make(chan *contract.Project, 10)
	chainHeadNotification := make(chan uint64, 10)

	projectNotifications := []chan<- *contract.Project{projectManagerNotification, schedulerNotification}
	chainHeadNotifications := []chan<- uint64{chainHeadNotification}

	local := conf.ProjectFileDirectory != ""

	var contractPersistence *contract.Contract
	if !local {
		contractPersistence, err = contract.New(conf.SchedulerEpoch, conf.ChainEndpoint, common.HexToAddress(conf.ProverContractAddress),
			common.HexToAddress(conf.ProjectContractAddress), common.HexToAddress(conf.BlockNumberContractAddress),
			common.HexToAddress(conf.MultiCallContractAddress), chainHeadNotifications, projectNotifications)
		if err != nil {
			log.Fatal(errors.Wrap(err, "failed to new contract persistence"))
		}
	}

	proverID := uint64(0)
	if !local {
		p := contractPersistence.Prover(proverOperatorAddress)
		if p == nil {
			log.Fatal(errors.New("failed to query operator's prover id"))
		}
		proverID = p.ID
	}
	slog.Info("my prover id", "prover_id", proverID)

	var projectManager *project.Manager
	if local {
		projectManager, err = project.NewLocalManager(conf.ProjectFileDirectory)
	} else {
		projectManager, err = project.NewManager(conf.ProjectCacheDirectory, conf.IPFSEndpoint, contractPersistence.LatestProject, projectManagerNotification)
	}
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new project manager"))
	}

	taskProcessor := processor.NewProcessor(vmHandler, projectManager.Project, sk, sequencerPubKey, proverID)

	pubSubs, err := p2p.NewPubSubs(taskProcessor.HandleP2PData, conf.BootNodeMultiAddr, conf.IoTeXChainID)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new pubsubs"))
	}

	if local {
		scheduler.RunLocal(pubSubs, taskProcessor.HandleProjectProvers, projectManager.ProjectIDs)
	} else {
		projectOffsets := scheduler.NewProjectEpochOffsets(conf.SchedulerEpoch, contractPersistence.LatestProjects, schedulerNotification)

		if err := scheduler.Run(conf.SchedulerEpoch, proverID, pubSubs, taskProcessor.HandleProjectProvers,
			chainHeadNotification, contractPersistence.Project, contractPersistence.Provers, projectOffsets); err != nil {
			log.Fatal(errors.Wrap(err, "failed to run scheduler"))
		}
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
