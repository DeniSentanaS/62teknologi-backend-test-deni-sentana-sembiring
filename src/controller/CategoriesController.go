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

type categoriesController struct {
	categoriesService service.CategoriesService
}

func NewCategoriesController(categoriesService service.CategoriesService) *categoriesController {
	return &categoriesController{categoriesService}
}

func (controller *categoriesController) CreateCategoriesController(c *gin.Context) {
	var categoriesRequest request.CategoriesRequest
	err := c.ShouldBindJSON(&categoriesRequest)

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

	categories, err := controller.categoriesService.Create(categoriesRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categoryResponse(categories),
	})
}

func (controller *categoriesController) GetCategoriesController(c *gin.Context) {
	categories, err := controller.categoriesService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var categoriesResponse []response.CategoriesResponse
	for _, b := range categories {
		categoryResponse := categoryResponse(b)
		categoriesResponse = append(categoriesResponse, categoryResponse)

	}

	c.JSON(http.StatusOK, gin.H{
		"data": categoriesResponse,
	})

}

func (controller *categoriesController) GetCategorieController(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	categories, err := controller.categoriesService.FindById(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	categoriesResponse := categoryResponse(categories)

	c.JSON(http.StatusOK, gin.H{
		"data": categoriesResponse,
	})
}

func (controller *categoriesController) UpdateCategoriesController(c *gin.Context) {
	var categoriesRequest request.CategoriesRequest
	err := c.ShouldBindJSON(&categoriesRequest)

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
	categories, err := controller.categoriesService.Update(id, categoriesRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categoryResponse(categories),
	})
}

func (controller *categoriesController) DeleteCategoriesController(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	categories, err := controller.categoriesService.Delete(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	categoriesResponse := categoryResponse(categories)

	c.JSON(http.StatusOK, gin.H{
		"data": categoriesResponse,
	})
}

func categoryResponse(r model.CategoriesEntity) response.CategoriesResponse {
	return response.CategoriesResponse{
		Id:    r.Id,
		Alias: r.Alias,
		Title: r.Title,
	}
}
