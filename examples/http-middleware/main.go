package main

import (
	"log"
	"net/http"
	"time"

	"aitigo/pkg/middleware/cors"
	"aitigo/pkg/middleware/ratelimit"
	recovermw "aitigo/pkg/middleware/recover"
	"aitigo/pkg/middleware/requestid"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	handler := requestid.Middleware(requestid.DefaultConfig())(mux)
	handler = recovermw.Middleware(recovermw.DefaultConfig())(handler)
	handler = cors.Middleware(cors.Config{
		AllowOrigins:  []string{"http://localhost:3000"},
		AllowMethods:  []string{http.MethodGet, http.MethodPost},
		AllowHeaders:  []string{"Content-Type"},
		MaxAgeSeconds: 600,
	})(handler)
	handler = ratelimit.Middleware(ratelimit.Config{
		Limit: ratelimit.Limit{Rate: 5, Burst: 10, Per: time.Second},
	})(handler)

	log.Println("listening on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
