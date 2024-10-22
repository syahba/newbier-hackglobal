package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name    string `json:"name"`
	Country string `json:"country"`
	Gender  string `json:"gender"`
}

func (User) TableName() string {
	return "users"
}

type Destination struct {
	gorm.Model
	Name    string `json:"name"`
	Type    string `json:"type"`
	Star    string `json:"star"`
	Address string `json:"address"`
	GmapUrl string `json:"gmap_url"`
	Image   string `json:"image"`
}

func (Destination) TableName() string {
	return "destinations"
}
