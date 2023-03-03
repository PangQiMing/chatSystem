package repository

import (
	"chat/entity"
	"gorm.io/gorm"
)

type GroupRepository interface {
	InsertGroup(group entity.Group) entity.Group
	DeleteGroup(groupID, groupLeaderID uint64) bool
	UpdateGroup(group entity.Group) (entity.Group, error)
	MyGroup(groupLeaderID uint64) []entity.Group
}

type groupRepository struct {
	connection *gorm.DB
}

func NewGroupRepository(db *gorm.DB) GroupRepository {
	return &groupRepository{connection: db}
}

func (db *groupRepository) InsertGroup(group entity.Group) entity.Group {
	db.connection.Save(&group)
	db.connection.Find(&group)
	return group

}

func (db *groupRepository) DeleteGroup(groupID, groupLeaderID uint64) bool {
	var group entity.Group
	if err := db.connection.Where("id = ? and group_leader_id = ?", groupID, groupLeaderID).Delete(&group).Error; err != nil {
		return false
	}
	return true
}

func (db *groupRepository) UpdateGroup(group entity.Group) (entity.Group, error) {
	if err := db.connection.Select("group_name", "notice").Where("group_leader_id = ?", group.GroupLeaderID).Save(&group).Error; err != nil {
		return entity.Group{}, err
	}
	db.connection.Find(&group)

	return group, nil
}

func (db *groupRepository) MyGroup(groupLeaderID uint64) []entity.Group {
	var group []entity.Group
	db.connection.Where("group_leader_id = ?", groupLeaderID).Find(&group)
	return group
}
