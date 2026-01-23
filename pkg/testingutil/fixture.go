package testingutil

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadJSONFixture(path string, target interface{}) error {
	payload, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read fixture: %w", err)
	}
	if err := json.Unmarshal(payload, target); err != nil {
		return fmt.Errorf("unmarshal fixture: %w", err)
	}
	return nil
}
