package http

import (
	"github.com/labstack/echo"
	"github.com/gwuhaolin/lightsocks/app/manager/model"
	"net/http"
	"github.com/gwuhaolin/lightsocks/app/manager/dao"
	"github.com/gwuhaolin/lightsocks/app/manager/service"
)

const CONTENT_USER = "user"

func init() {
	//身份验证拦截
	userGroup := App.Group("/user", func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, err := c.Request().Cookie(COOKIE_TOKEN)
			if err != nil {
				return ERR_AUTH_FAIL
			}
			user, err := dao.GetUserByToken(token.Value)
			if err != nil {
				return ERR_AUTH_FAIL
			}
			c.Set(CONTENT_USER, user)
			return next(c)
		}
	})

	// 获取当前用户信息
	userGroup.GET("/getMe", func(c echo.Context) error {
		user := c.Get(CONTENT_USER).(*model.User)
		return c.JSON(http.StatusOK, user)
	})

	//用户请求一个新的Service
	userGroup.POST("/applyService", func(c echo.Context) error {
		device := &model.DeviceInfo{}
		err := c.Bind(device)
		if err != nil {
			return ERR_INVALID_PARAMS
		}
		user := c.Get(CONTENT_USER).(*model.User)
		service, err := service.UserApplyService(user, device)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, service)
	})
}
