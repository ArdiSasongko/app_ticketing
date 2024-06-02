package middleware

import (
	"net/http"
	"os"
	"strconv"

	"github.com/ArdiSasongko/app_ticketing/helper"
	eventrepository "github.com/ArdiSasongko/app_ticketing/repository/event.repository"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTProtect() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(helper.CustomClaims)
		},
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, helper.ResponseToClient(http.StatusUnauthorized, "login needed", nil))
		},
	})
}

func AccessRole(role ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims, _ := user.Claims.(*helper.CustomClaims)
			userRole := claims.Role
			// checking role
			for _, r := range role {
				if userRole == r {
					return next(c)
				}
			}
			return c.JSON(http.StatusUnauthorized, helper.ResponseToClient(http.StatusUnauthorized, "unauthorized", nil))
		}
	}
}

func AccessID(r eventrepository.EventRepo) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			id, _ := strconv.Atoi(c.Param("id"))
			event, err := r.FetchEvent(id)

			if err != nil {
				return c.JSON(http.StatusNotFound, helper.ResponseToClient(http.StatusNotFound, err.Error(), nil))
			}

			user := c.Get("user").(*jwt.Token)
			claims, _ := user.Claims.(*helper.CustomClaims)

			if claims.UserID != event.SellerID {
				return c.JSON(http.StatusUnauthorized, helper.ResponseToClient(http.StatusUnauthorized, "Unauthorized", nil))
			}

			return next(c)
		}
	}
}
