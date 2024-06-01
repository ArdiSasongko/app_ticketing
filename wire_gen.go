// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/ArdiSasongko/app_ticketing/app"
	"github.com/ArdiSasongko/app_ticketing/controller/event.controller"
	"github.com/ArdiSasongko/app_ticketing/controller/user.controller"
	"github.com/ArdiSasongko/app_ticketing/db/conn"
	"github.com/ArdiSasongko/app_ticketing/helper"
	"github.com/ArdiSasongko/app_ticketing/repository/event.repository"
	"github.com/ArdiSasongko/app_ticketing/repository/user.repository"
	"github.com/ArdiSasongko/app_ticketing/repository/verification.repository"
	"github.com/ArdiSasongko/app_ticketing/service/event.service"
	"github.com/ArdiSasongko/app_ticketing/service/user.service"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

// Injectors from injector.go:

func StartServer() *echo.Echo {
	db := conn.DBConn()
	userRepo := userrepository.NewUserRepo(db)
	emailVerification := verificationrepository.NewEmailVerification(db)
	tokenUseCaseImpl := helper.NewTokenUseCase()
	userService := userservice.NewUserService(userRepo, emailVerification, tokenUseCaseImpl)
	userController := usercontroller.NewUserController(userService)
	eventRepo := eventrepository.NewEventRepo(db)
	eventService := eventservice.NewEventService(eventRepo)
	eventController := eventcontroller.NewEventController(eventService)
	echoEcho := app.Server(userController, eventController)
	return echoEcho
}

// injector.go:

var userSet = wire.NewSet(userrepository.NewUserRepo, wire.Bind(new(userrepository.UserRepositoryInterface), new(*userrepository.UserRepo)), verificationrepository.NewEmailVerification, wire.Bind(new(verificationrepository.VerificationEmailInterface), new(*verificationrepository.EmailVerification)), helper.NewTokenUseCase, wire.Bind(new(helper.TokenUseCaseInterface), new(*helper.TokenUseCaseImpl)), userservice.NewUserService, wire.Bind(new(userservice.UserServiceInterface), new(*userservice.UserService)), usercontroller.NewUserController, wire.Bind(new(usercontroller.UserControllerInterface), new(*usercontroller.UserController)))

var eventSet = wire.NewSet(eventrepository.NewEventRepo, wire.Bind(new(eventrepository.EventRepoInterface), new(*eventrepository.EventRepo)), eventservice.NewEventService, wire.Bind(new(eventservice.EventServiceInterface), new(*eventservice.EventService)), eventcontroller.NewEventController, wire.Bind(new(eventcontroller.EventControllerInterface), new(*eventcontroller.EventController)))
