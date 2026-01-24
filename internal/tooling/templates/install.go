package templates

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func Install(templateID, targetDir string) error {
	if templateID == "" {
		return fmt.Errorf("template id is required")
	}
	if targetDir == "" {
		return fmt.Errorf("target dir is required")
	}

	resolvedID := resolveTemplateID(templateID)
	manifestPath := path.Join("templates", resolvedID, "template.json")
	if _, err := LoadManifest(resolvedID); err != nil {
		return err
	}

	sourceRoot := path.Join("templates", resolvedID)
	if err := copyEmbeddedDir(sourceRoot, targetDir); err != nil {
		return err
	}

	metaDir := filepath.Join(targetDir, ".aitigo")
	if err := os.MkdirAll(metaDir, 0o755); err != nil {
		return err
	}
	manifestBytes, err := fs.ReadFile(embeddedTemplates, manifestPath)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(metaDir, "template.json"), manifestBytes, 0o644)
}

func copyEmbeddedDir(sourceRoot, targetDir string) error {
	return fs.WalkDir(embeddedTemplates, sourceRoot, func(entryPath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entryPath == sourceRoot {
			return nil
		}

		rel := strings.TrimPrefix(entryPath, sourceRoot+"/")
		if rel == entryPath {
			return fmt.Errorf("unexpected template path: %s", entryPath)
		}
		rel = normalizeTemplatePath(rel)
		targetPath := filepath.Join(targetDir, filepath.FromSlash(rel))

		if d.IsDir() {
			return os.MkdirAll(targetPath, 0o755)
		}

		data, err := fs.ReadFile(embeddedTemplates, entryPath)
		if err != nil {
			return err
		}
		if err := os.MkdirAll(filepath.Dir(targetPath), 0o755); err != nil {
			return err
		}
		return os.WriteFile(targetPath, data, 0o644)
	})
}

func normalizeTemplatePath(rel string) string {
	if strings.HasSuffix(rel, ".go.txt") {
		return strings.TrimSuffix(rel, ".txt")
	}
	if strings.HasSuffix(rel, "go.mod.txt") || strings.HasSuffix(rel, "go.sum.txt") {
		return strings.TrimSuffix(rel, ".txt")
	}
	return rel
}
