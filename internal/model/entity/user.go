package entity

import (
	"time"

	eventEntity "gilab.com/pragmaticreviews/golang-gin-poc/external/event/entity"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/model/entity/enum"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID           `gorm:"type:char(36);primaryKey" json:"id"`
	FirebaseUID string              `gorm:"type:varchar(255);not null" json:"-"`
	Username    string              `gorm:"type:varchar(255);unique;not null" json:"username"`
	Email       string              `gorm:"type:varchar(255);unique;not null" json:"email"`
	CreatedAt   time.Time           `gorm:"autoCreateTime" json:"created_at"`
	Role        enum.Role           `gorm:"type:enum('user', 'admin');not null;default:'user'" json:"role"`
	UserImage   *string             `gorm:"type:varchar(255);" json:"user_image"`
	Interests   []InterestType      `gorm:"many2many:user_interest_types;" json:"interests"`
	Events      []eventEntity.Event `gorm:"many2many:user_events;constraint:OnDelete:CASCADE" json:"events"`

	BuddyshipsUser1 []Buddyship `gorm:"foreignKey:User1ID;references:ID;constraint:OnDelete:CASCADE" json:"buddyships_user1"`
	BuddyshipsUser2 []Buddyship `gorm:"foreignKey:User2ID;references:ID;constraint:OnDelete:CASCADE" json:"buddyships_user2"`
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

func (u *User) LoadBuddyships(db *gorm.DB) *gorm.DB {
	return db.Preload("BuddyshipsUser1").Preload("BuddyshipsUser2").Model(u)
}
