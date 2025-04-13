package dto

import "gilab.com/pragmaticreviews/golang-gin-poc/external/event/entity"

type EventDetailDTO struct {
	ID                 string               `json:"id"`
	Name               string               `json:"name"`
	Start              string               `json:"start"`
	End                string               `json:"end"`
	City               string               `json:"city"`
	Country            string               `json:"country"`
	Locale             string               `json:"locale"`
	Distance           string               `json:"distance"`
	Images             []entity.Image       `json:"images"`
	VenueName          string               `json:"venue_name"`
	URL                string               `json:"ticket_url"`
	IsJoined           bool                 `json:"is_joined"`
	ParticipantAvatars []ParticipantsAvatar `json:"participant_avatars"`
}
