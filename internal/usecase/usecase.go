package usecase

import (
	"encoding/json"
	"fmt"
	internal_model "newbier-hackglobal/internal/model"
	chatgpt "newbier-hackglobal/pkg/chatGPT"
	"newbier-hackglobal/pkg/database/model"
	"slices"
	"strings"
	"time"

	"math/rand"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Usecase struct {
	db *gorm.DB
	ai *chatgpt.Model
}

func NewUsecase(db *gorm.DB, ai *chatgpt.Model) *Usecase {
	return &Usecase{db: db, ai: ai}
}

func (u *Usecase) GetUsecase() string {

	return "Hello World"
}

func (u *Usecase) GenerateItinerary(activity, trip string) (data []internal_model.GenerateItinerarySchema, err error) {
	table := model.Destination{}
	destinations := []map[string]any{}
	u.db.Select(table.ColumnName("id"), table.ColumnName("name"), table.ColumnName("type")).
		Model(table).
		Joins(`JOIN "destination_parameters" dp1 ON `+table.ColumnName("id")+` = dp1.destination_id AND dp1."name" = ?`, activity).
		Joins(`JOIN "destination_parameters" dp2 ON `+table.ColumnName("id")+` = dp2.destination_id AND dp2."name" = ?`, trip).
		Find(&destinations)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(destinations), func(i, j int) { destinations[i], destinations[j] = destinations[j], destinations[i] })

	var result, _ = u.ai.GenerateCustom(0.5, 700, GenerateItinerary(destinations, activity, trip))
	cleanedInput := strings.Trim(result, "```json")
	cleanedInput = strings.Trim(cleanedInput, "```")

	data = make([]internal_model.GenerateItinerarySchema, 0)
	json.Unmarshal([]byte(cleanedInput), &data)

	// using index looping to modify the actual value
	for index1 := 0; index1 < len(data); index1++ {
		for index2 := 0; index2 < len(data[index1].DestinationIDs); index2++ {
			destination, _ := u.GetDestinationById(data[index1].DestinationIDs[index2])
			data[index1].Destinations = append(data[index1].Destinations, destination)
		}
	}
	return
}

func (u *Usecase) GenerateItineraryWithDestination(destination, trip string) (data []internal_model.GenerateItinerarySchema, err error) {
	table := model.Destination{}
	destinations := []map[string]any{}
	u.db.Select(table.ColumnName("id"), table.ColumnName("name"), table.ColumnName("type")).
		Model(table).
		Joins(`JOIN "destination_parameters" dp1 ON `+table.ColumnName("id")+` = dp1.destination_id AND dp1."name" = ?`, destination).
		Joins(`JOIN "destination_parameters" dp2 ON `+table.ColumnName("id")+` = dp2.destination_id AND dp2."name" = ?`, trip).
		Find(&destinations)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(destinations), func(i, j int) { destinations[i], destinations[j] = destinations[j], destinations[i] })

	var result, _ = u.ai.GenerateCustom(0.5, 700, GenerateItineraryWithDestination(destinations, destination, trip))
	cleanedInput := strings.Trim(result, "```json")
	cleanedInput = strings.Trim(cleanedInput, "```")

	data = make([]internal_model.GenerateItinerarySchema, 0)
	json.Unmarshal([]byte(cleanedInput), &data)

	// using index looping to modify the actual value
	for index1 := 0; index1 < len(data); index1++ {
		for index2 := 0; index2 < len(data[index1].DestinationIDs); index2++ {
			destination1, _ := u.GetDestinationById(data[index1].DestinationIDs[index2])
			data[index1].Destinations = append(data[index1].Destinations, destination1)
		}
	}
	return
}

func (u *Usecase) GetItinerary() (data []model.Itinerary, err error) {
	data = make([]model.Itinerary, 0)

	err = u.db.
		Preload("ItineraryDestinations.Destination").
		Find(&data).Error
	if err != nil {
		return
	}

	return
}

func (u *Usecase) GetDestinations(name string) (data []model.Destination, err error) {
	data = make([]model.Destination, 0)

	query := u.db

	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	err = query.Find(&data).Error
	if err != nil {
		return
	}

	return
}

