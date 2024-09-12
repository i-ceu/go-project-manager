package models

type Invite struct {
	Base
	Firstname      string       `json:"firstName" gorm:"not null"`
	Lastname       string       `json:"lastName" gorm:"not null"`
	Email          string       `json:"email" gorm:"unique;not null"`
	OrganizationID string       `json:"-"`
	Organization   Organization `gorm:"constraint:OnDelete:SET NULL;"`
	RoleID         string       `json:"-"`
	Role           Role         `gorm:"constraint:OnDelete:SET NULL; foreignKey:RoleID;references:ID"`
	Status         string       `json:"status" gorm:"default:pending"`
	SentByID       string       `json:"-"`
	SentBy         User         `gorm:"constraint:OnDelete:SET NULL;"`
}
