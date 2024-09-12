package models

type User struct {
	Base
	Firstname string `json:"firstName" gorm:"not null"`
	Lastname  string `json:"lastName" gorm:"not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"-" gorm:"not null"`
	RoleID    string `json:"-"`
	Role      Role   `gorm:"constraint:OnDelete:SET NULL; foreignKey:RoleID;references:ID"`
	Status    string `json:"status" gorm:"default:not-verified"`
}
