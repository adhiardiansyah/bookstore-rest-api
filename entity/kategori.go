package entity

import "time"

type Kategori struct {
	IDKategori   int    `gorm:"primary_key:auto_increment" json:"id_kategori"`
	NamaKategori string `gorm:"type:varchar(255)" json:"nama_kategori"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Buku         *[]Buku `json:"buku,omitempty"`
}
