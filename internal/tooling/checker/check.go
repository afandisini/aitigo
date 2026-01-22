package checker

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"aitigo/internal/tooling/templates"
)

func RunChecks(projectDir string) error {
	manifestPath := filepath.Join(projectDir, ".aitigo", "template.json")
	manifestBytes, err := os.ReadFile(manifestPath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("template manifest not found at %s", manifestPath)
		}
		return err
	}

	var manifest templates.Manifest
	if err := json.Unmarshal(manifestBytes, &manifest); err != nil {
		return err
	}
	if len(manifest.Check.Commands) == 0 {
		return fmt.Errorf("no check commands defined in %s", manifestPath)
	}

	runner := Runner{}
	for _, cmd := range manifest.Check.Commands {
		if _, err := fmt.Fprintf(os.Stdout, "check %s\n", cmd.Name); err != nil {
			return err
		}
		if err := runner.Run(cmd.Run, projectDir); err != nil {
			if cmd.Optional {
				if _, warnErr := fmt.Fprintf(os.Stdout, "warn %s failed (optional)\n", cmd.Name); warnErr != nil {
					return warnErr
				}
				continue
			}
			return fmt.Errorf("%s failed: %w", cmd.Name, err)
		}
	}
	return nil
}
