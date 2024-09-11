package models

type Invite struct {
	Base
	Firstname      string       `json:"firstName" gorm:"not null"`
	Lastname       string       `json:"lastName" gorm:"not null"`
	Email          string       `json:"email" gorm:"unique;not null"`
	OrganizationID string       `json:"-"`
	Organization   Organization `gorm:"constraint:OnDelete:SET NULL;"`
	Role           string       `json:"role" gorm:"default:user"`
	Status         string       `json:"status"`
	SentByID       string       `json:"-"`
	SentBy         User         `gorm:"constraint:OnDelete:SET NULL;"`
}
