package main

import (
	"context"
	"log"
	"strings"
	"time"

	"aitigo/pkg/integrations/s3"
)

func main() {
	cfg, err := s3.FromEnv()
	if err != nil {
		log.Fatal(err)
	}

	client, err := s3.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Put(context.Background(), "hello.txt", strings.NewReader("hello"), int64(len("hello")), "text/plain"); err != nil {
		log.Fatal(err)
	}

	url, err := client.PresignGet(context.Background(), "hello.txt", 5*time.Minute)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("presigned url:", url.String())
}
