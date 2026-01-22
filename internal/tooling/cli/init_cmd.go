package cli

import (
	"fmt"
	"os"

	"aitigo/internal/tooling/templates"
)

func runInit(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: aitigo init <templateId> <dir>")
	}
	if len(args) > 2 {
		return fmt.Errorf("init accepts only <templateId> <dir>")
	}

	templateID := args[0]
	targetDir := args[1]

	if err := templates.Install(templateID, targetDir); err != nil {
		return err
	}
	_, err := fmt.Fprintf(os.Stdout, "init %s -> %s\n", templateID, targetDir)
	return err
}
