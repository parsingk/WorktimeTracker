package main

import (
	"github.com/tadvi/systray"
	"log"
)

var (
	winTrayMenu *systray.Systray
	mitem       []*systray.MenuItem
)

func TrayMode(m *WinResMgr) {
	var err error
	if winTrayMenu != nil {
		return
	}

	winTrayMenu, err = systray.New()

	if err != nil {
		log.Panic(err)
	}

	err = winTrayMenu.Show(2, "")
	if err != nil {
		log.Panic(err)
	}

	winTrayMenu.OnClick(func() {
		winTrayMenu.Menu = mitem
		m.window.SetVisible(true)
	})

	winTrayMenu.OnRightClick(func() {
		if len(winTrayMenu.Menu) > 0 {
			return
		}

		winTrayMenu.AppendMenu("열기", func() {
			m.window.SetVisible(true)
		})

		winTrayMenu.AppendMenu("종료", func() {
			err := m.window.Close()
			if err != nil {
				return
			}
		})
	})
}
