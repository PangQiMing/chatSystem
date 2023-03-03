package repository

import (
	"chat/entity"
	"gorm.io/gorm"
)

type MomentRepository interface {
	InsertBook(moment entity.Circle) entity.Circle
	UpdateBook(moment entity.Circle) entity.Circle
	DeleteBook(moment entity.Circle)
	AllBook(userID uint64) []entity.Circle
	FindBookByID(momentID uint64) entity.Circle
}

type momentRepository struct {
	connection *gorm.DB
}

func NewMomentRepository(db *gorm.DB) MomentRepository {
	return &momentRepository{connection: db}
}

func (db *momentRepository) InsertBook(moment entity.Circle) entity.Circle {
	db.connection.Save(&moment)
	db.connection.Preload("User").Find(&moment)
	return moment
}

func (db *momentRepository) UpdateBook(moment entity.Circle) entity.Circle {
	db.connection.Save(&moment)
	db.connection.Preload("User").Find(&moment)
	return moment
}

func (db *momentRepository) DeleteBook(moment entity.Circle) {
	db.connection.Delete(&moment)
}

func (db *momentRepository) AllBook(userID uint64) []entity.Circle {
	var moments []entity.Circle
	db.connection.Preload("User").Where("user_id = ?", userID).Find(&moments)
	return moments
}

func (db *momentRepository) FindBookByID(momentID uint64) entity.Circle {
	var moment entity.Circle
	db.connection.Preload("User").Find(&moment, momentID)
	return moment
}
