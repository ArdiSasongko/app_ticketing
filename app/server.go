package app

import (
	"log"

	"github.com/ArdiSasongko/app_ticketing/app/route"
	eventcontroller "github.com/ArdiSasongko/app_ticketing/controller/event.controller"
	usercontroller "github.com/ArdiSasongko/app_ticketing/controller/user.controller"
	"github.com/ArdiSasongko/app_ticketing/helper"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cV *CustomValidator) Validate(i interface{}) error {
	return cV.validator.Struct(i)
}

func Server(userController usercontroller.UserControllerInterface, eventController eventcontroller.EventControllerInterface) *echo.Echo {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	server := echo.New()
	server.Validator = &CustomValidator{validator: validator.New()}
	server.HTTPErrorHandler = helper.ValidateBind

	route.UserRoute(server, userController)
	route.EventRoute(server, eventController)
	return server
}
