package migration

import (
	"62teknologi-backend-test-deni-sentana-sembiring/src/database"
	"62teknologi-backend-test-deni-sentana-sembiring/src/model"
	"fmt"
)

func Migration() {
	err := database.Db.AutoMigrate(
		&model.BusinessesEntity{},
		&model.CategoriesEntity{},
		&model.LocationsEntity{},
		&model.TransactionsEntity{},
	)

	if err != nil {
		fmt.Println("can't run migration")
	}
	fmt.Println("migrated")
}
