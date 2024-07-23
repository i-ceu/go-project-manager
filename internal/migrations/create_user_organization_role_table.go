package migrations

import (
	"github.com/ubaniIsaac/go-project-manager/internal/config"
	"github.com/ubaniIsaac/go-project-manager/internal/models"
)

func CreateUserOrganizationRoleTable() {
	type UserOrganizationRole struct {
		UserID         int
		User           models.User `gorm:"constraint:OnDelete:SET NULL; foreignKey:UserID;references:ID"`
		RoleID         int
		Role           models.Role `gorm:"constraint:OnDelete:SET NULL; foreignKey:RoleID;references:ID"`
		OrganizationID int
		Organization   models.Organization `gorm:"constraint:OnDelete:SET NULL; foreignKey:OrganizationID;references:ID"`
	}

	if !config.DB.Migrator().HasTable(&UserOrganizationRole{}) {
		config.DB.Migrator().CreateTable(&UserOrganizationRole{})
	}
}
