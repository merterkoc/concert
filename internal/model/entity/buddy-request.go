package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BuddyRequest struct {
	ID         string        `gorm:"type:char(36);primaryKey"`
	SenderID   string        `gorm:"type:char(36);not null" json:"sender_id"`
	ReceiverID string        `gorm:"type:char(36);not null" json:"receiver_id"`
	EventID    string        `gorm:"type:char(36);not null" json:"event_id"`
	Status     RequestStatus `gorm:"type:enum('pending', 'accepted', 'rejected', 'blocked');not null;default:'pending'" json:"status"`
	CreatedAt  time.Time     `gorm:"default:current_timestamp" json:"created_at"`

	Sender   User `gorm:"foreignKey:SenderID;references:ID;constraint:OnDelete:CASCADE" json:"sender"`
	Receiver User `gorm:"foreignKey:ReceiverID;references:ID;constraint:OnDelete:CASCADE" json:"receiver"`
}

// BeforeCreate hook to generate UUID and set default role
func (u *BuddyRequest) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}

type RequestStatus string

const (
	Pending  RequestStatus = "pending"
	Accepted RequestStatus = "accepted"
	Rejected RequestStatus = "rejected"
	Blocked  RequestStatus = "blocked"
)
