package dto

//type RegisterDTO struct {
//	Name     string `json:"name" form:"name" binding:"required"`
//	Email    string `json:"email" form:"email" binding:"required,email"`
//	Password string `json:"password" form:"password" binding:"required"`
//}

type RegisterDTO struct {
	NickName string `json:"nick_name" form:"nick_name" binding:"required"`
	Avatar   string `json:"avatar" form:"avatar"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password,omitempty" form:"password,omitempty" binding:"required" validate:"min:6"`
	Age      string `json:"age,omitempty" form:"age,omitempty"`
	Sex      string `json:"sex,omitempty" form:"sex,omitempty"`
}
