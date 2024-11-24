package routes

import (
	"dibimbing/config"
	"dibimbing/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProdukRoutes(router *gin.Engine) {
	// Menambahkan produk baru
	router.POST("/produk", func(c *gin.Context) {
		var produk models.Produk
		if err := c.ShouldBindJSON(&produk); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		config.DB.Create(&produk)
		c.JSON(http.StatusCreated, produk)
	})

	// Lihat semua produk
	router.GET("/produk", func(c *gin.Context) {
		var produk []models.Produk
		if err := config.DB.Find(&produk).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
			return
		}
		c.JSON(http.StatusOK, produk)
	})

	// Lihat produk berdasarkan ID
	router.GET("/produk/:id", func(c *gin.Context) {
		var produk models.Produk
		id := c.Param("id")

		if err := config.DB.First(&produk, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
			return
		}
		c.JSON(http.StatusOK, produk)
	})

	// Lihat produk berdasarkan kategori
	router.GET("/produk/kategori/:kategori", func(c *gin.Context) {
		var produk []models.Produk
		kategori := c.Param("kategori")
		config.DB.Where("kategori = ?", kategori).Find(&produk)
		c.JSON(http.StatusOK, produk)
	})

	// Perbarui produk
	router.PUT("/produk/:id", func(c *gin.Context) {
		var produk models.Produk
		id := c.Param("id")
		if err := config.DB.First(&produk, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
			return
		}

		if err := c.ShouldBindJSON(&produk); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		config.DB.Save(&produk)
		c.JSON(http.StatusOK, produk)
	})

	// Hapus produk
	router.DELETE("/produk/:id", func(c *gin.Context) {
		var produk models.Produk
		id := c.Param("id")
		if err := config.DB.First(&produk, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
			return
		}
		config.DB.Delete(&produk)
		c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil dihapus"})
	})
}
