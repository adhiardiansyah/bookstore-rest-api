package service

import (
	"log"

	"github.com/adhiardiansyah/bookstore-rest-api/dto"
	"github.com/adhiardiansyah/bookstore-rest-api/entity"
	"github.com/adhiardiansyah/bookstore-rest-api/repository"
	"github.com/mashingan/smapping"
)

type UserService interface {
	Update(ID int, user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
	SaveImageUser(ID int, fileLocation string) entity.User
	GetAll() []entity.User
	GetUserByID(ID int) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(ID int, user dto.UserUpdateDTO) entity.User {
	userToUpdate := service.userRepository.FindByID(ID)
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Gagal mapping %v", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *userService) Profile(userID string) entity.User {
	return service.userRepository.ProfileUser(userID)
}

func (service *userService) SaveImageUser(ID int, fileLocation string) entity.User {
	user := service.userRepository.FindByID(ID)
	user.GambarUser = fileLocation
	updatedFoto := service.userRepository.UpdateImageUser(user)
	return updatedFoto
}

func (service *userService) GetAll() []entity.User {
	return service.userRepository.FindAll()
}

func (service *userService) GetUserByID(ID int) entity.User {
	user := service.userRepository.FindByID(ID)
	return user
}
