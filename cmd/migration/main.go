package main

import (
	"fmt"
	"log"
	"newbier-hackglobal/pkg/config"
	"newbier-hackglobal/pkg/database"

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

	m := gormigrate.New(db, gormigrate.DefaultOptions, migrations)

	if err := m.Migrate(); err != nil {
		panic(fmt.Sprintf("Could not apply migrations: %v", err))
	}

	log.Println("Migrations applied successfully")
}

var migrations = []*gormigrate.Migration{
	{
		ID: "1",
		Migrate: func(tx *gorm.DB) error {
			return tx.Exec(`
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    country VARCHAR(255),
    gender VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE destinations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(100),
    star VARCHAR(50),
    address TEXT,
    gmap_url TEXT,
    image TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE destination_products (
    id SERIAL PRIMARY KEY,
    destination_id INTEGER REFERENCES destinations(id),
    name VARCHAR(255) NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    unit VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE itineraries (
    id SERIAL PRIMARY KEY,
    destination VARCHAR(255),
    activity TEXT,
    date TIMESTAMP,
    vehicle VARCHAR(255),
    trip VARCHAR(255),
    created_by VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE itinerary_destinations (
    id SERIAL PRIMARY KEY,
    itinerary_id INTEGER REFERENCES itineraries(id),
    destination_id INTEGER REFERENCES destinations(id),
    time VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE itinerary_markets (
    id SERIAL PRIMARY KEY,
    itinerary_id INTEGER REFERENCES itineraries(id),
    destination_product_id INTEGER REFERENCES destination_products(id),
    amount INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE chat_rooms (
	id SERIAL PRIMARY KEY,
	created_by VARCHAR(255),
	message TEXT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP
);

CREATE TABLE itinerary_buddies (
    id SERIAL PRIMARY KEY,
    itinerary_id INTEGER REFERENCES itineraries(id),
    user_id INTEGER REFERENCES users(id),
    chat_room_id INTEGER REFERENCES chat_rooms(id),
    description TEXT,
    created_by VARCHAR(255),
    is_accept BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
            `).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return nil
		},
	},
	{
		ID: "2",
		Migrate: func(tx *gorm.DB) error {
			return tx.Exec(`
CREATE TABLE destination_parameters (
    id SERIAL PRIMARY KEY,
    destination_id INT NOT NULL,
    type VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);
            `).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return nil
		},
	},
	{
		ID: "3",
		Migrate: func(tx *gorm.DB) error {
			return tx.Exec(`
ALTER TABLE destinations
ADD COLUMN description TEXT;

ALTER TABLE destinations
ADD COLUMN best_product JSON DEFAULT '[]';
            `).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return nil
		},
	},
	{
		ID: "4",
		Migrate: func(tx *gorm.DB) error {
			return tx.Exec(`
ALTER TABLE itinerary_buddies 
DROP CONSTRAINT itinerary_buddies_chat_room_id_fkey;

ALTER TABLE itinerary_buddies 
DROP CONSTRAINT itinerary_buddies_itinerary_id_fkey;

ALTER TABLE itinerary_buddies 
DROP CONSTRAINT itinerary_buddies_user_id_fkey;

ALTER TABLE itinerary_buddies 
DROP COLUMN chat_room_id;

ALTER TABLE itinerary_buddies 
DROP COLUMN created_by;

ALTER TABLE itinerary_buddies 
DROP COLUMN is_accept;

CREATE TABLE itinerary_finders (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NULL DEFAULT NULL,
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    description VARCHAR(255) NOT NULL,
    activity VARCHAR(100) NOT NULL,
    date TIMESTAMP NOT NULL,
    trip VARCHAR(100) NOT NULL,
    created_by INT NOT NULL
);

CREATE TABLE itinerary_requests (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NULL DEFAULT NULL,
    updated_at TIMESTAMP NULL DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    description VARCHAR(255) NOT NULL,
    itinerary_buddy_id INT NOT NULL,
    itinerary_finder_id int NOT NULL,
    created_by INT NOT NULL,
    accepted BOOLEAN
);

            `).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return nil
		},
	},
}
