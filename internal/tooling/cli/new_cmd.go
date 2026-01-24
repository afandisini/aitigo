package cli

import (
	"flag"
	"fmt"
	"io"
	"os"

	"aitigo/internal/tooling/templates"
)

func runNew(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: aitigo new <name> --template <templateId>")
	}

	name := args[0]
	fs := flag.NewFlagSet("aitigo new", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	templateID := fs.String("template", "", "template id")
	if err := fs.Parse(args[1:]); err != nil {
		return err
	}
	if *templateID == "" {
		return fmt.Errorf("--template is required")
	}

	if err := templates.Install(*templateID, name); err != nil {
		return err
	}
	_, err := fmt.Fprintf(os.Stdout, "new %s -> %s\n", *templateID, name)
	return err
}
