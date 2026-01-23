package s3

import "testing"

func TestFromEnvParsesSSL(t *testing.T) {
	t.Setenv("S3_USE_SSL", "true")
	cfg, err := FromEnv()
	if err != nil {
		t.Fatalf("expected config, got error: %v", err)
	}
	if !cfg.UseSSL {
		t.Fatalf("expected ssl true")
	}
}
