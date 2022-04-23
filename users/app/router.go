package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"user-login-api/app/api"
	"user-login-api/core/common"
	"user-login-api/infrastructure/configs"
)

func UserRoute(e *echo.Echo) {
	//All routes related to users comes here
	api.RegisterAuthHandler(e, configs.DB)

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &common.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r := e.Group("", middleware.JWTWithConfig(config))
	api.Registerhandler(r, configs.DB)
}
