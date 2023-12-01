package redis

import (
	"github.com/go-redis/redis/v8"
)

func CreateClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,
	})
}
