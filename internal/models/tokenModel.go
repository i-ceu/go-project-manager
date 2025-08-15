package models

type Token struct {
	Base
	UserId   string `json:"userId"`
	User     User   `json:"user" gorm:"constraint:OnDelete:SET NULL;"`
	Token    int    `json:"token" gorm:"unique;not null"`
	Verified bool   `json:"verified"`
}
