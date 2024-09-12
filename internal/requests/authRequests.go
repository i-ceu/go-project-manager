package requests

type RegisterUserRequest struct {
	Firstname       string `validate:"required"`
	Lastname        string `validate:"required"`
	Email           string `validate:"required,email"`
	Password        string `validate:"required"`
	ConfirmPassword string `validate:"required,eqfield=Password"`
	RoleID          string
}

type SignInRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type AcceptInviteRequest struct {
	Password        string `validate:"required"`
	ConfirmPassword string `validate:"required,eqfield=Password"`
}
