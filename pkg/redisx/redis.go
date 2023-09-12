package redisx

import (
	"github.com/redis/go-redis/v9"
	"time"
)

var client *redis.Client

func GetClient() *redis.Client {
	return client
}

func New() {
	client = redis.NewClient(&redis.Options{
		Addr: "",
		DB:   0,

		DialTimeout:           10 * time.Second,
		ReadTimeout:           30 * time.Second,
		WriteTimeout:          30 * time.Second,
		ContextTimeoutEnabled: true,

		MaxRetries: -1,

		PoolSize:        10,
		PoolTimeout:     30 * time.Second,
		ConnMaxIdleTime: time.Minute,
	})
}
