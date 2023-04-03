package model

import "time"

type TransactionsEntity struct {
	Id          int    `json:"" form:"" gorm:"primaryKey"`
	Transaction string `json:"transaction" form:"transaction" gorm:"not null"`
	CreateAt    time.Time
	UpdateAt    time.Time
}
