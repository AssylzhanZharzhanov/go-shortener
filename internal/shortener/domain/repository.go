package domain

import "context"

// ShortenerRepository - provides access to storage
type ShortenerRepository interface {
	Create(link *Link) (*Link, error)
	Get(shortenURL string) (*Link, error)
}

// ShortenerRedisRepository - provides access to a redis storage
type ShortenerRedisRepository interface {
	Get(ctx context.Context, key string) (*Link, error)
	Set(ctx context.Context, key string, seconds int, resident *Link) error
	Delete(ctx context.Context, key string) error
}
