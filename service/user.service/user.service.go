package userservice

import (
	"github.com/ArdiSasongko/app_ticketing/db/model/web"
	"github.com/ArdiSasongko/app_ticketing/helper"
)

type UserServiceInterface interface {
	Create(user *web.UserRequest) (helper.CustomResponse, error)
	CreateSeller(user *web.UserRequest) (helper.CustomResponse, error)
	//FindByEmail(email, password string) (helper.CustomResponse, error)
	VerifyEmail(token string) (helper.CustomResponse, error)
}
