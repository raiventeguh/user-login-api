package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"user-login-api/core/common"
)

func Authorize(isAdminOnly bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*common.JwtCustomClaims)
			if isAdminOnly != claims.Admin {
				return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
			}

			return next(c)
		}
	}
}

func AuthorizeUserId(userIdPath string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*common.JwtCustomClaims)
			if claims.Admin {
				// By pass if admin
				return next(c)
			}

			if !strings.EqualFold(c.Param(userIdPath), claims.UserId) {
				return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
			}

			return next(c)
		}
	}
}
