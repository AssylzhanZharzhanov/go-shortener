package repository

import (
	"github.com/AssylzhanZharzhanov/go-shortener/internal/shortener/domain"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// NewShortenerRepository - returns a new "ShortenerRepository"
func NewShortenerRepository(db *gorm.DB) domain.ShortenerRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(link *domain.Link) (*domain.Link, error) {
	err := r.db.Create(&link).Error
	if err != nil {
		return nil, err
	}

	return link, nil
}

func (r *repository) Get(shortURL string) (*domain.Link, error) {
	var link *domain.Link

	err := r.db.Where("short_url = ?", shortURL).First(&link).Error
	if err != nil {
		return nil, err
	}

	return link, nil
}

func (r *repository) IsExist(shortURL string) (bool, error) {
	var isExist bool

	err := r.db.
		Model(&domain.Link{}).
		Select("count(*) > 0").
		Where("short_url = ?", shortURL).
		Find(&isExist).Error
	if err != nil {
		return false, err
	}

	return isExist, nil
}
