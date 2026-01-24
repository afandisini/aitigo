package cli

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type commandHandler func(args []string) error

const helpText = `AitiGo CLI

Usage:
  aitigo help
  aitigo check [dir]
  aitigo check:boundary
  aitigo migrate up|down|status [--dir dir] [--dsn dsn] [--driver driver]
  aitigo migrate create <name> [--dir dir]
  aitigo version
  aitigo make:crud <module> [--force]
  aitigo make:module <module>
  aitigo make:controller <Name> --module <module> [--force]
  aitigo make:service <Name> --module <module> [--force]
  aitigo make:repository <Name> --module <module> [--force]
  aitigo templates
  aitigo init <templateId> <dir>
  aitigo new <name> --template <templateId>

Examples:
  aitigo check
  aitigo check .
  aitigo check:boundary
  aitigo migrate up
  aitigo migrate status --dir ./migrations
  aitigo migrate create add_users_table
  aitigo version
  aitigo make:crud user
  aitigo make:crud user --force
  aitigo make:module user
  aitigo make:controller UserController --module user
  aitigo make:service UserService --module user --force
  aitigo make:repository UserRepository --module user
  aitigo templates
  aitigo init next-ts ./my-app
  aitigo new my-api --template gin-basic
`

func commands() map[string]commandHandler {
	return map[string]commandHandler{
		"check":           runCheck,
		"check:boundary":  runBoundaryCheck,
		"migrate":         runMigrate,
		"version":         runVersion,
		"make:crud":       makeCrud,
		"make:module":     makeModule,
		"make:controller": makeController,
		"make:service":    makeService,
		"make:repository": makeRepository,
		"templates":       runTemplates,
		"init":            runInit,
		"new":             runNew,
	}
}

func printHelp(w io.Writer) error {
	_, err := fmt.Fprint(w, helpText)
	return err
}

func makeModule(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("module name is required")
	}
	module := normalizeModule(args[0])
	return generateModule(module)
}

func makeController(args []string) error {
	name, module, force, err := parseNamedCommand(args)
	if err != nil {
		return err
	}
	return generateController(name, module, force)
}

func makeService(args []string) error {
	name, module, force, err := parseNamedCommand(args)
	if err != nil {
		return err
	}
	return generateService(name, module, force)
}

func makeRepository(args []string) error {
	name, module, force, err := parseNamedCommand(args)
	if err != nil {
		return err
	}
	return generateRepository(name, module, force)
}

func parseNamedCommand(args []string) (string, string, bool, error) {
	if len(args) < 1 {
		return "", "", false, fmt.Errorf("name is required")
	}

	name := args[0]
	fs := flag.NewFlagSet("aitigo", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	module := fs.String("module", "", "module name")
	force := fs.Bool("force", false, "overwrite files")
	if err := fs.Parse(args[1:]); err != nil {
		return "", "", false, err
	}
	if *module == "" {
		return "", "", false, fmt.Errorf("--module is required")
	}

	return name, normalizeModule(*module), *force, nil
}

func logCreate(path string) error {
	_, err := fmt.Fprintf(os.Stdout, "create %s\n", path)
	return err
}

func logOverwrite(path string) error {
	_, err := fmt.Fprintf(os.Stdout, "overwrite %s\n", path)
	return err
}

func logSkip(path string) error {
	_, err := fmt.Fprintf(os.Stdout, "skip %s (exists)\n", path)
	return err
}
