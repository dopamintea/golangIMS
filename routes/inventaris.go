package routes

import (
	"dibimbing/config"
	"dibimbing/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InventarisRoutes(router *gin.Engine) {
	// Melihat stok untuk produk berdasarkan ID
	router.GET("/inventaris/:id_produk", func(c *gin.Context) {
		var inventaris models.Inventaris
		id := c.Param("id_produk")

		if err := config.DB.First(&inventaris, "id_produk = ?", id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Stok untuk produk ini tidak ditemukan"})
			return
		}

		c.JSON(http.StatusOK, inventaris)
	})

	// Memperbarui tingkat stok
	router.PUT("/inventaris/:id_produk", func(c *gin.Context) {
		var inventaris models.Inventaris
		id := c.Param("id_produk")

		if err := config.DB.First(&inventaris, "id_produk = ?", id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Stok untuk produk ini tidak ditemukan"})
			return
		}

		var input struct {
			Jumlah int `json:"jumlah"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		inventaris.Jumlah += input.Jumlah

		if inventaris.Jumlah < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Jumlah stok tidak boleh negatif"})
			return
		}

		config.DB.Save(&inventaris)
		c.JSON(http.StatusOK, inventaris)
	})
}
