package s3

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	Region    string
	Bucket    string
	UseSSL    bool
}

func FromEnv() (Config, error) {
	cfg := Config{
		Endpoint:  os.Getenv("S3_ENDPOINT"),
		AccessKey: os.Getenv("S3_ACCESS_KEY"),
		SecretKey: os.Getenv("S3_SECRET_KEY"),
		Region:    os.Getenv("S3_REGION"),
		Bucket:    os.Getenv("S3_BUCKET"),
	}

	if raw := os.Getenv("S3_USE_SSL"); raw != "" {
		value, err := strconv.ParseBool(raw)
		if err != nil {
			return Config{}, fmt.Errorf("invalid S3_USE_SSL: %w", err)
		}
		cfg.UseSSL = value
	}

	return cfg, nil
}
