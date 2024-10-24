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
	Name    string `json:"name"`
	Type    string `json:"type"`
	Star    string `json:"star"`
	Address string `json:"address"`
	GmapUrl string `json:"gmap_url"`
	Image   string `json:"image"`
}

func (Destination) TableName() string {
	return "destinations"
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


func (ItineraryDestination) TableName() string {
	return "itinerary_destinations"
}

type ItineraryMarket struct {
	gorm.Model
	ItineraryID          int                `json:"itinerary"`
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
	ChatRoomID  string `json:"chat_room_id"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by"`
	IsAccept    bool   `json:"is_accept"`
	User        User   `gorm:"foreignKey:UserID;references:ID" json:"user"`
	CreateBy    User   `gorm:"foreignKey:CreatedBy;references:ID" json:"create_by"`
}


func (ItineraryBuddy) TableName() string {
	return "itinerary_buddies"
}

type ChatRoom struct {
	gorm.Model
	CreatedBy string `json:"created_by"`
	Message   string `json:"message"`
}

func (ChatRoom) TableName() string {
	return "chat_rooms"
}
