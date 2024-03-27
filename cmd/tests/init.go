package tests

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/machinefi/sprout/clients"
	"github.com/machinefi/sprout/cmd/coordinator/api"
	coordinatorconfig "github.com/machinefi/sprout/cmd/coordinator/config"
	proverconfig "github.com/machinefi/sprout/cmd/prover/config"
	"github.com/machinefi/sprout/datasource"
	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/persistence"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/scheduler"
	"github.com/machinefi/sprout/task"
	"github.com/machinefi/sprout/vm"
)

var (
	env = "INTEGRATION_TEST"
)

func init() {
	_ = os.Setenv("PROVER_ENV", env)
	_ = os.Setenv("COORDINATOR_ENV", env)

	proverConf, err := proverconfig.Get()
	if err != nil {
		os.Exit(-1)
	}
	coordinatorConf, err := coordinatorconfig.Get()
	if err != nil {
		os.Exit(-1)
	}

	go runProver(proverConf)
	go runCoordinator(coordinatorConf)

	// repeat 3 and duration 5s
	if err := checkLiveness(3, 5, func() error {
		if _, e := http.Get(fmt.Sprintf("http://localhost%s/live", coordinatorConf.ServiceEndpoint)); err != nil {
			return e
		}
		return nil
	}); err != nil {
		slog.Error("http server failed to start", "error", err)
	}
}

func checkLiveness(repeat int, duration int64, ping func() error) error {
	for i := 0; i < repeat; i++ {
		if err := ping; err != nil {
			slog.Warn("retry again", "duration", duration, "error", err)
			time.Sleep(time.Duration(duration) * time.Second)
			continue
		}
		break
	}
	return nil
}

func migrateDatabase(dsn string) error {
	var schema = `
	CREATE TABLE IF NOT EXISTS vms (
		id SERIAL PRIMARY KEY,
		project_name VARCHAR NOT NULL,
		elf TEXT NOT NULL,
		image_id VARCHAR NOT NULL
	  );
	  
	  CREATE TABLE IF NOT EXISTS proofs (
		id SERIAL PRIMARY KEY,
		image_id VARCHAR NOT NULL,
		private_input VARCHAR NOT NULL,
		public_input VARCHAR NOT NULL,
		receipt_type VARCHAR NOT NULL,
		receipt TEXT,
		status VARCHAR NOT NULL,
		create_at TIMESTAMP NOT NULL DEFAULT now()
	  );`

	slog.Debug("connecting database", "dsn", dsn)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return errors.Wrap(err, "connect db failed")
	}
	if _, err = db.Exec(schema); err != nil {
		return errors.Wrap(err, "migrate db failed")
	}
	return nil
}

func runProver(conf *proverconfig.Config) {
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

	taskProcessor := task.NewProcessor(vmHandler, projectConfigManager, conf.ProverID)

	pubSubs, err := p2p.NewPubSubs(taskProcessor.HandleP2PData, conf.BootNodeMultiAddr, conf.IoTeXChainID)
	if err != nil {
		log.Fatal(err)
	}

	if err := scheduler.Run(conf.SchedulerEpoch, conf.ChainEndpoint, conf.ProverContractAddress, conf.ProjectContractAddress, conf.ProverID, pubSubs, taskProcessor.HandleProjectProvers); err != nil {
		log.Fatal(err)
	}

	slog.Info("prover started")
}

func runCoordinator(conf *coordinatorconfig.Config) {
	pg, err := persistence.NewPostgres(conf.DatabaseDSN)
	if err != nil {
		log.Fatal(err)
	}

	_ = clients.NewManager()

	projectConfigManager, err := project.NewConfigManager(conf.ChainEndpoint, conf.ProjectContractAddress, conf.ProjectCacheDirectory, conf.IPFSEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	datasource, err := datasource.NewPostgres(conf.DatasourceDSN)
	if err != nil {
		log.Fatal(err)
	}

	nextTaskID, err := pg.FetchNextTaskID()
	if err != nil {
		log.Fatal(err)
	}

	dispatcher, err := task.NewDispatcher(pg, projectConfigManager, datasource, conf.BootNodeMultiAddr, conf.OperatorPrivateKey, conf.OperatorPrivateKeyED25519, conf.IoTeXChainID)
	if err != nil {
		log.Fatal(err)
	}
	go dispatcher.Dispatch(nextTaskID)

	go func() {
		if err := api.NewHttpServer(pg, conf).Run(conf.ServiceEndpoint); err != nil {
			log.Fatal(err)
		}
	}()

	slog.Info("coordinator started")
}
