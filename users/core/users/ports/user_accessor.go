package ports

import (
	"context"
	"github.com/labstack/echo/v4"
	"user-login-api/core/common"
)

type UserAccessor interface {
	Create(user common.User, ctx context.Context) *echo.Map
	Find(ctx context.Context, userId int) *echo.Map
	FindAll(ctx context.Context) *echo.Map
}
