package testingutil

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestMockClock(t *testing.T) {
	now := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	clock := NewMockClock(now)
	if got := clock.Now(); !got.Equal(now) {
		t.Fatalf("expected initial time")
	}
}

func TestLoadJSONFixture(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "fixture.json")
	if err := os.WriteFile(path, []byte(`{"name":"aitigo"}`), 0o644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	var payload struct {
		Name string `json:"name"`
	}
	if err := LoadJSONFixture(path, &payload); err != nil {
		t.Fatalf("load fixture: %v", err)
	}
	if payload.Name != "aitigo" {
		t.Fatalf("expected name")
	}
}
