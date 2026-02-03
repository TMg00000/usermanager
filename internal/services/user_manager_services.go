package services

import "usermanager/internal/domain/request"

type UsersManagerServices interface {
	Create(request.Users) error
	Get(email, password string) error
}
