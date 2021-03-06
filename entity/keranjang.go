package entity

import "time"

type Keranjang struct {
	IDKeranjang int `gorm:"primary_key:auto_increment" json:"id_keranjang"`
	BukuID      int `gorm:"not null" json:"buku_id"`
	UserID      int `gorm:"not null" json:"user_id"`
	Harga       int `gorm:"uniqueIndex" json:"harga"`
	Jumlah      int `json:"jumlah"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Buku        Buku `gorm:"foreignkey:BukuID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"buku"`
	User        User `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
