package requests

type CreateTeamRequest struct {
	Name        string `validate:"required"`
	Description string
	Size        string `validate:"required,is-valid-team-size"`
	Industry    string `validate:"required,is-valid-industry"`
}

type SendInviteRequest struct {
	Firstname string `validate:"required"`
	Lastname  string `validate:"required"`
	Email     string `validate:"required,email"`
	RoleID    string `validate:"required"`
}
