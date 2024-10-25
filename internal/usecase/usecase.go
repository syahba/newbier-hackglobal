package usecase

import (
	"fmt"
	chatgpt "newbier-hackglobal/pkg/chatGPT"
	"newbier-hackglobal/pkg/database/model"

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

func (u *Usecase) GenerateItinerary() (data []model.Destination, err error) {

	var ha, _ = u.ai.Generate(generateItineraryMessage())
	fmt.Println(ha)

	return
}

func (u *Usecase) GetDestinations() (data []model.Destination, err error) {

	data = make([]model.Destination, 0)

	err = u.db.Find(&data).Error
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
