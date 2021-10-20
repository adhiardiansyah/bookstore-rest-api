package dto

import (
	"github.com/adhiardiansyah/bookstore-rest-api/entity"
)

type AddToCartDTO struct {
	BukuID int `json:"buku_id" form:"buku_id" binding:"required,number"`
	Harga  int `json:"harga" form:"harga" binding:"required,number"`
	User   entity.User
}
