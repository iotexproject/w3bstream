package e2etest

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/machinefi/ioconnect-go/pkg/ioconnect"
	"github.com/machinefi/sprout/clients"
	"github.com/machinefi/sprout/cmd/coordinator/api"
	coordinatorconfig "github.com/machinefi/sprout/cmd/coordinator/config"
	proverconfig "github.com/machinefi/sprout/cmd/prover/config"
	seqapi "github.com/machinefi/sprout/cmd/sequencer/api"
	"github.com/machinefi/sprout/cmd/sequencer/persistence"
	"github.com/machinefi/sprout/datasource"
	"github.com/machinefi/sprout/p2p"
	"github.com/machinefi/sprout/persistence/postgres"
	"github.com/machinefi/sprout/project"
	"github.com/machinefi/sprout/scheduler"
	"github.com/machinefi/sprout/task/dispatcher"
	"github.com/machinefi/sprout/task/processor"
	"github.com/machinefi/sprout/vm"
)

var (
	env = "INTEGRATION_TEST"
)

type sequencerConf struct {
	aggregationAmount               uint
	address                         string
	coordinatorAddress              string
	databaseDSN                     string
	privateKey                      string
	jwkSecret                       string
	jwk                             *ioconnect.JWK
	ioIDRegistryEndpoint            string
	ioIDRegistryContractAddress     string
	projectClientContractAddress    string
	w3bstreamProjectContractAddress string
	chainEndpoint                   string
}

func seqConf(coordinatorEndpoint string) *sequencerConf {
	ret := &sequencerConf{
		aggregationAmount:               uint(1),
		address:                         ":19000",
		coordinatorAddress:              coordinatorEndpoint,
		databaseDSN:                     "postgres://test_user:test_passwd@localhost:15432/test?sslmode=disable",
		privateKey:                      "dbfe03b0406549232b8dccc04be8224fcc0afa300a33d4f335dcfdfead861c85",
		jwkSecret:                       "R3QNJihYLjtcaxALSTsKe1cYSX0pS28wZitFVXE4Y2klf2hxVCczYHw2dVg4fXJdSgdCcnM4PgV1aTo9DwYqEw==",
		ioIDRegistryContractAddress:     "0x06b3Fcda51e01EE96e8E8873F0302381c955Fddd",
		projectClientContractAddress:    "0xF4d6282C5dDD474663eF9e70c927c0d4926d1CEb",
		w3bstreamProjectContractAddress: "0x6AfCB0EB71B7246A68Bb9c0bFbe5cD7c11c4839f",
		chainEndpoint:                   "http://iotex.chainendpoint.io",
		ioIDRegistryEndpoint:            "did.iotex.me",
	}

	if ret.jwkSecret != "" {
		var (
			secrets = ioconnect.JWKSecrets{}
			jwk     *ioconnect.JWK
			err     error
		)
		if err = secrets.UnmarshalText([]byte(ret.jwkSecret)); err != nil {
			panic(errors.Wrap(err, "invalid jwk secrets from flag"))
		}
		if jwk, err = ioconnect.NewJWKBySecret(secrets); err != nil {
			panic(errors.Wrap(err, "failed to new jwk from secrets"))
		}
		ret.jwk = jwk
	}

	return ret
}

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

	conf := seqConf(coordinatorConf.ServiceEndpoint)
	go runSequencer(conf.privateKey, conf.databaseDSN, conf.coordinatorAddress, conf.address,
		conf.projectClientContractAddress, conf.ioIDRegistryContractAddress, conf.w3bstreamProjectContractAddress,
		conf.ioIDRegistryEndpoint, conf.chainEndpoint, conf.jwk)
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
		project_id VARCHAR NOT NULL,
		task_id VARCHAR NOT NULL,
		client_id VARCHAR NOT NULL,
		sequencer_sign VARCHAR NOT NULL,
		image_id VARCHAR NOT NULL,
		datas_input VARCHAR NOT NULL,
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

	projectManager, err := project.NewLocalManager(conf.ProjectFileDir)
	if err != nil {
		log.Fatal(err)
	}

	sk, err := crypto.HexToECDSA(conf.ProverOperatorPriKey)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed parse prover private key"))
	}

	sequencerPubKey, err := hexutil.Decode(conf.DefaultDatasourcePubKey)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to decode sequencer pubkey"))
	}

	taskProcessor := processor.NewProcessor(vmHandler, projectManager.Project, sk, sequencerPubKey, 1)

	pubSubs, err := p2p.NewPubSubs(taskProcessor.HandleP2PData, conf.BootNodeMultiAddr, conf.IoTeXChainID)
	if err != nil {
		log.Fatal(err)
	}

	scheduler.RunLocal(pubSubs, taskProcessor.HandleProjectProvers, projectManager.ProjectIDs)

	slog.Info("prover started")
}

func runCoordinator(conf *coordinatorconfig.Config) {
	pg, err := postgres.New(conf.DatabaseDSN)
	if err != nil {
		log.Fatal(err)
	}

	sequencerPubKey, err := hexutil.Decode(conf.DefaultDatasourcePubKey)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to decode sequencer pubkey"))
	}

	projectManager, err := project.NewLocalManager(conf.ProjectFileDir)
	if err != nil {
		log.Fatal(err)
	}

	datasourcePG := datasource.NewPostgres()

	taskDispatcher, err := dispatcher.NewLocal(pg, datasourcePG.New, projectManager, conf.DefaultDatasourceURI, conf.OperatorPriKey, conf.OperatorPriKeyED25519, conf.BootNodeMultiAddr, conf.ContractWhitelist, sequencerPubKey, conf.IoTeXChainID)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to new local dispatcher"))
	}
	taskDispatcher.Run()

	go func() {
		if err := api.NewHttpServer(pg, conf).Run(conf.ServiceEndpoint); err != nil {
			log.Fatal(err)
		}
	}()

	slog.Info("coordinator started")
}

func runSequencer(privateKey, databaseDSN, coordinatorAddress, address, projectClientContractAddress,
	ioIDRegistryContractAddress, w3bstreamProjectContractAddress, ioIDRegistryEndpoint, chainEndpoint string, key *ioconnect.JWK) {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.Level(int(slog.LevelDebug))})))

	sk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed parse private key"))
	}

	slog.Info("sequencer public key", "public_key", hexutil.Encode(crypto.FromECDSAPub(&sk.PublicKey)))

	manager, err := clients.NewManager(projectClientContractAddress, ioIDRegistryContractAddress,
		w3bstreamProjectContractAddress, ioIDRegistryEndpoint, chainEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	p, err := persistence.NewPersistence(databaseDSN)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := seqapi.NewHttpServer(p, uint(1), coordinatorAddress, sk, key, manager).Run(address); err != nil {
			log.Fatal(err)
		}
	}()

	slog.Info("sequencer started")
}
