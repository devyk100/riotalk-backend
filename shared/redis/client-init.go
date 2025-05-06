package redis

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

var RedisClient *redis.Client

func InitRedisClient(ctx context.Context) error {
	fmt.Println(os.Getenv("REDIS_PASSWORD"), os.Getenv("REDIS_URL"))
	rdb := redis.NewClient(&redis.Options{
		Addr:      os.Getenv("REDIS_URL"),
		Password:  os.Getenv("REDIS_PASSWORD"),
		TLSConfig: &tls.Config{},
		DB:        0,
	})
	// ctx := context.Background()
	// Ping the Redis server to check the connection
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	fmt.Println("Connected to Redis:", pong)
	RedisClient = rdb
	return nil
}
