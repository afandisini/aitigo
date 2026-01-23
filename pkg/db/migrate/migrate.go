package migrate

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"
)

const schemaTable = "schema_migrations"

type Migration struct {
	Version  int
	Name     string
	UpPath   string
	DownPath string
}

type Status struct {
	Migration Migration
	Applied   bool
	AppliedAt time.Time
}

type Runner struct {
	db  *sql.DB
	dir string
}

func NewRunner(db *sql.DB, dir string) *Runner {
	return &Runner{db: db, dir: dir}
}

func Load(dir string) ([]Migration, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	byVersion := make(map[int]*Migration)
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		parsed, ok := parseFilename(entry.Name())
		if !ok {
			continue
		}

		mig := byVersion[parsed.version]
		if mig == nil {
			mig = &Migration{Version: parsed.version, Name: parsed.name}
			byVersion[parsed.version] = mig
		}
		path := filepath.Join(dir, entry.Name())
		if parsed.direction == directionUp {
			mig.UpPath = path
		} else {
			mig.DownPath = path
		}
	}

	migrations := make([]Migration, 0, len(byVersion))
	for _, mig := range byVersion {
		if mig.UpPath == "" || mig.DownPath == "" {
			return nil, fmt.Errorf("migration %d missing up or down file", mig.Version)
		}
		migrations = append(migrations, *mig)
	}

	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Version < migrations[j].Version
	})
	return migrations, nil
}

func (r *Runner) Up(ctx context.Context) (int, error) {
	if err := r.ensureTable(ctx); err != nil {
		return 0, err
	}
	migrations, err := Load(r.dir)
	if err != nil {
		return 0, err
	}
	applied, err := r.appliedVersions(ctx)
	if err != nil {
		return 0, err
	}

	count := 0
	for _, mig := range migrations {
		if _, ok := applied[mig.Version]; ok {
			continue
		}
		if err := r.applyMigration(ctx, mig, directionUp); err != nil {
			return count, err
		}
		count++
	}
	return count, nil
}

func (r *Runner) Down(ctx context.Context) (int, error) {
	if err := r.ensureTable(ctx); err != nil {
		return 0, err
	}
	migrations, err := Load(r.dir)
	if err != nil {
		return 0, err
	}

	applied, err := r.appliedVersions(ctx)
	if err != nil {
		return 0, err
	}

	var latest *Migration
	for i := len(migrations) - 1; i >= 0; i-- {
		if _, ok := applied[migrations[i].Version]; ok {
			latest = &migrations[i]
			break
		}
	}
	if latest == nil {
		return 0, nil
	}

	if err := r.applyMigration(ctx, *latest, directionDown); err != nil {
		return 0, err
	}
	return 1, nil
}

func (r *Runner) Status(ctx context.Context) ([]Status, error) {
	if err := r.ensureTable(ctx); err != nil {
		return nil, err
	}
	migrations, err := Load(r.dir)
	if err != nil {
		return nil, err
	}
	applied, err := r.appliedVersions(ctx)
	if err != nil {
		return nil, err
	}

	statuses := make([]Status, 0, len(migrations))
	for _, mig := range migrations {
		status := Status{Migration: mig}
		if appliedAt, ok := applied[mig.Version]; ok {
			status.Applied = true
			status.AppliedAt = appliedAt
		}
		statuses = append(statuses, status)
	}
	return statuses, nil
}

func (r *Runner) ensureTable(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (version INTEGER PRIMARY KEY, name TEXT NOT NULL, applied_at TIMESTAMP NOT NULL)`, schemaTable))
	return err
}

func (r *Runner) appliedVersions(ctx context.Context) (map[int]time.Time, error) {
	rows, err := r.db.QueryContext(ctx, fmt.Sprintf(`SELECT version, applied_at FROM %s`, schemaTable))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	applied := make(map[int]time.Time)
	for rows.Next() {
		var version int
		var appliedAt time.Time
		if err := rows.Scan(&version, &appliedAt); err != nil {
			return nil, err
		}
		applied[version] = appliedAt
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return applied, nil
}

func (r *Runner) applyMigration(ctx context.Context, mig Migration, direction direction) error {
	path := mig.UpPath
	if direction == directionDown {
		path = mig.DownPath
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read migration %s: %w", path, err)
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx, string(content)); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("execute migration %s: %w", path, err)
	}

	if direction == directionUp {
		_, err = tx.ExecContext(ctx, fmt.Sprintf(`INSERT INTO %s (version, name, applied_at) VALUES ($1, $2, $3)`, schemaTable), mig.Version, mig.Name, time.Now().UTC())
	} else {
		_, err = tx.ExecContext(ctx, fmt.Sprintf(`DELETE FROM %s WHERE version = $1`, schemaTable), mig.Version)
	}
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
