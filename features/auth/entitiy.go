package auth

import (
	"clean-arch/features/users"
	"clean-arch/features/users/delivery"
)

type AuthServiceInterface interface {
	Login(delivery.UserRequest) (users.UserEntity, error)
	Register(request users.UserEntity) error
}

type AuthDataInterface interface {
	Login(delivery.UserRequest) (users.UserEntity, error)
	Register(request users.UserEntity) error
}
