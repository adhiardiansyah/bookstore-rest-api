package service

import (
	"log"

	"github.com/adhiardiansyah/bookstore-rest-api/dto"
	"github.com/adhiardiansyah/bookstore-rest-api/entity"
	"github.com/adhiardiansyah/bookstore-rest-api/repository"
	"github.com/mashingan/smapping"
)

type RoleService interface {
	CreateRole(r dto.CreateRoleDTO) entity.Role
	UpdateRole(IDRole int, r dto.UpdateRoleDTO) entity.Role
	DeleteRole(r entity.Role)
	GetAll() []entity.Role
	GetByID(IDRole int) entity.Role
}

type roleService struct {
	roleRepository repository.RoleRepository
}

func NewRoleService(roleRepo repository.RoleRepository) RoleService {
	return &roleService{
		roleRepository: roleRepo,
	}
}

func (service *roleService) CreateRole(r dto.CreateRoleDTO) entity.Role {
	role := entity.Role{}
	err := smapping.FillStruct(&role, smapping.MapFields(&r))
	if err != nil {
		log.Fatalf("Gagal mapping, error: %v", err)
	}
	res := service.roleRepository.SaveRole(role)
	return res
}

func (service *roleService) UpdateRole(IDRole int, r dto.UpdateRoleDTO) entity.Role {
	role := service.roleRepository.FindByID(IDRole)
	err := smapping.FillStruct(&role, smapping.MapFields(&r))
	if err != nil {
		log.Fatalf("Gagal mapping, error: %v", err)
	}
	res := service.roleRepository.UpdateRole(role)
	return res
}

func (service *roleService) DeleteRole(r entity.Role) {
	service.roleRepository.DeleteRole(r)
}

func (service *roleService) GetAll() []entity.Role {
	return service.roleRepository.FindAll()
}

func (service *roleService) GetByID(IDRole int) entity.Role {
	return service.roleRepository.FindByID(IDRole)
}
