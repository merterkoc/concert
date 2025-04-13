package dto

type EventDTO struct {
	ID                 string               `json:"id"`
	Name               string               `json:"name"`
	Start              string               `json:"start"`
	End                string               `json:"end"`
	Location           string               `json:"location"`
	Distance           string               `json:"distance"`
	Images             []string             `json:"images"`
	VenueName          string               `json:"venue_name"`
	City               string               `json:"city"`
	Country            string               `json:"country"`
	Locale             string               `json:"locale"`
	URL                string               `json:"ticket_url"`
	IsJoined           bool                 `json:"is_joined"`
	ParticipantAvatars []ParticipantsAvatar `json:"participant_avatars"`
}
