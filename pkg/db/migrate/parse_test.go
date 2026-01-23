package migrate

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadOrdersMigrations(t *testing.T) {
	dir := t.TempDir()
	files := []string{
		"0002_add_table.up.sql",
		"0002_add_table.down.sql",
		"0001_init.up.sql",
		"0001_init.down.sql",
	}
	for _, name := range files {
		if err := os.WriteFile(filepath.Join(dir, name), []byte("-- sql"), 0o644); err != nil {
			t.Fatalf("write file: %v", err)
		}
	}

	migrations, err := Load(dir)
	if err != nil {
		t.Fatalf("load migrations: %v", err)
	}
	if len(migrations) != 2 {
		t.Fatalf("expected 2 migrations, got %d", len(migrations))
	}
	if migrations[0].Version != 1 || migrations[1].Version != 2 {
		t.Fatalf("expected ordered versions")
	}
}

func TestCreateGeneratesFiles(t *testing.T) {
	dir := t.TempDir()
	mig, err := Create(dir, "Create Users")
	if err != nil {
		t.Fatalf("create migration: %v", err)
	}
	if _, err := os.Stat(mig.UpPath); err != nil {
		t.Fatalf("expected up file: %v", err)
	}
	if _, err := os.Stat(mig.DownPath); err != nil {
		t.Fatalf("expected down file: %v", err)
	}
}
