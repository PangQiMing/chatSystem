package repository

import (
	"chat/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type UserRepository interface {
	AddUser(user entity.User) entity.User
	UpdateUser(user entity.User) entity.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) entity.User
	ProfileUser(UserID string) entity.User
	ChangePass(changePass entity.User) entity.User
}

type userConn struct {
	connDB *gorm.DB
}

func NewUserConnection(db *gorm.DB) UserRepository {
	return &userConn{connDB: db}
}

// AddUser 注册用户
func (db *userConn) AddUser(user entity.User) entity.User {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connDB.Save(&user)
	return user
}

func (db *userConn) UpdateUser(user entity.User) entity.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser entity.User
		db.connDB.Find(&tempUser, user.UserId)
		user.Password = tempUser.Password
	}
	db.connDB.Save(&user)
	return user
}

// VerifyCredential 验证凭证
func (db *userConn) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	res := db.connDB.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

// IsDuplicateEmail 验证用户邮箱是否重复
func (db *userConn) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return db.connDB.Where("email = ?", email).Take(&user)
}

// FindByEmail 根据email找到用户
func (db *userConn) FindByEmail(email string) entity.User {
	var user entity.User
	db.connDB.Where("email = ?", email).Take(&user)
	return user
}

// ProfileUser 用户的详细信息
func (db *userConn) ProfileUser(userID string) entity.User {
	var user entity.User
	//db.connDB.Preload("Circle").Preload("Circle.User").Find(&user, userID)
	db.connDB.Find(&user, userID)
	return user
}

func (db *userConn) ChangePass(changePass entity.User) entity.User {
	changePass.Password = hashAndSalt([]byte(changePass.Password))
	db.connDB.Save(&changePass)
	return changePass
}

// 加密密码
func hashAndSalt(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Println("加密密码失败...")
	}
	return string(hash)
}
