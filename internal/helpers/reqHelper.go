package helpers

import (
	"github.com/go-playground/validator/v10"

	"github.com/ubaniIsaac/go-project-manager/internal/config"
	"github.com/ubaniIsaac/go-project-manager/internal/models"
)

func ValidateReq(req interface{}) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(req)
	return err
}

func emailExist(fl validator.FieldLevel) bool {
	// validate := validator.New()
	email := fl.Field().String()
	existingUser := config.DB.Find(models.User{}, "email = ?", email)
	return existingUser.RowsAffected > 0
}
