package cli

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func makeCrud(args []string) error {
	if len(args) < 1 {
		printCrudHelp(os.Stderr)
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
		printCrudHelp(os.Stderr)
		return fmt.Errorf("module is required")
	}

	return generateCrud(module, *force)
}

func generateCrud(module string, force bool) error {
	moduleName := toPascal(module)
	if err := generateModuleWithMode(module, force, true); err != nil {
		return err
	}
	if err := generateServiceWithMode(moduleName+"Service", module, force, true); err != nil {
		return err
	}
	if err := generateRepositoryWithMode(moduleName+"Repository", module, force, true); err != nil {
		return err
	}
	return nil
}

func printCrudHelp(w io.Writer) {
	fmt.Fprintln(w, "Usage:")
	fmt.Fprintln(w, "  aitigo make:crud <module> [--force]")
}
