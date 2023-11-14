package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/machinefi/w3bstream-mainnet/enums"
	"github.com/spf13/viper"
	"log/slog"
	"os"
)

func init() {
	initStdLogger()
	initEnvConfigBind()
	initDatabaseMigrating()
}

func initStdLogger() {
	var programLevel = slog.LevelDebug
	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))
}

func initEnvConfigBind() {
	viper.MustBindEnv(enums.EnvKeyServiceEndpoint)
	viper.MustBindEnv(enums.EnvKeyRisc0ServerEndpoint)
	viper.MustBindEnv(enums.EnvKeyHalo2ServerEndpoint)
	viper.MustBindEnv(enums.EnvKeyProjectConfigPath)
	viper.MustBindEnv(enums.EnvKeyChainEndpoint)
	viper.MustBindEnv(enums.EnvKeyOperatorPrivateKey)
	viper.MustBindEnv(enums.EnvKeyDatabaseDSN)
}

func initDatabaseMigrating() {
	// TODO use https://github.com/golang-migrate/migrate
	var schema = `
	CREATE TABLE vms (
		id SERIAL PRIMARY KEY,
		project_name VARCHAR NOT NULL,
		elf TEXT NOT NULL,
		image_id VARCHAR NOT NULL
	  );
	  
	  CREATE TABLE proofs (
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

	db, err := sqlx.Connect("postgres", viper.Get("DATABASE_URL").(string))
	if err != nil {
		slog.Error("connecting database: ", err)
		panic(err)
	}
	if _, err = db.Exec(schema); err != nil {
		slog.Error("migrating database: ", err)
		panic(err)
	}
}
