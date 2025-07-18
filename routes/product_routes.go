package routes

import (
	"backendGalonku/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoute(router *gin.Engine) {
	router.POST("/products", controllers.CreateProduct)
	router.GET("/products", controllers.GetAllProducts)
	// router.GET("/products/:id", controllers.GetProductByID)
	router.PUT("/products/:id", controllers.UpdateProduct)
	router.DELETE("/products/:id", controllers.DeleteProduct)
}
