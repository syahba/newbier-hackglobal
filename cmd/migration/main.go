package main

import (
	"fmt"
	"newbier-hackglobal/pkg/config"
	"newbier-hackglobal/pkg/database"
	"newbier-hackglobal/pkg/database/model"
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

	db.AutoMigrate(
		model.User{},
		model.Destination{},
		model.DestinationProduct{},
		model.Itinerary{},
		model.ItineraryDestination{},
		model.ItineraryMarket{},
		model.ItineraryBuddy{},
		model.ChatRoom{},
	)
}
