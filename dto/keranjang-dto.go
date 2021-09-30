package dto

import (
	"github.com/adhiardiansyah/bookstore-rest-api/entity"
)

type AddToCartDTO struct {
	BukuID int `json:"buku_id" form:"buku_id" binding:"required"`
	Harga  int `json:"harga" form:"harga" binding:"required"`
	Buku   entity.Buku
	User   entity.User
}
