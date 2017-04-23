package http

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/gwuhaolin/lightsocks/app/manager/service"
)

type loginForm struct {
	Email    string
	Password string
}

const COOKIE_TOKEN = "token"

func init() {
	//用户登入
	App.POST("login", func(c echo.Context) error {
		form := loginForm{}
		err := c.Bind(&form)
		if err != nil {
			return ERR_INVALID_PARAMS
		}
		user, err := service.UserAuth(form.Email, form.Password)
		if err != nil {
			return ERR_AUTH_FAIL
		}
		c.SetCookie(&http.Cookie{
			Name:  COOKIE_TOKEN,
			Value: user.Token,
		})
		return c.JSON(http.StatusOK, user)
	})
}
