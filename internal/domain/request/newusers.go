package request

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Users struct {
	Username string `json:"username" bson:"username" validate:"required,min=3,max=30"`
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=72"`
	Hash     string `json:"-"`
}
