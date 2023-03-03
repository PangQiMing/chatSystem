package entity

import (
	"database/sql"
	"time"
)

type Message struct {
	ID        uint64 `gorm:"primaryKey" json:"id"`
	FromId    uint64 `gorm:"default:null" json:"from_id"`
	ToId      uint64 `gorm:"default:null" json:"to_id"`
	Content   string `gorm:"type:text" json:"context"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}
