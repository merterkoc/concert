package dto

import "time"

type UserDto struct {
	Email     string    `json:"email"`
	UserName  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UserImage string    `json:"user_image"`
}
