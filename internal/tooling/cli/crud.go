package cli

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func makeCrud(args []string) error {
	if len(args) < 1 {
		if err := printCrudHelp(os.Stderr); err != nil {
			return err
		}
		return fmt.Errorf("module is required")
	}
	moduleArg := args[0]

	fs := flag.NewFlagSet("make:crud", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	force := fs.Bool("force", false, "overwrite files")
	if err := fs.Parse(args[1:]); err != nil {
		return err
	}

	module := normalizeModule(moduleArg)
	if module == "" {
		if err := printCrudHelp(os.Stderr); err != nil {
			return err
		}
		return fmt.Errorf("module is required")
	}

	return generateCrud(module, *force)
}

func generateCrud(module string, force bool) error {
	if err := generateModuleWithMode(module, force, true); err != nil {
		return err
	}
	if err := generateServiceWithMode("Service", module, force, true); err != nil {
		return err
	}
	if err := generateRepositoryWithMode("Repository", module, force, true); err != nil {
		return err
	}
	return nil
}

func printCrudHelp(w io.Writer) error {
	_, err := fmt.Fprintln(w, "Usage:")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(w, "  aitigo make:crud <module> [--force]")
	return err
}
