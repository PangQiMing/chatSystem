package dto

type UserUpdateDTO struct {
	Avatar   string `json:"avatar" form:"avatar"`
	NickName string `json:"nick_name" form:"nick_name"`
	Age      string `json:"age,omitempty" form:"age,omitempty"`
	Sex      string `json:"sex,omitempty" form:"sex,omitempty"`
}

type UserChangePass struct {
	Email       string `json:"email" form:"email" binding:"email"`
	Password    string `json:"password" form:"password,omitempty,"`
	NewPassword string `json:"new_password" form:"new_password,omitempty,"`
}
