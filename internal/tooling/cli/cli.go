package cli

import (
	"fmt"
	"os"
)

// Run executes the CLI command with the provided arguments.
func Run(args []string) error {
	if len(args) < 2 {
		if err := printHelp(os.Stdout); err != nil {
			return err
		}
		return nil
	}

	cmd := args[1]
	if cmd == "help" || cmd == "-h" || cmd == "--help" {
		if err := printHelp(os.Stdout); err != nil {
			return err
		}
		return nil
	}

	handler, ok := commands()[cmd]
	if !ok {
		if _, err := fmt.Fprintf(os.Stderr, "unknown command: %s\n", cmd); err != nil {
			return err
		}
		if err := printHelp(os.Stderr); err != nil {
			return err
		}
		return fmt.Errorf("unknown command")
	}

	if err := handler(args[2:]); err != nil {
		if _, printErr := fmt.Fprintln(os.Stderr, err); printErr != nil {
			return printErr
		}
		return err
	}
	return nil
}
