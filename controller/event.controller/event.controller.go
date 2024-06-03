package eventcontroller

import "github.com/labstack/echo/v4"

type EventControllerInterface interface {
	Create(c echo.Context) error
	FetchAll(c echo.Context) error
	FetchEvent(c echo.Context) error
	UpdateEvent(c echo.Context) error
	DeleteEvent(c echo.Context) error
	DeleteTicket(c echo.Context) error
	FetchEventBySellerID(c echo.Context) error
}
