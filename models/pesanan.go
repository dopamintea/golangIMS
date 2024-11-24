package models

import "time"

type Pesanan struct {
	IDPesanan      uint      `gorm:"primaryKey;autoIncrement" json:"id_pesanan"`
	IDProduk       uint      `gorm:"not null" json:"id_produk"`
	Jumlah         int       `gorm:"not null" json:"jumlah"`
	TanggalPesanan time.Time `gorm:"not null" json:"tanggal_pesanan"`

	Produk Produk `gorm:"foreignKey:IDProduk" json:"produk"`
}
