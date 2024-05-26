package e2etest

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/machinefi/ioconnect-go/pkg/ioconnect"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"

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
	endpoint              string
	privateKey            string
	databaseDSN           string
	coordinatorAddr       string
	didAuthServerAddr     string
	did                   bool
	jwkSecret             string
	jwk                   *ioconnect.JWK
	projectClientContract string
	ioIDContract          string
	ioIDContractAddress   string
	chainEndpoint         string
	ioIDRegistryEndpoint  string
}

func seqConf(coordinatorEndpoint, didAuthServerAddr string) *sequencerConf {
	ret := &sequencerConf{
		endpoint:             ":19000",
		privateKey:           "dbfe03b0406549232b8dccc04be8224fcc0afa300a33d4f335dcfdfead861c85",
		databaseDSN:          "postgres://test_user:test_passwd@localhost:15432/test?sslmode=disable",
		coordinatorAddr:      fmt.Sprintf("localhost%s", coordinatorEndpoint),
		didAuthServerAddr:    didAuthServerAddr,
		did:                  false,
		jwkSecret:            "JyBSf1Y9UFIbECdHJRkqZAkfaz16YGwRYyJpXSBwQXJ8RTsZcVMOWStaQxI8bUtafiMtYUE2KkgmI1wBGiQYMw==",
		ioIDContract:         "0xB63e6034361283dc8516480a46EcB9C122c983Bb",
		chainEndpoint:        "http://iotex.chainendpoint.io",
		ioIDRegistryEndpoint: "did.iotex.me",
	}

	jwk, err := ioconnect.NewJWKBySecretBase64(ret.jwkSecret)
	if err != nil {
		panic(errors.Wrap(err, "failed generate jwk from secret"))
	}
	ret.jwk = jwk
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

	conf := seqConf(coordinatorConf.ServiceEndpoint, coordinatorConf.DIDAuthServerEndpoint)
	go runSequencer(conf.privateKey, conf.databaseDSN, conf.coordinatorAddr, conf.didAuthServerAddr, conf.endpoint, conf.projectClientContract, conf.ioIDContractAddress, conf.chainEndpoint, conf.ioIDRegistryEndpoint, conf.jwk)
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

	vmHandler := vm.NewHandler(
		map[vm.Type]string{
			vm.Risc0:  conf.Risc0ServerEndpoint,
			vm.Halo2:  conf.Halo2ServerEndpoint,
			vm.ZKwasm: conf.ZKWasmServerEndpoint,
			vm.Wasm:   conf.WasmServerEndpoint,
		},
	)

	projectManager, err := project.NewLocalManager(conf.ProjectFileDirectory)
	if err != nil {
		log.Fatal(err)
	}

	sk, err := crypto.HexToECDSA(conf.ProverOperatorPrivateKey)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed parse prover private key"))
	}

	sequencerPubKey, err := hexutil.Decode(conf.SequencerPubKey)
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

	sequencerPubKey, err := hexutil.Decode(conf.SequencerPubKey)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to decode sequencer pubkey"))
	}

	projectManager, err := project.NewLocalManager(conf.ProjectFileDirectory)
	if err != nil {
		log.Fatal(err)
	}

	datasourcePG := datasource.NewPostgres()

	taskDispatcher, err := dispatcher.NewLocal(pg, datasourcePG.New, projectManager, conf.DefaultDatasourceURI, conf.OperatorPrivateKey, conf.OperatorPrivateKeyED25519, conf.BootNodeMultiAddr, sequencerPubKey, conf.IoTeXChainID)
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

func runSequencer(privateKey, databaseDSN, coordinatorAddress, didAuthServer, address, projectClientContract, ioIDContract, chainEndpoint, ioRegistryEndpoint string, key *ioconnect.JWK) {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.Level(int(slog.LevelDebug))})))

	sk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed parse private key"))
	}

	slog.Info("sequencer public key", "public_key", hexutil.Encode(crypto.FromECDSAPub(&sk.PublicKey)))

	manager, err := clients.NewManager(projectClientContract, ioIDContract, chainEndpoint, ioRegistryEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	p, err := persistence.NewPersistence(databaseDSN)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := seqapi.NewHttpServer(p, uint(1), coordinatorAddress, didAuthServer, sk, key, manager).Run(address); err != nil {
			log.Fatal(err)
		}
	}()

	slog.Info("sequencer started")
}
