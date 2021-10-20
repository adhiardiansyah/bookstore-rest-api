package service

import (
	"log"

	"github.com/adhiardiansyah/bookstore-rest-api/dto"
	"github.com/adhiardiansyah/bookstore-rest-api/entity"
	"github.com/adhiardiansyah/bookstore-rest-api/repository"
	"github.com/mashingan/smapping"
)

type KeranjangService interface {
	AddToCart(BukuID int, k dto.AddToCartDTO) entity.Keranjang
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

func (service *keranjangService) AddToCart(BukuID int, k dto.AddToCartDTO) entity.Keranjang {
	// keranjang1 := entity.Keranjang{}
	// keranjang2, ok := service.keranjangRepository.FindByBukuID(BukuID)

	keranjang := entity.Keranjang{}
	// if ok {
	// 	keranjang = keranjang2
	// } else {
	// 	keranjang = keranjang1
	// }

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
