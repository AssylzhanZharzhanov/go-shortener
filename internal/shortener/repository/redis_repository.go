package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/AssylzhanZharzhanov/go-shortener/internal/shortener/domain"

	"github.com/go-redis/redis/v8"
)

type redisRepository struct {
	client *redis.Client
}

// NewRedisRepository - create a new redis repository
func NewRedisRepository(client *redis.Client) domain.ShortenerRedisRepository {
	return &redisRepository{
		client: client,
	}
}

func (r *redisRepository) Set(ctx context.Context, key string, value *domain.Link, seconds int) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if err = r.client.Set(ctx, key, bytes, time.Second*time.Duration(seconds)).Err(); err != nil {
		return err
	}
	return nil
}

func (r *redisRepository) Get(ctx context.Context, key string) (*domain.Link, error) {
	bytes, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	value := &domain.Link{}
	if err = json.Unmarshal(bytes, value); err != nil {
		return nil, err
	}

	return value, nil
}

func (r *redisRepository) Delete(ctx context.Context, key string) error {
	if err := r.client.Del(ctx, key).Err(); err != nil {
		return err
	}
	return nil
}
