package services

import (
	"usermanager/internal/domain/request"
)

type UsersManagerServices interface {
	Create(request.Users) error
	Login(email, password string) error
	GetAllUsers() ([]request.Users, error)
}
