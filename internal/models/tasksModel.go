package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model   `json:"-"`
	Title        string    `json:"title"`
	Tag          string    `json:"tag" gorm:"not null"`
	Description  string    `json:"description"`
	Status       string    `json:"status" gorm:"default: todo"`
	DueDate      time.Time `json:"dueDate"`
	ProjectID    int       `json:"-"`
	Project      Project   `gorm:"constraint:OnDelete:SET NULL;"`
	AssignerID   int       `json:"-"`
	Assigner     User      `gorm:"constraint:OnDelete:SET NULL;"`
	AssignedToID int       `json:"-" gorm:"default: 1"`
	AssignedTo   User      `gorm:"constraint:OnDelete:SET NULL;"`
}
