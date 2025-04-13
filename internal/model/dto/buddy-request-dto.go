package dto

type BuddyRequestDTO struct {
	ID       string         `json:"id"`
	Sender   UserDto        `json:"sender"`
	Receiver UserDto        `json:"receiver"`
	Event    EventDetailDTO `json:"event"`
	Status   string         `json:"status"`
}
