package models

import "gorm.io/gorm"

type Organization struct {
	gorm.Model `json:"-"`
	Name       string `json:"organizationName" gorm:"not null"`
	Email      string `json:"email" gorm:"unique;not null"`
	Password   string `json:"-" gorm:"not null"`
}
