package handler

import (
	"net/http"
	"strings"
	"upsidr-coding-test/internal/auth"

	"github.com/labstack/echo/v4"
)

type User struct {
	UserId string
}

type UserFetcher interface {
	FetchUserFromToken(c echo.Context) (User, error)
}

type Fetcher struct{}

func (Fetcher) FetchUserFromToken(c echo.Context) (User, error) {
	handler := auth.NewHandler(c.Logger())

	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return User{}, auth.ErrorTokenVerification
	}

	// "Bearer" と "token" を分割
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		// Bearerトークンが正常にフォーマットされていない場合の処理
		return User{}, auth.ErrorTokenVerification
	}

	// Token取得
	token := parts[1]
	uid, err := handler.Verify(token)
	if err != nil {
		c.Logger().Error(err)
		return User{}, auth.ErrorTokenVerification
	}
	return User{UserId: uid}, nil
}

func UserMiddleware(fetcher UserFetcher) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, err := fetcher.FetchUserFromToken(c)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, ErrorResponse{
					Message: err.Error(),
				})
			}
			c.Set("user", user)
			return next(c)
		}
	}
}
