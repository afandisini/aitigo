package cli

import (
	"fmt"

	"aitigo/internal/tooling/checker"
)

func runCheck(args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("usage: aitigo check [dir]")
	}

	dir := "."
	if len(args) == 1 {
		dir = args[0]
	}
	return checker.RunChecks(dir)
}
