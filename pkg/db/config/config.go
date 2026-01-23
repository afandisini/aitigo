package config

import (
	"fmt"
	"os"
)

type Config struct {
	Driver string
	DSN    string
}

func FromEnv() (Config, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = os.Getenv("DB_DSN")
	}
	if dsn == "" {
		return Config{}, fmt.Errorf("DATABASE_URL or DB_DSN is required")
	}

	driver := os.Getenv("DB_DRIVER")
	if driver == "" {
		driver = "postgres"
	}
	return Config{Driver: driver, DSN: dsn}, nil
}
