package migrations

import (
	"github.com/ubaniIsaac/go-project-manager/internal/config"
	"github.com/ubaniIsaac/go-project-manager/internal/models"
)

func CreateStaffRoleTable() {
	type StaffRole struct {
		models.Base
		UserID         string
		User           models.User `gorm:"constraint:OnDelete:SET NULL; foreignKey:UserID;references:ID"`
		RoleID         string
		Role           models.Role `gorm:"constraint:OnDelete:SET NULL; foreignKey:RoleID;references:ID"`
		OrganizationID string
		Organization   models.Organization `gorm:"constraint:OnDelete:SET NULL; foreignKey:OrganizationID;references:ID"`
	}

	if !config.DB.Migrator().HasTable(&StaffRole{}) {
		config.DB.Migrator().CreateTable(&StaffRole{})
	}
}
