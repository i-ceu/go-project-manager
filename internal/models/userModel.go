package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	Firstname  string `json:"firstName" gorm:"not null"`
	Lastname   string `json:"lastName" gorm:"not null"`
	Email      string `json:"email" gorm:"unique;not null"`
	Password   string `json:"-" gorm:"not null"`
	Role       string `json:"role" gorm:"default:user"`
}
