package http

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/gwuhaolin/lightsocks/app/manager/dao"
)

type signupForm struct {
	Email    string
	Password string
}

func init() {
	//用户注册
	App.POST("/signup", func(c echo.Context) error {
		form := signupForm{}
		err := c.Bind(&form)
		if err != nil {
			return ERR_INVALID_PARAMS
		}
		err = dao.CreateUser(form.Email, form.Password)
		if err != nil {
			return ERR_EMAIL_EXIT
		}
		return c.JSON(http.StatusCreated, "")
	})
}
