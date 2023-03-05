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

type circleService struct {
	circleRepository repository.CircleRepository
}

func NewCircleService(circleRepository repository.CircleRepository) MomentService {
	return &circleService{circleRepository: circleRepository}
}

func (service *circleService) Insert(momentCreateDTO dto.CircleCreateDTO) entity.Circle {
	var circle entity.Circle
	err := smapping.FillStruct(&circle, smapping.MapFields(&momentCreateDTO))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	result := service.circleRepository.Insert(circle)
	return result
}

func (service *circleService) Update(circleUpdateDTO dto.CircleUpdateDTO) entity.Circle {
	var moment entity.Circle
	err := smapping.FillStruct(&moment, smapping.MapFields(&circleUpdateDTO))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	resMoment := service.circleRepository.Update(moment)
	return resMoment
}

func (service *circleService) Delete(circle entity.Circle) {
	service.circleRepository.Delete(circle)
}

func (service *circleService) All(userID uint64) []entity.Circle {
	return service.circleRepository.All(userID)
}

func (service *circleService) FindByID(circleID uint64) entity.Circle {
	return service.circleRepository.FindByID(circleID)
}

func (service *circleService) IsAllowedToEdit(userID string, circleID uint64) bool {
	circle := service.circleRepository.FindByID(circleID)
	id := fmt.Sprintf("%v", circle.UserID)
	return userID == id
}
