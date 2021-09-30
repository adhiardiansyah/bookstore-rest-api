package entity

import "time"

type Transaksi struct {
	IDTransaksi     int    `gorm:"primary_key:auto_increment" json:"id_transaksi"`
	UserID          int    `gorm:"not null" json:"user_id"`
	KodeTransaksi   string `gorm:"type:varchar(255)" json:"kode_transaksi"`
	Destinasi       string `gorm:"type:varchar(255)" json:"destinasi"`
	Status          string `gorm:"type:varchar(255)" json:"status"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	User            User               `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	DetailTransaksi *[]DetailTransaksi `json:"detail_transaksi,omitempty"`
}

type DetailTransaksi struct {
	IDDetailTransaksi int `gorm:"primary_key:auto_increment" json:"id_detailTransaksi"`
	TransaksiID       int `gorm:"not null" json:"transaksi_id"`
	BukuID            int `gorm:"not null" json:"buku_id"`
	Harga             int `gorm:"not null" json:"harga"`
	Jumlah            int `json:"jumlah"`
	Total             int `json:"total"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Transaksi         Transaksi `gorm:"foreignkey:TransaksiID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"transaksi"`
	Buku              Buku      `gorm:"foreignkey:BukuID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"buku"`
}
