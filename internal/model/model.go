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

type GenerateItinerarySchema struct {
	Time           string `json:"time"`
	DestinationIDs []int  `json:"destination_ids"`
	Destinations   []model.Destination
}

type Combo struct{
	Destination string 							`json:"destination"`// itinerary_destination
	Activity 	string 							`json:"activity"`	// itinerary
	Trip 		string 							`json:"trip"`		// itinerary
	User_id 	int 							`json:"user_id"`		// itinerary -> created_by
	Itinerary 	[]GenerateItinerarySchema 		`json:"itinerary"`	// itinerary_destination
	IsBuddy 	bool 							`json:"isBuddy"`	// itinerary_buddy
	Description string 							`json:"description"`// itinerary_buddy
}