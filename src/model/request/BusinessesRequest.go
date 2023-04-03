package request

import "encoding/json"

type BusinessesRequest struct {
	Alias          string      `json:"alias" binding:"required"`
	Name           string      `json:"name" binding:"required"`
	ImageUrl       string      `json:"imageUrl" binding:"required"`
	IsClosed       bool        `json:"isClosed" binding:"required"`
	Url            string      `json:"url" binding:"required"`
	ReviewCount    json.Number `json:"reviewCount" binding:"required,number"`
	Rating         float64     `json:"rating" binding:"required"`
	Phone          string      `json:"phone" binding:"required"`
	Distance       float64     `json:"distance" binding:"required"`
	CategoriesId   []int       `json:"categoriesId" binding:"required"`
	LocationsId    []int       `json:"locationsId" binding:"required"`
	TransactionsId []int       `json:"transactionsId" binding:"required"`
}
