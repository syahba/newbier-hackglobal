package internal_model

import (
	"newbier-hackglobal/pkg/database/model"
)

type ItineraryResponse struct {
	ItineraryID   int             `json:"itinerary_id"`
	ItineraryTime []ItineraryTime `json:"itinerary_time"`
}

type ItineraryTime struct {
	Time         string              `json:"time"`
	Destinations []model.Destination `json:"destination"`
}

type ItineraryFinderRequest struct {
	Activity string `json:"activity"`
	Date     string `json:"date"`
	Trip     string `json:"trip"`
}

type GenerateItinerarySchema struct {
	Time           string `json:"time"`
	DestinationIDs []int  `json:"destination_ids"`
	Destinations   []model.Destination
}

type Destination struct {
	Id          int                        `json:"id"`
	Name        string                     `json:"name"`
	Type        string                     `json:"type"`
	Star        string                     `json:"star"`
	Address     string                     `json:"address"`
	GmapUrl     string                     `json:"gmap_url"`
	Image       string                     `json:"image"`
	Description string                     `json:"description"`
	BestProduct []string                   `gorm:"serializer:json" json:"best_product"`
	Product     []model.DestinationProduct `json:"product"`
}

type Combo struct {
	Destination string                    `json:"destination"` // itinerary_destination
	Activity    string                    `json:"activity"`    // itinerary
	Trip        string                    `json:"trip"`        // itinerary
	User_id     int                       `json:"user_id"`     // itinerary -> created_by
	Itinerary   []GenerateItinerarySchema `json:"itinerary"`   // itinerary_destination
	IsBuddy     bool                      `json:"isBuddy"`     // itinerary_buddy
	Description string                    `json:"description"` // itinerary_buddy
}
type GetItineraryDestinationsResponse struct {
	Description string                    `json:"description"`
	User        model.User                `json:"user"`
	Itinerary   []GenerateItinerarySchema `json:"itinerary"`
}
