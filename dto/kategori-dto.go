package dto

type UpdateKategoriDTO struct {
	NamaKategori string `json:"nama_kategori" form:"nama_kategori" binding:"required"`
}

type CreateKategoriDTO struct {
	NamaKategori string `json:"nama_kategori" form:"nama_kategori" binding:"required"`
}
