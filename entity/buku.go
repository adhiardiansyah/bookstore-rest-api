package entity

import "time"

type Buku struct {
	IDBuku      int     `gorm:"primary_key:auto_increment" json:"id_buku"`
	Isbn        string  `gorm:"type:varchar(255)" json:"isbn"`
	Judul       string  `gorm:"type:varchar(255)" json:"judul"`
	Slug        string  `gorm:"type:varchar(255)" json:"slug"`
	KategoriID  int     `gorm:"not null" json:"-"`
	PengarangID int     `gorm:"not null" json:"-"`
	PenerbitID  int     `gorm:"not null" json:"-"`
	TahunTerbit string  `gorm:"type:varchar(255)" json:"tahun_terbit"`
	Stok        int     `json:"stok"`
	Berat       float64 `json:"berat"`
	Harga       int     `json:"harga"`
	Sinopsis    string  `gorm:"type:text" json:"sinopsis"`
	GambarBuku  string  `gorm:"type:varchar(255)" json:"gambar_buku"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Kategori    Kategori  `gorm:"foreignkey:KategoriID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"kategori"`
	Pengarang   Pengarang `gorm:"foreignkey:PengarangID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"pengarang"`
	Penerbit    Penerbit  `gorm:"foreignkey:PenerbitID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"penerbit"`
}

// type BukuGambar struct {
// 	IDGambar  int    `gorm:"primary_key:auto_increment" json:"id_gambar"`
// 	BukuID    int    `gorm:"not null" json:"-"`
// 	NamaFile  string `gorm:"type:varchar(255)" json:"nama_file"`
// 	IsPrimary int    `json:"is_primary"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	// Buku      Buku `gorm:"foreignkey:BukuID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"buku"`
// }
