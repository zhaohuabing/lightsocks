package util

import (
	"fmt"
	"path"
	"os"
	"path/filepath"
	"github.com/skratchdot/open-golang/open"
	"github.com/getlantern/pac"
	"github.com/gwuhaolin/lightsocks/app/manager/model"
	"github.com/satori/go.uuid"
)

const (
	LOCAL_HTTP_ADDR  = "127.0.0.1:15002"
	LOCAL_SOCKS_ADDR = "127.0.0.1:15003"
)

func OpenProxy() {
	pac.On(fmt.Sprintf("http://%s/proxy.pac", LOCAL_HTTP_ADDR))
}

func CloseProxy() {
	pac.Off(fmt.Sprintf("http://%s/proxy.pac", LOCAL_HTTP_ADDR))
}

func InitPacSysConfig() {
	pacPath := path.Join(os.TempDir(), "pac")
	iconFullPath, _ := filepath.Abs("./icon.png")
	err := pac.EnsureHelperToolPresent(pacPath, "需要帮你为受限的网络配置自动代理", iconFullPath)
	if err != nil {
	}
}

func OpenBrowser() {
	open.Run("http://localhost:8150")
}

func GetDeviceInfo() *model.DeviceInfo {
	return &model.DeviceInfo{
		Name:     "halwu's mac",
		Platform: "mac",
		UUID:     uuid.NewV4().String(),
	}
}
