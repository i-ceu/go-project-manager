package models

import (
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model `json:"-"`
	UserId     int  `json:"userId"`
	User       User `gorm:"constraint:OnDelete:SET NULL;"`
	Token      int  `json:"token" gorm:unique;not null`
	Verified   bool `json:"verified"`
}
