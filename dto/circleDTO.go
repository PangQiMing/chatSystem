package dto

type CircleCreateDTO struct {
	UserID      uint64 `json:"user_id" form:"user_id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
}
