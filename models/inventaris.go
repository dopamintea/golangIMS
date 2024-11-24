package models

type Inventaris struct {
	IDProduk uint   `gorm:"primaryKey" json:"id_produk"`
	Jumlah   int    `gorm:"not null" json:"jumlah"`
	Lokasi   string `gorm:"size:255" json:"lokasi"`
}
