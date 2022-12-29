package redis

import (
	"fmt"
	"time"

	"github.com/AssylzhanZharzhanov/go-shortener/internal/helpers"
	"github.com/go-redis/redis/v8"
)

// NewRedisClient - Returns new redis client
func NewRedisClient(cfg helpers.AppConfig) *redis.Client {
	redisHost := fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)

	client := redis.NewClient(&redis.Options{
		Addr:         redisHost,
		MinIdleConns: 200,
		PoolSize:     cfg.RedisPoolSize,
		PoolTimeout:  time.Duration(240) * time.Second,
		Password:     cfg.RedisPassword,
		DB:           cfg.RedisDB,
	})

	return client
}
