package usecase

import (
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

func (u *Usecase) GetDestinations() (data []model.Destination, err error) {

	data = make([]model.Destination, 0)



	err = u.db.Find(&data).Error
	if err != nil {
		return
	}

	return
}

func (u *Usecase) GetItternaryDestination() (data []model.ItineraryDestination, err error) {

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

func (u *Usecase) GetChat() (data []model.ChatRoom,err error) {

	data = make([]model.ChatRoom, 0)

	err = u.db.Find(&data).Error
	if err != nil {
		return
	}

	return
}

func (u *Usecase) CreateItternaryBuddy(newItternaryBuddy model.ItineraryBuddy) (err error) {

	err = u.db.Create(&newItternaryBuddy).Error
	if err != nil {
		return
	}

	return
}

func (u *Usecase) GetItternaryById(id int) (data model.Itinerary,err error) {
	err = u.db.Where("id = ?",id).Find(&data).Error
	if err != nil {
		return
	}

	return
}

func (u *Usecase) GetUserById(id int) (data model.User,err error) {
	err = u.db.Where("id = ?",id).Find(&data).Error
	if err != nil {
		return
	}

	return
}

func (u *Usecase) GetChatById(id int) (data model.ChatRoom,err error) {
	err = u.db.Where("id = ?",id).Find(&data).Error
	if err != nil {
		return
	}

	return
}