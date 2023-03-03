package entity

import (
	"database/sql"
	"time"
)

type Friend struct {
	ID             uint64 `gorm:"primaryKey" json:"id"`
	FriendNickName string `gorm:"not null" json:"friend_nick_name"`
	FriendEmail    string `gorm:"not null" json:"friend_email"`
	UserID         uint64 `gorm:"not null" json:"-"`
	User           User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      sql.NullTime `gorm:"index"`
}
