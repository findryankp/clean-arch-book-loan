package delivery

import (
	"clean-arch/features/auth"
	"clean-arch/features/users/data"
	"clean-arch/features/users/delivery"
	"clean-arch/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	Service auth.AuthServiceInterface
}

func New(s auth.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{
		Service: s,
	}
}

func (u *AuthHandler) Login(c echo.Context) error {
	loginRequest := LoginRequest{}
	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	user, err := u.Service.Login(loginRequest.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helpers.ResponseFail("User not found"))
	}

	//compare password
	if !helpers.CheckPasswordHash(loginRequest.Password, user.Password) {
		return c.JSON(http.StatusUnauthorized, helpers.ResponseFail("User password not match"))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", loginRequest))
}

func (u *AuthHandler) Register(c echo.Context) error {
	registerRequest := delivery.UserRequest{}
	if err := c.Bind(&registerRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}
	user := delivery.UserRequestToUserEntity(registerRequest)

	if err := u.Service.Register(user); err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", registerRequest))
}

func (u *AuthHandler) GetUserLogin(c echo.Context) error {
	var user data.User
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", user))
}
