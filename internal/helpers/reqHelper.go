package helpers

import (
	"github.com/go-playground/validator/v10"
	"github.com/ubaniIsaac/go-project-manager/internal/config"
	"github.com/ubaniIsaac/go-project-manager/internal/enums"
	"github.com/ubaniIsaac/go-project-manager/internal/models"
)

func ValidateReq(req interface{}) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	//rgeister valid organization size validation
	_ = validate.RegisterValidation("is-valid-organization-size", isValidOrganizationSize)

	//register is valid industry types validation
	_ = validate.RegisterValidation("is-valid-industry", isValidIndustry)

	err := validate.Struct(req)
	return err
}

func isValidIndustry(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	for _, indt := range enums.GetAllIndustries() {
		if value == indt {
			return true
		}
	}
	return false
}

func isValidOrganizationSize(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	for _, size := range enums.GetAllSizes() {
		if value == size {
			return true
		}
	}
	return false
}

func emailExist(fl validator.FieldLevel) bool {
	// validate := validator.New()
	email := fl.Field().String()
	existingUser := config.DB.Find(models.User{}, "email = ?", email)
	return existingUser.RowsAffected > 0
}
