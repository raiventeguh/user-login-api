package api

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"user-login-api/core/auth/services"
)

func RegisterAuthHandler(e *echo.Echo, db *mongo.Client) {
	//All routes related to users comes here
	e.POST(
		"/login",
		services.Login,
	)
}
