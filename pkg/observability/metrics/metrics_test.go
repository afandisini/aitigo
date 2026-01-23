package metrics

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddlewareRecordsMetrics(t *testing.T) {
	m, err := New(Config{})
	if err != nil {
		t.Fatalf("expected metrics, got error: %v", err)
	}

	handler := m.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	}))

	req := httptest.NewRequest(http.MethodPost, "/test", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	families, err := m.registry.Gather()
	if err != nil {
		t.Fatalf("expected gather, got error: %v", err)
	}

	found := false
	for _, family := range families {
		if family.GetName() == "aitigo_http_requests_total" {
			found = true
			if len(family.Metric) == 0 {
				t.Fatalf("expected metrics recorded")
			}
		}
	}
	if !found {
		t.Fatalf("expected requests metric family")
	}
}
