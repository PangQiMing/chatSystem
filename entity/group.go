package entity

import (
	"database/sql"
	"time"
)

type Group struct {
	ID            uint64 `gorm:"primaryKey" json:"id"`
	GroupName     string `gorm:"type:varchar(255)" json:"group_name"`
	GroupLeaderID uint64 `gorm:"not null" json:"group_leader_id"`
	Notice        string `gorm:"type:varchar(1024)" json:"notice"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     sql.NullTime `gorm:"index"`
}
