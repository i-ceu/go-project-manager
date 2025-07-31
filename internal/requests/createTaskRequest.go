package requests

type CreateTaskRequest struct {
	Title       string `validate:"required"`
	Tag         string
	Description string `validate:"required"`
	SprintID    string
	StartDate   string
	EndDate     string `validate:"is-valid-date-range"`
	Status      string
	AssignedTo  string
}

type UpdateTaskRequest struct {
	Title       string
	Description string
	StartDate   string
	EndDate     string `validate:"is-valid-date-range"`
	Status      string
}

type AssignTaskRequest struct {
	AssignedTo string
}
