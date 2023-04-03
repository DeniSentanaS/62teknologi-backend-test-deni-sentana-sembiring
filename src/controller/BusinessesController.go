package controller

import (
	"62teknologi-backend-test-deni-sentana-sembiring/src/database"
	"62teknologi-backend-test-deni-sentana-sembiring/src/model"
	"62teknologi-backend-test-deni-sentana-sembiring/src/model/join"
	"62teknologi-backend-test-deni-sentana-sembiring/src/model/request"
	"62teknologi-backend-test-deni-sentana-sembiring/src/response"
	"62teknologi-backend-test-deni-sentana-sembiring/src/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type businessesController struct {
	businessesService service.BusinessesService
}

func NewBusinessesController(businessesService service.BusinessesService) *businessesController {
	return &businessesController{businessesService}
}

func (controller *businessesController) CreateBusinessesController(c *gin.Context) {
	var businessRequest request.BusinessesRequest
	err := c.ShouldBindJSON(&businessRequest)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	business, err := controller.businessesService.Create(businessRequest)

	for _, categoriesId := range business.CategoriesId {
		bussinesCategories := new(join.BusinessesCategories)
		bussinesCategories.BusinessesEntityId = business.Id
		bussinesCategories.CategoriesEntityId = categoriesId
		database.Db.Create(bussinesCategories)
	}

	for _, locationsId := range business.LocationsId {
		bussinesLocations := new(join.BusinessesLocations)
		bussinesLocations.BusinessesEntityId = business.Id
		bussinesLocations.LocationsEntityId = locationsId
		database.Db.Create(bussinesLocations)
	}

	for _, transactionsId := range business.TransactionsId {
		bussinesTransaction := new(join.BusinessesTransactions)
		bussinesTransaction.BusinessesEntityId = business.Id
		bussinesTransaction.TransactionsEntityId = transactionsId
		database.Db.Create(bussinesTransaction)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": businessResponse(business),
	})
}

func (controller *businessesController) GetBusinessesController(c *gin.Context) {
	businesses, err := controller.businessesService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var businessesResponse []response.BusinessesResponse
	for _, b := range businesses {
		businessResponse := businessResponse(b)
		businessesResponse = append(businessesResponse, businessResponse)

	}

	c.JSON(http.StatusOK, gin.H{
		"data": businessesResponse,
	})

}

func (controller *businessesController) GetBusinessController(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	business, err := controller.businessesService.FindById(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	businessResponse := businessResponse(business)

	c.JSON(http.StatusOK, gin.H{
		"data": businessResponse,
	})
}

func (controller *businessesController) UpdateBusinessesController(c *gin.Context) {
	var businessRequest request.BusinessesRequest
	err := c.ShouldBindJSON(&businessRequest)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	business, err := controller.businessesService.Update(id, businessRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": businessResponse(business),
	})
}

func (controller *businessesController) DeleteBusinessesController(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	business, err := controller.businessesService.Delete(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	businessResponse := businessResponse(business)

	c.JSON(http.StatusOK, gin.H{
		"data": businessResponse,
	})
}

func businessResponse(r model.BusinessesEntity) response.BusinessesResponse {

	return response.BusinessesResponse{
		Id:          r.Id,
		Alias:       r.Alias,
		Name:        r.Name,
		ImageUrl:    r.ImageUrl,
		IsClosed:    r.IsClosed,
		Url:         r.Url,
		ReviewCount: r.ReviewCount,
		Rating:      r.Rating,
		Phone:       r.Phone,
		Distance:    r.Distance,
	}
}
