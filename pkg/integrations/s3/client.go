package s3

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Client struct {
	client *minio.Client
	bucket string
}

func New(cfg Config) (*Client, error) {
	if cfg.Endpoint == "" {
		return nil, fmt.Errorf("S3 endpoint is required")
	}
	if cfg.AccessKey == "" || cfg.SecretKey == "" {
		return nil, fmt.Errorf("S3 credentials are required")
	}
	if cfg.Bucket == "" {
		return nil, fmt.Errorf("S3 bucket is required")
	}

	minioClient, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure: cfg.UseSSL,
		Region: cfg.Region,
	})
	if err != nil {
		return nil, err
	}

	return &Client{client: minioClient, bucket: cfg.Bucket}, nil
}

func (c *Client) Put(ctx context.Context, objectName string, reader io.Reader, size int64, contentType string) error {
	if c == nil || c.client == nil {
		return fmt.Errorf("s3 client is nil")
	}
	_, err := c.client.PutObject(ctx, c.bucket, objectName, reader, size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	return err
}

func (c *Client) Get(ctx context.Context, objectName string) (io.ReadCloser, error) {
	if c == nil || c.client == nil {
		return nil, fmt.Errorf("s3 client is nil")
	}
	object, err := c.client.GetObject(ctx, c.bucket, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return object, nil
}

func (c *Client) PresignGet(ctx context.Context, objectName string, expiry time.Duration) (*url.URL, error) {
	if c == nil || c.client == nil {
		return nil, fmt.Errorf("s3 client is nil")
	}
	return c.client.PresignedGetObject(ctx, c.bucket, objectName, expiry, nil)
}
