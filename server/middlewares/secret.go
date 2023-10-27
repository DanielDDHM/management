package middlewares

import (
	"net/http"

	"github.com/DanielDDHM/management/config"
	"github.com/gin-gonic/gin"
)

func Secret() gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := c.GetHeader("x-api-secret")

		if secret == "" || config.GetConfig().ApiSecret != secret {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
