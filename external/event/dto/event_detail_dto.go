package dto

import "gilab.com/pragmaticreviews/golang-gin-poc/external/event/entity"

type EventDetailDTO struct {
	ID                 string         `json:"id"`
	Name               string         `json:"name"`
	Start              string         `json:"start"`
	End                string         `json:"end"`
	Location           string         `json:"location"`
	Distance           string         `json:"distance"`
	Images             []entity.Image `json:"images"`
	URL                string         `json:"ticket_url"`
	IsJoined           bool           `json:"is_joined"`
	ParticipantAvatars []*string      `json:"participant_avatars"`
}
