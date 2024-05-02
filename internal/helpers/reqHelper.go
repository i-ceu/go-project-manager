package helpers

import (
	"github.com/go-playground/validator/v10"
)

func ValidateReq(req interface{}) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(req)
	return err
}
