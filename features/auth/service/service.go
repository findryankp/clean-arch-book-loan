package service

import (
	"clean-arch/features/auth"
	"clean-arch/features/users"
)

type authService struct {
	data auth.AuthDataInterface
}

func New(u auth.AuthDataInterface) auth.AuthServiceInterface {
	return &authService{
		data: u,
	}
}

func (u *authService) Login(email string) (users.UserEntity, error) {
	return u.data.GetUserByEmail(email)
}

func (u *authService) Register(request users.UserEntity) error {
	return u.data.Register(request)
}
