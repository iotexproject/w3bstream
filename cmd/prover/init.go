package main

import (
	"log/slog"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// TODO it's risc0 depend tables, will move to risc0
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
