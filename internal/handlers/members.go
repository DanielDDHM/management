package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListMembers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "Return members",
	})
}
