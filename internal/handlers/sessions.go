package handlers

import (
	"net/http"
	"strconv"

	"github.com/DanielDDHM/management/internal/repositories"
	"github.com/DanielDDHM/management/internal/services"
	"github.com/gin-gonic/gin"
)

func ListSessiosByUser(c *gin.Context) {
	id := c.Param("userId")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	repository := &repositories.SessionRepository{}

	service := services.NewSessionService(repository)

	sessions, err := service.ListSessionsByUser(newid)

	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, successResponse(sessions))
}
