package internal_model

import "newbier-hackglobal/pkg/database/model"

type Itinerary struct {
	ItineraryID   string
	ItineraryTime []Itinerary
}

type ItineraryTime struct {
	Time         string              `json:"time"`
	Destinations []model.Destination `json:"destination"`
}
