package cli

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type commandHandler func(args []string) error

func commands() map[string]commandHandler {
	return map[string]commandHandler{
		"check":           runCheck,
		"make:crud":       makeCrud,
		"make:module":     makeModule,
		"make:controller": makeController,
		"make:service":    makeService,
		"make:repository": makeRepository,
	}
}

func printHelp(w io.Writer) {
	fmt.Fprintln(w, "AitiGo CLI")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Usage:")
	fmt.Fprintln(w, "  aitigo help")
	fmt.Fprintln(w, "  aitigo check")
	fmt.Fprintln(w, "  aitigo make:crud <module> [--force]")
	fmt.Fprintln(w, "  aitigo make:module <module>")
	fmt.Fprintln(w, "  aitigo make:controller <Name> --module <module> [--force]")
	fmt.Fprintln(w, "  aitigo make:service <Name> --module <module> [--force]")
	fmt.Fprintln(w, "  aitigo make:repository <Name> --module <module> [--force]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Examples:")
	fmt.Fprintln(w, "  aitigo check")
	fmt.Fprintln(w, "  aitigo make:crud user")
	fmt.Fprintln(w, "  aitigo make:crud user --force")
	fmt.Fprintln(w, "  aitigo make:module user")
	fmt.Fprintln(w, "  aitigo make:controller UserController --module user")
	fmt.Fprintln(w, "  aitigo make:service UserService --module user --force")
	fmt.Fprintln(w, "  aitigo make:repository UserRepository --module user")
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

func logCreate(path string) {
	fmt.Fprintf(os.Stdout, "create %s\n", path)
}

func logOverwrite(path string) {
	fmt.Fprintf(os.Stdout, "overwrite %s\n", path)
}

func logSkip(path string) {
	fmt.Fprintf(os.Stdout, "skip %s (exists)\n", path)
}
