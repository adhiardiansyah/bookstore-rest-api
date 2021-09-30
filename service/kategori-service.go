package service

import (
	"log"

	"github.com/adhiardiansyah/bookstore-rest-api/dto"
	"github.com/adhiardiansyah/bookstore-rest-api/entity"
	"github.com/adhiardiansyah/bookstore-rest-api/repository"
	"github.com/mashingan/smapping"
)

type KategoriService interface {
	CreateKategori(b dto.CreateKategoriDTO) entity.Kategori
	UpdateKategori(IDKategori int, b dto.UpdateKategoriDTO) entity.Kategori
	DeleteKategori(b entity.Kategori)
	GetAll() []entity.Kategori
	GetByID(IDKategori int) entity.Kategori
}

type kategoriService struct {
	kategoriRepository repository.KategoriRepository
}

func NewKategoriService(kategoriRepo repository.KategoriRepository) KategoriService {
	return &kategoriService{
		kategoriRepository: kategoriRepo,
	}
}

func (service *kategoriService) CreateKategori(b dto.CreateKategoriDTO) entity.Kategori {
	kategori := entity.Kategori{}
	err := smapping.FillStruct(&kategori, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Gagal mapping, error: %v", err)
	}
	res := service.kategoriRepository.SaveKategori(kategori)
	return res
}

func (service *kategoriService) UpdateKategori(IDKategori int, b dto.UpdateKategoriDTO) entity.Kategori {
	kategori := service.kategoriRepository.FindByID(IDKategori)
	err := smapping.FillStruct(&kategori, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Gagal mapping, error: %v", err)
	}
	res := service.kategoriRepository.UpdateKategori(kategori)
	return res
}

func (service *kategoriService) DeleteKategori(b entity.Kategori) {
	service.kategoriRepository.DeleteKategori(b)
}

func (service *kategoriService) GetAll() []entity.Kategori {
	return service.kategoriRepository.FindAll()
}

func (service *kategoriService) GetByID(IDKategori int) entity.Kategori {
	return service.kategoriRepository.FindByID(IDKategori)
}
