package usercontroller

import (
	"net/http"

	"github.com/ArdiSasongko/app_ticketing/db/model/web"
	"github.com/ArdiSasongko/app_ticketing/helper"
	userservice "github.com/ArdiSasongko/app_ticketing/service/user.service"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	service userservice.UserServiceInterface
}

func NewUserController(service userservice.UserServiceInterface) *UserController {
	return &UserController{
		service: service,
	}
}

func (controller *UserController) Create(c echo.Context) error {
	user := new(web.UserRequest)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(user); err != nil {
		return err
	}

	result, errCreated := controller.service.Create(user)

	if errCreated != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseToClient(http.StatusInternalServerError, errCreated.Error(), nil))
	}

	return c.JSON(http.StatusCreated, helper.ResponseToClient(http.StatusCreated, "User created", result))
}

func (controller *UserController) CreateSeller(c echo.Context) error {
	user := new(web.UserRequest)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(user); err != nil {
		return err
	}

	result, errCreated := controller.service.CreateSeller(user)

	if errCreated != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseToClient(http.StatusInternalServerError, errCreated.Error(), nil))
	}

	return c.JSON(http.StatusCreated, helper.ResponseToClient(http.StatusCreated, "User created", result))
}

func (controller *UserController) VerifyEmail(c echo.Context) error {
	token := new(web.TokenRequest)

	if err := c.Bind(token); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(token); err != nil {
		return err
	}

	result, errVerify := controller.service.VerifyEmail(token.Token)

	if errVerify != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseToClient(http.StatusInternalServerError, errVerify.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseToClient(http.StatusOK, "Email verified", result))
}

func (controller *UserController) Login(c echo.Context) error {
	user := new(web.UserLoginRequest)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(user); err != nil {
		return err
	}

	result, err := controller.service.Login(user.Email, user.Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseToClient(http.StatusInternalServerError, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseToClient(http.StatusOK, "Login success", result))
}
