package dto

type GroupMembersCreateDTO struct {
	ID      uint64 `json:"id" form:"id"`
	GroupID uint64 `json:"group_id" form:"group_id"`
	UserID  uint64 `json:"user_id" form:"user_id"`
}
