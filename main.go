package main

import (
	"62teknologi-backend-test-deni-sentana-sembiring/src/database"
	"62teknologi-backend-test-deni-sentana-sembiring/src/database/migration"
	"62teknologi-backend-test-deni-sentana-sembiring/src/routes"
)

func main() {
	// Initialization Database
	database.DatabaseInit()

	// Migration Database
	migration.Migration()

	routes.RoutesInit()
}

// func CORS(c *gin.Context) {
// 	c.Header("Access-Control-Allow-Origin", "http://localhost:4200")
// 	c.Header("Access-Control-Allow-Methods", "*")
// 	c.Header("Access-Control-Allow-Headers", "*")
// 	c.Header("Content-Type", "application/json")

// 	if c.Request.Method != "OPTIONS" {

// 		c.Next()

// 	} else {
// 		c.AbortWithStatus(http.StatusOK)
// 	}
// }
