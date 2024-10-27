package internal_model

import "newbier-hackglobal/pkg/database/model"

type ItineraryResponse struct {
	ItineraryID   int             `json:"itinerary_id"`
	ItineraryTime []ItineraryTime `json:"itinerary_time"`
}

type ItineraryTime struct {
	Time         string              `json:"time"`
	Destinations []model.Destination `json:"destination"`
}
