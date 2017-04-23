package http

import (
	"github.com/gwuhaolin/lightsocks/app/desktop/util"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"github.com/gwuhaolin/lightsocks/app/manager/model"
	"github.com/gwuhaolin/lightsocks/app/desktop/lightsocks"
	"net"
)

type ClientStatus struct {
	// 当前客户端的lightsocks服务是否正在运行
	Lightsocks bool `json:"lightsocks"`
	// 当前客户端的pac配置是否已经下载好
	Pac        bool `json:"pac"`
	DeviceInfo *model.DeviceInfo `json:"deviceInfo"`
}

var clientStatus = &ClientStatus{
	DeviceInfo: util.GetDeviceInfo(),
}

func ListenHTTP(succ func()) {
	app := echo.New()
	app.Debug = true
	app.Use(middleware.Recover())
	// TODO 只在开发模式下使用
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8150"},
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowCredentials: true,
	}))
	app.Static("/", cacheFolder.Path)
	go ensurePac()

	// 获取当前客户端运行状态
	app.GET("/getStatus", func(c echo.Context) error {
		return c.JSON(http.StatusOK, clientStatus)
	})

	// 更新客户端的lightsocks服务配置，让客户端重启lightsocks服务
	app.POST("/setService", func(c echo.Context) error {
		service := &model.Service{}
		err := c.Bind(service)
		if err != nil {
			return err
		}
		return lightsocks.ListenLightsocks(service, func(listenAddr net.Addr) {
			clientStatus.Lightsocks = true
			c.JSON(http.StatusOK, listenAddr)
		})
	})

	succ()
	app.Start(util.LOCAL_HTTP_ADDR)
}
