package main

import (
	"backendGalonku/config"
	"backendGalonku/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	//koneksi ke mangodb
	config.ConnectDB()

	//inisiasi router gin
	router := gin.Default()

	//setup routes
	routes.SetupRoutes(router)

	//jalankan server
	router.Run(":8080")
}