func (u *Usecase) GetDestinationById(id int) (data model.Destination, err error) {
	err = u.db.Where("id = ?", id).Find(&data).Error
	if err != nil {
		return
	}
	return
}

func (u *Usecase) GetDestinationByItineraryId(id int) (respondData []internal_model.Destination, err error) {

	data := make([]model.ItineraryDestination, 0)

	err = u.db.Where("itinerary_id = ?", id).Find(&data).Error
	var destinationIdArr []int
	for _, element := range data {
		destinationIdArr = append(destinationIdArr, element.DestinationID)
	}

	respondData = make([]internal_model.Destination, 0)

	for _, id := range destinationIdArr {
		newDestination, _ := u.GetDestinationById(id)

		newData := internal_model.Destination{
			Id:          int(newDestination.Model.ID),
			Name:        newDestination.Name,
			Type:        newDestination.Type,
			Star:        newDestination.Star,
			Address:     newDestination.Address,
			GmapUrl:     newDestination.GmapUrl,
			Image:       newDestination.Image,
			Description: newDestination.Description,
			BestProduct: newDestination.BestProduct,
		}
		respondData = append(respondData, newData)
	}

	for index := 0; index < len(respondData); index++ {
		var listProduct []model.DestinationProduct
		err = u.db.Where("destination_id = ?", respondData[index].Id).Find(&listProduct).Error
		if err != nil {
			return
		}

		if len(listProduct) == 0 {
			respondData = append(respondData[:index], respondData[index+1:]...)
			index--
			continue
		}

		respondData[index].Product = listProduct
	}

	if err != nil {
		return
	}

	return
}

func (u *Usecase) GetItineraryDestination() (data []model.ItineraryDestination, err error) {

	data = make([]model.ItineraryDestination, 0)

	err = u.db.Preload("Destination").Order("time ASC").Find(&data).Error
	if err != nil {
		return
	}

	return
}

func (u *Usecase) CreateChat(newChat model.ChatRoom) (err error) {

	err = u.db.Create(&newChat).Error
	if err != nil {
		return
	}

	return
}

func (u *Usecase) GetChat() (data []model.ChatRoom, err error) {

	data = make([]model.ChatRoom, 0)

	err = u.db.Find(&data).Error
	if err != nil {
		return
	}

	return
}

func (u *Usecase) CreateItineraryBuddy(newItineraryBuddy model.ItineraryBuddy) (err error) {

	err = u.db.Create(&newItineraryBuddy).Error
	if err != nil {
		return
	}

	return
}

func (u *Usecase) GetItineraryById(id int) (data model.Itinerary, err error) {
	err = u.db.Where("id = ?", id).Find(&data).Error
	if err != nil {
		return
	}

	return
}

func (u *Usecase) CreateItineraryAndItineraryDestinationAndBuddy(data internal_model.Combo) error {
	newItinerary := model.Itinerary{
		Destination: &data.Description,
		Activity:    data.Activity,
		Vehicle:     "empty",
		Trip:        data.Trip,
		CreatedBy:   fmt.Sprintf("%v", data.User_id),
	}

	err := u.db.Create(&newItinerary).Error
	if err != nil {
		return err
	}

	var user model.User
	err = u.db.Where("id = ?", data.User_id).Find(&user).Error
	if err != nil {
		return err
	}

	newBuddy := model.ItineraryBuddy{
		ItineraryID: int(newItinerary.Model.ID),
		UserID:      data.User_id,
		Description: data.Description,
	}

	err = u.db.Create(&newBuddy).Error
	if err != nil {
		return err
	}

	var itineraryDestination []model.ItineraryDestination
	for _, element := range data.Itinerary {
		for _, destination1 := range element.Destinations {

			newItineraryDestination := model.ItineraryDestination{
				ItineraryID:   int(newItinerary.Model.ID),
				Time:          element.Time,
				DestinationID: int(destination1.ID),
			}
			itineraryDestination = append(itineraryDestination, newItineraryDestination)
		}
	}

	err = u.db.Create(&itineraryDestination).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *Usecase) GetUserById(id int) (data model.User, err error) {
	err = u.db.Where("id = ?", id).Find(&data).Error
	if err != nil {
		return
	}

	return
}

