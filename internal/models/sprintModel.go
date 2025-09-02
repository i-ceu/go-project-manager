package models

import "time"

type Sprint struct {
	Base
	Name        string     `json:"name"`
	Status      string     `json:"status"`
	StartDate   *time.Time `json:"startDate" gorm:"default:NULL"`
	EndDate     *time.Time `json:"endDate" gorm:"default:NULL"`
	ProjectID   string     `json:"-"`
	Project     *Project   `json:"project,omitempty" gorm:"constraint:OnDelete:SET NULL;"`
	CreatedByID string     `json:"-"`
	CreatedBy   *User      `json:"createdBy,omitempty" gorm:"constraint:OnDelete:SET NULL;"`
}
