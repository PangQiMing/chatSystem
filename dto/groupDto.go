package dto

type GroupCreateDTO struct {
	GroupName     string `json:"group_name" form:"group_name" binding:"required"`
	GroupLeaderID uint64 `json:"group_leader_id" form:"group_leader_id"`
	Notice        string `json:"notice" form:"group_name" binding:"required"`
}

type GroupUpdateDTO struct {
	ID            uint64 `json:"id" form:"id"`
	GroupName     string `json:"group_name" form:"group_name"`
	GroupLeaderID uint64 `json:"group_leader_id" form:"group_leader_id"`
	Notice        string `json:"notice" form:"notice"`
}
