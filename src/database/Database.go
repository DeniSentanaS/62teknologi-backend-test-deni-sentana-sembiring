package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func DatabaseInit() {
	var err error
	dsn := "host=localhost user=postgres password=root dbname=enamdua port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection failed")
		return
	}
	fmt.Println("Connected to Database")
}
