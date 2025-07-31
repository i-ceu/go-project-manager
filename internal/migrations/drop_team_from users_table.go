package migrations

import (
	"github.com/i-ceu/go-project-manager/internal/config"
	"github.com/i-ceu/go-project-manager/internal/models"
)

func DropTeamFromUsersTable() {
	if config.DB.Migrator().HasConstraint(&models.User{}, "fk_users_team") {
		config.DB.Migrator().DropConstraint(&models.User{}, "fk_users_team")
	}
	if config.DB.Migrator().HasColumn(&models.User{}, "team_id") {
		config.DB.Migrator().DropColumn(&models.User{}, "team_id")
	}

}
func DropEmailFromTeamTable() {
	if config.DB.Migrator().HasColumn(&models.Team{}, "email") {
		config.DB.Migrator().DropColumn(&models.Team{}, "email")
	}

}
