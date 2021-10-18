package dto

type UpdatePenerbitDTO struct {
	NamaPenerbit string `json:"nama_penerbit" form:"nama_penerbit" binding:"required"`
	KotaPenerbit string `json:"kota_penerbit" form:"kota_penerbit" binding:"required"`
}

type CreatePenerbitDTO struct {
	NamaPenerbit string `json:"nama_penerbit" form:"nama_penerbit" binding:"required"`
	KotaPenerbit string `json:"kota_penerbit" form:"kota_penerbit" binding:"required"`
}
