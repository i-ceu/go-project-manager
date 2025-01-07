package requests

type CreateProjectRequest struct {
	Title          string `validate:"required"`
	Tag            string `validate:"required"`
	Description    string `validate:"required"`
	OrganizationID string `validate:"required"`
	Status         string
	DeliveryDate   string `validate:"required"`
}
