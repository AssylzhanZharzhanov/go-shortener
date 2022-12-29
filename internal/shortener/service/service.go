package service

import (
	"context"

	"github.com/AssylzhanZharzhanov/go-shortener/internal/shortener/domain"
)

const (
	cacheDuration = 3600
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
		url := shortUrl[i : i+7]
		isExist, err := s.repository.IsExist(url)
		if err != nil {
			return nil, err
		}
		if isExist {
			continue
		} else {
			shortUrl = url
			break
		}
	}

	link.Hash = shortUrl
	return s.repository.Create(link)
}

func (s *service) Get(ctx context.Context, shortenURL string) (*domain.Link, error) {
	cached, err := s.redisRepository.Get(ctx, shortenURL)
	if err != nil {
		return nil, err
	}
	if cached != nil {
		return cached, nil
	}

	response, err := s.repository.Get(shortenURL)
	if err != nil {
		return nil, err
	}

	err = s.redisRepository.Set(ctx, shortenURL, response, cacheDuration)
	if err != nil {
		return nil, err
	}
	return response, nil
}
