package main

import (
	"os"

	"aitigo/internal/tooling/cli"
)

func main() {
	if err := cli.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
