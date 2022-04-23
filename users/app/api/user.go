package api

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"user-login-api/app/middleware"
	"user-login-api/core/users/services"
)

func Registerhandler(e *echo.Group, db *mongo.Client) {
	//All routes related to users comes here
	e.POST(
		"/user",
		services.CreateUser,
		middleware.Authorize(true),
	)

	e.GET("/user", services.GetAllUser, middleware.Authorize(true))
	e.GET("/user/:userId", services.GetAUser, middleware.AuthorizeUserId("userId"))
	e.PUT("/user/:userId", services.EditAUser, middleware.AuthorizeUserId("userId"))
	e.DELETE("/user/:userId", services.DeleteAUser, middleware.AuthorizeUserId("userId"))
}
