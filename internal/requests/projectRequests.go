package requests

type CreateProjectRequest struct {
	Title        string `validate:"required"`
	Description  string `validate:"required"`
	Status       string
	StartDate    string   `validate:"required"`
	DeliveryDate string   `validate:"required,is-valid-date-range"`
	WorkFlow     []string `validate:"required"`
}
