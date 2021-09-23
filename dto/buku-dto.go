package dto

type UpdateBukuDTO struct {
	Isbn        string  `json:"isbn" form:"isbn" binding:"required"`
	Judul       string  `json:"judul" form:"judul" binding:"required"`
	KategoriID  int     `json:"kategori_id" form:"kategori_id" binding:"required,number"`
	PengarangID int     `json:"pengarang_id" form:"pengarang_id" binding:"required,number"`
	PenerbitID  int     `json:"penerbit_id" form:"penerbit_id" binding:"required,number"`
	TahunTerbit string  `json:"tahun_terbit" form:"tahun_terbit" binding:"required,number"`
	Stok        int     `json:"stok" form:"stok" binding:"required,number"`
	Berat       float64 `json:"berat" form:"berat" binding:"required,number"`
	Harga       int     `json:"harga" form:"harga" binding:"required,number"`
	Sinopsis    string  `json:"sinopsis" form:"sinopsis" binding:"required"`
}

type CreateBukuDTO struct {
	Isbn        string  `json:"isbn" form:"isbn" binding:"required"`
	Judul       string  `json:"judul" form:"judul" binding:"required"`
	KategoriID  int     `json:"kategori_id" form:"kategori_id" binding:"required,number"`
	PengarangID int     `json:"pengarang_id" form:"pengarang_id" binding:"required,number"`
	PenerbitID  int     `json:"penerbit_id" form:"penerbit_id" binding:"required,number"`
	TahunTerbit string  `json:"tahun_terbit" form:"tahun_terbit" binding:"required,number"`
	Stok        int     `json:"stok" form:"stok" binding:"required,number"`
	Berat       float64 `json:"berat" form:"berat" binding:"required,number"`
	Harga       int     `json:"harga" form:"harga" binding:"required,number"`
	Sinopsis    string  `json:"sinopsis" form:"sinopsis" binding:"required"`
}
