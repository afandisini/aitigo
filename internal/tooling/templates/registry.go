package templates

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"path"
	"sort"
)

var ErrTemplateNotFound = fmt.Errorf("template not found")

var templateAliases = map[string]string{
	"gin": "gin-basic",
}

func List() ([]Manifest, error) {
	paths, err := fs.Glob(embeddedTemplates, "templates/*/template.json")
	if err != nil {
		return nil, err
	}

	manifests := make([]Manifest, 0, len(paths))
	for _, manifestPath := range paths {
		manifest, err := readManifest(manifestPath)
		if err != nil {
			return nil, err
		}
		manifests = append(manifests, manifest)
	}

	sort.Slice(manifests, func(i, j int) bool {
		return manifests[i].ID < manifests[j].ID
	})
	return manifests, nil
}

func FindByID(id string) (Manifest, error) {
	manifest, err := LoadManifest(id)
	if err != nil {
		return Manifest{}, err
	}
	if manifest.ID == "" {
		return Manifest{}, ErrTemplateNotFound
	}
	return manifest, nil
}

func LoadManifest(id string) (Manifest, error) {
	id = resolveTemplateID(id)
	if id == "" {
		return Manifest{}, ErrTemplateNotFound
	}
	manifestPath := path.Join("templates", id, "template.json")
	manifest, err := readManifest(manifestPath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return Manifest{}, ErrTemplateNotFound
		}
		return Manifest{}, err
	}
	return manifest, nil
}

func resolveTemplateID(id string) string {
	if id == "" {
		return id
	}
	if resolved, ok := templateAliases[id]; ok {
		return resolved
	}
	return id
}

func readManifest(manifestPath string) (Manifest, error) {
	data, err := fs.ReadFile(embeddedTemplates, manifestPath)
	if err != nil {
		return Manifest{}, err
	}
	var manifest Manifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		return Manifest{}, err
	}
	return manifest, nil
}
