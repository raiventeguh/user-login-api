package services

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
	"user-login-api/core/common"
	"user-login-api/infrastructure/users"
)

type UserResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    *echo.Map `json:"data"`
}

var validate = validator.New()

// Create User godoc
// @Summary Create User
// @Description Create a user item
// @Tags users
// @Accept json,xml
// @Produce json
// @Param user body common.User true "User"
// @Param Authorization header string true "Insert your access token"
// @Success 200 {object} UserResponse
// @Router /user [post]
// @Security ApiKeyAuth
func CreateUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user common.User
	defer cancel()

	//validate the request body
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, UserResponse{
			Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.JSON(http.StatusBadRequest, UserResponse{
			Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}

	newUser := common.User{
		Name:  user.Name,
		Email: user.Email,
		Age:   user.Age,
		Admin: user.Admin,
	}
	accessor := users.UserAccessor{}
	result := accessor.Create(newUser, ctx)

	return c.JSON(http.StatusCreated, UserResponse{
		Status: http.StatusCreated, Message: "success", Data: result})
}

// Find All User godoc
// @Summary Find All User
// @Description Find all a user item
// @Tags users
// @Accept json,xml
// @Produce json
// @Param Authorization header string true "Insert your access token"
// @Success 200 {object} UserResponse
// @Router /user [get]
// @Security ApiKeyAuth
func GetAllUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	accessor := users.UserAccessor{}
	result := accessor.FindAll(ctx)
	return c.JSON(http.StatusCreated, UserResponse{
		Status: http.StatusCreated, Message: "success", Data: result})
}

// Find User godoc
// @Summary Find a User
// @Description Find a user item
// @Tags users
// @Accept json,xml
// @Produce json
// @Param userId path string true "User Id"
// @Success 200 {object} UserResponse
// @Router /user/{userId} [get]
// @Security ApiKeyAuth
func GetAUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Param("userId")
	defer cancel()

	accessor := users.UserAccessor{}
	result := accessor.Find(ctx, userId)
	return c.JSON(http.StatusCreated, UserResponse{
		Status: http.StatusCreated, Message: "success", Data: result})
}

// Edit User godoc
// @Summary Edit a User
// @Description Edit a user item
// @Tags users
// @Accept json,xml
// @Produce json
// @Param userId path string true "User Id"
// @Param user body common.User true "User"
// @Param Authorization header string true "Insert your access token"
// @Success 200 {object} UserResponse
// @Router /user/{userId} [put]
// @Security ApiKeyAuth
func EditAUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Param("userId")
	var user common.User
	defer cancel()

	//validate the request body
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.JSON(http.StatusBadRequest, UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}

	update := bson.M{"name": user.Name, "age": user.Age, "email": user.Email, "admin": user.Admin}
	accessor := users.UserAccessor{}
	result := accessor.Update(ctx, userId, update)
	return c.JSON(http.StatusCreated, UserResponse{
		Status: http.StatusCreated, Message: "success", Data: result})

}

// Delete User godoc
// @Summary Delete a User
// @Description Delete a user item
// @Tags users
// @Accept json,xml
// @Produce json
// @Param userId path string true "User Id"
// @Param Authorization header string true "Insert your access token"
// @Success 200 {object} UserResponse
// @Router /user/{userId} [delete]
func DeleteAUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Param("userId")
	defer cancel()

	accessor := users.UserAccessor{}
	result := accessor.Delete(ctx, userId)
	return c.JSON(http.StatusCreated, UserResponse{
		Status: http.StatusCreated, Message: "success", Data: result})
}
