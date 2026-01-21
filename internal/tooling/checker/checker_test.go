package checker

import "testing"

func TestDummyChecker(t *testing.T) {
	// test sengaja simple dulu
	if true != true {
		t.Fatal("unreachable, but keeps go test honest")
	}
}
