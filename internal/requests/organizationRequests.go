package requests

type CreateOrganizationRequest struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Size     string `validate:"required,is-valid-organization-size"`
	Industry string `validate:"required,is-valid-industry"`
}

type SendInviteRequest struct {
	Firstname string `validate:"required"`
	Lastname  string `validate:"required"`
	Email     string `validate:"required,email"`
	RoleID    string `validate:"required"`
}
