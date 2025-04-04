package entity

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/entity/enum"
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
	Role        enum.Role `gorm:"type:enum('user', 'admin');not null;default:'user'" json:"role"`
	UserImage   *string   `gorm:"type:varchar(255);" json:"user_image"`
}

// BeforeCreate hook to generate UUID and set default role
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	if u.Role.String() == "" {
		u.Role = enum.User
	}
	return
}
