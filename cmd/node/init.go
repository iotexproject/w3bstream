package main

import (
	"log/slog"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	"github.com/machinefi/sprout/enums"
)

func init() {
	initStdLogger()
	initEnvConfigBind()
	LogAllSettings()
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
	viper.MustBindEnv(enums.EnvKeyProjectFileDirectory)
	viper.MustBindEnv(enums.EnvKeyChainEndpoint)
	viper.MustBindEnv(enums.EnvKeyProjectContractAddress)
	viper.MustBindEnv(enums.EnvKeyDatabaseDSN)

	viper.BindEnv(enums.EnvKeyOperatorPrivateKey)
}

func LogAllSettings() {
	settings := viper.AllSettings()
	slog.Debug("--------------")
	for key, value := range settings {
		slog.Debug("SETTING:", key, value)
	}
	slog.Debug("--------------")
}

func initDatabaseMigrating() {
	// TODO use https://github.com/golang-migrate/migrate
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

	dsn := viper.GetString(enums.EnvKeyDatabaseDSN)
	slog.Debug("connecting database", "dsn", dsn)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		slog.Error("connecting database: ", err)
		panic(err)
	}
	if _, err = db.Exec(schema); err != nil {
		slog.Error("migrating database: ", err)
		panic(err)
	}
}
