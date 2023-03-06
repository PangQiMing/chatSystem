package service

import (
	"chat/dto"
	"chat/entity"
	"chat/repository"
	"fmt"
	"github.com/mashingan/smapping"
	"log"
)

type CircleService interface {
	Insert(circleCreateDTO dto.CircleCreateDTO) entity.Circle
	Delete(circle entity.Circle)
	All(userID uint64) []entity.Circle
	IsAllowedToDelete(userID string, circleID uint64) bool
}

type circleService struct {
	circleRepository repository.CircleRepository
}

func NewCircleService(circleRepository repository.CircleRepository) CircleService {
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

func (service *circleService) Delete(circle entity.Circle) {
	service.circleRepository.Delete(circle)
}

func (service *circleService) All(userID uint64) []entity.Circle {
	return service.circleRepository.All(userID)
}

func (service *circleService) IsAllowedToDelete(userID string, circleID uint64) bool {
	circle := service.circleRepository.FindCircleByID(circleID)
	id := fmt.Sprintf("%v", circle.ID)
	return userID == id
}
