package service

import (
	"log"

	"github.com/adhiardiansyah/bookstore-rest-api/dto"
	"github.com/adhiardiansyah/bookstore-rest-api/entity"
	"github.com/adhiardiansyah/bookstore-rest-api/repository"
	"github.com/mashingan/smapping"
)

type PenerbitService interface {
	CreatePenerbit(b dto.CreatePenerbitDTO) entity.Penerbit
	UpdatePenerbit(IDPenerbit int, b dto.UpdatePenerbitDTO) entity.Penerbit
	DeletePenerbit(b entity.Penerbit)
	GetAll() []entity.Penerbit
	GetByID(IDPenerbit int) entity.Penerbit
}

type penerbitService struct {
	penerbitRepository repository.PenerbitRepository
}

func NewPenerbitService(penerbitRepo repository.PenerbitRepository) PenerbitService {
	return &penerbitService{
		penerbitRepository: penerbitRepo,
	}
}

func (service *penerbitService) CreatePenerbit(b dto.CreatePenerbitDTO) entity.Penerbit {
	penerbit := entity.Penerbit{}
	err := smapping.FillStruct(&penerbit, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Gagal mapping, error: %v", err)
	}
	res := service.penerbitRepository.SavePenerbit(penerbit)
	return res
}

func (service *penerbitService) UpdatePenerbit(IDPenerbit int, b dto.UpdatePenerbitDTO) entity.Penerbit {
	penerbit := service.penerbitRepository.FindByID(IDPenerbit)
	err := smapping.FillStruct(&penerbit, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Gagal mapping, error: %v", err)
	}
	res := service.penerbitRepository.UpdatePenerbit(penerbit)
	return res
}

func (service *penerbitService) DeletePenerbit(b entity.Penerbit) {
	service.penerbitRepository.DeletePenerbit(b)
}

func (service *penerbitService) GetAll() []entity.Penerbit {
	return service.penerbitRepository.FindAll()
}

func (service *penerbitService) GetByID(IDPenerbit int) entity.Penerbit {
	return service.penerbitRepository.FindByID(IDPenerbit)
}
