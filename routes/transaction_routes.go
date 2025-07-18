package routes

import (
	"backendGalonku/controllers" // ganti dengan nama modul kamu

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(router *gin.Engine) {
	transaction := router.Group("/api/transactions")
	{
		transaction.POST("/", controllers.CreateTransaction)

	}
}
