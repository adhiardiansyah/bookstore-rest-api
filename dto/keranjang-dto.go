package dto

import (
	"github.com/adhiardiansyah/bookstore-rest-api/entity"
)

type AddToCartDTO struct {
	Buku entity.Buku
	User entity.User
}
