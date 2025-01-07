package requests

type CreateSprintRequest struct {
	Name      string `validate:"required"`
	Project   string `validate:"required"`
	StartDate string
	EndDate   string
	Status    string
	CreatedBy string
}
