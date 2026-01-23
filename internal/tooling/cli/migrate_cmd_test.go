package cli

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestMigrateHelp(t *testing.T) {
	output, err := captureStdout(t, func() error {
		return runMigrate([]string{"--help"})
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !strings.Contains(output, "AitiGo migrate") {
		t.Fatalf("expected migrate help output, got %q", output)
	}
}

func TestMigrateUpHelp(t *testing.T) {
	output, err := captureStdout(t, func() error {
		return runMigrate([]string{"up", "--help"})
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !strings.Contains(output, "AitiGo migrate up") {
		t.Fatalf("expected migrate up help output, got %q", output)
	}
}

func TestMigrateUnknownSubcommand(t *testing.T) {
	err := runMigrate([]string{"foo"})
	if err == nil {
		t.Fatalf("expected error for unknown subcommand")
	}
	if !strings.Contains(err.Error(), "unknown migrate subcommand") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func captureStdout(t *testing.T, fn func() error) (string, error) {
	t.Helper()

	orig := os.Stdout
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("pipe: %v", err)
	}
	os.Stdout = writer

	runErr := fn()
	_ = writer.Close()
	os.Stdout = orig

	output, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	_ = reader.Close()
	return string(output), runErr
}
