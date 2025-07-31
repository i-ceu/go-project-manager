package models

type User struct {
	Base
	Firstname   string        `json:"firstName" gorm:"not null"`
	Lastname    string        `json:"lastName" gorm:"not null"`
	Email       string        `json:"email" gorm:"unique;not null"`
	Password    string        `json:"-" gorm:"not null"`
	Status      string        `json:"status" gorm:"default:not-verified"`
	MemberRoles *[]MemberRole `json:"memberRoles,omitempty"`
}
