package handlers

import (
	"net/http"
	"strconv"

	"github.com/DanielDDHM/management/internal/repositories"
	"github.com/DanielDDHM/management/internal/services"
	"github.com/gin-gonic/gin"
)

func ListDeposits(c *gin.Context) {
	id := c.Param("userId")

	page, size, _, _, _, _, _ := getPagination(c)
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	repository := &repositories.DepositRepository{}

	service := services.NewDepositService(repository)

	curr_type := c.Query("type")
	deposits, err := service.ListDeposits(newid, curr_type, page, size)

	c.JSON(http.StatusOK, successResponse(deposits))
}
