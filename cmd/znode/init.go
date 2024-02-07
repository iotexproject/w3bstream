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

func initConfig() {
	viper.MustBindEnv(Risc0ServerEndpoint)
	viper.MustBindEnv(Halo2ServerEndpoint)
	viper.MustBindEnv(ZkwasmServerEndpoint)
	viper.MustBindEnv(ChainEndpoint)
	viper.MustBindEnv(ProjectContractAddress)
	viper.MustBindEnv(ZNodeContractAddress)
	viper.MustBindEnv(DatabaseDSN)
	viper.MustBindEnv(BootNodeMultiaddr)
	viper.MustBindEnv(IotexChainID)
	viper.MustBindEnv(IPFSEndpoint)
	viper.MustBindEnv(IoID)

	viper.SetDefault(Risc0ServerEndpoint, "risc0:4001")
	viper.SetDefault(Halo2ServerEndpoint, "halo2:4001")
	viper.SetDefault(ZkwasmServerEndpoint, "zkwasm:4001")
	viper.SetDefault(ChainEndpoint, "https://babel-api.testnet.iotex.io")
	viper.SetDefault(ProjectContractAddress, "0x02feBE78F3A740b3e9a1CaFAA1b23a2ac0793D26")
	viper.SetDefault(ZNodeContractAddress, "0x45fe67CB442B2e88Ab18229a1992AA134C05c7C9")
	viper.SetDefault(DatabaseDSN, "postgres://test_user:test_passwd@postgres:5432/test?sslmode=disable")
	viper.SetDefault(BootNodeMultiaddr, "/dns4/bootnode-0.testnet.iotex.one/tcp/4689/ipfs/12D3KooWFnaTYuLo8Mkbm3wzaWHtUuaxBRe24Uiopu15Wr5EhD3o")
	viper.SetDefault(IotexChainID, 2)
	viper.SetDefault(IPFSEndpoint, "ipfs.mainnet.iotex.io")
	viper.SetDefault(IoID, "did:key:z6MkmF1AgufHf8ASaxDcCR8iSZjEsEbJMp7LkqyEHw6SNgp8")
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
