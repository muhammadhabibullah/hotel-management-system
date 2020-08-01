package validators

import "github.com/go-playground/validator/v10"

func InitValidators() *validator.Validate {
	v := validator.New()

	return v
}
