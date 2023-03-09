package entity

import (
	"database/sql"
	"time"
)

type Friend struct {
	ID             uint64 `gorm:"primaryKey" json:"id"`
	FriendAvatar   string `gorm:"type:varchar(255)" json:"friend_avatar"`
	FriendNickName string `gorm:"not null" json:"friend_nick_name"`
	FriendEmail    string `gorm:"not null" json:"friend_email"`
	FriendStatus   uint64 `gorm:"not null" json:"friend_status"` //0表示新好友，1表示正常，2表示黑名单
	UserID         uint64 `gorm:"not null" json:"-"`
	User           User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      sql.NullTime `gorm:"index"`
}
