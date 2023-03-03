package repository

import (
	"chat/entity"
	"gorm.io/gorm"
)

type GroupMemberRepository interface {
	InsertGroupMember(groupMember entity.GroupMembers) entity.GroupMembers
	DeleteGroupMember(groupMember entity.GroupMembers)
}

type groupMembersRepository struct {
	connection *gorm.DB
}

func NewGroupMembersRepository(db *gorm.DB) GroupMemberRepository {
	return &groupMembersRepository{connection: db}
}

func (db *groupMembersRepository) InsertGroupMember(groupMember entity.GroupMembers) entity.GroupMembers {
	db.connection.Save(&groupMember)
	db.connection.Preload("User").Preload("Group").Find(&groupMember)
	return groupMember
}

func (db *groupMembersRepository) DeleteGroupMember(groupMember entity.GroupMembers) {
	db.connection.Where("group_id = ? and user_id = ?", groupMember.GroupID, groupMember.UserID).Delete(&groupMember)
}
