package dto

type UserUpdateDTO struct {
	UserId   uint64 `json:"user_id" form:"user_id"`
	Avatar   string `json:"avatar" form:"avatar"`
	NickName string `json:"nick_name" form:"nick_name"`
	Email    string `json:"email" form:"email" binding:"email"`
	Password string `json:"password" form:"password,omitempty"`
	Age      string `json:"age,omitempty" form:"age,omitempty"`
	Sex      string `json:"sex,omitempty" form:"sex,omitempty"`
}

//type UserCreateDTO struct {
//	Avatar   string `json:"avatar" form:"avatar"`
//	NickName string `json:"nick_name" form:"nick_name" binding:"required"`
//	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
//	Password string `json:"password,omitempty" form:"password,omitempty" binding:"required" validate:"min:6"`
//	Age      uint64 `json:"age,omitempty" form:"age,omitempty"`
//	Sex      uint64 `json:"sex,omitempty" form:"sex,omitempty"`
//}
