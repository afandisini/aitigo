package cli

import (
	"fmt"
	"os"
	"path/filepath"
)

func generateModule(module string) error {
	return generateModuleWithMode(module, false, false)
}

func generateController(name, module string, force bool) error {
	return generateControllerWithMode(name, module, force, false)
}

func generateService(name, module string, force bool) error {
	return generateServiceWithMode(name, module, force, false)
}

func generateRepository(name, module string, force bool) error {
	return generateRepositoryWithMode(name, module, force, false)
}

func generateModuleWithMode(module string, force bool, skipIfExists bool) error {
	for _, f := range moduleFiles(module) {
		status, err := writeFileWithMode(f.path, f.content, force, skipIfExists)
		if err != nil {
			return err
		}
		if err := logWrite(f.path, status); err != nil {
			return err
		}
	}
	return nil
}

func generateControllerWithMode(name, module string, force bool, skipIfExists bool) error {
	dir := filepath.Join("internal", "app", "http", "controller")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	fileName := toSnake(name) + ".go"
	path := filepath.Join(dir, fileName)
	content := templateController(toPascal(name), module)
	status, err := writeFileWithMode(path, content, force, skipIfExists)
	if err != nil {
		return err
	}
	if err := logWrite(path, status); err != nil {
		return err
	}
	return nil
}

func generateServiceWithMode(name, module string, force bool, skipIfExists bool) error {
	dir := filepath.Join("internal", "domain", module)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	fileName := toSnake(name) + ".go"
	path := filepath.Join(dir, fileName)
	content := templateService(module, toPascal(name))
	status, err := writeFileWithMode(path, content, force, skipIfExists)
	if err != nil {
		return err
	}
	if err := logWrite(path, status); err != nil {
		return err
	}
	return nil
}

func generateRepositoryWithMode(name, module string, force bool, skipIfExists bool) error {
	domainDir := filepath.Join("internal", "domain", module)
	if err := os.MkdirAll(domainDir, 0o755); err != nil {
		return err
	}
	infraDir := filepath.Join("internal", "infra", "repository")
	if err := os.MkdirAll(infraDir, 0o755); err != nil {
		return err
	}

	snakeName := toSnake(name)
	pascalName := toPascal(name)
	interfacePath := filepath.Join(domainDir, snakeName+".go")
	implPath := filepath.Join(infraDir, snakeName+"_impl.go")

	status, err := writeFileWithMode(interfacePath, templateRepositoryInterface(module, pascalName), force, skipIfExists)
	if err != nil {
		return err
	}
	if err := logWrite(interfacePath, status); err != nil {
		return err
	}

	status, err = writeFileWithMode(implPath, templateRepositoryImpl(pascalName+"Impl"), force, skipIfExists)
	if err != nil {
		return err
	}
	if err := logWrite(implPath, status); err != nil {
		return err
	}
	return nil
}

func writeFileWithMode(path, content string, force bool, skipIfExists bool) (string, error) {
	exists := false
	if _, err := os.Stat(path); err == nil {
		exists = true
	} else if !os.IsNotExist(err) {
		return "", err
	}

	if exists && !force {
		if skipIfExists {
			return "skip", nil
		}
		return "", fmt.Errorf("file exists: %s (use --force to overwrite)", path)
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return "", err
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		return "", err
	}
	if exists {
		return "overwrite", nil
	}
	return "create", nil
}

type fileSpec struct {
	path    string
	content string
}

func moduleFiles(module string) []fileSpec {
	domainDir := filepath.Join("internal", "domain", module)
	controllerDir := filepath.Join("internal", "app", "http", "controller")
	moduleName := toPascal(module)

	return []fileSpec{
		{filepath.Join(domainDir, "entity.go"), templateEntity(module, moduleName)},
		{filepath.Join(domainDir, "repository.go"), templateRepositoryInterface(module, "Repository")},
		{filepath.Join(domainDir, "service.go"), templateService(module, "Service")},
		{filepath.Join(controllerDir, module+"_controller.go"), templateController(moduleName+"Controller", module)},
	}
}

func logWrite(path, status string) error {
	switch status {
	case "overwrite":
		return logOverwrite(path)
	case "skip":
		return logSkip(path)
	default:
		return logCreate(path)
	}
}
