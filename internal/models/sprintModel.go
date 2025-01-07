package models

import "time"

type Sprint struct {
	Base
	Name        string    `json:"name"`
	Status      string    `json:"status"`
	StartDate   time.Time `json:"startDate" gorm:"type:datetime;default:NULL"`
	EndDate     time.Time `json:"endDate" gorm:"type:datetime;default:NULL"`
	ProjectID   string    `json:"-"`
	Project     Project   `gorm:"constraint:OnDelete:SET NULL;"`
	CreatedByID string    `json:"-"`
	CreatedBy   User      `gorm:"constraint:OnDelete:SET NULL;"`
}
