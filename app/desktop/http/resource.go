package http

import (
	"github.com/shibukawa/configdir"
	"github.com/gwuhaolin/gfwlist4go/gfwlist"
	"fmt"
	"github.com/gwuhaolin/lightsocks/app/desktop/util"
	"log"
)

const (
	FILENAME_PAC = "proxy.pac"
)

var (
	cacheFolder *configdir.Config
)

func init() {
	cacheFolder = configdir.New("lightsocks.net", "lightsocks").QueryCacheFolder()
}

func downloadPac() {
	bs, err := gfwlist.Pac(fmt.Sprintf("SOCKS5 %s", util.LOCAL_SOCKS_ADDR))
	if err != nil {
		log.Println("downloadPac", err)
		return
	}
	err = cacheFolder.WriteFile(FILENAME_PAC, bs)
	if err != nil {
		log.Println("downloadPac", err)
		return
	}
	clientStatus.Pac = true
}

func ensurePac() {
	if cacheFolder.Exists(FILENAME_PAC) {
		go downloadPac()
	} else {
		downloadPac()
	}
}
