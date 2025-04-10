package dto

type BuddyRequestDTO struct {
	ID         string `json:"id"`
	SenderID   string `json:"sender_id"`
	ReceiverID string `json:"receiver_id"`
	EventID    string `json:"event_id"`
	Status     string `json:"status"`
}
