package route

import (
	"github.com/ArdiSasongko/app_ticketing/app/middleware"
	eventcontroller "github.com/ArdiSasongko/app_ticketing/controller/event.controller"
	"github.com/labstack/echo/v4"
)

func EventRoute(e *echo.Echo, controller eventcontroller.EventControllerInterface) {
	apiv1 := e.Group("/api/v1")

	apiv1.POST("/event/create", controller.Create, middleware.JWTProtect(), middleware.AccessRole("seller"))
	apiv1.GET("/events", controller.FetchAll)
}
