package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

// TODO use https://github.com/golang-migrate/migrate
func dbMigrate() {
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
		log.Fatalln(err)
	}
	db.Exec(schema)
}
