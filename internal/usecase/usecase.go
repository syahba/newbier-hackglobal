package usecase

import (
	"newbier-hackglobal/pkg/database/model"

	"gorm.io/gorm"
)

type Usecase struct {
	db *gorm.DB
}

func NewUsecase(db *gorm.DB) *Usecase {
	return &Usecase{db: db}
}

func (u *Usecase) GetUsecase() string {
	
	return "Hello World"
}

func (u *Usecase) GetDestinations() ([]model.Destination,error){
	var destinationList []model.Destination
	err := u.db.Find(&destinationList)

	if err != nil{
		return destinationList,err.Error
	}

	return destinationList,nil
}