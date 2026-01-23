package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"aitigo/pkg/middleware/requestid"
	"aitigo/pkg/observability/logging"
	"aitigo/pkg/observability/metrics"
	"aitigo/pkg/observability/otel"
)

func main() {
	logger := logging.New(logging.Config{Level: "info"})
	logging.SetDefault(logger)

	shutdown, err := otel.Setup(context.Background(), otel.Config{
		ServiceName:    "aitigo-example",
		ServiceVersion: "0.2.0",
		Environment:    "dev",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = shutdown(context.Background()) }()

	metricsSvc, err := metrics.New(metrics.Config{})
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", metricsSvc.Handler())
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		logging.With(r.Context()).Info("hello request")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("hello"))
	})

	handler := requestid.Middleware(requestid.DefaultConfig())(mux)
	handler = metricsSvc.Middleware(handler)
	handler = otel.Middleware("http-server")(handler)

	server := &http.Server{
		Addr:              ":8081",
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("listening on :8081")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
