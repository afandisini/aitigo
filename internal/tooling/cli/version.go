package cli

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime/debug"
)

var (
	Version   = "dev"
	Commit    = "none"
	BuildTime = "unknown"
)

func runVersion(args []string) error {
	fs := flag.NewFlagSet("aitigo version", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	help := fs.Bool("help", false, "show help")
	helpShort := fs.Bool("h", false, "show help")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *help || *helpShort {
		return printVersionHelp(os.Stdout)
	}

	version, commit, buildTime := resolveBuildInfo()
	_, err := fmt.Fprintf(os.Stdout, "version=%s commit=%s build_time=%s\n", version, commit, buildTime)
	return err
}

func resolveBuildInfo() (string, string, string) {
	version := Version
	commit := Commit
	buildTime := BuildTime

	if info, ok := debug.ReadBuildInfo(); ok {
		if version == "" || version == "dev" {
			if info.Main.Version != "" {
				version = info.Main.Version
			}
		}
		for _, setting := range info.Settings {
			switch setting.Key {
			case "vcs.revision":
				if commit == "" || commit == "none" {
					commit = setting.Value
				}
			case "vcs.time":
				if buildTime == "" || buildTime == "unknown" {
					buildTime = setting.Value
				}
			}
		}
	}

	return version, commit, buildTime
}

func printVersionHelp(w io.Writer) error {
	_, err := fmt.Fprint(w, versionHelpText)
	return err
}

const versionHelpText = `AitiGo version

Usage:
  aitigo version

Flags:
  -h, --help  show help
`
