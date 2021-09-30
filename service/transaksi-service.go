package service

import (
	"log"

	"github.com/adhiardiansyah/bookstore-rest-api/dto"
	"github.com/adhiardiansyah/bookstore-rest-api/entity"
	"github.com/adhiardiansyah/bookstore-rest-api/repository"
	"github.com/mashingan/smapping"
)

type TransaksiService interface {
	CreateTransaksi(t dto.CreateTransaksiDTO, dt dto.CreateDetailTransaksiDTO) (entity.Transaksi, entity.DetailTransaksi)
}

type transaksiService struct {
	transaksiRepository repository.TransaksiRepository
}

func NewTransaksiService(transaksiRepo repository.TransaksiRepository) TransaksiService {
	return &transaksiService{
		transaksiRepository: transaksiRepo,
	}
}

func (service *transaksiService) CreateTransaksi(t dto.CreateTransaksiDTO, dt dto.CreateDetailTransaksiDTO) (entity.Transaksi, entity.DetailTransaksi) {
	transaksi := entity.Transaksi{}
	err := smapping.FillStruct(&transaksi, smapping.MapFields(&t))

	transaksi.UserID = t.User.ID
	transaksi.Status = "Pending"

	if err != nil {
		log.Fatalf("Gagal mapping, error: %v", err)
	}

	res1 := service.transaksiRepository.SaveTransaksi(transaksi)

	detailtransaksi := entity.DetailTransaksi{}
	detailtransaksi.TransaksiID = 1
	detailtransaksi.BukuID = 1
	detailtransaksi.Harga = 120000
	detailtransaksi.Jumlah = 1
	detailtransaksi.Total = detailtransaksi.Harga * detailtransaksi.Jumlah

	res2 := service.transaksiRepository.SaveDetailTransaksi(detailtransaksi)

	return res1, res2
}
