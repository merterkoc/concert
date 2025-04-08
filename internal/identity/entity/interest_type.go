package entity

type InterestType struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"unique;not null" json:"name"`

	Users []User `gorm:"many2many:user_interest_types;" json:"-"`
}
