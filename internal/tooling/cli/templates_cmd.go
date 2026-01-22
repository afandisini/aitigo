package cli

import (
	"fmt"
	"os"

	"aitigo/internal/tooling/templates"
)

func runTemplates(args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("templates does not accept arguments")
	}

	manifests, err := templates.List()
	if err != nil {
		return err
	}

	for _, manifest := range manifests {
		if _, err := fmt.Fprintf(os.Stdout, "%s - %s\n", manifest.ID, manifest.Description); err != nil {
			return err
		}
	}
	return nil
}
