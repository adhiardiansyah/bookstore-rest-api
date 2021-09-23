package service

import (
	"fmt"
	"log"

	"github.com/adhiardiansyah/bookstore-rest-api/dto"
	"github.com/adhiardiansyah/bookstore-rest-api/entity"
	"github.com/adhiardiansyah/bookstore-rest-api/repository"
	"github.com/gosimple/slug"
	"github.com/mashingan/smapping"
)

type BukuService interface {
	CreateBuku(b dto.CreateBukuDTO) entity.Buku
	SaveImageBuku(IDBuku int, fileLocation string) entity.Buku
	UpdateBuku(IDBuku int, b dto.UpdateBukuDTO) entity.Buku
	DeleteBuku(b entity.Buku)
	GetAll() []entity.Buku
	GetByID(IDBuku int) entity.Buku
}

type bukuService struct {
	bukuRepository repository.BukuRepository
}

func NewBukuService(bukuRepo repository.BukuRepository) BukuService {
	return &bukuService{
		bukuRepository: bukuRepo,
	}
}

func (service *bukuService) CreateBuku(b dto.CreateBukuDTO) entity.Buku {
	buku := entity.Buku{}
	err := smapping.FillStruct(&buku, smapping.MapFields(&b))

	slugCandidate := fmt.Sprintf("%s %d", b.Judul, b.PengarangID)
	buku.Slug = slug.Make(slugCandidate)

	if err != nil {
		log.Fatalf("Gagal mapping, error: %v", err)
	}
	res := service.bukuRepository.SaveBuku(buku)
	return res
}

func (service *bukuService) SaveImageBuku(IDBuku int, fileLocation string) entity.Buku {
	buku := service.bukuRepository.FindByID(IDBuku)
	buku.GambarBuku = fileLocation
	updatedFoto := service.bukuRepository.UpdateBuku(buku)
	return updatedFoto
}

func (service *bukuService) UpdateBuku(IDBuku int, b dto.UpdateBukuDTO) entity.Buku {
	buku := service.bukuRepository.FindByID(IDBuku)
	err := smapping.FillStruct(&buku, smapping.MapFields(&b))

	slugCandidate := fmt.Sprintf("%s %d", b.Judul, b.PengarangID)
	buku.Slug = slug.Make(slugCandidate)

	if err != nil {
		log.Fatalf("Gagal mapping, error: %v", err)
	}
	res := service.bukuRepository.UpdateBuku(buku)
	return res
}

func (service *bukuService) DeleteBuku(b entity.Buku) {
	service.bukuRepository.DeleteBuku(b)
}

func (service *bukuService) GetAll() []entity.Buku {
	return service.bukuRepository.FindAll()
}

func (service *bukuService) GetByID(IDBuku int) entity.Buku {
	return service.bukuRepository.FindByID(IDBuku)
}
