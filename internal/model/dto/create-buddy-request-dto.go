package dto

type CreateBuddyRequestDTO struct {
	ReceiverID string `json:"receiver_id"`
	EventID    string `json:"event_id"`
}
