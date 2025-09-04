package migrations

import (
	"github.com/i-ceu/go-project-manager/internal/config"
)

func CreateDBIndexes() {
	config.DB.Exec("CREATE INDEX IF NOT EXISTS idx_member_roles_user_id ON member_roles(user_id);")
	config.DB.Exec("CREATE INDEX IF NOT EXISTS idx_member_roles_team_id ON member_roles(team_id);")
	config.DB.Exec("CREATE INDEX IF NOT EXISTS idx_member_roles_role_id ON member_roles(role_id);")

	config.DB.Exec("CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);")

	config.DB.Exec("CREATE INDEX IF NOT EXISTS idx_project_team_id ON projects(team_id);")
	config.DB.Exec("CREATE INDEX IF NOT EXISTS idx_tasks_project_id ON tasks(project_id);")
	config.DB.Exec("CREATE INDEX IF NOT EXISTS idx_tasks_assigned_to ON tasks(assigned_to_id);")
}
