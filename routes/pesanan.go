package routes

import (
	"dibimbing/config"
	"dibimbing/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PesananRoutes(router *gin.Engine) {
	// Membuat pesanan baru
	router.POST("/pesanan", func(c *gin.Context) {
		var pesanan models.Pesanan
		if err := c.ShouldBindJSON(&pesanan); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Cek apakah produk ada
		var produk models.Produk
		if err := config.DB.First(&produk, pesanan.IDProduk).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
			return
		}

		// Simpan pesanan
		pesanan.TanggalPesanan = time.Now() // Set tanggal pesanan ke waktu saat ini
		config.DB.Create(&pesanan)

		c.JSON(http.StatusCreated, pesanan)
	})

	// Mengambil detail pesanan berdasarkan ID
	router.GET("/pesanan/:id", func(c *gin.Context) {
		var pesanan models.Pesanan
		id := c.Param("id")

		if err := config.DB.Preload("Produk").First(&pesanan, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Pesanan tidak ditemukan"})
			return
		}

		c.JSON(http.StatusOK, pesanan)
	})
}
