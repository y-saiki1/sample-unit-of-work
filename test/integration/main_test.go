package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"upsidr-coding-test/internal/payment/handler"
	"upsidr-coding-test/internal/rdb"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/labstack/echo/v4"
)

var e *echo.Echo

func TestMain(m *testing.M) {
	setUp()
	defer tearDown()

	code := m.Run()

	tearDown()

	os.Exit(code)
}

func setUp() {
	e = echo.New()
	isDebugMode, _ := strconv.ParseBool(os.Getenv("IS_DEBUG"))
	e.Debug = isDebugMode
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Recover())

	if err := rdb.Init(
		true,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		fmt.Sprintf("%s_test", os.Getenv("DB_NAME")),
	); err != nil {
		e.Logger.Fatal(err)
	}
}

func tearDown() {
	if err := truncateAllTables(); err != nil {
		e.Logger.Fatal(err)
	}
	if err := rdb.Close(); err != nil {
		e.Logger.Fatal(err)
	}
}

func truncateAllTables() error {
	if err := rdb.DB.Exec("SET FOREIGN_KEY_CHECKS = 0").Error; err != nil {
		return err
	}

	var tableNames []string
	err := rdb.DB.Raw("SHOW TABLES").Scan(&tableNames).Error
	if err != nil {
		return err
	}

	for _, tableName := range tableNames {
		err = rdb.DB.Exec(fmt.Sprintf("TRUNCATE TABLE %s", tableName)).Error
		if err != nil {
			return err
		}
	}

	if err = rdb.DB.Exec("SET FOREIGN_KEY_CHECKS = 1").Error; err != nil {
		return err
	}
	return nil
}

func initServer(url string, r any, user handler.User) (echo.Context, handler.Server, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(http.MethodPost, url, toIoReader(r))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("user", user)
	s := handler.Server{}
	return c, s, rec
}

func toIoReader(r any) io.Reader {
	b, err := json.Marshal(r)
	if err != nil {
		e.Logger.Error(err)
	}
	return bytes.NewBuffer(b)
}
