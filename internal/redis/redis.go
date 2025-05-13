package redis

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client
var Ctx = context.Background()

func InitRedis() {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}

	Client = redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0,
	})

	if err := Client.Ping(Ctx).Err(); err != nil {
		log.Fatalf("Redis connection error: %v", err)
	}
}
