package main

import (
	"fmt"
	"os"
	"strconv"
	"upsidr-coding-test/internal/payment/handler"
	"upsidr-coding-test/internal/rdb"

	"github.com/labstack/echo/v4"
	echoMidd "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type MockUserFetcher struct{}

func (m *MockUserFetcher) FetchUserFromToken(c echo.Context) (handler.User, error) {
	testUser := handler.User{
		UserId: "user1",
	}

	return testUser, nil
}

func main() {
	e := echo.New()
	e.Use(echoMidd.Recover())
	mockUserFetcher := MockUserFetcher{}
	e.Use(handler.UserMiddleware(&mockUserFetcher))

	isDebugMode, _ := strconv.ParseBool(os.Getenv("IS_DEBUG"))
	e.Debug = isDebugMode
	e.Logger.SetLevel(log.INFO)

	if err := initDB(isDebugMode); err != nil {
		e.Logger.Fatal(err)
	}
	defer func() {
		err := rdb.Close()
		if err != nil {
			e.Logger.Fatal(err)
		}
	}()

	s := handler.Server{}
	handler.RegisterHandlers(e, s)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))

}

func initDB(isDebugMode bool) error {
	return rdb.Init(
		isDebugMode,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
}
