package models

import (
	"time"
)

type Task struct {
	Base
	Title        string    `json:"title"`
	Tag          string    `json:"tag" gorm:"not null"`
	Description  string    `json:"description"`
	Status       string    `json:"status" gorm:"default: todo"`
	DueDate      time.Time `json:"dueDate"`
	ProjectID    string    `json:"-"`
	Project      Project   `gorm:"constraint:OnDelete:SET NULL;"`
	AssignerID   string    `json:"-"`
	Assigner     User      `gorm:"constraint:OnDelete:SET NULL;"`
	AssignedToID string    `json:"-" gorm:"default: 1"`
	AssignedTo   User      `gorm:"constraint:OnDelete:SET NULL;"`
}
