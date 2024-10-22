package main

import (
	"fmt"
	"log"
	"newbier-hackglobal/pkg/config"
	"newbier-hackglobal/pkg/database"
	"newbier-hackglobal/pkg/database/model"

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

	if err = seedUser(db); err != nil {
		log.Print("Failed to seeding: seedUser")
		return
	}

	log.Print("seeding: success")

}

// Seed User Data
func seedUser(db *gorm.DB) error {
	return db.Create([]*model.User{
		{
			Model:   gorm.Model{ID: 1},
			Name:    "Natasha",
			Country: "Indonesia",
			Gender:  "female",
		},
		{
			Model:   gorm.Model{ID: 2},
			Name:    "Elok",
			Country: "Indonesia",
			Gender:  "female",
		},
	}).Error
}
