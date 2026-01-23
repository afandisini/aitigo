# Tracing (OpenTelemetry)

Package: `pkg/observability/otel`

Behavior:

- Stdout exporter by default.
- OTLP exporter enabled when `OTEL_EXPORTER_OTLP_ENDPOINT` is set.
- HTTP middleware uses `otelhttp`.

Usage:

```go
shutdown, err := otel.Setup(context.Background(), otel.Config{
	ServiceName: "my-service",
})
defer shutdown(context.Background())

handler := otel.Middleware("http-server")(mux)
```
