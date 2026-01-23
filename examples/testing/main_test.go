package testingutil_test

import (
	"net/http"
	"testing"
	"time"

	"aitigo/pkg/testingutil"
)

func TestExampleHelpers(t *testing.T) {
	clock := testingutil.NewMockClock(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC))
	testingutil.RequireTrue(t, clock.Now().Year() == 2024, "clock should be deterministic")

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	req := testingutil.NewRequest(http.MethodGet, "/health", nil)
	rec := testingutil.ExecuteRequest(handler, req)
	testingutil.RequireEqual(t, rec.Code, http.StatusOK, "response status")
}
