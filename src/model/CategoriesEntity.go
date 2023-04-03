package model

import "time"

type CategoriesEntity struct {
	Id       int    `json:"id" form:"id" gorm:"primaryKey"`
	Alias    string `json:"alias" form:"alias" gorm:"not null"`
	Title    string `json:"title" form:"title" gorm:"not null"`
	CreateAt time.Time
	UpdateAt time.Time
}
