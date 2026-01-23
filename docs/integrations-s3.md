# S3 / MinIO Integration

Package: `pkg/integrations/s3`

Env:

- `S3_ENDPOINT`
- `S3_ACCESS_KEY`
- `S3_SECRET_KEY`
- `S3_REGION`
- `S3_BUCKET`
- `S3_USE_SSL`

Usage:

```go
cfg, _ := s3.FromEnv()
client, _ := s3.New(cfg)
_ = client.Put(ctx, "file.txt", reader, size, "text/plain")
```
