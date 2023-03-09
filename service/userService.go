package service

import (
	"chat/entity"
	"chat/repository"
)

type UserService interface {
	UpdateUserStatus(user entity.User) entity.User
	Update(user entity.User) entity.User
	Profile(userID string) entity.User
	ChangePass(userChangePass entity.User) entity.User
	FindByEmail(email string) entity.User
	FindUserByID(id uint64) entity.User
	VerifyCredential(email string, password string) interface{}
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(service repository.UserRepository) UserService {
	return &userService{userRepository: service}
}

func (service *userService) UpdateUserStatus(user entity.User) entity.User {
	return service.userRepository.UpdateUserStatus(user)
}

func (service *userService) Update(user entity.User) entity.User {
	updateUser := service.userRepository.UpdateUser(user)
	return updateUser
}

func (service *userService) Profile(userID string) entity.User {
	return service.userRepository.ProfileUser(userID)
}

func (service *userService) FindByEmail(email string) entity.User {
	return service.userRepository.FindByEmail(email)
}

func (service *userService) FindUserByID(id uint64) entity.User {
	return service.userRepository.FindUserByID(id)
}

func (service *userService) VerifyCredential(email string, password string) interface{} {
	return service.userRepository.VerifyCredential(email, password)
}

func (service *userService) ChangePass(UserChangePass entity.User) entity.User {
	changePass := service.userRepository.ChangePass(UserChangePass)
	return changePass
}
