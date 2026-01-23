package metrics

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Config struct {
	Namespace string
	Registry  *prometheus.Registry
}

type Metrics struct {
	registry *prometheus.Registry

	requests *prometheus.CounterVec
	duration *prometheus.HistogramVec
	inFlight prometheus.Gauge
}

func New(cfg Config) (*Metrics, error) {
	registry := cfg.Registry
	if registry == nil {
		registry = prometheus.NewRegistry()
	}
	if cfg.Namespace == "" {
		cfg.Namespace = "aitigo"
	}

	requests := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: cfg.Namespace,
		Name:      "http_requests_total",
		Help:      "Total number of HTTP requests.",
	}, []string{"method", "path", "status"})
	duration := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: cfg.Namespace,
		Name:      "http_request_duration_seconds",
		Help:      "Duration of HTTP requests.",
	}, []string{"method", "path", "status"})
	inFlight := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: cfg.Namespace,
		Name:      "http_in_flight",
		Help:      "Number of HTTP requests in flight.",
	})

	if err := registry.Register(requests); err != nil {
		return nil, err
	}
	if err := registry.Register(duration); err != nil {
		return nil, err
	}
	if err := registry.Register(inFlight); err != nil {
		return nil, err
	}

	return &Metrics{
		registry: registry,
		requests: requests,
		duration: duration,
		inFlight: inFlight,
	}, nil
}

func (m *Metrics) Handler() http.Handler {
	return promhttp.HandlerFor(m.registry, promhttp.HandlerOpts{})
}

func (m *Metrics) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		m.inFlight.Inc()
		defer m.inFlight.Dec()

		rr := &responseRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rr, r)

		labels := prometheus.Labels{
			"method": r.Method,
			"path":   r.URL.Path,
			"status": strconv.Itoa(rr.status),
		}
		m.requests.With(labels).Inc()
		m.duration.With(labels).Observe(time.Since(start).Seconds())
	})
}

type responseRecorder struct {
	http.ResponseWriter
	status int
}

func (r *responseRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}
