package entity

import "time"

type Pengarang struct {
	IDPengarang   int    `gorm:"primary_key:auto_increment" json:"id_pengarang"`
	NamaPengarang string `gorm:"type:varchar(255)" json:"nama_pengarang"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Buku          *[]Buku `json:"buku,omitempty"`
}
