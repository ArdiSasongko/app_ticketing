package ordercontroller

import "github.com/labstack/echo/v4"

type OrderControllerInterface interface {
	Create(c echo.Context) error
	PayOrder(c echo.Context) error
}
