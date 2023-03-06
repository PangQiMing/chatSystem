package entity

import (
	"database/sql"
	"time"
)

type Circle struct {
	ID         uint64   `gorm:"primaryKey" json:"id"`
	Content    string   `gorm:"type:text" json:"content"`
	PictureUrl []string `gorm:"type:text" json:"picture_url"`
	UserID     uint64   `gorm:"not null" json:"-"`
	User       User     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"circle"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  sql.NullTime `gorm:"index"`
}
