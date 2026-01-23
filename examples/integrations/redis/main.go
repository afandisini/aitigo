package main

import (
	"context"
	"log"

	"aitigo/pkg/integrations/redis"
)

func main() {
	cfg, err := redis.FromEnv()
	if err != nil {
		log.Fatal(err)
	}

	client := redis.NewClient(cfg)
	if err := redis.Ping(context.Background(), client); err != nil {
		log.Fatal(err)
	}
	log.Println("redis ok")
}
