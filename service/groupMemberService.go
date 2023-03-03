package service

import (
	"chat/dto"
	"chat/entity"
	"chat/repository"
	"github.com/mashingan/smapping"
	"log"
)

type GroupMembersService interface {
	Insert(groupMemberDTO dto.GroupMembersCreateDTO) entity.GroupMembers
	Delete(members entity.GroupMembers)
}

type groupMembersService struct {
	groupMembersRepository repository.GroupMemberRepository
}

func NewGroupMembersService(groupMemberRepository repository.GroupMemberRepository) GroupMembersService {
	return &groupMembersService{groupMembersRepository: groupMemberRepository}
}

func (service *groupMembersService) Insert(groupMemberDTO dto.GroupMembersCreateDTO) entity.GroupMembers {
	var groupMembers entity.GroupMembers
	err := smapping.FillStruct(&groupMembers, smapping.MapFields(&groupMemberDTO))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	return service.groupMembersRepository.InsertGroupMember(groupMembers)
}

func (service *groupMembersService) Delete(members entity.GroupMembers) {
	service.groupMembersRepository.DeleteGroupMember(members)
}
