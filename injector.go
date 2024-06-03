//go:build wireinject
// +build wireinject

package main

import (
	"github.com/ArdiSasongko/app_ticketing/app"
	eventcontroller "github.com/ArdiSasongko/app_ticketing/controller/event.controller"
	ordercontroller "github.com/ArdiSasongko/app_ticketing/controller/order.controller"
	usercontroller "github.com/ArdiSasongko/app_ticketing/controller/user.controller"
	"github.com/ArdiSasongko/app_ticketing/db/conn"
	"github.com/ArdiSasongko/app_ticketing/helper"
	eventrepository "github.com/ArdiSasongko/app_ticketing/repository/event.repository"
	orderrepository "github.com/ArdiSasongko/app_ticketing/repository/order.repository"
	userrepository "github.com/ArdiSasongko/app_ticketing/repository/user.repository"
	verificationrepository "github.com/ArdiSasongko/app_ticketing/repository/verification.repository"
	eventservice "github.com/ArdiSasongko/app_ticketing/service/event.service"
	orderservice "github.com/ArdiSasongko/app_ticketing/service/order.service"
	userservice "github.com/ArdiSasongko/app_ticketing/service/user.service"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

var userSet = wire.NewSet(
	userrepository.NewUserRepo,
	wire.Bind(new(userrepository.UserRepositoryInterface), new(*userrepository.UserRepo)),
	verificationrepository.NewEmailVerification,
	wire.Bind(new(verificationrepository.VerificationEmailInterface), new(*verificationrepository.EmailVerification)),
	helper.NewTokenUseCase,
	wire.Bind(new(helper.TokenUseCaseInterface), new(*helper.TokenUseCaseImpl)),
	userservice.NewUserService,
	wire.Bind(new(userservice.UserServiceInterface), new(*userservice.UserService)),
	usercontroller.NewUserController,
	wire.Bind(new(usercontroller.UserControllerInterface), new(*usercontroller.UserController)),
)

var eventSet = wire.NewSet(
	eventrepository.NewEventRepo,
	wire.Bind(new(eventrepository.EventRepoInterface), new(*eventrepository.EventRepo)),
	eventservice.NewEventService,
	wire.Bind(new(eventservice.EventServiceInterface), new(*eventservice.EventService)),
	eventcontroller.NewEventController,
	wire.Bind(new(eventcontroller.EventControllerInterface), new(*eventcontroller.EventController)),
)

var orderSet = wire.NewSet(
	orderrepository.NewOrderRepo,
	wire.Bind(new(orderrepository.OrderRepositoryInterface), new(*orderrepository.OrderRepo)),
	orderservice.NewOrderService,
	wire.Bind(new(orderservice.OrderServiceInterface), new(*orderservice.OrderService)),
	ordercontroller.NewOrderController,
	wire.Bind(new(ordercontroller.OrderControllerInterface), new(*ordercontroller.OrderController)),
)

func StartServer() *echo.Echo {
	wire.Build(
		conn.DBConn,
		userSet,
		eventSet,
		orderSet,
		app.Server,
	)
	return nil
}
