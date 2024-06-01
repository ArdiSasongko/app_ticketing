package usercontroller

import "github.com/labstack/echo/v4"

type UserControllerInterface interface {
	Create(c echo.Context) error
	CreateSeller(c echo.Context) error
	VerifyEmail(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
}
