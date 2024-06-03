package route

import (
	"github.com/ArdiSasongko/app_ticketing/app/middleware"
	eventcontroller "github.com/ArdiSasongko/app_ticketing/controller/event.controller"
	"github.com/ArdiSasongko/app_ticketing/db/conn"
	eventrepository "github.com/ArdiSasongko/app_ticketing/repository/event.repository"
	"github.com/labstack/echo/v4"
)

func EventRoute(e *echo.Echo, controller eventcontroller.EventControllerInterface) {
	DB := conn.DBConn()
	event := eventrepository.NewEventRepo(DB)

	apiv1 := e.Group("/api/v1")

	apiv1.POST("/event/create", controller.Create, middleware.JWTProtect(), middleware.AccessRole("seller"))
	apiv1.GET("/events", controller.FetchAll)
	apiv1.GET("/event/:id", controller.FetchEvent)
	apiv1.PUT("/event/:id", controller.UpdateEvent, middleware.JWTProtect(), middleware.AccessRole("seller"), middleware.AccessID(*event))
	apiv1.DELETE("/event/:id", controller.DeleteEvent, middleware.JWTProtect(), middleware.AccessRole("seller"), middleware.AccessID(*event))
	apiv1.DELETE("/ticket/:id", controller.DeleteTicket, middleware.JWTProtect(), middleware.AccessRole("seller"), middleware.AccessTicketID(*event))
	apiv1.GET("/events/seller", controller.FetchEventBySellerID, middleware.JWTProtect(), middleware.AccessRole("seller"))
}
