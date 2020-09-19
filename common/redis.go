package common

import (
	"os"

	"github.com/go-redis/redis"
)

var client *redis.Client

func RedisInit() {
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}

	client = redis.NewClient(&redis.Options{
		Addr: dsn,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func GetClient() *redis.Client {
	return client
}
