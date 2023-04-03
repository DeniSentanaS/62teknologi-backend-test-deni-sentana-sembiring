package model

import "time"

type BusinessesEntity struct {
	Id             int     `json:"id" form:"id" gorm:"primaryKey"`
	Alias          string  `json:"alias" form:"alias"`
	Name           string  `json:"name" form:"name"`
	ImageUrl       string  `json:"imageUrl" form:"imageUrl"`
	IsClosed       bool    `json:"isClosed" form:"isClosed"`
	Url            string  `json:"url" form:"url"`
	ReviewCount    int     `json:"reviewCount" form:"reviewCount"`
	Rating         float64 `json:"rating" form:"rating"`
	Phone          string  `json:"phone" form:"phone"`
	Distance       float64 `json:"distance" form:"distance"`
	CreateAt       time.Time
	UpdateAt       time.Time
	Categories     []CategoriesEntity   `json:"categoriess" gorm:"many2many:businesses_categories"`
	Locations      []LocationsEntity    `json:"locations" gorm:"many2many:businesses_locations"`
	Transactions   []TransactionsEntity `json:"transactions" gorm:"many2many:businesses_transactions"`
	CategoriesId   []int                `json:"categoriesEntityId" form:"categoriesId" gorm:"-"`
	LocationsId    []int                `json:"locationsEntityId" form:"locationsId" gorm:"-"`
	TransactionsId []int                `json:"transactionsEntityId" form:"transactionsId" gorm:"-"`
}
