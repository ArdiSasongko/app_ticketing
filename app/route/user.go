package route

import (
	"github.com/ArdiSasongko/app_ticketing/app/middleware"
	usercontroller "github.com/ArdiSasongko/app_ticketing/controller/user.controller"
	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo, controller usercontroller.UserControllerInterface) {
	apiv1 := e.Group("/api/v1")

	apiv1.POST("/user", controller.Create)
	apiv1.POST("/user/seller", controller.CreateSeller)
	apiv1.POST("/user/verify-email", controller.VerifyEmail)
	apiv1.POST("/user/login", controller.Login)
	apiv1.POST("/user/logout", controller.Logout, middleware.JWTProtect())
	apiv1.GET("/user/order", controller.GetOrder, middleware.JWTProtect())
}
