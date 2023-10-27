package handlers

import (
	"net/http"
	"strconv"

	"github.com/DanielDDHM/management/internal/repositories"
	"github.com/DanielDDHM/management/internal/services"
	"github.com/gin-gonic/gin"
)

func ListTransactions(c *gin.Context) {
	id := c.Param("userId")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	page, size, _, sortOrder, sortOrderName, startDate, endDate := getPagination(c)

	repository := &repositories.TransactionRepository{}

	service := services.NewTransactionService(repository)

	transactions, err := service.GetAll(newid, page, size, sortOrder,
		sortOrderName, startDate, endDate)

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, successResponse(transactions))
}
