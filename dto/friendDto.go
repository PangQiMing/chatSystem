package dto

type FriendDTO struct {
	UserID      uint64 `json:"user_id" form:"user_id"`
	FriendEmail string `json:"friend_email" form:"friend_email" binding:"required,email"`
}
