package dto

type UpdateRoleDTO struct {
	NamaRole      string `json:"nama_role" form:"nama_role" binding:"required"`
	DeskripsiRole string `json:"deskripsi_role" form:"deskripsi_role" binding:"required"`
}

type CreateRoleDTO struct {
	NamaPenerbit  string `json:"nama_role" form:"nama_role" binding:"required"`
	DeskripsiRole string `json:"deskripsi_role" form:"deskripsi_role" binding:"required"`
}
