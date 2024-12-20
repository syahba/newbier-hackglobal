package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"newbier-hackglobal/cmd/generate/schema"
	"newbier-hackglobal/pkg/config"
	"newbier-hackglobal/pkg/database"
	"newbier-hackglobal/pkg/database/model"
	"os"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
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
	{
		ID: "2-seed",
		Migrate: func(tx *gorm.DB) error {
			dataList := []string{
				"destinations",
			}

			var ID int
			for _, list := range dataList {
				file, err := os.Open(fmt.Sprintf("cmd/seeder/data/scrap/%s.json", list))
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
				for idx := range destinations {
					ID = ID + 1
					destinations[idx].ID = uint(ID)
				}

				if err := tx.Create(&destinations).Error; err != nil {
					return fmt.Errorf("failed to seed destinations: %w", err)
				}
			}

			if err := seedItinerary(tx); err != nil {
				return err
			}

			if err := seedItineraryDestination(tx); err != nil {
				return err
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return nil
		},
	},
	{
		ID: "3-seed",
		Migrate: func(tx *gorm.DB) error {
			dataList := []string{
				"destination_additional",
				"destination_additional2",
			}

			for _, list := range dataList {
				file, err := os.Open(fmt.Sprintf("cmd/seeder/data/destination-additional/%s.json", list))
				if err != nil {
					return fmt.Errorf("failed to open destinations JSON file: %w", err)
				}
				defer file.Close()

				byteValue, err := io.ReadAll(file)
				if err != nil {
					return fmt.Errorf("failed to read destinations JSON file: %w", err)
				}

				var destinations = []schema.Response[schema.DestinationAdditional]{}
				err = json.Unmarshal(byteValue, &destinations)
				if err != nil {
					return fmt.Errorf("failed to unmarshal destinations data: %w", err)
				}

				var destinationParam = []*model.DestinationParameter{}
				var DestinationProduct = []*model.DestinationProduct{}
				for _, elm := range destinations {
					for _, el := range elm.Data.Activity {
						destinationParam = append(destinationParam, &model.DestinationParameter{
							DestinationID: elm.ID,
							Type:          "activity",
							Name:          el,
						})
					}

					for _, el := range elm.Data.Trip {
						destinationParam = append(destinationParam, &model.DestinationParameter{
							DestinationID: elm.ID,
							Type:          "trip",
							Name:          el,
						})
					}

					for _, el := range elm.Data.Product {
						DestinationProduct = append(DestinationProduct, &model.DestinationProduct{
							DestinationID: elm.ID,
							Name:          el.Name,
							Unit:          el.Unit,
							Price:         el.Price,
						})
					}
				}

				if err := tx.Create(&destinationParam).Error; err != nil {
					return fmt.Errorf("failed to seed destination_parameter: %w", err)
				}
				if err := tx.Create(&DestinationProduct).Error; err != nil {
					return fmt.Errorf("failed to seed destination_products: %w", err)
				}
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return nil
		},
	},
	{
		ID: "4-seed",
		Migrate: func(tx *gorm.DB) error {
			dataList := []string{
				"destination_detail",
				"destination_detail2",
			}

			for _, list := range dataList {
				file, err := os.Open(fmt.Sprintf("cmd/seeder/data/destination-additional/%s.json", list))
				if err != nil {
					return fmt.Errorf("failed to open destinations JSON file: %w", err)
				}
				defer file.Close()

				byteValue, err := io.ReadAll(file)
				if err != nil {
					return fmt.Errorf("failed to read destinations JSON file: %w", err)
				}

				var destinations = []schema.Response[schema.DestinationDetail]{}
				err = json.Unmarshal(byteValue, &destinations)
				if err != nil {
					return fmt.Errorf("failed to unmarshal destinations data: %w", err)
				}

				for _, elm := range destinations {
					if err := tx.Updates(model.Destination{
						Model: gorm.Model{
							ID: uint(elm.ID),
						},
						Description: elm.Data.Description,
						BestProduct: elm.Data.Product,
					}).Error; err != nil {
						return err
					}
				}

			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return nil
		},
	},
}

func seedItinerary(tx *gorm.DB) error {
	return tx.Create([]*model.Itinerary{
		{
			Destination: ptr("Singapore"),
			Activity:    "Museum Tour",
			Date:        time.Date(2024, 11, 1, 10, 0, 0, 0, time.UTC),
			Vehicle:     "Bus",
			Trip:        "Cultural Exploration",
			CreatedBy:   "user1",
		},
		{
			Destination: ptr("Singapore"),
			Activity:    "Heritage Walk",
			Date:        time.Date(2024, 12, 5, 9, 0, 0, 0, time.UTC),
			Vehicle:     "Walk",
			Trip:        "Local Heritage",
			CreatedBy:   "user2",
		},
		{
			Destination: ptr("Singapore"),
			Activity:    "Family Fun",
			Date:        time.Date(2024, 12, 20, 15, 0, 0, 0, time.UTC),
			Vehicle:     "Car",
			Trip:        "Family Trip",
			CreatedBy:   "user3",
		},
		{
			Destination: ptr("Singapore"),
			Activity:    "Beach & Relaxation",
			Date:        time.Date(2025, 1, 15, 11, 0, 0, 0, time.UTC),
			Vehicle:     "Bicycle",
			Trip:        "Outdoor Escape",
			CreatedBy:   "user4",
		},
		{
			Destination: ptr("Singapore"),
			Activity:    "Adventure Day",
			Date:        time.Date(2025, 2, 10, 14, 0, 0, 0, time.UTC),
			Vehicle:     "Scooter",
			Trip:        "Urban Adventure",
			CreatedBy:   "user5",
		},
	}).Error
}

func ptr(s string) *string {
	return &s
}

func seedItineraryDestination(tx *gorm.DB) error {
	return tx.Create([]*model.ItineraryDestination{
		{
			ItineraryID:   1,
			DestinationID: 1, // ArtScience Museum
			Time:          "10:00 AM",
		},
		{
			ItineraryID:   1,
			DestinationID: 2, // National Museum of Singapore
			Time:          "02:00 PM",
		},
		{
			ItineraryID:   2,
			DestinationID: 3, // National Gallery Singapore
			Time:          "11:00 AM",
		},
		{
			ItineraryID:   2,
			DestinationID: 4, // Peranakan Museum
			Time:          "03:00 PM",
		},
		{
			ItineraryID:   3,
			DestinationID: 5, // Children Little Museum
			Time:          "09:00 AM",
		},
		{
			ItineraryID:   3,
			DestinationID: 7, // Asian Civilisations Museum
			Time:          "01:00 PM",
		},
		{
			ItineraryID:   4,
			DestinationID: 10, // Indian Heritage Centre
			Time:          "08:00 AM",
		},
		{
			ItineraryID:   4,
			DestinationID: 12, // Fort Siloso
			Time:          "04:00 PM",
		},
		{
			ItineraryID:   5,
			DestinationID: 20, // Singapore Air Force Museum
			Time:          "07:00 AM",
		},
		{
			ItineraryID:   5,
			DestinationID: 25, // Friends of the Museums
			Time:          "05:00 PM",
		},
	}).Error
}
