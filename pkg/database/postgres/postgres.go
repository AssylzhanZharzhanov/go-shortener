package postgres

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewConnection creates a new database connection.
func NewConnection(dsn string) (*gorm.DB, error) {
	gormDB, err := gorm.Open(postgres.Open(dsn),
		&gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Fatal("Failed to setup sql database:", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return gormDB, nil
}
