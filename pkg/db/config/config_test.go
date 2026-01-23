package config

import "testing"

func TestFromEnv(t *testing.T) {
	t.Setenv("DATABASE_URL", "postgres://user:pass@localhost/db")
	t.Setenv("DB_DRIVER", "postgres")

	cfg, err := FromEnv()
	if err != nil {
		t.Fatalf("expected config, got error: %v", err)
	}
	if cfg.DSN == "" || cfg.Driver == "" {
		t.Fatalf("expected dsn and driver")
	}
}
