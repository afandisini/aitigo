# Metrics (Prometheus)

Package: `pkg/observability/metrics`

Metrics:

- `http_requests_total`
- `http_request_duration_seconds`
- `http_in_flight`

Usage:

```go
metricsSvc, _ := metrics.New(metrics.Config{})
mux.Handle("/metrics", metricsSvc.Handler())
handler := metricsSvc.Middleware(mux)
```
