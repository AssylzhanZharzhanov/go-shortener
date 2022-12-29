package repository

import (
	"context"

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

func (r redisRepository) Get(ctx context.Context, key string) (*domain.Link, error) {
	//TODO implement me
	panic("implement me")
}

func (r redisRepository) Set(ctx context.Context, key string, seconds int, resident *domain.Link) error {
	//TODO implement me
	panic("implement me")
}

func (r redisRepository) Delete(ctx context.Context, key string) error {
	//TODO implement me
	panic("implement me")
}
