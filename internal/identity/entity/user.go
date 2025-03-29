package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	FirebaseUID string    `gorm:"type:varchar(255);not null" json:"-"`
	Username    string    `gorm:"type:varchar(255);unique;not null" json:"username"`
	Email       string    `gorm:"type:varchar(255);unique;not null" json:"email"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}

// BeforeCreate hook to generate UUID before inserting a new record
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	return
}
