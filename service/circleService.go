package service

import (
	"chat/dto"
	"chat/entity"
	"chat/repository"
	"fmt"
	"github.com/mashingan/smapping"
	"log"
)

type MomentService interface {
	Insert(momentCreateDTO dto.CircleCreateDTO) entity.Circle
	Update(momentUpdateDTO dto.CircleUpdateDTO) entity.Circle
	Delete(moment entity.Circle)
	All(userID uint64) []entity.Circle
	FindByID(momentID uint64) entity.Circle
	IsAllowedToEdit(userID string, momentID uint64) bool
}

type momentService struct {
	momentRepository repository.MomentRepository
}

func NewBookService(momentRepository repository.MomentRepository) MomentService {
	return &momentService{momentRepository: momentRepository}
}

func (service *momentService) Insert(momentCreateDTO dto.CircleCreateDTO) entity.Circle {
	var moment entity.Circle
	err := smapping.FillStruct(&moment, smapping.MapFields(&momentCreateDTO))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	resMoment := service.momentRepository.InsertBook(moment)
	return resMoment
}

func (service *momentService) Update(momentUpdateDTO dto.CircleUpdateDTO) entity.Circle {
	var moment entity.Circle
	err := smapping.FillStruct(&moment, smapping.MapFields(&momentUpdateDTO))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	resMoment := service.momentRepository.UpdateBook(moment)
	return resMoment
}

func (service *momentService) Delete(moment entity.Circle) {
	service.momentRepository.DeleteBook(moment)
}

func (service *momentService) All(userID uint64) []entity.Circle {
	return service.momentRepository.AllBook(userID)
}

func (service *momentService) FindByID(momentID uint64) entity.Circle {
	return service.momentRepository.FindBookByID(momentID)
}

func (service *momentService) IsAllowedToEdit(userID string, momentID uint64) bool {
	moment := service.momentRepository.FindBookByID(momentID)
	id := fmt.Sprintf("%v", moment.UserID)
	return userID == id
}
