package entity

import "time"

type Penerbit struct {
	IDPenerbit   int    `gorm:"primary_key:auto_increment" json:"id_penerbit"`
	NamaPenerbit string `gorm:"type:varchar(255)" json:"nama_penerbit"`
	KotaPenerbit string `gorm:"type:varchar(255)" json:"kota_penerbit"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Buku         *[]Buku `json:"buku,omitempty"`
}
