package entity

import (
	"github.com/google/uuid"
)

type UserInterestType struct {
	UserID         uuid.UUID `gorm:"type:char(36);primaryKey" json:"user_id"`
	InterestTypeID int       `gorm:"primaryKey" json:"interest_type_id"`
	User           User      `gorm:"foreignKey:UserID" json:"user"`
}
