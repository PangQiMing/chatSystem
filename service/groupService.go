package service

import (
	"chat/dto"
	"chat/entity"
	"chat/repository"
	"github.com/mashingan/smapping"
	"log"
)

type GroupService interface {
	Insert(groupCreateDTO dto.GroupCreateDTO) entity.Group
	Delete(groupID, groupLeaderID uint64) bool
	Update(groupUpdateDTO dto.GroupUpdateDTO) (entity.Group, error)
	MyGroup(groupLeaderID uint64) []entity.Group
}

type groupService struct {
	groupRepository repository.GroupRepository
}

func NewGroupService(groupRepository repository.GroupRepository) GroupService {
	return &groupService{groupRepository: groupRepository}
}

func (service *groupService) Insert(groupCreateDTO dto.GroupCreateDTO) entity.Group {
	var group entity.Group
	err := smapping.FillStruct(&group, smapping.MapFields(&groupCreateDTO))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	resGroup := service.groupRepository.InsertGroup(group)
	return resGroup
}

func (service *groupService) Delete(groupID, groupLeaderID uint64) bool {
	return service.groupRepository.DeleteGroup(groupID, groupLeaderID)
}

func (service *groupService) Update(groupUpdateDTO dto.GroupUpdateDTO) (entity.Group, error) {
	var group entity.Group
	err := smapping.FillStruct(&group, smapping.MapFields(&groupUpdateDTO))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	resGroup, updateErr := service.groupRepository.UpdateGroup(group)
	return resGroup, updateErr
}

func (service *groupService) MyGroup(groupLeaderID uint64) []entity.Group {
	return service.groupRepository.MyGroup(groupLeaderID)
}
