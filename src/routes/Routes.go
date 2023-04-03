package routes

import (
	"62teknologi-backend-test-deni-sentana-sembiring/src/controller"
	"62teknologi-backend-test-deni-sentana-sembiring/src/database"
	"62teknologi-backend-test-deni-sentana-sembiring/src/repository"
	"62teknologi-backend-test-deni-sentana-sembiring/src/service"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RoutesInit() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:4200"
		},
		MaxAge: 12 * time.Hour,
	}))

	v1 := router.Group("/api/v1")

	db := database.Db
	v1.GET("/")

	/**
	Business
	*/
	businessesRepository := repository.NewBusinessesRepository(db)
	businessesService := service.NewBusinessesService(businessesRepository)
	businessesController := controller.NewBusinessesController(businessesService)

	v1.GET("/businesses", businessesController.GetBusinessesController)
	v1.GET("/business/:id", businessesController.GetBusinessController)
	v1.POST("/businesses", businessesController.CreateBusinessesController)
	v1.PUT("/businesses/:id", businessesController.UpdateBusinessesController)
	v1.DELETE("/businesses/:id", businessesController.DeleteBusinessesController)

	/**
	Categories
	*/
	categoriesRepository := repository.NewCategoriesRepository(db)
	categoriesService := service.NewCategoriesService(categoriesRepository)
	categoriesController := controller.NewCategoriesController(categoriesService)

	v1.GET("/categories", categoriesController.GetCategoriesController)
	v1.GET("/category/:id", categoriesController.GetCategorieController)
	v1.POST("/categories", categoriesController.CreateCategoriesController)
	v1.PUT("/categories/:id", categoriesController.UpdateCategoriesController)
	v1.DELETE("/categories/:id", categoriesController.DeleteCategoriesController)

	/**
	Locations
	*/
	locationsRepository := repository.NewLocationsRepository(db)
	locationsService := service.NewLocationsService(locationsRepository)
	locationsController := controller.NewLocationsController(locationsService)

	v1.GET("/locations", locationsController.GetLocationsController)
	v1.GET("/location/:id", locationsController.GetLocationController)
	v1.POST("/locations", locationsController.CreateLocationsController)
	v1.PUT("/locations/:id", locationsController.UpdateLocationsController)
	v1.DELETE("/locations/:id", locationsController.DeleteLocationsController)

	/**
	Transactions
	*/
	transactiosRepository := repository.NewTransactionsRepository(db)
	transactiosService := service.NewTransactionsService(transactiosRepository)
	transactiosController := controller.NewTransactionsController(transactiosService)

	v1.GET("/transactions", transactiosController.GetTransactionsController)
	v1.GET("/transaction/:id", transactiosController.GetTransactionController)
	v1.POST("/transactions", transactiosController.CreateTransactionsController)
	v1.PUT("/transactions/:id", transactiosController.UpdateTransactionsController)
	v1.DELETE("/transactions/:id", transactiosController.DeleteTransactionsController)

	router.Run()
}
