package routes

import (
	"backendGalonku/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/register", controllers.RegisterUser)
		
		// Tambahkan route untuk produk
		api.POST("/products", controllers.CreateProduct)
		api.GET("/products", controllers.GetAllProducts)
		api.PUT("/products/:id", controllers.UpdateProduct)
		api.DELETE("/products/:id", controllers.DeleteProduct)

		// Tambahkan route untuk transaksi
		api.POST("/transactions", controllers.CreateTransaction)
	}

}