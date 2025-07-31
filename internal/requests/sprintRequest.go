package requests

type CreateSprintRequest struct {
	Name      string `validate:"required"`
	StartDate string
	EndDate   string `validate:"is-valid-date-range"`
	Status    string
}
