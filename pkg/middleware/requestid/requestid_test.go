package requestid

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddlewareGeneratesID(t *testing.T) {
	cfg := DefaultConfig()
	cfg.Generator = func() string { return "test-id" }

	handler := Middleware(cfg)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, ok := FromContext(r.Context())
		if !ok || id != "test-id" {
			t.Fatalf("expected request id in context")
		}
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if got := rec.Header().Get(cfg.HeaderName); got != "test-id" {
		t.Fatalf("expected header set, got %q", got)
	}
}

func TestMiddlewareRespectsExistingID(t *testing.T) {
	cfg := DefaultConfig()
	cfg.Generator = func() string { return "generated" }

	handler := Middleware(cfg)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, ok := FromContext(r.Context())
		if !ok || id != "client-id" {
			t.Fatalf("expected client id in context")
		}
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(cfg.HeaderName, "client-id")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if got := rec.Header().Get(cfg.HeaderName); got != "client-id" {
		t.Fatalf("expected header preserved, got %q", got)
	}
}
