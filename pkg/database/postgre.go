package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(dbUrl string) (*gorm.DB, error) {
	gormConfig := &gorm.Config{
		// TranslateError: true,
		// Logger:         logger.Default.LogMode(logger.Info), // Set the desired log level
	}

	db, err := gorm.Open(postgres.Open(dbUrl), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil

}
