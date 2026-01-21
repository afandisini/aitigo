package cli

import (
	"fmt"
	"os"
)

func Run(args []string) error {
	if len(args) < 2 {
		printHelp(os.Stdout)
		return nil
	}

	cmd := args[1]
	if cmd == "help" || cmd == "-h" || cmd == "--help" {
		printHelp(os.Stdout)
		return nil
	}

	handler, ok := commands()[cmd]
	if !ok {
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", cmd)
		printHelp(os.Stderr)
		return fmt.Errorf("unknown command")
	}

	if err := handler(args[2:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	return nil
}
