# Migrations

Package: `pkg/db/migrate`

Format:

- `0001_name.up.sql`
- `0001_name.down.sql`

CLI:

```sh
aitigo migrate create add_users_table
aitigo migrate up
aitigo migrate status
aitigo migrate down
```

Programmatic usage:

```go
db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
runner := migrate.NewRunner(db, "migrations")
_, _ = runner.Up(context.Background())
```

Notes:

- Default driver is `postgres` and SQL placeholders use `$1` style.
