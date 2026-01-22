package checker

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type Runner struct{}

func (Runner) Run(command, dir string) error {
	if command == "" {
		return fmt.Errorf("command is required")
	}

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd.exe", "/C", command)
	} else {
		cmd = exec.Command("sh", "-lc", command)
	}
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
