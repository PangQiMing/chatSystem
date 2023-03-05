package service

import (
	"chat/dto"
	"chat/entity"
	"chat/repository"
	"github.com/mashingan/smapping"
	"log"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
	ChangePass(userChangePass entity.User) entity.User
	FindByEmail(email string) entity.User
	VerifyCredential(email string, password string) interface{}
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(service repository.UserRepository) UserService {
	return &userService{userRepository: service}
}

func (service *userService) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	updateUser := service.userRepository.UpdateUser(userToUpdate)
	return updateUser
}

func (service *userService) Profile(userID string) entity.User {
	return service.userRepository.ProfileUser(userID)
}

func (service *userService) FindByEmail(email string) entity.User {
	return service.userRepository.FindByEmail(email)
}

func (service *userService) VerifyCredential(email string, password string) interface{} {
	return service.userRepository.VerifyCredential(email, password)
}

func (service *userService) ChangePass(UserChangePass entity.User) entity.User {
	changePass := service.userRepository.ChangePass(UserChangePass)
	return changePass
}
