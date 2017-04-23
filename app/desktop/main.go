package main

import (
	"github.com/getlantern/systray"
	"github.com/gwuhaolin/lightsocks/app/desktop/http"
	"github.com/gwuhaolin/lightsocks/app/desktop/util"
	"os"
	"github.com/getlantern/systray/example/icon"
)

type Menu struct {
	open   *systray.MenuItem
	exit   *systray.MenuItem
	OnInit func()
	OnOpen func()
	OnExit func()
}

func (menu *Menu) Run() {
	systray.SetIcon(icon.Data)
	systray.SetTooltip("lightsocks")
	menu.open = systray.AddMenuItem("控制面板", "")
	menu.exit = systray.AddMenuItem("退出", "")
	go func() {
		for {
			select {
			case <-menu.open.ClickedCh:
				menu.OnOpen()
			case <-menu.exit.ClickedCh:
				menu.OnExit()
			}
		}
	}()
	menu.OnInit()
}

func main() {
	App := &Menu{
		OnInit: func() {
			http.ListenHTTP(func() {
				util.OpenBrowser()
			})
		},
		OnExit: func() {
			util.CloseProxy()
			os.Exit(0)
		},
		OnOpen: func() {
			util.OpenBrowser()
		},
	}
	systray.Run(App.Run)
}
