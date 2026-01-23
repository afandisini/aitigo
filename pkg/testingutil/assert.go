package testingutil

import "testing"

func RequireNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func RequireTrue(t *testing.T, condition bool, message string) {
	t.Helper()
	if !condition {
		t.Fatalf("assertion failed: %s", message)
	}
}

func RequireEqual[T comparable](t *testing.T, got, want T, message string) {
	t.Helper()
	if got != want {
		t.Fatalf("%s: got %v, want %v", message, got, want)
	}
}
