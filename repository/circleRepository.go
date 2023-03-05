package repository

import (
	"chat/entity"
	"gorm.io/gorm"
)

type CircleRepository interface {
	Insert(circle entity.Circle) entity.Circle
	Update(circle entity.Circle) entity.Circle
	Delete(circle entity.Circle)
	All(userID uint64) []entity.Circle
	FindByID(circleID uint64) entity.Circle
}

type circleRepository struct {
	connection *gorm.DB
}

func NewCircleRepository(db *gorm.DB) CircleRepository {
	return &circleRepository{connection: db}
}

func (db *circleRepository) Insert(circle entity.Circle) entity.Circle {
	db.connection.Save(&circle)
	db.connection.Preload("User").Find(&circle)
	return circle
}

func (db *circleRepository) Update(circle entity.Circle) entity.Circle {
	db.connection.Save(&circle)
	db.connection.Preload("User").Find(&circle)
	return circle
}

func (db *circleRepository) Delete(circle entity.Circle) {
	db.connection.Delete(&circle)
}

func (db *circleRepository) All(userID uint64) []entity.Circle {
	var circles []entity.Circle
	db.connection.Preload("User").Where("user_id = ?", userID).Find(&circles)
	return circles
}

func (db *circleRepository) FindByID(circleID uint64) entity.Circle {
	var circle entity.Circle
	db.connection.Preload("User").Find(&circle, circleID)
	return circle
}
