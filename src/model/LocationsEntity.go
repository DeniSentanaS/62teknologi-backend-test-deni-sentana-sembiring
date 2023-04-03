package model

import (
	"time"
)

type LocationsEntity struct {
	Id        int     `json:"id" gorm:"primaryKey" gorm:"primaryKey"`
	Address1  string  `json:"address1" form:"address1" gorm:"not null"`
	Address2  string  `json:"address2" form:"address2"`
	Address3  string  `json:"address3" form:"address3"`
	City      string  `json:"city" form:"city" gorm:"not null"`
	ZipCode   string  `json:"zipCode" form:"zipCode" gorm:"not null"`
	Country   string  `json:"country" form:"country" gorm:"not null"`
	State     string  `json:"state" form:"state" gorm:"not null"`
	Latitude  float64 `json:"latitude" form:"latitude" gorm:"not null"`
	Longitude float64 `json:"longitude" form:"longitude" gorm:"not null"`
	CreateAt  time.Time
	UpdateAt  time.Time
}
