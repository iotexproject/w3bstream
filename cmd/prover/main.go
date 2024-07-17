package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/cockroachdb/pebble"
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

	sk, err := crypto.HexToECDSA(conf.ProverOperatorPriKey)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to parse prover private key"))
	}
	proverOperatorAddress := crypto.PubkeyToAddress(sk.PublicKey)

	defaultDatasourcePubKey, err := hexutil.Decode(conf.DefaultDatasourcePubKey)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to decode default datasource public key"))
	}

	vmHandler, err := vm.NewHandler(
		map[vm.Type]string{
			vm.Risc0:  conf.Risc0ServerEndpoint,
			vm.Halo2:  conf.Halo2ServerEndpoint,
			vm.ZKwasm: conf.ZKWasmServerEndpoint,
			vm.Wasm:   conf.WasmServerEndpoint,
		},
	)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new vm handler"))
	}

	projectManagerNotification := make(chan uint64, 10)
	schedulerNotification := make(chan uint64, 10)
	chainHeadNotification := make(chan uint64, 10)

	projectNotifications := []chan<- uint64{projectManagerNotification, schedulerNotification}
	chainHeadNotifications := []chan<- uint64{chainHeadNotification}

	local := conf.ProjectFileDir != ""

	var contractPersistence *contract.Contract
	if !local {
		db, err := pebble.Open(conf.LocalDBDir, &pebble.Options{})
		if err != nil {
			log.Fatal(errors.Wrap(err, "failed to open pebble db"))
		}
		defer db.Close()

		contractPersistence, err = contract.New(db, conf.SchedulerEpoch, conf.BeginningBlockNumber,
			conf.ChainEndpoint, common.HexToAddress(conf.ProverContractAddr),
			common.HexToAddress(conf.ProjectContractAddr), chainHeadNotifications, projectNotifications)
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
		projectManager, err = project.NewLocalManager(conf.ProjectFileDir)
	} else {
		projectManager, err = project.NewManager(conf.ProjectCacheDir, conf.IPFSEndpoint, contractPersistence.LatestProject, projectManagerNotification)
	}
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new project manager"))
	}

	taskProcessor := processor.NewProcessor(vmHandler, projectManager.Project, sk, defaultDatasourcePubKey, proverID)

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
