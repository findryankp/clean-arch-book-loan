package service

import (
	"clean-arch/features/auth"
	"clean-arch/features/users"
	"clean-arch/features/users/delivery"
)

type authService struct {
	data auth.AuthDataInterface
}

func New(u auth.AuthDataInterface) auth.AuthServiceInterface {
	return &authService{
		data: u,
	}
}

func (u *authService) Login(a delivery.UserRequest) (users.UserEntity, error) {
	return u.data.Login(a)
}

func (u *authService) Register(request users.UserEntity) error {
	return u.data.Register(request)
}
