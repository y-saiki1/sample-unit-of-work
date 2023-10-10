package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"testing"
	"time"
	"upsidr-coding-test/internal/auth"
	"upsidr-coding-test/internal/payment/handler"
	"upsidr-coding-test/internal/rdb"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/labstack/echo/v4"
)

var (
	e        *echo.Echo
	testPort = "81"
)

func TestMain(m *testing.M) {
	setUp()
	go func() {
		e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", testPort)))
	}()
	defer tearDown()
	time.Sleep(1 * time.Second)

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
	e.Use(handler.UserMiddleware(handler.Fetcher{}))

	if err := rdb.Init(
		false,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		fmt.Sprintf("%s_test", os.Getenv("DB_NAME")),
	); err != nil {
		e.Logger.Fatal(err)
	}
	s := handler.Server{}
	handler.RegisterHandlers(e, s)
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

func postRequest(link, method string, r any, token string) ([]byte, int, error) {
	reader, err := toIoReader(r)
	if err != nil {
		return nil, 0, err
	}

	req, err := http.NewRequest(method, fmt.Sprintf("http://localhost:%s%s", testPort, link), reader)
	if err != nil {
		return nil, 0, err
	}
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "Bearer "+token)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, 0, err
	}

	return body, res.StatusCode, nil
}

func getRequest(link, method string, params map[string]string, token string) ([]byte, int, error) {
	u, err := url.Parse(fmt.Sprintf("http://localhost:%s%s", testPort, link))
	if err != nil {
		return nil, 0, err
	}

	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, 0, err
	}

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "Bearer "+token)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, 0, err
	}

	return body, res.StatusCode, nil
}

func toIoReader(r any) (io.Reader, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(b), nil
}

func hashPassword(rawPass string) (string, error) {
	handler := auth.NewHandler(e.Logger)
	return handler.HashPassword(rawPass)
}

func getToken(uid string) (string, error) {
	handler := auth.NewHandler(e.Logger)
	return handler.CreateToken(uid)
}
