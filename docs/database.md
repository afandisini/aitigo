# Database Config

AitiGo reads DB config from environment variables:

- `DATABASE_URL` or `DB_DSN` (required)
- `DB_DRIVER` (optional, default `postgres`)

SSL mode guidance (Postgres):

- Local dev on private networks: `sslmode=disable`
- Cloud or shared networks: `sslmode=require` or stronger
