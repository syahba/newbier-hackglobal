package main

import (
	"newbier-hackglobal/pkg/config"
	"newbier-hackglobal/pkg/database"
	"newbier-hackglobal/pkg/database/model"
	"github.com/go-gormigrate/gormigrate/v2"
	"encoding/json"
	"gorm.io/gorm"
	"fmt"
	"log"
	"io"
	"os"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load configuration: %v", err))
	}

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
		ID: "1-seedUserAndDestination",
		Migrate: func(tx *gorm.DB) error {
			if err := seedUsers(tx); err != nil {
				return err
			}

			if err := seedDestinations(tx); err != nil {
				return err
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return nil
		},
	},
}

func seedUsers(tx *gorm.DB) error {
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
}

func seedDestinations(tx *gorm.DB) error {

	dataList := []string{
		"scrap museum",
		"scrap park",
		"scrap restaurant",
		"scrap shop",
		"scrap tourist places",
	}

	for _,list := range dataList{
		file, err := os.Open(fmt.Sprintf("pkg/dataset/%s.json",list))
		if err != nil {
			return fmt.Errorf("failed to open destinations JSON file: %w", err)
		}
		defer file.Close()

		byteValue, err := io.ReadAll(file)
		if err != nil {
			return fmt.Errorf("failed to read destinations JSON file: %w", err)
		}

		var destinations []model.Destination
		err = json.Unmarshal(byteValue, &destinations)
		if err != nil {
			return fmt.Errorf("failed to unmarshal destinations data: %w", err)
		}

		if err := tx.Create(&destinations).Error; err != nil {
			return fmt.Errorf("failed to seed destinations: %w", err)
		}
	}

	return nil
}