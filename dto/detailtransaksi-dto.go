package dto

import (
	"github.com/adhiardiansyah/bookstore-rest-api/entity"
)

type CreateDetailTransaksiDTO struct {
	Transaksi entity.Transaksi
	Keranjang entity.Keranjang
}
