package domain

import "context"

// ShortenerService - provides access to a business logic.
type ShortenerService interface {
	CreateShortenURL(
		ctx context.Context,
		link *Link,
	) (*Link, error)

	Get(
		ctx context.Context,
		shortenURL string,
	) (*Link, error)
}
