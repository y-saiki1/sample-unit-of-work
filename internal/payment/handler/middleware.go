package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	UserId string
}

type UserFetcher interface {
	FetchUserFromToken(c echo.Context) (User, error)
}

func UserMiddleware(fetcher UserFetcher) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, err := fetcher.FetchUserFromToken(c)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, nil)
			}
			c.Set("user", user)
			return next(c)
		}
	}
}
