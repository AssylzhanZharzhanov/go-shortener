package domain

import "fmt"

// LinkID - id of the link
type LinkID int64

// Link - link struct
type Link struct {
	ID          LinkID `json:"id" gorm:"column:id"`
	OriginalURL string `json:"original_url" gorm:"column:original_url"`
	Hash        string `json:"hash" gorm:"column:hash"`
	CreatedAt   int64  `json:"created_at" gorm:"not null;autoCreateTime;column:created_at"`
}

// Validate - validates an option.
func (l *Link) Validate() error {
	if len(l.OriginalURL) == 0 {
		return fmt.Errorf("invalid original url")
	}
	return nil
}
