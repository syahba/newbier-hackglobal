package usecase

import (
	"encoding/json"
	"fmt"
	chatgpt "newbier-hackglobal/pkg/chatGPT"
	"newbier-hackglobal/pkg/database/model"
	"strings"
	"time"

	"math/rand"

	"gorm.io/gorm"
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

func (u *Usecase) GenerateItinerary(activity,trip string) (data []GenerateItinerarySchema, err error) {
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

	data = make([]GenerateItinerarySchema,0)
	json.Unmarshal([]byte(cleanedInput), &data)

	// using index looping to modify the actual value
	for index1:=0;index1<len(data);index1++{
		for index2:=0;index2<len(data[index1].DestinationIDs);index2++{
			destination,_ := u.GetDestinationById(data[index1].DestinationIDs[index2])
			data[index1].Destinations = append(data[index1].Destinations, destination)
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

func (u *Usecase) GetDestinationById(id int) (data model.Destination,err error){
	err = u.db.Where("id = ?",id).Find(&data).Error
	if err != nil{
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

	ItineraryBuddy.IsAccept = true

	err = u.db.Save(&ItineraryBuddy).Error
	if err != nil {
		return err
	}

	return nil
}
