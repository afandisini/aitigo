package otel

import (
	"context"
	"os"
	"testing"
)

func TestResolveExporterDefaultsToStdout(t *testing.T) {
	t.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "")

	exporter, err := resolveExporter(context.Background())
	if err != nil {
		t.Fatalf("expected stdout exporter, got error: %v", err)
	}
	if exporter == nil {
		t.Fatalf("expected exporter")
	}
}

func TestResolveExporterUsesOtlpWhenConfigured(t *testing.T) {
	t.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "localhost:4317")
	defer os.Unsetenv("OTEL_EXPORTER_OTLP_ENDPOINT")

	exporter, err := resolveExporter(context.Background())
	if err != nil {
		t.Fatalf("expected otlp exporter, got error: %v", err)
	}
	if exporter == nil {
		t.Fatalf("expected exporter")
	}
}
