package redis

import "testing"

func TestFromEnvDefaults(t *testing.T) {
	t.Setenv("REDIS_ADDR", "")
	t.Setenv("REDIS_PASSWORD", "")
	t.Setenv("REDIS_DB", "")

	cfg, err := FromEnv()
	if err != nil {
		t.Fatalf("expected config, got error: %v", err)
	}
	if cfg.Addr == "" {
		t.Fatalf("expected default addr")
	}
}
