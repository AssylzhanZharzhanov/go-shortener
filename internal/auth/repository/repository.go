package repository

import (
	"github.com/AssylzhanZharzhanov/go-shortener/internal/auth/domain"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// NewAuthRepository - provides access to a storage.
func NewAuthRepository(db *gorm.DB) domain.AuthRepository {
	return &repository{
		db: db,
	}
}
