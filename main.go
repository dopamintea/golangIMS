package main

import (
	"dibimbing/config"
	"dibimbing/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDB()

	// Inisialisasi router Gin
	router := gin.Default()

	routes.ProdukRoutes(router)
	routes.InventarisRoutes(router)
	routes.PesananRoutes(router)

	// Jalankan server
	router.Run(":8080")
}
