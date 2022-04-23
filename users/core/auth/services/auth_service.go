package services

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
	"user-login-api/core/common"
	"user-login-api/infrastructure/users"
)

type LoginResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    *echo.Map `json:"data"`
}

var validate = validator.New()

// Login godoc
// @Summary Login
// @Description Login
// @Tags login
// @Accept json,xml
// @Produce json
// @Param user body common.Auth true "Auth"
// @Success 200 {object} LoginResponse
// @Router /login [post]
func Login(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var auth common.Auth
	defer cancel()

	//validate the request body
	if err := c.Bind(&auth); err != nil {
		return c.JSON(http.StatusBadRequest, LoginResponse{
			Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&auth); validationErr != nil {
		return c.JSON(http.StatusBadRequest, LoginResponse{
			Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}

	accessor := users.UserAccessor{}
	result, err := accessor.FindByEmail(ctx, auth.Email)
	if result.Password != auth.Password {
		return c.JSON(http.StatusBadRequest, LoginResponse{
			Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": "User/Pass not matched"}})
	}

	// Set custom claims
	claims := &common.JwtCustomClaims{
		Name:   result.Name,
		Admin:  result.Admin,
		UserId: result.Id.Hex(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, LoginResponse{
			Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err}})
	}

	return c.JSON(http.StatusCreated, LoginResponse{
		Status: http.StatusCreated, Message: "success", Data: &echo.Map{
			"token": t,
		}})
}

func logout() {

}
