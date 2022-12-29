package service

import "github.com/AssylzhanZharzhanov/go-shortener/internal/auth/domain"

type service struct {
	repository domain.AuthService
}

// NewAuthService - returns a new "AuthService"
func NewAuthService(repository domain.AuthService) domain.AuthService {
	return &service{
		repository: repository,
	}
}
