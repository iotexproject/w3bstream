package main

import (
	"log/slog"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func initLogger() {
	var programLevel = slog.LevelDebug
	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))
}

func bindEnvConfig() {
	viper.MustBindEnv(Risc0ServerEndpoint)
	viper.MustBindEnv(Halo2ServerEndpoint)
	viper.MustBindEnv(ZkwasmServerEndpoint)
	viper.MustBindEnv(ProjectFileDirectory)
	viper.MustBindEnv(ChainEndpoint)
	viper.MustBindEnv(ProjectContractAddress)
	viper.MustBindEnv(ZNodeContractAddress)
	viper.MustBindEnv(DatabaseDSN)
	viper.MustBindEnv(BootNodeMultiaddr)
	viper.MustBindEnv(IotexChainID)
	viper.MustBindEnv(IPFSEndpoint)
	viper.MustBindEnv(IoID)

	viper.BindEnv(OperatorPrivateKey)
	viper.BindEnv(OperatorPrivateKeyED25519)

	viper.SetDefault(IPFSEndpoint, gDefaultIPFSEndpoint)
}

// TODO it's risc0 depend tables, will move to risc0
func migrateDatabase() error {
	var schema = `
	CREATE TABLE IF NOT EXISTS vms (
		id SERIAL PRIMARY KEY,
		project_name VARCHAR NOT NULL,
		elf TEXT NOT NULL,
		image_id VARCHAR NOT NULL
	  );
	  
	  CREATE TABLE IF NOT EXISTS proofs (
		id SERIAL PRIMARY KEY,
		name VARCHAR NOT NULL,
		template_name VARCHAR NOT NULL,
		image_id VARCHAR NOT NULL,
		private_input VARCHAR NOT NULL,
		public_input VARCHAR NOT NULL,
		receipt_type VARCHAR NOT NULL,
		receipt TEXT,
		status VARCHAR NOT NULL,
		create_at TIMESTAMP NOT NULL DEFAULT now()
	  );`

	dsn := viper.GetString(DatabaseDSN)
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
