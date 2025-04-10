package dto

import "mime/multipart"

type CreateUserRequest struct {
	Email    string                `json:"email" binding:"required"`
	Password string                `json:"password" binding:"required"`
	Username string                `json:"username" binding:"required"`
	Image    *multipart.FileHeader `form:"image"`
}
