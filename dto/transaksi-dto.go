package dto

import (
	"github.com/adhiardiansyah/bookstore-rest-api/entity"
)

type CreateTransaksiDTO struct {
	IDTransaksi   int
	KodeTransaksi string `json:"kode_transaksi" form:"kode_transaksi" binding:"required"`
	Destinasi     string `json:"destinasi" form:"destinasi" binding:"required"`
	User          entity.User
}
