package service

import (
	"context"

	"github.com/AssylzhanZharzhanov/go-shortener/internal/shortener/domain"
)

type service struct {
	repository      domain.ShortenerRepository
	redisRepository domain.ShortenerRedisRepository
}

// NewShortenerService return a new "ShortenerService"
func NewShortenerService(repository domain.ShortenerRepository, redisRepository domain.ShortenerRedisRepository) domain.ShortenerService {
	return &service{
		repository:      repository,
		redisRepository: redisRepository,
	}
}

func (s *service) CreateShortenURL(ctx context.Context, link *domain.Link) (*domain.Link, error) {

	var (
		shortUrl = Shorten(link.OriginalURL)
	)

	for i := 0; i < len(shortUrl)-7; i++ {
		isExist, err := s.repository.IsExist(shortUrl)
		if err != nil {
			return nil, err
		}
		if isExist {
			continue
		} else {
			break
		}
	}

	link.Hash = shortUrl
	return s.repository.Create(link)
}

func (s *service) Get(ctx context.Context, shortenURL string) (*domain.Link, error) {
	return s.repository.Get(shortenURL)
}
