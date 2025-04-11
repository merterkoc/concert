package entity

import (
	"time"

	"gorm.io/gorm"
)

type Buddyship struct {
	ID        string    `gorm:"type:char(36);primaryKey;default:uuid();not null" json:"id"`
	User1ID   string    `gorm:"type:char(36);not null" json:"user1_id"`
	User2ID   string    `gorm:"type:char(36);not null" json:"user2_id"`
	EventID   string    `gorm:"type:char(36);not null" json:"event_id"`
	MatchedAt time.Time `gorm:"default:current_timestamp;not null" json:"matched_at"`

	User1 User `gorm:"foreignKey:User1ID;references:ID;constraint:OnDelete:CASCADE" json:"user1"`
	User2 User `gorm:"foreignKey:User2ID;references:ID;constraint:OnDelete:CASCADE" json:"user2"`
}

// BeforeCreate hook to generate UUID before inserting into the database
func (uc *Buddyship) BeforeCreate(tx *gorm.DB) (err error) {
	uc.MatchedAt = time.Now()
	return
}
