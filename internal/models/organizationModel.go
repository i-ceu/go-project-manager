package models

type Organization struct {
	Base
	Name     string `json:"organizationName" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Size     string `json:"size"`
	Industry string `json:"industry"`
}
