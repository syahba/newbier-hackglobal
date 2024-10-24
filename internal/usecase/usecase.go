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