func (u *Usecase) GetChatById(id int) (data model.ChatRoom, err error) {
	err = u.db.Where("id = ?", id).Find(&data).Error
	if err != nil {
		return
	}

	return
}

func (u *Usecase) GetDestinationProductById(id int) (data model.DestinationProduct, err error) {
	err = u.db.Where("id = ?", id).Find(&data).Error
	if err != nil {
		return
	}

	return
}

func (u *Usecase) CreateItternaryMarkets(data model.ItineraryMarket) (err error) {
	err = u.db.Create(&data).Error
	if err != nil {
		return
	}

	return
}

func (u *Usecase) GetItternaryMarketByItternaryId(id int) (data []model.ItineraryMarket, err error) {
	data = make([]model.ItineraryMarket, 0)

	err = u.db.Preload("DestinationProduct").Where("itinerary_id = ?", id).Find(&data).Error
	if err != nil {
		return
	}

	return
}

func (u *Usecase) JoinItineraryBuddy(userId, itineraryId int) (err error) {
	var ItineraryBuddy model.ItineraryBuddy
	err = u.db.Where("itinerary_id = ? AND user_id = ?", itineraryId, userId).First(&ItineraryBuddy).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("itinerary buddy not found")
		}
		return err
	}

	// ItineraryBuddy.IsAccept = true

	err = u.db.Save(&ItineraryBuddy).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *Usecase) GetFinderByItineraryId(itineraryId int) (data model.ItineraryFinder, err error) {

	var buddy model.ItineraryBuddy
	err = u.db.Where("itinerary_id = ?", itineraryId).Find(&buddy).Error
	if err != nil {
		return
	}

	var itineraryRequest model.ItineraryRequest
	err = u.db.Where("itinerary_buddy_id = ?", buddy.ID).Find(&itineraryRequest).Error
	if err != nil {
		return
	}

	err = u.db.Where("id = ?", itineraryRequest.ItineraryFinderID).Find(&data).Error
	if err != nil {
		return
	}

	return
}
func (u *Usecase) ItineraryFinder(payload *model.ItineraryFinder) (err error) {
	return u.db.Create(payload).Error
}

func (u *Usecase) GetItineraryDestinations(payload *internal_model.ItineraryFinderRequest) (result internal_model.GetItineraryDestinationsResponse) {
	var data = model.Itinerary{}
	u.db.Where("(activity = ? OR  trip = ?) AND  date = ?", payload.Activity, payload.Trip, payload.Date).
		Where("id in (SELECT * FROM itinerary_buddies WHERE deleted_at is NULL)").
		Find(&data)

	var itineararyDestinations = []model.ItineraryDestination{}
	u.db.Where("itinerary_id = ?", data.ID).Preload("Destination").Find(&itineararyDestinations)

	var timeFormat = []string{"morning", "afternoon", "evening", "night"}
	var itinerary = []internal_model.GenerateItinerarySchema{
		{Time: "morning"}, {Time: "afternoon"}, {Time: "evening"}, {Time: "night"},
	}
	for _, elm := range itineararyDestinations {
		idx := slices.Index(timeFormat, elm.Time)
		itinerary[idx].DestinationIDs = append(itinerary[idx].DestinationIDs, elm.DestinationID)
		itinerary[idx].Destinations = append(itinerary[idx].Destinations, elm.Destination)
	}

	var user = model.User{}
	u.db.Where("id = ?", data.CreatedBy).Find(&user)

	var buddy = model.ItineraryBuddy{}
	u.db.Where("itinerary_id = ?", data.ID).Find(&buddy)

	return internal_model.GetItineraryDestinationsResponse{
		Description: buddy.Description,
		User:        user,
		Itinerary:   itinerary,
	}
}

func (u *Usecase) ItineraryJoin(payload *model.ItineraryRequest) (err error) {
	return u.db.Omit(clause.Associations).Create(payload).Error
}

func (u *Usecase) GetItineraryJoin(ID int) (err model.ItineraryRequest) {
	data := model.ItineraryRequest{}
	u.db.Where("id = ?", ID).Find(&data)
	return data
}
