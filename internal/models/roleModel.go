package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model `json:"-"`
	Name       string `json:"name"`
}
