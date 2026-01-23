package main

import (
	"context"
	"database/sql"
	"log"

	dbconfig "aitigo/pkg/db/config"
	"aitigo/pkg/db/migrate"
)

func main() {
	cfg, err := dbconfig.FromEnv()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open(cfg.Driver, cfg.DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	runner := migrate.NewRunner(db, "migrations")
	if _, err := runner.Up(context.Background()); err != nil {
		log.Fatal(err)
	}
}
