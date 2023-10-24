package redisx

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var client *redis.Client

type Config struct {
	Host                  string
	Password              string
	DB                    int
	DialTimeout           time.Duration
	ReadTimeout           time.Duration
	WriteTimeout          time.Duration
	ContextTimeoutEnabled bool
	MaxRetries            int
	PoolSize              int
	PoolTimeout           time.Duration
	ConnMaxIdleTime       time.Duration
}

func GetClient() *redis.Client {
	return client
}

func New(config *Config) error {
	client = redis.NewClient(&redis.Options{
		Addr:                  config.Host,
		DB:                    config.DB,
		DialTimeout:           config.DialTimeout,
		ReadTimeout:           config.ReadTimeout,
		WriteTimeout:          config.WriteTimeout,
		ContextTimeoutEnabled: config.ContextTimeoutEnabled,
		MaxRetries:            config.MaxRetries,
		PoolSize:              config.PoolSize,
		PoolTimeout:           config.PoolTimeout,
		ConnMaxIdleTime:       config.ConnMaxIdleTime,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return client.Ping(ctx).Err()
}
