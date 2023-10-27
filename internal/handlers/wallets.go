package handlers

import (
	"net/http"
	"strconv"

	"github.com/DanielDDHM/management/internal/repositories"
	"github.com/DanielDDHM/management/internal/services"

	"github.com/gin-gonic/gin"
)

func GetWallets(c *gin.Context) {
	id := c.Param("userId")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	repository := &repositories.WalletRepository{}

	service := services.NewWalletService(repository)

	wallets, err := service.GetWallets(newid)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, successResponse(wallets))
}
