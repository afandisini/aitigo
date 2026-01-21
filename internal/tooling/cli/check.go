package cli

import (
	"fmt"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func runCheck(args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("check does not accept arguments")
	}

	var violations []string

	domainForbidden := []string{
		"aitigo/internal/infra",
		"internal/infra",
		"aitigo/internal/app",
		"internal/app",
		"net/http",
		"github.com/gin-gonic",
		"github.com/labstack/echo",
		"github.com/gofiber",
		"gorm.io/gorm",
		"database/sql",
		"github.com/jmoiron/sqlx",
	}

	controllerForbidden := []string{
		"aitigo/internal/infra",
		"internal/infra",
	}

	domainPath := filepath.Join("internal", "domain")
	if _, err := os.Stat(domainPath); err == nil {
		v, err := scanImports(domainPath, domainForbidden)
		if err != nil {
			return err
		}
		violations = append(violations, v...)
	}

	controllerPath := filepath.Join("internal", "app", "http", "controller")
	if _, err := os.Stat(controllerPath); err == nil {
		v, err := scanImports(controllerPath, controllerForbidden)
		if err != nil {
			return err
		}
		violations = append(violations, v...)
	}

	if len(violations) == 0 {
		if _, err := fmt.Fprintln(os.Stdout, "AitiGo check: OK"); err != nil {
			return err
		}
		return nil
	}

	for _, v := range violations {
		if _, err := fmt.Fprintln(os.Stdout, v); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprintf(os.Stdout, "AitiGo check: %d violation(s)\n", len(violations)); err != nil {
		return err
	}
	return fmt.Errorf("boundary check failed")
}

func scanImports(root string, forbidden []string) ([]string, error) {
	var violations []string

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			name := d.Name()
			if name == "vendor" || name == ".git" || name == "node_modules" {
				return filepath.SkipDir
			}
			return nil
		}
		if filepath.Ext(path) != ".go" {
			return nil
		}

		fileSet := token.NewFileSet()
		file, parseErr := parser.ParseFile(fileSet, path, nil, parser.ImportsOnly)
		if parseErr != nil {
			return parseErr
		}

		for _, spec := range file.Imports {
			if spec.Path == nil {
				continue
			}
			importPath, unquoteErr := strconv.Unquote(spec.Path.Value)
			if unquoteErr != nil {
				continue
			}
			if hit := matchForbidden(importPath, forbidden); hit != "" {
				violations = append(violations, fmt.Sprintf("%s: forbidden import %q", toSlash(path), importPath))
			}
		}
		return nil
	})

	return violations, err
}

func matchForbidden(importPath string, forbidden []string) string {
	for _, rule := range forbidden {
		if strings.Contains(importPath, rule) {
			return rule
		}
	}
	return ""
}

func toSlash(path string) string {
	return filepath.ToSlash(path)
}
