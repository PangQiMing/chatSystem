package dto

type CircleCreateDTO struct {
	UserID     uint64   `json:"user_id" form:"user_id"`
	Content    string   `json:"content" form:"content"`
	PictureUrl []string `json:"picture_url" form:"picture_url"`
}
