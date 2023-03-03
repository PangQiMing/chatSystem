package entity

import (
	"database/sql"
	"time"
)

type User struct {
	UserId     uint64 `gorm:"primaryKey" json:"user_id"`                  //ID
	Avatar     string `gorm:"type:varchar(255)" json:"avatar"`            //头像
	NickName   string `gorm:"type:varchar(255)" json:"nick_name"`         //昵称
	Email      string `gorm:"uniqueIndex;type:varchar(255)" json:"email"` //邮箱
	Password   string `gorm:"->;<-;not null" json:"-"`                    //密码
	Age        string `gorm:"default:0" json:"age"`                       //年龄
	Sex        string `gorm:"default:0" json:"sex"`                       //性别 0表示女，1表示男
	Token      string `gorm:"-" json:"token,omitempty"`                   //token
	UserStatus uint64 `gorm:"default:0" json:"user_status"`               //用户在线状态 0表示不在线，1表示在线
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  sql.NullTime `gorm:"index"`
}
