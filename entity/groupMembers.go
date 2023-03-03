package entity

import (
	"database/sql"
	"time"
)

type GroupMembers struct {
	ID        uint64 `gorm:"primaryKey" json:"id"`
	GroupID   uint64 `gorm:"not null" json:"-"`
	UserID    uint64 `gorm:"not null" json:"-"`
	NikeName  string `gorm:"varchar(250)" json:"nike_name"`
	Group     Group  `gorm:"foreignKey:GroupID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"group"`
	User      User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}
