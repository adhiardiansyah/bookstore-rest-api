package dto

type UpdatePengarangDTO struct {
	NamaPengarang string `json:"nama_pengarang" form:"nama_pengarang" binding:"required"`
}

type CreatePengarangDTO struct {
	NamaPengarang string `json:"nama_pengarang" form:"nama_pengarang" binding:"required"`
}
