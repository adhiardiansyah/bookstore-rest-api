package service

import (
	"log"

	"github.com/adhiardiansyah/bookstore-rest-api/dto"
	"github.com/adhiardiansyah/bookstore-rest-api/entity"
	"github.com/adhiardiansyah/bookstore-rest-api/repository"
	"github.com/mashingan/smapping"
)

type PengarangService interface {
	CreatePengarang(p dto.CreatePengarangDTO) entity.Pengarang
	UpdatePengarang(IDPengarang int, p dto.UpdatePengarangDTO) entity.Pengarang
	DeletePengarang(p entity.Pengarang)
	GetAll() []entity.Pengarang
	GetByID(IDPengarang int) entity.Pengarang
}

type pengarangService struct {
	pengarangRepository repository.PengarangRepository
}

func NewPengarangService(pengarangRepo repository.PengarangRepository) PengarangService {
	return &pengarangService{
		pengarangRepository: pengarangRepo,
	}
}

func (service *pengarangService) CreatePengarang(p dto.CreatePengarangDTO) entity.Pengarang {
	pengarang := entity.Pengarang{}
	err := smapping.FillStruct(&pengarang, smapping.MapFields(&p))
	if err != nil {
		log.Fatalf("Gagal mapping, error: %v", err)
	}
	res := service.pengarangRepository.SavePengarang(pengarang)
	return res
}

func (service *pengarangService) UpdatePengarang(IDPengarang int, p dto.UpdatePengarangDTO) entity.Pengarang {
	pengarang := service.pengarangRepository.FindByID(IDPengarang)
	err := smapping.FillStruct(&pengarang, smapping.MapFields(&p))
	if err != nil {
		log.Fatalf("Gagal mapping, error: %v", err)
	}
	res := service.pengarangRepository.UpdatePengarang(pengarang)
	return res
}

func (service *pengarangService) DeletePengarang(p entity.Pengarang) {
	service.pengarangRepository.DeletePengarang(p)
}

func (service *pengarangService) GetAll() []entity.Pengarang {
	return service.pengarangRepository.FindAll()
}

func (service *pengarangService) GetByID(IDPengarang int) entity.Pengarang {
	return service.pengarangRepository.FindByID(IDPengarang)
}
