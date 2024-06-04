package route

import (
	"github.com/ArdiSasongko/app_ticketing/app/middleware"
	ordercontroller "github.com/ArdiSasongko/app_ticketing/controller/order.controller"
	"github.com/labstack/echo/v4"
)

func OrderRoute(e *echo.Echo, controller ordercontroller.OrderControllerInterface) {
	apiv1 := e.Group("api/v1")

	apiv1.POST("/event/:event_id/ticket/:ticket_id/order", controller.Create, middleware.JWTProtect(), middleware.AccessRole("buyer"))
	apiv1.POST("/user/order/:order_id/payment", controller.PayOrder, middleware.JWTProtect(), middleware.AccessRole("buyer"))
}
