package dto

type FriendDTO struct {
	Email string `json:"email" form:"email"`
}

type AddFriendDTO struct {
	UserID         uint64 `json:"user_id" form:"user_id"`
	FriendAvatar   string `json:"friend_avatar" form:"friend_avatar"`
	FriendNickName string `json:"friend_nick_name" form:"friend_nick_name"`
	FriendEmail    string `json:"friend_email" form:"friend_email"`
	FriendStatus   uint64 `json:"friend_status" form:"friend_status"`
}

type UpdateFriendStatus struct {
	UserID       uint64 `json:"user_id" form:"user_id"`
	FriendEmail  string `json:"friend_email" form:"friend_email"`
	FriendStatus uint64 `json:"friend_status" form:"friend_status"`
}
