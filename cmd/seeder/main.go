package main

import (
	"fmt"
	"log"
	"newbier-hackglobal/pkg/config"
	"newbier-hackglobal/pkg/database"
	"newbier-hackglobal/pkg/database/model"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load configuration: %v", err))
	}

	// Package
	db, err := database.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		panic(fmt.Sprintf("Failed to load database: %v", err))
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, seeds)

	if err := m.Migrate(); err != nil {
		panic(fmt.Sprintf("Could not apply migrations: %v", err))
	}

	log.Println("Seeding applied successfully")
}

var seeds = []*gormigrate.Migration{
	{
		ID: "1-seedUser",
		Migrate: func(tx *gorm.DB) error {
			return tx.Create([]*model.User{
				{
					Name:    "Natasha",
					Country: "Indonesia",
					Gender:  "female",
				},
				{
					Name:    "Elok",
					Country: "Indonesia",
					Gender:  "female",
				},
			}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return nil
		},
	},
}
