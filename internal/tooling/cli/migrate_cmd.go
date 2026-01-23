package cli

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"io"
	"os"

	dbconfig "aitigo/pkg/db/config"
	"aitigo/pkg/db/migrate"
)

func runMigrate(args []string) error {
	fs := flag.NewFlagSet("aitigo migrate", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	help := fs.Bool("help", false, "show help")
	helpShort := fs.Bool("h", false, "show help")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *help || *helpShort || fs.NArg() == 0 {
		return printMigrateHelp(os.Stdout)
	}

	switch fs.Arg(0) {
	case "up":
		return runMigrateUp(fs.Args()[1:])
	case "down":
		return runMigrateDown(fs.Args()[1:])
	case "status":
		return runMigrateStatus(fs.Args()[1:])
	case "create":
		return runMigrateCreate(fs.Args()[1:])
	default:
		_ = printMigrateHelp(os.Stderr)
		return fmt.Errorf("unknown migrate subcommand: %s", fs.Arg(0))
	}
}

func runMigrateUp(args []string) error {
	flags, err := parseMigrateFlags("up", args)
	if err != nil {
		return err
	}
	if flags.showHelp {
		return printMigrateUpHelp(os.Stdout)
	}
	cfg := flags.config
	dir := flags.dir
	db, err := sql.Open(cfg.Driver, cfg.DSN)
	if err != nil {
		return err
	}
	defer db.Close()

	runner := migrate.NewRunner(db, dir)
	count, err := runner.Up(context.Background())
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(os.Stdout, "applied %d migrations\n", count)
	return err
}

func runMigrateDown(args []string) error {
	flags, err := parseMigrateFlags("down", args)
	if err != nil {
		return err
	}
	if flags.showHelp {
		return printMigrateDownHelp(os.Stdout)
	}
	cfg := flags.config
	dir := flags.dir
	db, err := sql.Open(cfg.Driver, cfg.DSN)
	if err != nil {
		return err
	}
	defer db.Close()

	runner := migrate.NewRunner(db, dir)
	count, err := runner.Down(context.Background())
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(os.Stdout, "rolled back %d migrations\n", count)
	return err
}

func runMigrateStatus(args []string) error {
	flags, err := parseMigrateFlags("status", args)
	if err != nil {
		return err
	}
	if flags.showHelp {
		return printMigrateStatusHelp(os.Stdout)
	}
	cfg := flags.config
	dir := flags.dir
	db, err := sql.Open(cfg.Driver, cfg.DSN)
	if err != nil {
		return err
	}
	defer db.Close()

	runner := migrate.NewRunner(db, dir)
	statuses, err := runner.Status(context.Background())
	if err != nil {
		return err
	}
	return printStatus(os.Stdout, statuses)
}

func runMigrateCreate(args []string) error {
	fs := flag.NewFlagSet("aitigo migrate create", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	help := fs.Bool("help", false, "show help")
	helpShort := fs.Bool("h", false, "show help")
	dir := fs.String("dir", "migrations", "migrations directory")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *help || *helpShort {
		return printMigrateCreateHelp(os.Stdout)
	}
	if fs.NArg() < 1 {
		return fmt.Errorf("migration name is required")
	}

	mig, err := migrate.Create(*dir, fs.Arg(0))
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(os.Stdout, "create %s\ncreate %s\n", mig.UpPath, mig.DownPath)
	return err
}

type migrateFlags struct {
	config   dbconfig.Config
	dir      string
	showHelp bool
}

func parseMigrateFlags(command string, args []string) (migrateFlags, error) {
	fs := flag.NewFlagSet("aitigo migrate "+command, flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	help := fs.Bool("help", false, "show help")
	helpShort := fs.Bool("h", false, "show help")
	dir := fs.String("dir", "migrations", "migrations directory")
	dsn := fs.String("dsn", "", "database dsn")
	driver := fs.String("driver", "", "database driver")
	if err := fs.Parse(args); err != nil {
		return migrateFlags{}, err
	}
	if *help || *helpShort {
		return migrateFlags{dir: *dir, showHelp: true}, nil
	}

	cfg, err := dbconfig.FromEnv()
	if err != nil && *dsn == "" {
		return migrateFlags{}, err
	}
	if *dsn != "" {
		cfg.DSN = *dsn
	}
	if *driver != "" {
		cfg.Driver = *driver
	}

	return migrateFlags{config: cfg, dir: *dir}, nil
}

func printStatus(w io.Writer, statuses []migrate.Status) error {
	for _, status := range statuses {
		line := "pending"
		if status.Applied {
			line = "applied"
		}
		if _, err := fmt.Fprintf(w, "%s %04d_%s\n", line, status.Migration.Version, status.Migration.Name); err != nil {
			return err
		}
	}
	return nil
}

func printMigrateHelp(w io.Writer) error {
	_, err := fmt.Fprint(w, migrateHelpText)
	return err
}

func printMigrateUpHelp(w io.Writer) error {
	_, err := fmt.Fprint(w, migrateUpHelpText)
	return err
}

func printMigrateDownHelp(w io.Writer) error {
	_, err := fmt.Fprint(w, migrateDownHelpText)
	return err
}

func printMigrateStatusHelp(w io.Writer) error {
	_, err := fmt.Fprint(w, migrateStatusHelpText)
	return err
}

func printMigrateCreateHelp(w io.Writer) error {
	_, err := fmt.Fprint(w, migrateCreateHelpText)
	return err
}

const migrateHelpText = `AitiGo migrate

Usage:
  aitigo migrate up [--dir dir] [--dsn dsn] [--driver driver]
  aitigo migrate down [--dir dir] [--dsn dsn] [--driver driver]
  aitigo migrate status [--dir dir] [--dsn dsn] [--driver driver]
  aitigo migrate create <name> [--dir dir]

Flags:
  -h, --help  show help
`

const migrateUpHelpText = `AitiGo migrate up

Usage:
  aitigo migrate up [--dir dir] [--dsn dsn] [--driver driver]

Flags:
  --dir      migrations directory (default "migrations")
  --dsn      database dsn (DATABASE_URL or DB_DSN)
  --driver   database driver (default "postgres")
  -h, --help show help
`

const migrateDownHelpText = `AitiGo migrate down

Usage:
  aitigo migrate down [--dir dir] [--dsn dsn] [--driver driver]

Flags:
  --dir      migrations directory (default "migrations")
  --dsn      database dsn (DATABASE_URL or DB_DSN)
  --driver   database driver (default "postgres")
  -h, --help show help
`

const migrateStatusHelpText = `AitiGo migrate status

Usage:
  aitigo migrate status [--dir dir] [--dsn dsn] [--driver driver]

Flags:
  --dir      migrations directory (default "migrations")
  --dsn      database dsn (DATABASE_URL or DB_DSN)
  --driver   database driver (default "postgres")
  -h, --help show help
`

const migrateCreateHelpText = `AitiGo migrate create

Usage:
  aitigo migrate create <name> [--dir dir]

Flags:
  --dir      migrations directory (default "migrations")
  -h, --help show help
`
