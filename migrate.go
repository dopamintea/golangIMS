package main

import (
	"dibimbing/config"
	"dibimbing/models"
	"log"
)

func main() {
	// Inisialisasi koneksi database
	config.ConnectDB()

	// Migrasi model
	err := config.DB.AutoMigrate(&models.Produk{}, &models.Inventaris{}, &models.Pesanan{})
	if err != nil {
		log.Fatal("Gagal melakukan migrasi:", err)
	}

	log.Println("Migrasi selesai.")

	insertSampleData()
}

func insertSampleData() {
	// Insert sample data for Produk
	products := []models.Produk{
		{NamaProduk: "Laptop", Deskripsi: "Laptop gaming dengan spesifikasi tinggi", Harga: 15000000.00, Kategori: "Elektronik"},
		{NamaProduk: "Smartphone", Deskripsi: "Smartphone dengan kamera terbaik", Harga: 7000000.00, Kategori: "Elektronik"},
		{NamaProduk: "Meja Kerja", Deskripsi: "Meja kerja kayu solid", Harga: 3000000.00, Kategori: "Furniture"},
	}
	for _, product := range products {
		config.DB.Create(&product)
	}

	// Insert sample data for Inventaris
	inventories := []models.Inventaris{
		{IDProduk: 1, Jumlah: 50, Lokasi: "Gudang Utama"},
		{IDProduk: 2, Jumlah: 30, Lokasi: "Gudang Utama"},
		{IDProduk: 3, Jumlah: 20, Lokasi: "Gudang Cabang"},
	}
	for _, inventory := range inventories {
		config.DB.Create(&inventory)
	}

	// Insert sample data for Pesanan
	// orders := []models.Pesanan{
	// 	{IDProduk: 1, Jumlah: 2, TanggalPesanan: "2024-11-01"},
	// 	{IDProduk: 2, Jumlah: 1, TanggalPesanan: "2024-11-10"},
	// 	{IDProduk: 1, Jumlah: 1, TanggalPesanan: "2024-11-15"},
	// 	{IDProduk: 3, Jumlah: 5, TanggalPesanan: "2024-11-20"},
	// }
	// for _, order := range orders {
	// 	config.DB.Create(&order)
	// }

	log.Println("Sample data inserted successfully.")
}
