package helper

import "github.com/golang-jwt/jwt/v5"

type CustomResponse map[string]interface{}

var BlockedToken map[string]jwt.NumericDate

type ResponseClientModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseToClient(code int, message string, data interface{}) *ResponseClientModel {
	return &ResponseClientModel{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
