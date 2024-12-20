package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string `json:"name"`
	Country string `json:"country"`
	Gender  string `json:"gender"`
}

func (User) TableName() string {
	return "users"
}

type Destination struct {
	gorm.Model
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	Star        string   `json:"star"`
	Address     string   `json:"address"`
	GmapUrl     string   `json:"gmap_url"`
	Image       string   `json:"image"`
	Description string   `json:"description"`
	BestProduct []string `gorm:"serializer:json" json:"best_product"`
}

func (Destination) TableName() string {
	return "destinations"
}

func (Destination) ColumnName(column string) string {
	return "destinations." + column
}

type DestinationParameter struct {
	gorm.Model
	DestinationID int    `json:"destination_id"`
	Type          string `json:"type"`
	Name          string `json:"name"`
}

func (DestinationParameter) TableName() string {
	return "destination_parameters"
}

type DestinationProduct struct {
	gorm.Model
	DestinationID int     `json:"destination_id"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	Unit          string  `json:"unit"`
}

func (DestinationProduct) TableName() string {
	return "destination_products"
}

type Itinerary struct {
	gorm.Model
	Destination           *string                `json:"destination"`
	Activity              string                 `json:"activity"`
	Date                  time.Time              `json:"date"`
	Vehicle               string                 `json:"vehicle"`
	Trip                  string                 `json:"trip"`
	CreatedBy             string                 `json:"created_by"`
	ItineraryDestinations []ItineraryDestination `gorm:"-:migration;foreignKey:ItineraryID" json:"itinerary_destination"`
	ItineraryBuddy        *ItineraryBuddy        `gorm:"-:migration;foreignKey:ItineraryID" json:"itinerary_buddy"`
}

func (Itinerary) TableName() string {
	return "itineraries"
}

type ItineraryDestination struct {
	gorm.Model
	ItineraryID   int         `json:"itinerary"`
	DestinationID int         `json:"destination_id"`
	Time          string      `json:"time"`
	Destination   Destination `gorm:"foreignKey:DestinationID;references:ID" json:"destination"`
}

type ItineraryDestinations struct {
	gorm.Model
	ItineraryID   int         `json:"itinerary"`
	DestinationID int         `json:"destination_id"`
	Time          string      `json:"time"`
	Destination   []Destination `gorm:"foreignKey:DestinationID;references:ID" json:"destination"`
}

func (ItineraryDestination) TableName() string {
	return "itinerary_destinations"
}

type ItineraryMarket struct {
	gorm.Model
	ItineraryID          int                `json:"itinerary_id"`
	DestinationProductID int                `json:"destination_product_id"`
	Amount               int                `json:"amount"`
	DestinationProduct   DestinationProduct `gorm:"foreignKey:DestinationProductID;references:ID" json:"destination_product"`
}

func (ItineraryMarket) TableName() string {
	return "itinerary_markets"
}

type ItineraryBuddy struct {
	gorm.Model
	ItineraryID int    `json:"itinerary_id"`
	UserID      int    `json:"user_id"`
	Description string `json:"description"`
}

func (ItineraryBuddy) TableName() string {
	return "itinerary_buddies"
}

type ItineraryFinder struct {
	gorm.Model
	Description string    `json:"description"`
	Activity    string    `json:"activity"`
	Date        time.Time `json:"date"`
	Trip        string    `json:"trip"`
	CreatedBy   int       `json:"created_by"`
	User        *User     `gorm:"foreignKey:CreatedBy;references:ID" json:"user"`
}

func (ItineraryFinder) TableName() string {
	return "itinerary_finders"
}

type ItineraryRequest struct {
	gorm.Model
	Description       string `json:"description"`
	ItineraryID       int    `json:"itinerary_id"`
	ItineraryFinderID int    `json:"itinerary_finder_id"`
	CreatedBy         int    `json:"created_by"`
	Accepted          *bool  `json:"accepted"`
}

func (ItineraryRequest) TableName() string {
	return "itinerary_requests"
}

type ChatRoom struct {
	gorm.Model
	ItineraryID int    `json:"itinerary_id"`
	CreatedBy   string `json:"created_by"`
	Message     string `json:"message"`
}

func (ChatRoom) TableName() string {
	return "chat_rooms"
}
