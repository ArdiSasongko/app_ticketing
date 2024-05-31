// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/ArdiSasongko/app_ticketing/app"
	"github.com/ArdiSasongko/app_ticketing/controller/user.controller"
	"github.com/ArdiSasongko/app_ticketing/db/conn"
	"github.com/ArdiSasongko/app_ticketing/repository/user.repository"
	"github.com/ArdiSasongko/app_ticketing/repository/verification.repository"
	"github.com/ArdiSasongko/app_ticketing/service/user.service"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

// Injectors from injector.go:

func StartServer() *echo.Echo {
	db := conn.DBConn()
	userRepo := userrepository.NewUserRepo(db)
	emailVerification := verificationrepository.NewEmailVerification(db)
	userService := userservice.NewUserService(userRepo, emailVerification)
	userController := usercontroller.NewUserController(userService)
	echoEcho := app.Server(userController)
	return echoEcho
}

// injector.go:

var userSet = wire.NewSet(userrepository.NewUserRepo, wire.Bind(new(userrepository.UserRepositoryInterface), new(*userrepository.UserRepo)), verificationrepository.NewEmailVerification, wire.Bind(new(verificationrepository.VerificationEmailInterface), new(*verificationrepository.EmailVerification)), userservice.NewUserService, wire.Bind(new(userservice.UserServiceInterface), new(*userservice.UserService)), usercontroller.NewUserController, wire.Bind(new(usercontroller.UserControllerInterface), new(*usercontroller.UserController)))
