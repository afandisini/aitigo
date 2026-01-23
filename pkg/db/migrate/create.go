package migrate

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func Create(dir, name string) (Migration, error) {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return Migration{}, err
	}

	migrations, err := Load(dir)
	if err != nil && !os.IsNotExist(err) {
		return Migration{}, err
	}

	nextVersion := 1
	if len(migrations) > 0 {
		nextVersion = migrations[len(migrations)-1].Version + 1
	}

	slug := slugify(name)
	if slug == "" {
		return Migration{}, fmt.Errorf("invalid migration name")
	}

	base := fmt.Sprintf("%04d_%s", nextVersion, slug)
	upPath := filepath.Join(dir, base+".up.sql")
	downPath := filepath.Join(dir, base+".down.sql")

	if err := os.WriteFile(upPath, []byte("-- up\n"), 0o644); err != nil {
		return Migration{}, err
	}
	if err := os.WriteFile(downPath, []byte("-- down\n"), 0o644); err != nil {
		return Migration{}, err
	}

	return Migration{
		Version:  nextVersion,
		Name:     slug,
		UpPath:   upPath,
		DownPath: downPath,
	}, nil
}

func slugify(input string) string {
	input = strings.ToLower(strings.TrimSpace(input))
	input = strings.ReplaceAll(input, " ", "_")
	re := regexp.MustCompile(`[^a-z0-9_-]`)
	return re.ReplaceAllString(input, "")
}
