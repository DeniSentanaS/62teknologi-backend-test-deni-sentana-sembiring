package controller

import (
	"62teknologi-backend-test-deni-sentana-sembiring/src/model"
	"62teknologi-backend-test-deni-sentana-sembiring/src/model/request"
	"62teknologi-backend-test-deni-sentana-sembiring/src/response"
	"62teknologi-backend-test-deni-sentana-sembiring/src/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type locationController struct {
	locationsService service.LocationsService
}

func NewLocationsController(locationsService service.LocationsService) *locationController {
	return &locationController{locationsService}
}

func (controller *locationController) CreateLocationsController(c *gin.Context) {
	var locationsRequest request.LocationsRequest
	err := c.ShouldBindJSON(&locationsRequest)

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

	locations, err := controller.locationsService.Create(locationsRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": locationResponse(locations),
	})
}

func (controller *locationController) GetLocationsController(c *gin.Context) {
	locations, err := controller.locationsService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var locationsResponse []response.LocationsResponse
	for _, b := range locations {
		locationResponse := locationResponse(b)
		locationsResponse = append(locationsResponse, locationResponse)

	}

	c.JSON(http.StatusOK, gin.H{
		"data": locationsResponse,
	})

}

func (controller *locationController) GetLocationController(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	locations, err := controller.locationsService.FindById(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	locationsResponse := locationResponse(locations)

	c.JSON(http.StatusOK, gin.H{
		"data": locationsResponse,
	})
}

func (controller *locationController) UpdateLocationsController(c *gin.Context) {
	var locationsRequest request.LocationsRequest
	err := c.ShouldBindJSON(&locationsRequest)

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
	locations, err := controller.locationsService.Update(id, locationsRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": locationResponse(locations),
	})
}

func (controller *locationController) DeleteLocationsController(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	locations, err := controller.locationsService.Delete(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	locationsResponse := locationResponse(locations)

	c.JSON(http.StatusOK, gin.H{
		"data": locationsResponse,
	})
}

func locationResponse(r model.LocationsEntity) response.LocationsResponse {
	return response.LocationsResponse{
		Id:        r.Id,
		Address1:  r.Address1,
		Address2:  r.Address2,
		Address3:  r.Address3,
		City:      r.City,
		ZipCode:   r.ZipCode,
		State:     r.State,
		Latitude:  r.Latitude,
		Longitude: r.Longitude,
	}
}
