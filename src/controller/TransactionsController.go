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

type transactionsController struct {
	transactionsService service.TransactionsService
}

func NewTransactionsController(transactionsService service.TransactionsService) *transactionsController {
	return &transactionsController{transactionsService}
}

func (controller *transactionsController) CreateTransactionsController(c *gin.Context) {
	var transactionsRequest request.TransactionsRequest
	err := c.ShouldBindJSON(&transactionsRequest)

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

	transactions, err := controller.transactionsService.Create(transactionsRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": transactionResponse(transactions),
	})
}

func (controller *transactionsController) GetTransactionsController(c *gin.Context) {
	transactions, err := controller.transactionsService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var transactionsResponse []response.TransactionsResponse
	for _, b := range transactions {
		transactionResponse := transactionResponse(b)
		transactionsResponse = append(transactionsResponse, transactionResponse)

	}

	c.JSON(http.StatusOK, gin.H{
		"data": transactionsResponse,
	})

}

func (controller *transactionsController) GetTransactionController(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	transactions, err := controller.transactionsService.FindById(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	transactionsResponse := transactionResponse(transactions)

	c.JSON(http.StatusOK, gin.H{
		"data": transactionsResponse,
	})
}

func (controller *transactionsController) UpdateTransactionsController(c *gin.Context) {
	var transactionsRequest request.TransactionsRequest
	err := c.ShouldBindJSON(&transactionsRequest)

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
	transactions, err := controller.transactionsService.Update(id, transactionsRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": transactionResponse(transactions),
	})
}

func (controller *transactionsController) DeleteTransactionsController(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	transactions, err := controller.transactionsService.Delete(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	transactionsResponse := transactionResponse(transactions)

	c.JSON(http.StatusOK, gin.H{
		"data": transactionsResponse,
	})
}

func transactionResponse(r model.TransactionsEntity) response.TransactionsResponse {
	return response.TransactionsResponse{
		Id:          r.Id,
		Transaction: r.Transaction,
	}
}
