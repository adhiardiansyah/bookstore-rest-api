package service

import (
	"log"

	"github.com/adhiardiansyah/bookstore-rest-api/dto"
	"github.com/adhiardiansyah/bookstore-rest-api/entity"
	"github.com/adhiardiansyah/bookstore-rest-api/repository"
	"github.com/mashingan/smapping"
)

type KeranjangService interface {
	AddToCart(k dto.AddToCartDTO, b entity.Buku) entity.Keranjang
	GetCartByUserID(UserID int) entity.Keranjang
}

type keranjangService struct {
	keranjangRepository repository.KeranjangRepository
}

func NewKeranjangService(keranjangRepo repository.KeranjangRepository) KeranjangService {
	return &keranjangService{
		keranjangRepository: keranjangRepo,
	}
}

func (service *keranjangService) AddToCart(k dto.AddToCartDTO, b entity.Buku) entity.Keranjang {
	keranjang := entity.Keranjang{}
	err := smapping.FillStruct(&keranjang, smapping.MapFields(&k))

	keranjang.BukuID = b.IDBuku
	keranjang.UserID = k.User.ID
	keranjang.Harga = b.Harga
	keranjang.Jumlah = keranjang.Jumlah + 1

	if err != nil {
		log.Fatalf("Gagal mapping, error: %v", err)
	}
	res := service.keranjangRepository.SaveKeranjang(keranjang)
	return res
}

func (service *keranjangService) GetCartByUserID(UserID int) entity.Keranjang {
	keranjang := service.keranjangRepository.FindByUserID(UserID)
	return keranjang
}
