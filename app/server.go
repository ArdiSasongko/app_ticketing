package app

import (
	"log"

	"github.com/ArdiSasongko/app_ticketing/app/route"
	eventcontroller "github.com/ArdiSasongko/app_ticketing/controller/event.controller"
	ordercontroller "github.com/ArdiSasongko/app_ticketing/controller/order.controller"
	usercontroller "github.com/ArdiSasongko/app_ticketing/controller/user.controller"
	"github.com/ArdiSasongko/app_ticketing/helper"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/robfig/cron/v3"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cV *CustomValidator) Validate(i interface{}) error {
	return cV.validator.Struct(i)
}

// making alias
// type userCon usercontroller.UserControllerInterface
// type eventCon eventcontroller.EventControllerInterface
// type orderCon ordercontroller.OrderControllerInterface

// server func declaration
func Server(
	user usercontroller.UserControllerInterface,
	event eventcontroller.EventControllerInterface,
	order ordercontroller.OrderControllerInterface,
	cronJob *cron.Cron,
) *echo.Echo {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	server := echo.New()
	server.Validator = &CustomValidator{validator: validator.New()}
	server.HTTPErrorHandler = helper.ValidateBind

	route.UserRoute(server, user)
	route.EventRoute(server, event)
	route.OrderRoute(server, order)
	return server
}
