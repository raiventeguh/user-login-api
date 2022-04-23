package main

import (
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	"user-login-api/app"
	_ "user-login-api/docs"
	"user-login-api/infrastructure/configs"
)

// @title User Application
// @description User Application Login & User
// @version 1.0
// @host localhost:8081
// @BasePath /
func main() {
	e := echo.New()

	configs.ConnectDB()
	app.UserRoute(e)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":8081"))
}
