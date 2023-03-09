package repository

import (
	"chat/entity"
	"gorm.io/gorm"
)

type FriendRepository interface {
	Insert(friend entity.Friend) entity.Friend
	AllFriends(userID uint64) []entity.Friend
	Delete(friend entity.Friend, userID uint64)
	IsFriendAlready(friend entity.Friend) bool
	FriendNotInDB(friend entity.Friend) bool
	ShowAddFriendList(friendEmail string) []entity.Friend
}

type friendConnection struct {
	connection *gorm.DB
}

func NewFriendRepository(db *gorm.DB) FriendRepository {
	return &friendConnection{connection: db}
}

func (db *friendConnection) Insert(friend entity.Friend) entity.Friend {
	db.connection.Save(&friend)
	db.connection.Preload("User").Find(&friend)
	return friend
}

//func (db *friendConnection) FindFriendByEmail(friendEmail string) entity.Friend {
//	db.connection.Preload("User").Where("")
//}

func (db *friendConnection) AllFriends(userID uint64) []entity.Friend {
	var friend []entity.Friend
	db.connection.Preload("User").Where("user_id = ?", userID).Find(&friend)
	return friend
}

func (db *friendConnection) Delete(friend entity.Friend, userID uint64) {
	db.connection.Where("user_id = ?", userID).Delete(&friend)
}

func (db *friendConnection) ShowAddFriendList(email string) []entity.Friend {
	var friend []entity.Friend
	db.connection.Preload("User").Where("friend_email = ? && friend_status = ?", email, 0).Find(&friend)
	return friend
}

func (db *friendConnection) IsFriendAlready(friend entity.Friend) bool {
	var count int64
	db.connection.Preload("User").Where("friend_email = ?", friend.FriendEmail).Count(&count)
	if count > 0 {
		return true
	}
	return false
}

func (db *friendConnection) FriendNotInDB(friend entity.Friend) bool {
	var count int64
	db.connection.Model(&entity.User{}).Where("email = ?", friend.FriendEmail).Count(&count)
	if count > 0 {
		return false
	}
	return true
}
