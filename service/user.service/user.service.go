package userservice

import (
	entityuser "github.com/ArdiSasongko/app_ticketing/db/model/entity/entity.user"
	"github.com/ArdiSasongko/app_ticketing/db/model/web"
	"github.com/ArdiSasongko/app_ticketing/helper"
)

type UserServiceInterface interface {
	Create(user *web.UserRequest) (helper.CustomResponse, error)
	CreateSeller(user *web.UserRequest) (helper.CustomResponse, error)
	Login(email, password string) (helper.CustomResponse, error)
	Logout(token string) (helper.CustomResponse, error)
	VerifyEmail(token string) (helper.CustomResponse, error)
	GetOrder(userID int) (entityuser.UserEntityOrder, error)
}
