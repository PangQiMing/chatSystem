package service

import (
	"chat/dto"
	"chat/entity"
	"chat/repository"
	"github.com/mashingan/smapping"
	"log"
)

type FriendService interface {
	Insert(addFriendDTO dto.AddFriendDTO) (entity.Friend, int64)
	AllFriend(userID uint64) []entity.Friend
	Delete(friend entity.Friend, userID uint64)
	ShowAddFriendList(email string) []entity.Friend
	isFriendAlready(friend entity.Friend) bool
	friendNotInDB(friend entity.Friend) bool
}

type friendService struct {
	friendRepository repository.FriendRepository
}

func NewFriendService(friendRepository repository.FriendRepository) FriendService {
	return &friendService{friendRepository: friendRepository}
}

func (service *friendService) Insert(addFriendDTO dto.AddFriendDTO) (entity.Friend, int64) {
	var friend entity.Friend
	err := smapping.FillStruct(&friend, smapping.MapFields(&addFriendDTO))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	if service.isFriendAlready(friend) {
		return entity.Friend{}, 1
	} else if service.friendNotInDB(friend) {
		return entity.Friend{}, -1
	} else {
		resFriend := service.friendRepository.Insert(friend)
		return resFriend, 0
	}
}

func (service *friendService) AllFriend(userID uint64) []entity.Friend {
	return service.friendRepository.AllFriends(userID)
}

func (service *friendService) Delete(friend entity.Friend, userID uint64) {
	service.friendRepository.Delete(friend, userID)
}

func (service *friendService) ShowAddFriendList(email string) []entity.Friend {
	return service.friendRepository.ShowAddFriendList(email)
}

func (service *friendService) isFriendAlready(friend entity.Friend) bool {
	return service.friendRepository.IsFriendAlready(friend)
}
func (service *friendService) friendNotInDB(friend entity.Friend) bool {
	return service.friendRepository.FriendNotInDB(friend)
}
