package migrations

import (
	"github.com/i-ceu/go-project-manager/internal/config"
	"github.com/i-ceu/go-project-manager/internal/models"
)

func DropOrganizationFromUsersTable() {
	if config.DB.Migrator().HasConstraint(&models.User{}, "fk_users_organization") {
		config.DB.Migrator().DropConstraint(&models.User{}, "fk_users_organization")
	}
	if config.DB.Migrator().HasColumn(&models.User{}, "organization_id") {
		config.DB.Migrator().DropColumn(&models.User{}, "organization_id")
	}

}
