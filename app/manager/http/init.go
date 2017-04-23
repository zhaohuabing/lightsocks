package http

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"fmt"
)

type ErrorResponse struct {
	Code    int `json:"code"`
	Message string `json:"message"`
}

func (err *ErrorResponse) Error() string {
	return fmt.Sprintf("%d:%s", err.Code, err.Message)
}

const (
	LISTEN_ADDR = ":1500"
)

var (
	App                *echo.Echo
	ERR_INVALID_PARAMS = &ErrorResponse{1000, "参数不合法"}
	ERR_EMAIL_EXIT     = &ErrorResponse{1001, "该邮箱已经注册"}
	ERR_AUTH_FAIL      = &ErrorResponse{1002, "身份验证不通过"}
)

func init() {
	App = echo.New()
	App.Debug = true
	App.Use(middleware.Recover())
	// TODO 只在开发模式下使用
	App.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8150"},
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowCredentials: true,
	}))
	App.HTTPErrorHandler = func(err error, c echo.Context) {
		c.JSON(http.StatusInternalServerError, err)
	}
}

func ListenHTTP() {
	log.Fatal(App.Start(LISTEN_ADDR))
}
