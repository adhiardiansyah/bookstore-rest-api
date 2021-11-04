package service

import (
	"log"

	"github.com/adhiardiansyah/bookstore-rest-api/dto"
	"github.com/adhiardiansyah/bookstore-rest-api/entity"
	"github.com/adhiardiansyah/bookstore-rest-api/repository"
	"github.com/mashingan/smapping"
)

type KeranjangService interface {
	AddToCart(k dto.AddToCartDTO) entity.Keranjang
	UpdateCart(BukuID int, k dto.AddToCartDTO) entity.Keranjang
	GetCartByUserID(UserID int) entity.Keranjang
	FindByBukuID(BukuID int) bool
}

type keranjangService struct {
	keranjangRepository repository.KeranjangRepository
}

func NewKeranjangService(keranjangRepo repository.KeranjangRepository) KeranjangService {
	return &keranjangService{
		keranjangRepository: keranjangRepo,
	}
}

func (service *keranjangService) AddToCart(k dto.AddToCartDTO) entity.Keranjang {
	keranjang := entity.Keranjang{}

	err := smapping.FillStruct(&keranjang, smapping.MapFields(&k))

	keranjang.UserID = k.User.ID
	keranjang.Jumlah = keranjang.Jumlah + 1

	if err != nil {
		log.Fatalf("Gagal mapping, error: %v", err)
	}
	res := service.keranjangRepository.SaveKeranjang(keranjang)
	return res
}

func (service *keranjangService) UpdateCart(BukuID int, k dto.AddToCartDTO) entity.Keranjang {
	keranjang := service.keranjangRepository.FindByBukuID(BukuID)
	err := smapping.FillStruct(&keranjang, smapping.MapFields(&k))
	keranjang.UserID = k.User.ID
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

func (service *keranjangService) FindByBukuID(BukuID int) bool {
	res := service.keranjangRepository.FindByBukuID2(BukuID)
	return !(res.Error == nil)
}
