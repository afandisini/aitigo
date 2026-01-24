package templates

import "testing"

func TestGinTemplateRegistered(t *testing.T) {
	manifests, err := List()
	if err != nil {
		t.Fatalf("list templates: %v", err)
	}

	for _, manifest := range manifests {
		if manifest.ID == "gin-basic" {
			return
		}
	}

	t.Fatal("gin-basic template not registered")
}

func TestGinAliasResolves(t *testing.T) {
	manifest, err := LoadManifest("gin")
	if err != nil {
		t.Fatalf("load gin alias: %v", err)
	}
	if manifest.ID != "gin-basic" {
		t.Fatalf("expected gin-basic manifest, got %q", manifest.ID)
	}
}
