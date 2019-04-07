package validators

import (
	"gopkg.in/go-playground/validator.v9"
	"regexp"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()
	Validate.RegisterValidation("empty", emptyStringValidator)
}

// Valdator to check if the string only has space
func emptyStringValidator(fl validator.FieldLevel) bool {
	space := regexp.MustCompile(`\s+`)
	text := space.ReplaceAllString(fl.Field().String(), "")
	if len(text) <= 0 {
		return false
	}
	return true
}
