package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewPostgresDB(dbUrl string) (*gorm.DB, error) {
	gormConfig := &gorm.Config{}

	db, err := gorm.Open(postgres.Open(dbUrl), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}
