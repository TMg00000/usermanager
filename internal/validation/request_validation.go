package validation

import (
	"usermanager/internal/domain/request"

	"github.com/go-playground/validator/v10"
)

func ValidateNewUser(u request.Users) error {
	err := validator.New().Struct(u)
	if err != nil {
		return err
	}
	return nil
}
