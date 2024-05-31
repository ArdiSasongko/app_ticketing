//go:build wireinject
// +build wireinject

package main

import (
	"github.com/ArdiSasongko/app_ticketing/app"
	usercontroller "github.com/ArdiSasongko/app_ticketing/controller/user.controller"
	"github.com/ArdiSasongko/app_ticketing/db/conn"
	userrepository "github.com/ArdiSasongko/app_ticketing/repository/user.repository"
	verificationrepository "github.com/ArdiSasongko/app_ticketing/repository/verification.repository"
	userservice "github.com/ArdiSasongko/app_ticketing/service/user.service"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

var userSet = wire.NewSet(
	userrepository.NewUserRepo,
	wire.Bind(new(userrepository.UserRepositoryInterface), new(*userrepository.UserRepo)),
	verificationrepository.NewEmailVerification,
	wire.Bind(new(verificationrepository.VerificationEmailInterface), new(*verificationrepository.EmailVerification)),
	userservice.NewUserService,
	wire.Bind(new(userservice.UserServiceInterface), new(*userservice.UserService)),
	usercontroller.NewUserController,
	wire.Bind(new(usercontroller.UserControllerInterface), new(*usercontroller.UserController)),
)

func StartServer() *echo.Echo {
	wire.Build(
		conn.DBConn,
		userSet,
		app.Server,
	)
	return nil
}
