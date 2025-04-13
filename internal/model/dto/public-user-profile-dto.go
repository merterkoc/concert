package dto

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/model/entity"
	"github.com/google/uuid"
	"time"
)

type PublicUserProfileDTO struct {
	ID        uuid.UUID             `json:"id"`
	UserName  string                `json:"username"`
	CreatedAt time.Time             `json:"created_at"`
	UserImage string                `json:"user_image"`
	Interests []entity.InterestType `json:"interests"`
	Events    []entity.UserEvents   `json:"events"`
}
