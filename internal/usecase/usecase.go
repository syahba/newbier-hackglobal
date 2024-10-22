package usecase

import (
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
