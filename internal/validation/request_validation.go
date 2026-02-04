package validation

import (
	"github.com/go-playground/validator/v10"
)

func ValidateData(u interface{}) error {
	err := validator.New().Struct(u)
	if err != nil {
		return err
	}
	return nil
}
