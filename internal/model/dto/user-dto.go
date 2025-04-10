package dto

import (
	"time"

	"gilab.com/pragmaticreviews/golang-gin-poc/internal/model/entity"
)

type UserDto struct {
	Email     string                `json:"email"`
	UserName  string                `json:"username"`
	CreatedAt time.Time             `json:"created_at"`
	UserImage string                `json:"user_image"`
	Interests []entity.InterestType `gorm:"many2many:user_interest_types;" json:"interests"`
}
