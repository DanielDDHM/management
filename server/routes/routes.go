package routes

import (
	"github.com/DanielDDHM/management/internal/handlers"
	"github.com/DanielDDHM/management/server/middlewares"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1", middlewares.Secret())
	{
		sessions := main.Group("sessions")
		{
			sessions.GET("/:userId", handlers.ListSessiosByUser)
		}
		wallets := main.Group("wallets")
		{
			wallets.GET("/:userId", handlers.GetWallets)
		}
		transactions := main.Group("transactions")
		{
			transactions.GET("/:userId", handlers.ListTransactions)
		}
		deposits := main.Group("deposits")
		{
			deposits.GET("/:userId", handlers.ListDeposits)
		}
	}
	return router
}
