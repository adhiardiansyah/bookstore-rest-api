package dto

type UserUpdateDTO struct {
	ID       int    `json:"id" form:"id"`
	Nama     string `json:"nama" form:"nama" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password,omitempty" form:"password,omitempty"`
}
