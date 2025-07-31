package helpers

import (
	"errors"
	"reflect"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	tr_en "github.com/go-playground/validator/v10/translations/en"
	"github.com/i-ceu/go-project-manager/internal/enums"
)

func registerValidations(validate *validator.Validate, trans ut.Translator) {
	validations := map[string]struct {
		fn      validator.Func
		message string
	}{
		"is-valid-team-size":  {isValidTeamSize, "Invalid Team Size"},
		"is-valid-industry":   {isValidIndustry, "Industry must be a valid option"},
		"is-valid-date-range": {isValidDateRange, "End date must be after or equal to start date"},
	}

	for tag, validation := range validations {
		// Register the validation function
		_ = validate.RegisterValidation(tag, validation.fn)

		// Register the translation
		_ = validate.RegisterTranslation(tag, trans,
			func(ut ut.Translator) error {
				return ut.Add(tag, validation.message, true)
			},
			func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T(tag)
				return t
			},
		)
	}
}

func ValidateReq(req interface{}) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	_ = tr_en.RegisterDefaultTranslations(validate, trans)

	// Register all validations at once
	registerValidations(validate, trans)

	err := validate.Struct(req)
	if err != nil {
		// Translate the error
		translations := err.(validator.ValidationErrors).Translate(trans)
		for _, msg := range translations {
			return errors.New(msg)
		}
	}

	return nil
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

func isValidTeamSize(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	for _, size := range enums.GetAllSizes() {
		if value == size {
			return true
		}
	}
	return false
}

// func emailExist(fl validator.FieldLevel) bool {
// 	email := fl.Field().String()
// 	existingUser := config.DB.Find(models.User{}, "email = ?", email)
// 	return existingUser.RowsAffected > 0
// }

func isValidDateRange(fl validator.FieldLevel) bool {
	endDate := fl.Field()
	if endDate.String() == "" {
		return true // Not time.Time fields, skip validation
	}

	parent := fl.Parent()

	var startDate reflect.Value

	startDateFields := []string{"StartDate", "Start", "start_date", "start"}

	for _, fieldName := range startDateFields {
		startDate = parent.FieldByName(fieldName)
		if startDate.IsValid() {
			break
		}

		parentType := parent.Type()
		for i := 0; i < parentType.NumField(); i++ {
			field := parentType.Field(i)
			jsonTag := field.Tag.Get("json")
			if jsonTag == fieldName {
				startDate = parent.Field(i)
				break
			}
		}
		if startDate.IsValid() {
			break
		}
	}
	if !startDate.IsValid() {
		return true
	}

	startTime, ok1 := time.Parse(enums.Date_format, startDate.String())
	endTime, ok2 := time.Parse(enums.Date_format, endDate.String())

	if ok1 != nil || ok2 != nil {
		return true // No time.Time fields, skip validation
	}

	return !endTime.Before(startTime)
}
