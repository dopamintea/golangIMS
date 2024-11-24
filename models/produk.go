package models

type Produk struct {
	IDProduk   uint    `gorm:"primaryKey;autoIncrement" json:"id_produk"`
	NamaProduk string  `gorm:"size:255;not null" json:"nama_produk"`
	Deskripsi  string  `gorm:"type:text" json:"deskripsi"`
	Harga      float64 `gorm:"type:decimal(10,2)" json:"harga"`
	Kategori   string  `gorm:"size:255" json:"kategori"`
}
