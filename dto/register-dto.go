package dto

type RegisterDTO struct {
	Nama     string  `json:"nama" form:"nama" binding:"required"`
	Email    string  `json:"email" form:"email" binding:"required,email"`
	Password *string `json:"password" binding:"required"`
	RoleID   int     `json:"role_id,omitempty" form:"role_id,omitempty"`
}
