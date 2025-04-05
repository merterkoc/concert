package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserEvents struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey;default:uuid_generate_v4()"`
	UserID    uuid.UUID `gorm:"type:char(36);not null" json:"user_id"` // ForeignKey for User
	EventID   string    `gorm:"type:varchar(255);not null"`
	Status    string    `gorm:"type:enum('interested', 'going');not null"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	User      User      `gorm:"foreignKey:UserID;references:ID" json:"user"`
}

// BeforeCreate hook to generate UUID before inserting into the database
func (uc *UserEvents) BeforeCreate(tx *gorm.DB) (err error) {
	if uc.ID == uuid.Nil {
		uc.ID = uuid.New()
	}
	return
}
