package dto

import "github.com/google/uuid"

type ParticipantsAvatar struct {
	ID        uuid.UUID `json:"user_id"`
	UserImage *string   `json:"user_image"`
}
