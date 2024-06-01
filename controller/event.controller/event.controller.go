package eventcontroller

import "github.com/labstack/echo/v4"

type EventControllerInterface interface {
	Create(c echo.Context) error
	FetchAll(c echo.Context) error
}
